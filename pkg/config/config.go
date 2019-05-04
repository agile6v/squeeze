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

type ProtoConfigArgs struct {
	ID          int             // The ID of the task
	HttpAddr    string          // Usually used to save the address of the master
	Callback    string          // If it is asynchronous mode, the response
								// will be sent to the address specified by Callback
	Options     interface{}
}

func NewConfigArgs(opts interface{}) *ProtoConfigArgs {
	return &ProtoConfigArgs{Options: opts}
}

// WebOptions contains options of the web command
type WebOptions struct {
	DSN     string      `json:"dsn,omitempty"`
	File    string      `json:"file,omitempty"`
	Type    string      `json:"type,omitempty"`
}

func NewWebOptions() *WebOptions {
	return &WebOptions{}
}
