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

import (
	"fmt"
	"net"
	"errors"
	"net/url"
	"github.com/agile6v/squeeze/pkg/util"
)

type ProtoConfigArgs struct {
	ID          int             // The ID of the task
	HttpAddr    string          // Usually used to save the address of the master
	Callback    string          // If it is asynchronous mode, the response
								// will be sent to the address specified by Callback
	WsOpts      WsOptions       // Parameters of the WEBSOCKET protocol
	Options     interface{}
}

func NewConfigArgs(opts interface{}) *ProtoConfigArgs {
	return &ProtoConfigArgs{Options: opts}
}

// HttpOptions contains http protocol runtime parameters
type HttpOptions struct {
	URL              string         `json:"url,omitempty"`
	HTTP2            bool           `json:"http2,omitempty"`
	Requests         int            `json:"requests,omitempty"`
	Method           string         `json:"method,omitempty"`
	ProxyAddr        string         `json:"proxyAddr,omitempty"`
	Headers          []string       `json:"headers,omitempty"`
	Concurrency      int            `json:"concurrency,omitempty"`
	RateLimit        int            `json:"rateLimit,omitempty"`
	Timeout          int            `json:"timeout,omitempty"`
	Duration         int            `json:"duration,omitempty"`
	Body             string         `json:"body,omitempty"`
	BodyFile         string         `json:"bodyFile,omitempty"`
	ContentType      string         `json:"contentType,omitempty"`
	MaxResults       int            `json:"maxResults,omitempty"`
	DisableKeepAlive bool           `json:"disableKeepAlive,omitempty"`
	DisableCompression bool         `json:"disableCompression,omitempty"`
}

func NewHttpOptions() *HttpOptions {
	return &HttpOptions{}
}

func (httpOpts *HttpOptions) Validate(args []string) error {
	// Check the validity of the concurrency
	if httpOpts.Concurrency < 1 {
		return fmt.Errorf("option --concurrency must be greater than 0.")
	}

	// Check if the options are missing
	if httpOpts.Requests == 0 && httpOpts.Duration == 0 {
		return fmt.Errorf("option --requests or --duration must be specified one of them.")
	}

	//
	if httpOpts.Duration == 0 {
		if httpOpts.Requests < httpOpts.Concurrency {
			return fmt.Errorf("option --concurrecny must be greater than --requests.")
		}
	}

	// Check if the format of http headers' is vaild
	if len(httpOpts.Headers) > 0 {
		for _, h := range httpOpts.Headers {
			_, err := util.ParseHTTPHeader(h)
			if err != nil {
				return fmt.Errorf("HTTP Header format is invalid, %v", err)
			}
		}
	}

	// Check the validity of the target URL
	u, err := url.ParseRequestURI(args[0])
	if err != nil {
		return err
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return fmt.Errorf("Please specify the url scheme, like http://abc.com or https://abc.com")
	}

	if httpOpts.ProxyAddr != "" {
		_, err := url.Parse(httpOpts.ProxyAddr)
		if err != nil {
			return fmt.Errorf("invalid argument %s: %s", httpOpts.ProxyAddr, err.Error())
		}
	}

	return nil
}

type UDPOptions struct {
	Addr        string
	Requests    int
	Concurrency int
	Timeout     int
	Duration    int
	MsgLength   int
	MaxResults  int
}

func NewUDPOptions() *UDPOptions {
	return &UDPOptions{}
}

func (udpOptions *UDPOptions) Validate(args []string) error {
	if udpOptions.Concurrency < 1 {
		return fmt.Errorf("option --concurrency must be greater than 0.")
	}

	// Check the validity of the target address
	_, _, err := net.SplitHostPort(args[0])
	if err != nil {
		return err
	}

	udpOptions.Addr = args[0]
	return nil
}

// WsOptions contains websocket protocol runtime parameters
type WsOptions struct {
	Scheme      string      `json:"scheme,omitempty"`
	Host        string      `json:"host,omitempty"`
	Path        string      `json:"path,omitempty"`
	Requests    int         `json:"requests,omitempty"`
	Concurrency int         `json:"concurrency,omitempty"`
	Timeout     int         `json:"timeout,omitempty"`
	Duration    int         `json:"duration,omitempty"`
	Body        string      `json:"body,omitempty"`
	MaxResults  int         `json:"maxResults,omitempty"`
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

// WebOptions contains options of the web command
type WebOptions struct {
	DSN     string      `json:"dsn,omitempty"`
	File    string      `json:"file,omitempty"`
	Type    string      `json:"type,omitempty"`
}

func NewWebOptions() *WebOptions {
	return &WebOptions{}
}
