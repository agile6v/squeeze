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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/agile6v/squeeze/pkg/pb"
	"github.com/agile6v/squeeze/pkg/proto/builder"
	"github.com/agile6v/squeeze/pkg/util"
	log "github.com/golang/glog"
	"github.com/golang/protobuf/jsonpb"
	protobuf "github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"sync"
	"github.com/agile6v/squeeze/pkg/proto"
)

var currentReq *pb.ExecuteTaskRequest

type MasterServer struct {
	ServerBase
	results chan protobuf.Message
}

func (m *MasterServer) Initialize(args *ServerArgs) error {
	m.ServerBase.Initialize(args)
	m.Mode = Master

	http.HandleFunc("/info", m.handleInfo)
	http.HandleFunc("/task/start", m.handleTask)
	http.HandleFunc("/task/stop", m.handleTask)

	// Create the gRPC server
	// TODO: create the grpc options
	m.grpcServer = grpc.NewServer(grpc.MaxConcurrentStreams(256), grpc.MaxMsgSize(1024*1024))
	pb.RegisterSqueezeServiceServer(m.grpcServer, m)

	return nil
}

func (m *MasterServer) startTask(taskReq *pb.ExecuteTaskRequest, conns []*SlaveConn) (interface{}, error) {
	count := len(conns)
	var wg sync.WaitGroup
	m.results = make(chan protobuf.Message, count)
	mergedResults := make(chan interface{}, 1)

	go func(mergedResults chan interface{}) {
		m.runCollector(mergedResults, taskReq)
	}(mergedResults)

	if taskReq.Callback != "" {
		err := m.dispatchTask(taskReq, conns, &wg)
		if err != nil {
			return nil, err
		}

		go func() {
			wg.Wait()
			close(m.results)

			// Wait for the collector goroutine to merge the results
			// and then read the merged results.
			response := <-mergedResults

			data, err := json.Marshal(map[string]interface{} {
				"error": "",
				"data": response,
			})
			if err != nil {
				log.Errorf("unable to marshal data : %v", err)
				return
			}

			resp, err := util.DoRequest("POST", taskReq.Callback, string(data), 30)
			if err != nil {
				log.Errorf("Failed to send results to callback address: %s, %s", err.Error(), resp)
				return
			}
			log.Infof("Send results to callback address successfully: %s", resp)
		}()

		return "success", nil
	} else {
		err := m.dispatchTask(taskReq, conns, &wg)
		if err != nil {
			return nil, err
		}

		wg.Wait()
		close(m.results)

		// Wait for the collector goroutine to merge the results
		// and then read the merged results.
		response := <-mergedResults
		return response, nil
	}
}

func (m *MasterServer) stopTask(conns []*SlaveConn) error {
	var wg sync.WaitGroup
	count := len(conns)
	wg.Add(count)
	for _, conn := range conns {
		log.Infof("Cancel task on slave %s", conn.PeerAddr)
		var err error
		go func(conn *SlaveConn, errP *error) {
			defer wg.Done()
			slaveAddr, err := util.BuildHostname(conn.PeerAddr, strconv.Itoa(conn.GrpcPort))
			if err != nil {
				log.Error("failed to build slave hostname: %s", err.Error())
				*errP = err
				return
			}

			ret, err := m.DoExecuteTask(slaveAddr, &pb.ExecuteTaskRequest{
				Cmd: pb.ExecuteTaskRequest_STOP},
			)
			if err != nil {
				*errP = err
				log.Error("Failed to execute cancel task :%s", err.Error())
				return
			}

			if ret.Status != pb.ExecuteTaskResponse_SUCC {
				*errP = fmt.Errorf("Failed to execute task with command stop.")
				return
			}

			log.Infof("Dispatch task to %s return %d", slaveAddr, ret.Status)
		}(conn, &err)
	}

	log.Infof("Waiting for all slaves to stop.")
	wg.Wait()
	log.Infof("Slaves have stopped.")

	return nil
}

func (m *MasterServer) dispatchTask(taskReq *pb.ExecuteTaskRequest, conns []*SlaveConn, wg *sync.WaitGroup) error {
	reqs := builder.NewBuilder(taskReq.Protocol).Split(taskReq, len(conns))
	wg.Add(util.Min(len(conns), len(reqs)))
	log.V(2).Infof("connections: %d, requests: %d", len(conns), len(reqs))

	for i, conn := range conns {
		go func(conn *SlaveConn, index int) {
			defer wg.Done()
			log.V(2).Infof("slave address: %s", conn.PeerAddr)
			slaveAddr, err := util.BuildHostname(conn.PeerAddr, strconv.Itoa(conn.GrpcPort))
			if err != nil {
				log.Errorf("failed to build slave hostname: %s", err.Error())
				return
			}

			log.V(2).Infof("dispatch to slave: %s", slaveAddr)

			ret, err := m.DoExecuteTask(slaveAddr, reqs[index])
			if err != nil {
				log.Errorf("failed to dispatch task to %s: %s", slaveAddr, err.Error())
				return
			}
			log.Infof("dispatch task to %s, return %d, %s", slaveAddr, ret.Status, ret.Error)
			ret.Addr = slaveAddr

			m.results <- ret
		}(conn, i)

		if (len(reqs) - 1) == i {
			break
		}
	}
	return nil
}

func (m *MasterServer) DoExecuteTask(address string, request *pb.ExecuteTaskRequest) (*pb.ExecuteTaskResponse, error) {
	// Set up a connection to the slave server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Errorf("can not connect with slave server %v", err)
		return nil, err
	}

	// create stream
	client := pb.NewSqueezeServiceClient(conn)
	resp, err := client.ExecuteTask(context.Background(), request)
	if err != nil {
		log.Errorf("open stream error %v", err)
		return nil, err
	}

	return resp, nil
}

func (m *MasterServer) handleTask(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	r.Body.Close()

	taskReq := &pb.ExecuteTaskRequest{}
	err = jsonpb.Unmarshal(bytes.NewReader(body), taskReq)
	if err != nil {
		log.Errorf("unable to decode json : %s", err)
		util.RespondWithError(w, http.StatusBadRequest, "Unable to decode request")
		return
	}

	// Check if it can run execute request.
	if currentReq != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "There are task in progress, please try again later.")
		return
	}

	currentReq = taskReq
	defer func() {
		currentReq = nil
	}()

	slaveConns := GetConnections()
	if len(slaveConns) == 0 {
		util.RespondWithError(w, http.StatusInternalServerError, "No slave available.")
		return
	}

	if taskReq.Cmd == pb.ExecuteTaskRequest_STOP {
		err := m.stopTask(slaveConns)
		if err != nil {
			util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		} else {
			util.RespondWithJSON(w, http.StatusOK, "success")
		}
		return
	}

	data, err := m.startTask(taskReq, slaveConns)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, data)
	return
}

// runCollector is used to collect results which all slaves generated.
func (m *MasterServer) runCollector(aggregation chan interface{}, taskReq *pb.ExecuteTaskRequest) {
	builder := builder.NewBuilder(taskReq.Protocol)
	results := make([]string, 0)
	SqueezeResult := &proto.SqueezeResult{
		ID: taskReq.Id,
		Result: nil,
	}

	for r := range m.results {
		res, ok := r.(*pb.ExecuteTaskResponse)
		if !ok {
			log.Errorf("Incorrect message type, expected: ExecuteTaskResponse, but got: %T", r)
			continue
		}

		SqueezeResult.AgentStats = append(SqueezeResult.AgentStats, proto.SqueezeStats{
			Addr: res.Addr,
			Status: int32(res.Status),
			Error: res.Error},
		)

		if res.Status == pb.ExecuteTaskResponse_FAIL {
			continue
		}

		results = append(results, res.Data)
	}

	if len(results) != 0 {
		ret, err := builder.Merge(results)
		if err != nil {
			log.Error("failed to merge result, ", err.Error())
			return
		}
		SqueezeResult.Result = ret
	}

	aggregation <- SqueezeResult
}

func (m *MasterServer) handleInfo(w http.ResponseWriter, r *http.Request) {
	resp := []AgentStatusResp{}
	slaveConnsMutex.Lock()
	for connID, c := range slaveConns {
		resp = append(resp, AgentStatusResp{ConnID: connID, Addr: c.PeerAddr, Status: c.Status})
	}
	slaveConnsMutex.Unlock()

	util.RespondWithJSON(w, http.StatusOK, resp)
}

func (m *MasterServer) ExecuteTask(ctx context.Context, in *pb.ExecuteTaskRequest) (*pb.ExecuteTaskResponse, error) {
	return nil, nil
}

func (m *MasterServer) HeartBeat(stream pb.SqueezeService_HeartBeatServer) error {
	peerInfo, ok := peer.FromContext(stream.Context())
	peerAddr := "0.0.0.0"
	if ok {
		peerAddr = peerInfo.Addr.String()
	}

	log.V(2).Infof("Heartbeat ... from %s", peerAddr)

	conn := newSlaveConn(peerAddr, stream)

	var receiveError error
	reqChannel := make(chan *pb.HeartBeatRequest, 1)
	go recvThread(conn, reqChannel, &receiveError)

	for {
		select {
		case req, ok := <-reqChannel:
			if !ok {
				return nil
			}

			if req.Task != nil {
				log.V(2).Infof("recv heartbeat(%s) from slave %s.",
					pb.HeartBeatRequest_Task_Status_name[int32(req.Task.Status)], peerAddr)
			}

			if !conn.added {
				conn.added = true
				conn.ConnID = conn.connectionID()
				conn.GrpcPort = int(req.Info.GrpcPort)
				m.addConn(conn.ConnID, conn)
				defer m.removeConn(conn.ConnID, conn)
			}

			m.updateConn(conn.ConnID, pb.HeartBeatRequest_Task_Status_name[int32(req.Task.Status)])

			if err := stream.Send(&pb.HeartBeatResponse{}); err != nil {
				log.Error("send failed.")
			}
		}
	}

	return nil
}

func recvThread(con *SlaveConn, reqChannel chan *pb.HeartBeatRequest, errP *error) {
	defer close(reqChannel)
	for {
		// receive data from stream
		req, err := con.stream.Recv()
		if err != nil {
			if status.Code(err) == codes.Canceled || err == io.EOF {
				return
			}
			*errP = err
			log.Errorf("%q terminated with errors %v", con.PeerAddr, err)
			return
		}
		reqChannel <- req
	}
}

func (m *MasterServer) Start(stopChan <-chan struct{}) error {
	// http server
	go func() {
		log.Infof("http listening on %s", m.httpServer.Addr)
		err := m.httpServer.ListenAndServe()
		if err != nil {
			log.Errorf("Start http server error: %s", err.Error())
		}
	}()

	// grpc server
	go func() {
		log.Infof("grpc listening on %s", m.args.GRPCAddr)
		listener, err := net.Listen("tcp", m.args.GRPCAddr)
		if err != nil {
			log.Errorf("GRPC failed to listen: %v", err)
			return
		}

		if err = m.grpcServer.Serve(listener); err != nil {
			log.Errorf("Start grpc server error: %s", err.Error())
		}
	}()

	go func() {
		<-stopChan
		err := m.httpServer.Close()
		if err != nil {
			log.Error(err)
		}
		m.grpcServer.Stop()
	}()
	return nil
}
