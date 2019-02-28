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

package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"time"
	"net/http"
	"regexp"
	"encoding/json"
	"github.com/agile6v/squeeze/pkg/version"
)

var headerRegexp = `^([\w-]+):\s*(.+)`

func BuildHostname(ip, port string) (string, error) {
	host, _, err := GetHostPort(ip)
	if err != nil {
		return "", err
	}

	return host + ":" + port, nil
}

func GetHostPort(addr string) (host string, port string, err error) {
	host, port, err = net.SplitHostPort(addr)
	if err != nil {
		return "", "", fmt.Errorf("unable to parse address %q: %v", addr, err)
	}
	return host, port, nil
}

func ParseHTTPHeader(in string) ([]string, error) {
	return parseInputWithRegexp(in, headerRegexp)
}

func parseInputWithRegexp(input, regx string) ([]string, error) {
	re := regexp.MustCompile(regx)
	matches := re.FindStringSubmatch(input)
	if len(matches) < 1 {
		return nil, fmt.Errorf("could not parse the provided input; input = %v", input)
	}
	return matches, nil
}

func DoRequest(method, host, body string, timeout time.Duration) (string, error) {
	req, err := http.NewRequest(method, host, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return "", err
	}

	header := make(http.Header)
	header.Set("User-Agent", version.GetVersion())
	req.Header = header

	// Send request
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return string(respBytes), fmt.Errorf("Received Non-200 status code %v.", resp.StatusCode)
	}

	return string(respBytes), nil
}

// cloneRequest returns a clone of the provided *http.Request.
// The clone is a shallow copy of the struct and its Header map.
func CloneRequest(r *http.Request, body []byte) *http.Request {
	// shallow copy of the struct
	r2 := new(http.Request)
	*r2 = *r
	// deep copy of the Header
	r2.Header = make(http.Header, len(r.Header))
	for k, s := range r.Header {
		r2.Header[k] = append([]string(nil), s...)
	}
	if len(body) > 0 {
		r2.Body = ioutil.NopCloser(bytes.NewReader(body))
	}
	return r2
}

func ReadBody(r *http.Request, obj interface{}) error {
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return err
	}

	// Unmarshal
	err = json.Unmarshal(b, obj)
	if err != nil {
		return err
	}

	return nil
}