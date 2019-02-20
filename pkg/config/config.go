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

package config

var ConfigArgs ProtoConfigArgs

type ProtoConfigArgs struct {
	HttpAddr string
	Callback string
	HttpOpts HttpOptions
	WsOpts   WsOptions
}

// HttpOptions contains http protocol runtime parameters
type HttpOptions struct {
	URL              string
	Requests         int
	Method           string
	ProxyAddr        string
	Headers          []string
	Concurrency      int
	RateLimit        int
	Timeout          int
	Duration         int
	Body             string
	ContentType      string
	BodyFile         string
	DisableKeepAlive bool
}

// WsOptions contains websocket protocol runtime parameters
type WsOptions struct {
	Scheme      string
	Host        string
	Path        string
	Requests    int
	Concurrency int
	Timeout     int
	Duration    int
	Body        string
}
