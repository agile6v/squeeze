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

package http

import (
	"fmt"
	"io"
	"math"
	"sort"
	"time"
	"context"
	"crypto/tls"
	"net/url"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"net/http/httptrace"
	"golang.org/x/net/http2"
	"github.com/agile6v/squeeze/pkg/config"
	"github.com/agile6v/squeeze/pkg/pb"
	"github.com/agile6v/squeeze/pkg/util"
	"github.com/agile6v/squeeze/pkg/version"
	"github.com/golang/protobuf/jsonpb"
)

type ElapsedInfo struct {
	Max float64     `json:"max,omitempty"`
	Min float64     `json:"min,omitempty"`
	Avg float64     `json:"avg,omitempty"`
}

type LatencyDistribution struct {
	Percentage  uint32
    Latency     float64
}

type HttpStats struct {
	TotalRequests       int64       `json:"totalRequests,omitempty"`
	// Total time for running
	Duration            float64     `json:"duration,omitempty"`
	FastestReqTime      float64     `json:"fastestReqTime,omitempty"`
	SlowestReqTime      float64     `json:"slowestReqTime,omitempty"`
	AvgReqTime          float64     `json:"avgReqTime,omitempty"`
	// Average response size per request
	AvgSize             int64       `json:"avgSize,omitempty"`
	// The sum of all response sizes
	TotalSize           int64       `json:"totalSize,omitempty"`
	// Requests per second
	Rps                 float64     `json:"rps,omitempty"`
	Dns                 ElapsedInfo `json:"dns,omitempty"`
	Delay               ElapsedInfo `json:"delay,omitempty"`
	Resp                ElapsedInfo `json:"resp,omitempty"`
	Conn                ElapsedInfo `json:"conn,omitempty"`
	Req                 ElapsedInfo `json:"req,omitempty"`
	StatusCodes         map[uint32]uint32
	ErrMap              map[string]uint32
	ConnDuration        float64      `json:"connDuration,omitempty"`
	DnsDuration         float64      `json:"dnsDuration,omitempty"`
	ReqDuration         float64      `json:"reqDuration,omitempty"`
	RespDuration        float64      `json:"respDuration,omitempty"`
	DelayDuration       float64      `json:"delayDuration,omitempty"`
	// Total number of requests
	Requests            int64        `json:"requests,omitempty"`
	TotalDuration       float64      `json:"totalDuration,omitempty"`
	LatencyDistribution []LatencyDistribution `json:latencyDistribution,omitempty`
	// time spent per request
	Lats                []float64    `json:"latencies,omitempty"`
}

type httpResult struct {
	Err           error
	StatusCode    int
	Offset        time.Duration
	Duration      time.Duration
	ConnDuration  time.Duration // connection setup(DNS lookup + Dial up) duration
	DnsDuration   time.Duration // dns lookup duration
	ReqDuration   time.Duration // request "write" duration
	ResDuration   time.Duration // response "read" duration
	DelayDuration time.Duration // delay between response and request
	ContentLength int64
}

type httpReport struct {
	result      *HttpStats
	connLats    []float64
	dnsLats     []float64
	reqLats     []float64
	resLats     []float64
	delayLats   []float64
	offsets     []float64
	statusCodes []int
}

func newHttpReport(n int) *httpReport {
	cap := n
	return &httpReport{
		result: &HttpStats{
			ErrMap:   make(map[string]uint32),
			Lats:        make([]float64, 0, cap),
		},
		connLats:    make([]float64, 0, cap),
		dnsLats:     make([]float64, 0, cap),
		reqLats:     make([]float64, 0, cap),
		resLats:     make([]float64, 0, cap),
		delayLats:   make([]float64, 0, cap),
		statusCodes: make([]int, 0, cap),
	}
}

func latencies(stats *HttpStats) []LatencyDistribution {
	pctls := []uint32{10, 25, 50, 75, 90, 95, 99}
	data := make([]float64, len(pctls))
	j := 0
	for i := 0; i < len(stats.Lats) && j < len(pctls); i++ {
		current := i * 100 / len(stats.Lats)
		if uint32(current) >= pctls[j] {
			data[j] = stats.Lats[i]
			j++
		}
	}

	res := make([]LatencyDistribution, len(pctls))
	for i := 0; i < len(pctls); i++ {
		if data[i] > 0 {
			res[i].Percentage = pctls[i]
			res[i].Latency = data[i]
		}
	}

	return res
}

type HttpBuilder struct {
	report     *httpReport
	HttpReq    *http.Request
	HttpClient *http.Client
}

func NewBuilder() *HttpBuilder {
	return &HttpBuilder{}
}

func (builder *HttpBuilder) CreateTask(ConfigArgs *config.ProtoConfigArgs) (string, error) {
	if ConfigArgs.HttpOpts.Duration > 0 {
		ConfigArgs.HttpOpts.Requests = math.MaxInt32
	}

	req := &pb.ExecuteTaskRequest{
		Cmd:      pb.ExecuteTaskRequest_START,
		Protocol: pb.Protocol_HTTP,
		Callback: ConfigArgs.Callback,
		Duration: uint32(ConfigArgs.HttpOpts.Duration),
		Task: &pb.TaskRequest{
			Requests:    uint32(ConfigArgs.HttpOpts.Requests),
			Concurrency: uint32(ConfigArgs.HttpOpts.Concurrency),
			RateLimit:   uint32(ConfigArgs.HttpOpts.RateLimit),
			Type: &pb.TaskRequest_Http{
				Http: &pb.HttpTask{
					Url:               ConfigArgs.HttpOpts.URL,
					Method:            ConfigArgs.HttpOpts.Method,
					Body:              ConfigArgs.HttpOpts.Body,
					Timeout:           uint32(ConfigArgs.HttpOpts.Timeout),
					DisableKeepalives: ConfigArgs.HttpOpts.DisableKeepAlive,
					Headers:           ConfigArgs.HttpOpts.Headers,
					ProxyAddr:         ConfigArgs.HttpOpts.ProxyAddr,
					ContentType:       ConfigArgs.HttpOpts.ContentType,
					MaxResults:        int32(ConfigArgs.HttpOpts.MaxResults),
				},
			},
		},
	}

	m := jsonpb.Marshaler{}
	jsonStr, err := m.MarshalToString(req)
	if err != nil {
		return "", err
	}

	resp, err := util.DoRequest("POST", ConfigArgs.HttpAddr+"/task/start", string(jsonStr), 0)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (builder *HttpBuilder) Split(request *pb.ExecuteTaskRequest, count int) []*pb.ExecuteTaskRequest {
	var requests []*pb.ExecuteTaskRequest

	if count > int(request.Task.Concurrency) {
		count = int(request.Task.Concurrency)
	} else if (count > int(request.Task.Requests)) {
		count = int(request.Task.Requests)
	}

	for i := 1; i <= count; i++ {
		req := new(pb.ExecuteTaskRequest)
		*req = *request
		task := new(pb.HttpTask)
		*task = *request.Task.GetHttp()

		req.Task = &pb.TaskRequest{
			Type: &pb.TaskRequest_Http{
				Http: task,
			},
		}

		if count != i {
			req.Task.Requests = request.Task.Requests / uint32(count)
			req.Task.RateLimit = request.Task.RateLimit / uint32(count)
			req.Task.Concurrency = request.Task.Concurrency / uint32(count)
		} else {
			req.Task.Requests = request.Task.Requests/uint32(count) + request.Task.Requests%uint32(count)
			req.Task.RateLimit = request.Task.RateLimit/uint32(count) + request.Task.RateLimit%uint32(count)
			req.Task.Concurrency = request.Task.Concurrency/uint32(count) + request.Task.Concurrency%uint32(count)
		}

		requests = append(requests, req)
	}

	return requests
}

func (builder *HttpBuilder) Init(ctx context.Context, taskReq *pb.TaskRequest) error {
	task := taskReq.GetHttp()

	builder.report = newHttpReport(util.Min(int(taskReq.Requests), int(task.MaxResults)))
	httpReq, err := http.NewRequest(task.Method, task.Url, nil)
	if err != nil {
		return err
	}

	// copy headers
	header := make(http.Header)
	if len(task.Headers) > 0 {
		for _, h := range task.Headers {
			matched, err := util.ParseHTTPHeader(h)
			if err != nil {
				return fmt.Errorf("HTTP Header format is invalid, %v", err)
			}
			header.Set(matched[1], matched[2])
		}
	}
	httpReq.Header = header

	// user agent
	ua := httpReq.UserAgent()
	if ua == "" {
		ua = version.GetVersion()
	} else {
		ua += " " + version.GetVersion()
	}
	header.Set("User-Agent", ua)

	// content-type
	header.Set("Content-Type", task.ContentType)

	if len(taskReq.GetHttp().Body) > 0 {
		httpReq.ContentLength = int64(len(taskReq.GetHttp().Body))
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         httpReq.Host,
		},
		MaxIdleConnsPerHost: int(taskReq.Concurrency),
		DisableCompression:  task.DisableCompression,
		DisableKeepAlives:   task.DisableKeepalives,
	}

	if task.ProxyAddr != "" {
		proxyURL, err := url.Parse(task.ProxyAddr)
		if err != nil {
			return fmt.Errorf("invalid argument %s: %s", task.ProxyAddr, err.Error())
		}
		tr.Proxy = http.ProxyURL(proxyURL)
	}

	if task.Http2 {
		http2.ConfigureTransport(tr)
	} else {
		tr.TLSNextProto = make(map[string]func(string, *tls.Conn) http.RoundTripper)
	}
	client := &http.Client{Transport: tr, Timeout: time.Duration(task.Timeout) * time.Second}

	if task.DisableRedirects {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	builder.HttpReq = httpReq
	builder.HttpClient = client
	return nil
}

func (builder *HttpBuilder) PreRequest(taskReq *pb.TaskRequest) (interface{}, interface{}) {
	return nil, nil
}

func (builder *HttpBuilder) Request(ctx context.Context, obj interface{}, taskReq *pb.TaskRequest) interface{} {
	s := util.Now()
	var size int64
	var code int
	var dnsStart, connStart, resStart, reqStart, delayStart time.Duration
	var dnsDuration, connDuration, resDuration, reqDuration, delayDuration time.Duration
	req := util.CloneRequest(builder.HttpReq, []byte(taskReq.GetHttp().Body))
	trace := &httptrace.ClientTrace{
		DNSStart: func(info httptrace.DNSStartInfo) {
			dnsStart = util.Now()
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			dnsDuration = util.Now() - dnsStart
		},
		GetConn: func(h string) {
			connStart = util.Now()
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			if !connInfo.Reused {
				connDuration = util.Now() - connStart
			}
			reqStart = util.Now()
		},
		WroteRequest: func(w httptrace.WroteRequestInfo) {
			reqDuration = util.Now() - reqStart
			delayStart = util.Now()
		},
		GotFirstResponseByte: func() {
			delayDuration = util.Now() - delayStart
			resStart = util.Now()
		},
	}

	req = req.WithContext(ctx)
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	resp, err := builder.HttpClient.Do(req)
	if err == nil {
		size = resp.ContentLength
		code = resp.StatusCode
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}
	t := util.Now()
	resDuration = t - resStart
	finish := t - s

	return &httpResult{
		Offset:        s,
		StatusCode:    code,
		Duration:      finish,
		Err:           err,
		ContentLength: size,
		ConnDuration:  connDuration,
		DnsDuration:   dnsDuration,
		ReqDuration:   reqDuration,
		ResDuration:   resDuration,
		DelayDuration: delayDuration,
	}
}

func (builder *HttpBuilder) PostRequest(result interface{}) error {
	res, ok := result.(*httpResult)
	if !ok {
		return fmt.Errorf("Expected httpResult type, but got %T", result)
	}

	report := builder.report
	report.result.TotalRequests++

	if res.Err != nil {
		report.result.ErrMap[res.Err.Error()]++
	} else {
		report.result.Requests++
		report.result.TotalDuration += res.Duration.Seconds()
		report.result.ConnDuration += res.ConnDuration.Seconds()
		report.result.DnsDuration += res.DnsDuration.Seconds()
		report.result.ReqDuration += res.ReqDuration.Seconds()
		report.result.RespDuration += res.ResDuration.Seconds()
		report.result.DelayDuration += res.DelayDuration.Seconds()
		if res.ContentLength > 0 {
			report.result.TotalSize += res.ContentLength
		}

		if len(report.resLats) < cap(report.resLats) {
			report.result.Lats = append(report.result.Lats, res.Duration.Seconds())
			report.connLats = append(report.connLats, res.ConnDuration.Seconds())
			report.dnsLats = append(report.dnsLats, res.DnsDuration.Seconds())
			report.reqLats = append(report.reqLats, res.ReqDuration.Seconds())
			report.delayLats = append(report.delayLats, res.DelayDuration.Seconds())
			report.resLats = append(report.resLats, res.ResDuration.Seconds())
			report.statusCodes = append(report.statusCodes, res.StatusCode)
			report.offsets = append(report.offsets, res.Offset.Seconds())
		}
	}
	return nil
}

func (builder *HttpBuilder) Done(total time.Duration) (interface{}, error) {
	report := builder.report
	report.result.Duration = total.Seconds()

	statusCodes := make(map[uint32]uint32, len(builder.report.statusCodes))
	for _, statusCode := range builder.report.statusCodes {
		statusCodes[uint32(statusCode)]++
	}

	report.result.StatusCodes = statusCodes
	if len(report.result.Lats) == 0 {
		return report.result, nil
	}

	sort.Float64s(report.result.Lats)
	sort.Float64s(report.connLats)
	sort.Float64s(report.dnsLats)
	sort.Float64s(report.delayLats)
	sort.Float64s(report.reqLats)
	sort.Float64s(report.resLats)

	report.result.Dns.Max = report.dnsLats[len(report.dnsLats)-1]
	report.result.Dns.Min = report.dnsLats[0]
	report.result.Delay.Max = report.delayLats[len(report.delayLats)-1]
	report.result.Delay.Min = report.delayLats[0]
	report.result.Resp.Max = report.resLats[len(report.resLats)-1]
	report.result.Resp.Min = report.resLats[0]
	report.result.Conn.Max = report.connLats[len(report.connLats)-1]
	report.result.Conn.Min = report.connLats[0]
	report.result.Req.Max = report.reqLats[len(report.reqLats)-1]
	report.result.Req.Min = report.reqLats[0]

	report.result.FastestReqTime = report.result.Lats[0]
	report.result.SlowestReqTime = report.result.Lats[len(report.result.Lats)-1]

	return report.result, nil
}

func (builder *HttpBuilder) Merge(messages []string) (interface{}, error) {
	stats := &HttpStats{}
	stats.StatusCodes = make(map[uint32]uint32, 100)
	stats.ErrMap = make(map[string]uint32)

	for _, message := range messages {
		r := &HttpStats{}
		err := json.Unmarshal([]byte(message), r)
		if err != nil {
			return nil, fmt.Errorf("cannot cast to websocketStats: %#v", message)
		}

		if stats.Duration < r.Duration {
			stats.Duration = r.Duration
		}
		stats.TotalRequests += r.TotalRequests
		stats.TotalDuration += r.TotalDuration
		stats.Requests += r.Requests
		stats.Rps += r.Rps
		stats.TotalSize += r.TotalSize
		stats.FastestReqTime = r.FastestReqTime
		stats.SlowestReqTime = r.SlowestReqTime

		stats.ConnDuration += r.ConnDuration
		stats.DnsDuration += r.DnsDuration
		stats.ReqDuration += r.ReqDuration
		stats.RespDuration += r.RespDuration
		stats.DelayDuration += r.DelayDuration

		stats.Req.Max = math.Max(stats.Req.Max, r.Req.Max)
		stats.Req.Min = math.Min(stats.Req.Min, r.Req.Min)

		stats.Conn.Max = math.Max(stats.Conn.Max, r.Conn.Max)
		stats.Conn.Min = math.Min(stats.Conn.Min, r.Conn.Min)

		stats.Delay.Max = math.Max(stats.Delay.Max, r.Delay.Max)
		stats.Delay.Min = math.Min(stats.Delay.Min, r.Delay.Min)

		stats.Resp.Max = math.Max(stats.Resp.Max, r.Resp.Max)
		stats.Resp.Min = math.Min(stats.Resp.Min, r.Resp.Min)

		stats.Dns.Max = math.Max(stats.Dns.Max, r.Dns.Max)
		stats.Dns.Min = math.Min(stats.Dns.Min, r.Dns.Min)

		for k, v := range r.StatusCodes {
			if _, ok := stats.StatusCodes[k]; ok {
				stats.StatusCodes[k] += v
			} else {
				stats.StatusCodes[k] = v
			}
		}

		for k, v := range r.ErrMap {
			if _, ok := stats.ErrMap[k]; ok {
				stats.ErrMap[k] += v
			} else {
				stats.ErrMap[k] = v
			}
		}
		stats.Lats = append(stats.Lats, r.Lats...)
	}

	if stats.Requests > 0 {
		sort.Float64s(stats.Lats)
		stats.LatencyDistribution = latencies(stats)
		stats.AvgReqTime = stats.TotalDuration / float64(stats.Requests)
		stats.AvgSize = stats.TotalSize / stats.Requests
		stats.Req.Avg = stats.ReqDuration / float64(stats.Requests)
		stats.Dns.Avg = stats.DnsDuration / float64(stats.Requests)
		stats.Conn.Avg = stats.ConnDuration / float64(stats.Requests)
		stats.Resp.Avg = stats.RespDuration / float64(stats.Requests)
		stats.Delay.Avg = stats.DelayDuration / float64(stats.Requests)
		stats.Rps = float64(stats.TotalRequests) / stats.Duration
		stats.Lats = nil
	}

	return stats, nil
}

var (
	ResultTmpl = `
Summary:
  Requests:	{{ formatNumberInt64 .TotalRequests }}
  Total:	{{ formatNumber .Duration }} secs
  Slowest:	{{ formatNumber .SlowestReqTime }} secs
  Fastest:	{{ formatNumber .FastestReqTime }} secs
  Average:	{{ formatNumber .AvgReqTime }} secs
  Requests/sec:	{{ formatNumber .Rps }}
  {{ if gt .TotalSize 0 }}
  Total data:	{{ .TotalSize }} bytes
  Size/request:	{{ .AvgSize }} bytes{{ end }}

Latency distribution:{{ range .LatencyDistribution }}
  {{ .Percentage }}% in {{ formatNumber .Latency }} secs{{ end }}

Details (average, fastest, slowest):
  DNS+dialup:	{{ formatNumber .AvgReqTime }} secs, {{ formatNumber .FastestReqTime }} secs, {{ formatNumber .SlowestReqTime }} secs
  DNS-lookup:	{{ formatNumber .Dns.Avg }} secs, {{ formatNumber .Dns.Min }} secs, {{ formatNumber .Dns.Max }} secs
  req write:	{{ formatNumber .Req.Avg }} secs, {{ formatNumber .Req.Min }} secs, {{ formatNumber .Req.Max }} secs
  resp wait:	{{ formatNumber .Delay.Avg }} secs, {{ formatNumber .Delay.Min }} secs, {{ formatNumber .Delay.Max }} secs
  resp read:	{{ formatNumber .Resp.Avg }} secs, {{ formatNumber .Resp.Min }} secs, {{ formatNumber .Resp.Max }} secs

Status code distribution:{{ range $code, $num := .StatusCodes }}
  [{{ $code }}]	{{ $num }} responses{{ end }}

{{ if gt (len .ErrMap) 0 }}Error distribution:{{ range $err, $num := .ErrMap }}
  [{{ $num }}]	{{ $err }}{{ end }}{{ end }}
`
)
