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

package tcp

import (
	"fmt"
	"time"
	"net"
	"context"
	"encoding/json"
	log "github.com/golang/glog"
	"github.com/golang/protobuf/jsonpb"
	"github.com/agile6v/squeeze/pkg/config"
	"github.com/agile6v/squeeze/pkg/pb"
	"github.com/agile6v/squeeze/pkg/util"
)

type TCPStats struct {
	TotalSize       int64       `json:"totalSize,omitempty"`
	Rps             float64     `json:"rps,omitempty"`
	Duration        float64     `json:"duration,omitempty"`
	TotalDuration   float64     `json:"totalDuration,omitempty"`
	Requests        int64       `json:"requests,omitempty"`
	TotalRequests   int64       `json:"totalRequests,omitempty"`
	TotalResponses  int64       `json:"totalResponses,omitempty"`
	AvgSize         int64       `json:"avgSize,omitempty"`
	ErrMap          map[string]uint32 `json:"errMap,omitempty"`
}

type tcpResult struct {
	Err           error
	StatusCode    int
	Offset        time.Duration
	Duration      time.Duration
	ContentLength int64
}

type tcpReport struct {
	result *TCPStats
	lats   []float64 // time spent per request
}

func newTCPReport(n int) *tcpReport {
	cap := n
	return &tcpReport{
		result: &TCPStats{
			ErrMap: make(map[string]uint32),
		},
		lats: make([]float64, 0, cap),
	}
}

type TCPBuilder struct {
	report  *tcpReport
	options *TCPOptions
}

func NewBuilder() *TCPBuilder {
	return &TCPBuilder{}
}

func (builder *TCPBuilder) CreateTask(configArgs *config.ProtoConfigArgs) (string, error) {
	tcpOptions, ok := configArgs.Options.(*TCPOptions)
	if !ok {
		return "", fmt.Errorf("Expected tcpOptions type, but got %T", configArgs.Options)
	}

	data, err := json.Marshal(tcpOptions)
	if err != nil {
		log.Errorf("could not marshal message : %v", err)
		return "", err
	}

	req := &pb.ExecuteTaskRequest{
		Id:       uint32(configArgs.ID),
		Cmd:      pb.ExecuteTaskRequest_START,
		Protocol: pb.Protocol_TCP,
		Callback: configArgs.Callback,
		Duration: uint32(tcpOptions.Duration),
		Requests: uint32(tcpOptions.Requests),
		Concurrency: uint32(tcpOptions.Concurrency),
		//RateLimit:
		Data: string(data),
	}

	m := jsonpb.Marshaler{}
	jsonStr, err := m.MarshalToString(req)
	if err != nil {
		return "", err
	}

	resp, err := util.DoRequest("POST", configArgs.HttpAddr+"/task/start", string(jsonStr), 0)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (builder *TCPBuilder) Split(request *pb.ExecuteTaskRequest, count int) []*pb.ExecuteTaskRequest {
	var requests []*pb.ExecuteTaskRequest

	if count > int(request.Concurrency) {
		count = int(request.Concurrency)
	} else if count > int(request.Requests) {
		count = int(request.Requests)
	}

	for i := 1; i <= count; i++ {
		req := new(pb.ExecuteTaskRequest)
		*req = *request

		if count != i {
			req.Requests = request.Requests / uint32(count)
			req.RateLimit = request.RateLimit / uint32(count)
			req.Concurrency = request.Concurrency / uint32(count)
		} else {
			req.Requests = request.Requests/uint32(count) + request.Requests%uint32(count)
			req.RateLimit = request.RateLimit/uint32(count) + request.RateLimit%uint32(count)
			req.Concurrency = request.Concurrency/uint32(count) + request.Concurrency%uint32(count)
		}

		requests = append(requests, req)
	}

	return requests
}

func (builder *TCPBuilder) Init(ctx context.Context, taskReq *pb.ExecuteTaskRequest) error {
	var options TCPOptions
	err := json.Unmarshal([]byte(taskReq.Data), &options)
	if err != nil {
		return err
	}

	builder.options = &options
	return nil
}

func (builder *TCPBuilder) PreRequest(taskReq *pb.ExecuteTaskRequest) (interface{}, interface{}) {
	builder.report = newTCPReport(util.Min(int(taskReq.Requests), int(builder.options.MaxResults)))

	conn, err := net.Dial("tcp", builder.options.Addr)
	if err != nil {
		return nil, &tcpResult{
			Err: fmt.Errorf("connect to server %v failed : %v", builder.options.Addr, err.Error()),
		}
	}

	return conn, nil
}

func (builder *TCPBuilder) Request(ctx context.Context, obj interface{}, taskReq *pb.ExecuteTaskRequest) interface{} {
	s := util.Now()
	conn, ok := obj.(*net.TCPConn)
	if !ok {
		return fmt.Errorf("Expected TCPConn type, but got %T", obj)
	}

	content := make([]byte, builder.options.MsgLength)
	_, err := conn.Write(content)
	if err != nil {
		log.Infof(">>>>> %s", err)
		return err
	}

	t := util.Now()
	finish := t - s

	return &tcpResult{
		Duration:      finish,
		Err:           err,
	}
}

func (builder *TCPBuilder) PostRequest(result interface{}) error {
	res, ok := result.(*tcpResult)
	if !ok {
		return fmt.Errorf("Expected tcpResult type, but got %T", result)
	}

	report := builder.report
	report.result.TotalRequests++

	if res.Err != nil {
		report.result.ErrMap[res.Err.Error()]++
	} else {
		report.result.Requests++
		report.result.TotalDuration += res.Duration.Seconds()
		if res.ContentLength > 0 {
			report.result.TotalSize += res.ContentLength
		}

		if len(report.lats) < cap(report.lats) {
			report.lats = append(report.lats, res.Duration.Seconds())
		}
	}
	return nil
}

func (builder *TCPBuilder) Destroy(obj interface{}) error {
	conn, ok := obj.(*net.TCPConn)
	if !ok {
		return fmt.Errorf("Expected TCPConn type, but got %T", obj)
	}

	conn.Close()
	return nil
}

func (builder *TCPBuilder) Done(total time.Duration) (interface{}, error) {
	report := builder.report
	report.result.Duration = total.Seconds()

	return report.result, nil
}

func (builder *TCPBuilder) Merge(messages []string) (interface{}, error) {
	stats := &TCPStats{}
	stats.ErrMap = make(map[string]uint32)

	for _, message := range messages {
		r := &TCPStats{}
		err := json.Unmarshal([]byte(message), r)
		if err != nil {
			return nil, fmt.Errorf("cannot cast to TCPStats: %#v", message)
		}

		if stats.Duration < r.Duration {
			stats.Duration = r.Duration
		}
		stats.TotalRequests += r.TotalRequests
		stats.TotalDuration += r.TotalDuration
		stats.Requests += r.Requests
		stats.Rps += r.Rps
		stats.TotalSize += r.TotalSize

		for k, v := range r.ErrMap {
			if _, ok := stats.ErrMap[k]; ok {
				stats.ErrMap[k] += v
			} else {
				stats.ErrMap[k] = v
			}
		}
	}

	if stats.Requests > 0 {
		stats.AvgSize = stats.TotalSize / stats.Requests
		stats.Rps = float64(stats.TotalRequests) / stats.Duration
	}

	return stats, nil
}

var (
	ResultTmpl = `
Summary:
  Requests:	{{ formatNumberInt64 .TotalRequests }}
  Total:	{{ formatNumber .Duration }} secs
  Requests/sec:	{{ formatNumber .Rps }}
  {{ if gt .TotalSize 0 }}
  Total data:	{{ .TotalSize }} bytes
  Size/request:	{{ .AvgSize }} bytes{{ end }}

{{ if gt (len .ErrMap) 0 }}Error distribution:{{ range $err, $num := .ErrMap }}
  [{{ $num }}]	{{ $err }}{{ end }}{{ end }}
`
)
