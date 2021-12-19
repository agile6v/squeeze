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
	"github.com/agile6v/squeeze/pkg/version"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Proxy struct {
	target *url.URL
	proxy  *httputil.ReverseProxy
}

func NewProxy(target string) (*Proxy, error) {
	url, err := url.Parse(fmt.Sprintf("http://%v/", target))
	if err != nil {
		return nil, err
	}

	return &Proxy{target: url, proxy: httputil.NewSingleHostReverseProxy(url)}, nil
}

func (p *Proxy) handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Forwarded-P", version.GetVersion())

	p.proxy.ServeHTTP(w, r)
}
