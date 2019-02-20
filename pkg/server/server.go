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
	"fmt"
	"github.com/agile6v/squeeze/pkg/util"
	log "github.com/golang/glog"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
	"time"
)

var SrvArgs ServerArgs

// NodeType indicates the kind of the server.
type NodeType int

const (
	//	Client represents the client mode
	Client NodeType = iota
	//	Slave represents the slave mode
	Slave
	//	Master represents the master mode
	Master
)

func (t NodeType) String() string {
	switch t {
	case Client:
		return "client"
	case Slave:
		return "slave"
	case Master:
		return "master"
	default:
		return fmt.Sprintf("%d", t)
	}
}

type AgentStatusResp struct {
	ConnID string `json:"id"`
	Addr   string `json:"addr"`
}

type ServerArgs struct {
	HTTPAddr       string
	GrpcAddr       string

	MasterAddr     string
	GrpcMasterAddr string
	ReportInterval time.Duration
	ResultCapacity int
}

type Server interface {
	Initialize(args ServerArgs) error
	Start(stopChan <-chan struct{}) error
}

type ServerBase struct {
	args       ServerArgs
	Mode       NodeType
	httpServer *http.Server
	grpcServer *grpc.Server
	grpcPort   int
}

func NewServer(nodeType NodeType) Server {
	if nodeType == Master {
		return &MasterServer{}
	} else {
		return &SlaveServer{}
	}
}

func (s *ServerBase) Initialize(args ServerArgs) error {
	s.args = args
	//s.Mode = nodeType
	s.httpServer = &http.Server{
		Addr: args.HTTPAddr,
	}

	_, port, err := util.GetHostPort(s.args.GrpcAddr)
	if err != nil {
		return err
	}

	s.grpcPort, err = strconv.Atoi(port)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServerBase) addConn(connID string, conn *SlaveConn) {
	slaveConnsMutex.Lock()
	defer slaveConnsMutex.Unlock()

	slaveConns[connID] = conn
}

func (s *ServerBase) removeConn(connID string, conn *SlaveConn) {
	slaveConnsMutex.Lock()
	defer slaveConnsMutex.Unlock()

	if slaveConns[connID] == nil {
		log.Errorf("Removing connection for non-existing node %v.", s)
	}
	delete(slaveConns, connID)
}
