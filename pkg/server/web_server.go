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
	"net/http"
	"sync"
	"github.com/agile6v/squeeze/pkg/pb"
)

type WebServer struct {
	ServerBase
	lastTaskReq *pb.ExecuteTaskRequest
	mutex       sync.RWMutex
}

func (s *WebServer) Initialize(args ServerArgs) error {
	s.ServerBase.Initialize(args)
	s.Mode = Web

	http.HandleFunc("/", s.handleInfo)
	http.HandleFunc("/api/create", s.handleInfo)
	http.HandleFunc("/api/delete", s.handleInfo)
	http.HandleFunc("/api/search", s.handleInfo)
	http.HandleFunc("/api/list", s.handleInfo)
	http.HandleFunc("/api/start", s.handleInfo)
	http.HandleFunc("/api/stop", s.handleInfo)

	return nil
}

func (s *WebServer) Start(stopChan <-chan struct{}) error {
	return nil
}

func (m *WebServer) handleInfo(writer http.ResponseWriter, request *http.Request) {

}