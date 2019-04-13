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

package websocket

import (
	"fmt"
	"time"
	"net/url"
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/golang/protobuf/jsonpb"
	"github.com/agile6v/squeeze/pkg/config"
	"github.com/agile6v/squeeze/pkg/pb"
	"github.com/agile6v/squeeze/pkg/util"
	log "github.com/golang/glog"
)

type WebsocketStats struct {
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

type wsResult struct {
	Err           error
	StatusCode    int
	Offset        time.Duration
	Duration      time.Duration
	ContentLength int64
}

type wsReport struct {
	result *WebsocketStats
	lats   []float64 // time spent per request
}

func newWsReport(n int) *wsReport {
	cap := n
	return &wsReport{
		result: &WebsocketStats{
			ErrMap: make(map[string]uint32),
		},
		lats: make([]float64, 0, cap),
	}
}

type WebSocketBuilder struct {
	Conn    *websocket.Conn
	report  *wsReport
	options *config.WsOptions
}

func NewBuilder() *WebSocketBuilder {
	return &WebSocketBuilder{}
}

func (builder *WebSocketBuilder) CreateTask(configArgs *config.ProtoConfigArgs) (string, error) {
	wsOptions, ok := configArgs.Options.(*config.WsOptions)
	if !ok {
		return "", fmt.Errorf("Expected WsOptions type, but got %T", configArgs.Options)
	}

	data, err := json.Marshal(wsOptions)
	if err != nil {
		log.Errorf("could not marshal message : %v", err)
		return "", err
	}

	req := &pb.ExecuteTaskRequest{
		Id:       uint32(configArgs.ID),
		Cmd:      pb.ExecuteTaskRequest_START,
		Protocol: pb.Protocol_WEBSOCKET,
		Callback: configArgs.Callback,
		Duration: uint32(wsOptions.Duration),
		Requests: uint32(wsOptions.Requests),
		Concurrency: uint32(wsOptions.Concurrency),
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

func (builder *WebSocketBuilder) Split(request *pb.ExecuteTaskRequest, count int) []*pb.ExecuteTaskRequest {
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

func (builder *WebSocketBuilder) Init(ctx context.Context, taskReq *pb.ExecuteTaskRequest) error {
	var options config.WsOptions
	err := json.Unmarshal([]byte(taskReq.Data), &options)
	if err != nil {
		return err
	}

	builder.options = &options
	return nil
}

func (builder *WebSocketBuilder) PreRequest(taskReq *pb.ExecuteTaskRequest) (interface{}, interface{}) {
	builder.report = newWsReport(util.Min(int(taskReq.Requests), int(builder.options.MaxResults)))
	dialer := websocket.Dialer{
		HandshakeTimeout: time.Duration(builder.options.Timeout) * time.Second,
	}

	u := url.URL{Scheme: builder.options.Scheme, Host: builder.options.Host, Path: builder.options.Path}
	conn, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		return nil,  &wsResult{Err: err}
	}

	return conn, nil
}

func (builder *WebSocketBuilder) Request(ctx context.Context, obj interface{}, taskReq *pb.ExecuteTaskRequest) interface{} {
	s := util.Now()
	conn, _ := obj.(*websocket.Conn)

	var resp []byte
	err := conn.WriteMessage(websocket.TextMessage, []byte(builder.options.Body))
	if err == nil {
		// Since our goal is to press server-side capabilities, so we only
		// consider synchronization scenario. If needed, we will support
		// asynchronous scenario in the future.
		_, resp, err = conn.ReadMessage()
		if err == nil {
			log.V(3).Infof("read message %s from target.", string(resp))
		}
	}

	t := util.Now()
	finish := t - s

	return &wsResult{
		Duration:      finish,
		ContentLength: int64(len([]byte(resp))),
		Err:           err,
	}
}

func (builder *WebSocketBuilder) PostRequest(result interface{}) error {
	res, ok := result.(*wsResult)
	if !ok {
		return fmt.Errorf("Expected wsResult type, but got %T", result)
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

func (builder *WebSocketBuilder) Destroy(obj interface{}) error {
	return nil
}

func (builder *WebSocketBuilder) Done(total time.Duration) (interface{}, error) {
	report := builder.report
	report.result.Duration = total.Seconds()

	if len(report.lats) == 0 {
		return report.result, nil
	}

	// TODO

	// Sending a close message to close the connection.
	err := builder.Conn.WriteMessage(websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		return report.result, err
	}

	builder.Conn.Close()

	return report.result, nil
}

func (builder *WebSocketBuilder) Merge(messages []string) (interface{}, error) {
	stats := &WebsocketStats{}
	stats.ErrMap = make(map[string]uint32)

	for _, message := range messages {
		r := &WebsocketStats{}
		err := json.Unmarshal([]byte(message), r)
		if err != nil {
			return nil, fmt.Errorf("cannot cast to WebsocketStats: %#v", message)
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
