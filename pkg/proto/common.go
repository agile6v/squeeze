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

package proto

import (
	"time"
	"bytes"
	"context"
	"encoding/json"
	"github.com/agile6v/squeeze/pkg/config"
	"github.com/agile6v/squeeze/pkg/pb"
	"github.com/agile6v/squeeze/pkg/util"
	"github.com/golang/protobuf/jsonpb"
)

type SqueezeStats struct {
	Addr    string  `json:"addr"`
	Status  int32   `json:"status"`
	Error   string  `json:"error"`
}

type SqueezeResponse struct {
	Data   *SqueezeResult `json:"data"`
	Error   string        `json:"error"`
}

type SqueezeResult struct {
	ID         uint32         `json:"id"`
	AgentStats []SqueezeStats `json:"agent_stats"`
	Result     interface{}    `json:"result"`
}

type ProtoBuilder interface {
	// slave side
	// These functions are executed in the following order
	Init(context.Context, *pb.TaskRequest) error
	PreRequest(*pb.TaskRequest) (interface{}, interface{})
	Request(context.Context, interface{}, *pb.TaskRequest) interface{}
	PostRequest(interface{}) error
	Done(time.Duration) (interface{}, error)

	// master side
	Split(*pb.ExecuteTaskRequest, int) []*pb.ExecuteTaskRequest
	Merge([]string) (interface{}, error)

	// client side
	CreateTask(*config.ProtoConfigArgs) (string, error)
}

type ProtoBuilderBase struct {
	ProtoBuilder
	Template *string
	Stats    interface{}
}

func (proto *ProtoBuilderBase) CancelTask(ConfigArgs *config.ProtoConfigArgs) (string, error) {
	req := &pb.ExecuteTaskRequest{
		Cmd:      pb.ExecuteTaskRequest_STOP,
		Callback: ConfigArgs.Callback,
	}

	m := jsonpb.Marshaler{}
	jsonStr, err := m.MarshalToString(req)
	if err != nil {
		return "", err
	}

	resp, err := util.DoRequest("POST", ConfigArgs.HttpAddr+"/task/stop", string(jsonStr), 0)
	if err != nil {
		return "", err
	}
	return resp, nil
}

func (proto *ProtoBuilderBase) Render(data string) (string, error) {
	response := &SqueezeResponse{
		Data: &SqueezeResult{
			Result: proto.Stats,
		},
		Error: "",
	}

	err := json.Unmarshal([]byte(data), response)
	if err != nil {
		return "", err
	}

	buf := &bytes.Buffer{}
	if response.Data.Result == nil {
		if err := util.NewTemplate(errorTemplate).Execute(buf, response); err != nil {
			return "", err
		}
	} else {
		if err := util.NewTemplate(*proto.Template).Execute(buf, proto.Stats); err != nil {
			return "", err
		}
	}

	return buf.String(), nil
}

var (
	errorTemplate = `
Summary:
{{ range .AgentStats }}
  Agent: {{ .Addr }}, {{ if eq .Status 0 }}SUCCESS{{ else }}FAILED{{ end }}, {{ .Error }}
{{ end }}
`
)
