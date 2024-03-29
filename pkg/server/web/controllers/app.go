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

package controllers

import (
	"encoding/json"
	"github.com/agile6v/squeeze/pkg/config"
	"github.com/agile6v/squeeze/pkg/pb"
	"github.com/agile6v/squeeze/pkg/proto"
	"github.com/agile6v/squeeze/pkg/proto/builder"
	"github.com/agile6v/squeeze/pkg/proto/http"
	"github.com/agile6v/squeeze/pkg/proto/websocket"
	"github.com/agile6v/squeeze/pkg/server/web/dao"
	log "github.com/golang/glog"
	"strings"
)

type CreateTask struct {
	Protocol string          `json:"protocol"`
	Data     json.RawMessage `json:"data"`
}

func (c *CreateTask) Handle(data string) (*dao.Task, error) {
	// check the validity of the data
	createTask := &CreateTask{}
	err := json.Unmarshal([]byte(data), createTask)
	if err != nil {
		return nil, err
	}

	ret, err := dao.CreateTask(data)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

type GenericTask struct {
	ID int `json:"id"`
}

func (task *GenericTask) Delete() error {
	err := dao.DeleteTask(task.ID)
	if err != nil {
		return err
	}
	return nil
}

func (g *GenericTask) Start(masterAddr, webAddr string) error {
	task, err := dao.SearchTask(g.ID)
	if err != nil {
		return err
	}

	createTask := &CreateTask{}
	err = json.Unmarshal([]byte(task.Request), createTask)
	if err != nil {
		return err
	}

	protocol := pb.Protocol(pb.Protocol_value[strings.ToUpper(createTask.Protocol)])
	builder := builder.NewBuilder(protocol)

	var options interface{}
	if protocol == pb.Protocol_HTTP {
		options = http.NewHttpOptions()
		err = json.Unmarshal(createTask.Data, options)
	} else if protocol == pb.Protocol_WEBSOCKET {
		options = websocket.NewWsOptions()
		err = json.Unmarshal(createTask.Data, options)
	} else {
		// TODO: error
	}

	if err != nil {
		return err
	}

	// update status to "START"
	err = dao.UpdateTaskByStatus(task.Id, dao.STATUS_START)
	if err != nil {
		return err
	}

	args := config.NewConfigArgs(options)
	args.Callback = "http://" + webAddr + "/api/callback"
	args.HttpAddr = masterAddr
	args.ID = task.Id

	resp, err := builder.CreateTask(args)
	if err != nil {
		return err
	}

	log.Infof("start task returns %s", resp)
	return nil
}

func (g *GenericTask) Stop(masterAddr string) error {
	task, err := dao.SearchTask(g.ID)
	if err != nil {
		return err
	}

	createTask := &CreateTask{}
	err = json.Unmarshal([]byte(task.Request), createTask)
	if err != nil {
		return err
	}

	protocol := pb.Protocol(pb.Protocol_value[strings.ToUpper(createTask.Protocol)])
	builder := builder.NewBuilder(protocol)

	var options interface{}
	if protocol == pb.Protocol_HTTP {
		options = http.NewHttpOptions()
		err = json.Unmarshal(createTask.Data, options)
	} else if protocol == pb.Protocol_WEBSOCKET {
		options = websocket.NewWsOptions()
		err = json.Unmarshal(createTask.Data, options)
	} else {
		// TODO: error
	}

	if err != nil {
		return err
	}

	args := config.NewConfigArgs(options)
	args.HttpAddr = masterAddr

	resp, err := builder.CancelTask(args)
	if err != nil {
		return err
	}

	err = dao.UpdateTaskByStatus(task.Id, dao.STATUS_STOP)
	if err != nil {
		return err
	}

	log.Infof("stop task returns %s", resp)
	return nil
}

func (g *GenericTask) Search() (*dao.Task, error) {
	return dao.SearchTask(g.ID)
}

func ListTask() ([]dao.Task, error) {
	return dao.ListTask()
}

func HandleCallback(data string) error {
	response := &proto.SqueezeResponse{}
	err := json.Unmarshal([]byte(data), response)
	if err != nil {
		return err
	}

	id := response.Data.ID

	// update status to "STOP"
	err = dao.UpdateTaskByStatus(int(id), dao.STATUS_STOP)
	if err != nil {
		return err
	}

	err = dao.UpdateTaskByResponse(int(id), data)
	if err != nil {
		return err
	}

	return nil
}
