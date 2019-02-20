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
	"github.com/agile6v/squeeze/pkg/pb"
	"google.golang.org/grpc"
	"strconv"
	"sync"
)

var (
	slaveConns      = map[string]*SlaveConn{}
	slaveConnsMutex sync.RWMutex

	connectionNumber = int64(0)
	connectionMutex  sync.Mutex
)

type SlaveStream interface {
	Send(*pb.HeartBeatResponse) error
	Recv() (*pb.HeartBeatRequest, error)
	grpc.ServerStream
}

type SlaveConn struct {
	PeerAddr string
	stream   SlaveStream
	added    bool
	ConnID   string
	GrpcPort int
}

func newSlaveConn(addr string, stream SlaveStream) *SlaveConn {
	return &SlaveConn{
		PeerAddr: addr,
		stream:   stream,
	}
}

func (conn *SlaveConn) connectionID() string {
	connectionMutex.Lock()
	connectionNumber++
	c := connectionNumber
	connectionMutex.Unlock()
	return conn.PeerAddr + "-" + strconv.Itoa(int(c))
}

func SlaveCount() int {
	var n int
	slaveConnsMutex.RLock()
	n = len(slaveConns)
	slaveConnsMutex.RUnlock()
	return n
}

func GetConnections() []*SlaveConn {
	conns := []*SlaveConn{}
	slaveConnsMutex.RLock()
	for _, conn := range slaveConns {
		conns = append(conns, conn)
	}
	slaveConnsMutex.RUnlock()
	return conns
}
