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

package builder

import (
	"github.com/agile6v/squeeze/pkg/pb"
	"github.com/agile6v/squeeze/pkg/proto"
	"github.com/agile6v/squeeze/pkg/proto/http"
	"github.com/agile6v/squeeze/pkg/proto/websocket"
	"github.com/agile6v/squeeze/pkg/proto/udp"
	"github.com/agile6v/squeeze/pkg/proto/tcp"
)

func NewBuilder(protocol pb.Protocol) *proto.ProtoBuilderBase {
	switch protocol {
	case pb.Protocol_HTTP:
		return &proto.ProtoBuilderBase{http.NewBuilder(), &http.ResultTmpl, &http.HttpStats{}}
	case pb.Protocol_WEBSOCKET:
		return &proto.ProtoBuilderBase{websocket.NewBuilder(), &websocket.ResultTmpl, &websocket.WebsocketStats{}}
	case pb.Protocol_UDP:
		return &proto.ProtoBuilderBase{udp.NewBuilder(), &udp.ResultTmpl, &udp.UDPStats{}}
	case pb.Protocol_TCP:
		return &proto.ProtoBuilderBase{tcp.NewBuilder(), &tcp.ResultTmpl, &tcp.TCPStats{}}
	}
	return nil
}
