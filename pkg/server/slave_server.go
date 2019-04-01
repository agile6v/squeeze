// Copyright 2019 Squeeze Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
	"encoding/json"

	"google.golang.org/grpc"
	"github.com/agile6v/squeeze/pkg/pb"
	log "github.com/golang/glog"
	"github.com/agile6v/squeeze/pkg/proto/builder"
	"github.com/agile6v/squeeze/pkg/util"
)

type SlaveServer struct {
	ServerBase
	work        *Work
	lastTaskReq *pb.ExecuteTaskRequest
	mutex       sync.RWMutex
}

func (s *SlaveServer) Initialize(args *ServerArgs) error {
	s.ServerBase.Initialize(args)
	s.Mode = Slave

	// check if the HttpMasterAddr is valid
	_, _, err := util.GetHostPort(s.args.HttpMasterAddr)
	if err != nil {
		return err
	}

	// check if the GrpcMasterAddr is valid
	_, _, err = util.GetHostPort(s.args.GrpcMasterAddr)
	if err != nil {
		return err
	}

	// Create the gRPC server
	// TODO: create the grpc options
	s.grpcServer = grpc.NewServer(grpc.MaxConcurrentStreams(256), grpc.MaxMsgSize(1024*1024))
	pb.RegisterSqueezeServiceServer(s.grpcServer, s)

	proxy, err := NewProxy(s.args.HttpMasterAddr)
	if err != nil {
		return err
	}

	http.HandleFunc("/", proxy.handle)
	return nil
}

func (s *SlaveServer) Start(stopChan <-chan struct{}) error {
	// grpc server
	go func() {
		log.Infof("grpc listening on %s", s.args.GRPCAddr)
		listener, err := net.Listen("tcp", s.args.GRPCAddr)
		if err != nil {
			log.Fatalf("GRPC failed to listen: %v", err)
		}

		if err = s.grpcServer.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		<-stopChan
		s.grpcServer.Stop()
	}()

	go func() {
		for {
			err := s.doHeartBeat(stopChan, s.args.ReportInterval*time.Second)
			if err != nil {
				log.Errorf("report task : %v", err)
				time.Sleep(s.args.ReportInterval * time.Second)
			}
		}
	}()

	return nil
}

func (s *SlaveServer) HeartBeat(stream pb.SqueezeService_HeartBeatServer) error {
	return nil
}

func (s *SlaveServer) ExecuteTask(ctx context.Context, req *pb.ExecuteTaskRequest) (*pb.ExecuteTaskResponse, error) {
	log.V(2).Infof("Execute Task ... %s: %v",
		pb.ExecuteTaskRequest_Command_name[int32(req.Cmd)], req)

	if req.Cmd == pb.ExecuteTaskRequest_STOP {
		if s.work != nil {
			s.work.Stop()
		}

		return &pb.ExecuteTaskResponse{Status: pb.ExecuteTaskResponse_SUCC}, nil
	}

	if req.Concurrency < 1 {
		return &pb.ExecuteTaskResponse{Status: pb.ExecuteTaskResponse_FAIL,
			Error: fmt.Sprintf("Concurrency is invalid")}, nil
	}

	s.work = &Work{
		Req:       req,
		Builder:   builder.NewBuilder(req.Protocol),
		Requests:  int(req.Requests),
		Workers:   int(req.Concurrency),
		ResultCapacity: s.args.ResultCapacity,
		RateLimit: float64(req.RateLimit),
	}

	if req.Duration > 0 {
		s.work.Ctx, s.work.Cancel = context.WithTimeout(ctx, time.Duration(req.Duration)*time.Second)
	} else {
		s.work.Ctx, s.work.Cancel = context.WithCancel(ctx)
	}

	err := s.addTask(req)
	if err != nil {
		log.Error("Failed to add task : ", err)
		return &pb.ExecuteTaskResponse{Status: pb.ExecuteTaskResponse_FAIL, Error: err.Error()}, nil
	}

	elapsed, err := s.work.Run(s.work.Ctx)
	s.delTask()

	if err != nil {
		log.Error("Failed to run task: ", err)
		return &pb.ExecuteTaskResponse{Status: pb.ExecuteTaskResponse_FAIL, Error: err.Error()}, nil
	}

	msg, err := s.work.Builder.Done(elapsed)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(msg)
	if err != nil {
		log.Errorf("could not marshal message : %v", err)
		return nil, err
	}

	resp := &pb.ExecuteTaskResponse{
		Status: pb.ExecuteTaskResponse_SUCC,
		Data: string(data),
	}

	return resp, nil
}

func (s *SlaveServer) doHeartBeat(stopChan <-chan struct{}, frequency time.Duration) error {
	conn, err := grpc.Dial(s.args.GrpcMasterAddr, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("dial with master server, %s", err.Error())
	}
	defer conn.Close()

	// create stream
	client := pb.NewSqueezeServiceClient(conn)
	stream, err := client.HeartBeat(context.Background())
	if err != nil {
		return fmt.Errorf("connect with master server, %s", err.Error())
	}

	errorChan := make(chan error, 1)

	go func() {
		for {
			resp, err := stream.Recv()
			if err != nil {
				errorChan <- err
				return
			}

			if len(resp.Tasks) == 0 {
				log.V(2).Info("No tasks received.")
			} else {
				log.Infof("recv %d tasks.", len(resp.Tasks))
			}
		}
	}()

	ticker := time.NewTicker(frequency)
	defer ticker.Stop()
	for {
		select {
		case err := <-errorChan:
			return fmt.Errorf("recv meet error :%s", err.Error())
		case <-stopChan:
			if err := stream.CloseSend(); err != nil {
				return fmt.Errorf("failed to close stream, error: %s", err.Error())
			}
			log.Info("Finish sending heartbeat to master node.")
			return nil
		case <-ticker.C:
			log.V(2).Info("Send heartbeat to master node.")

			task := s.getTask()
			status := pb.HeartBeatRequest_Task_DONE
			if task != nil {
				status = pb.HeartBeatRequest_Task_RUNNING
			}

			err := stream.Send(&pb.HeartBeatRequest{
				Task: &pb.HeartBeatRequest_Task{Id: 0, Status: status},
				Info: &pb.HeartBeatRequest_SlaveInfo{GrpcPort: uint32(s.grpcPort)},
			})
			if err != nil {
				return fmt.Errorf("failed to send stream to master, error: %s", err.Error())
			}
		}
	}
}

func (s *SlaveServer) addTask(req *pb.ExecuteTaskRequest) error {
	log.V(3).Infof("Save the last task.")
	var err error
	s.mutex.Lock()
	if s.lastTaskReq == nil {
		s.lastTaskReq = req
	} else {
		err = fmt.Errorf("Task is not empty, cannot add it.")
	}
	s.mutex.Unlock()
	return err
}

func (s *SlaveServer) delTask() {
	log.V(3).Infof("Cleanup the last task.")
	s.mutex.Lock()
	s.lastTaskReq = nil
	s.mutex.Unlock()
}

func (s *SlaveServer) getTask() *pb.ExecuteTaskRequest {
	s.mutex.RLock()
	req := s.lastTaskReq
	s.mutex.RUnlock()
	return req
}
