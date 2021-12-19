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
	"errors"
	"fmt"
	"net/url"
)

// WsOptions contains websocket protocol runtime parameters
type WsOptions struct {
	Scheme      string `json:"scheme,omitempty"`
	Host        string `json:"host,omitempty"`
	Path        string `json:"path,omitempty"`
	Requests    int    `json:"requests,omitempty"`
	Concurrency int    `json:"concurrency,omitempty"`
	Timeout     int    `json:"timeout,omitempty"`
	Duration    int    `json:"duration,omitempty"`
	Body        string `json:"body,omitempty"`
	MaxResults  int    `json:"maxResults,omitempty"`
}

func NewWsOptions() *WsOptions {
	return &WsOptions{}
}

func (wsOptions *WsOptions) Validate(args []string) error {
	if wsOptions.Concurrency < 1 {
		return fmt.Errorf("option --concurrency must be greater than 0.")
	}

	// Check the validity of the target URL
	u, err := url.Parse(args[0])
	if err != nil {
		return err
	}

	if u.Scheme == "" || u.Path == "" {
		return errors.New("URL Scheme or Path cannot be empty.")
	}

	wsOptions.Scheme = u.Scheme
	wsOptions.Host = u.Host
	wsOptions.Path = u.Path

	return nil
}
