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
	log "github.com/golang/glog"
	"github.com/agile6v/squeeze/pkg/server/web/api"
	"github.com/agile6v/squeeze/pkg/util"
)

type WebServer struct {
	ServerBase
}

func (s *WebServer) Initialize(args *ServerArgs) error {
	s.ServerBase.Initialize(args)
	s.Mode = Web

	ip, err := util.ExternalIP()
	if err != nil {
		return err
	}

	_, port, err := util.GetHostPort(args.HTTPAddr)
	if err != nil {
		return err
	}

	api := &api.AppAPI{
		MasterAddr: s.args.MasterAddr,
		HTTPAddr: s.args.HTTPAddr,
		LocalAddr: ip + ":" + port,
	}

	api.Init()

	return nil
}

func (s *WebServer) Start(stopChan <-chan struct{}) error {
	// http server
	go func() {
		log.Infof("http listening on %s", s.httpServer.Addr)
		err := s.httpServer.ListenAndServe()
		if err != nil {
			log.Errorf("Start http server error: %s", err.Error())
		}
	}()

	go func() {
		<-stopChan
		err := s.httpServer.Close()
		if err != nil {
			log.Error(err)
		}
	}()
	return nil
}