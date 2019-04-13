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
	"context"
	"sync"
	"time"

	"github.com/agile6v/squeeze/pkg/pb"
	build "github.com/agile6v/squeeze/pkg/proto"
	"github.com/agile6v/squeeze/pkg/util"
	log "github.com/golang/glog"
)

type Work struct {
	// Protocol constructor
	Builder build.ProtoBuilder

	// Request task to be executed
	Req *pb.ExecuteTaskRequest

	// Requests is the total number of requests to make.
	Requests int

	// Number of goroutines to run (Concurrent connections)
	Workers int

	Ctx    context.Context
	Cancel context.CancelFunc

	// RateLimit is the rate limit in queries per second.
	RateLimit float64

	// The capacity of the result of the collector channel
	ResultCapacity int

	// Used to aggregate the result of each request
	results chan interface{}

	// Use it to check if collecotr goroutine has exited
	done chan bool
}

// Run starts collecotr & worker goroutine. It blocks until
// all work is done or receive cancel signal.
func (w *Work) Run(ctx context.Context) (time.Duration, error) {
	// Initialization before calling the request handler
	err := w.Builder.Init(ctx, w.Req)
	if err != nil {
		log.Infof("failed to execute Init: %s", err.Error())
		return 0, err
	}

	w.results = make(chan interface{}, w.Workers * w.ResultCapacity)
	w.done = make(chan bool, 1)
	start := util.Now()

	// Start collector goroutine to collect the results
	// produced by the worker goroutines
	go func() {
		w.runCollector()
	}()

	// Start all worker goroutines
	w.runWorkers(ctx)
	return w.Finish(start), nil
}

// Stop will terminate all worker goroutines & collector groutine
func (w *Work) Stop() {
	log.Infof("stop workers(%d) goroutines & collecotr goroutine", w.Workers)
	w.Cancel()
}

// Finish waits for worker goroutines & collecotr goroutine to exit and return
// time spent on pressure measurement.
func (w *Work) Finish(start time.Duration) time.Duration {
	close(w.results)
	total := util.Now() - start

	// Wait until the collector is done.
	<-w.done
	return total
}

// runWorker starts to press target n times and limit the frequency of pressing
func (w *Work) runWorker(ctx context.Context, n int) {
	var throttle <-chan time.Time
	if w.RateLimit > 0 {
		throttle = time.Tick(time.Duration(1e6/(w.RateLimit)) * time.Microsecond)
	}

	obj, result := w.Builder.PreRequest(w.Req)
	if result != nil {
		w.results <- result
		return
	}
	for i := 0; i < n; i++ {
		select {
		case <-ctx.Done():
			log.V(2).Infof("worker goroutine exited.")
			return
		default:
			if w.RateLimit > 0 {
				<-throttle
			}

			// Initiate a request
			ret := w.Builder.Request(ctx, obj, w.Req)

			// The result of the request will be sent to collector goroutine for summary
			w.results <- ret
		}
	}

	err := w.Builder.Destroy(obj)
	if err != nil {
		// TODO: exception handing
	}
}

// runWorkers starts all workers(groutines) to press target
func (w *Work) runWorkers(ctx context.Context) {
	var wg sync.WaitGroup
	wg.Add(w.Workers)

	// Start goroutines according to the number of concurrency.
	// Note: Master specifies how many concurrent numbers per slave.
	for i := 0; i < w.Workers; i++ {
		go func(i int) {
			n := w.Requests / w.Workers
			if (i + 1) == w.Workers {
				n += w.Requests % w.Workers
			}
			w.runWorker(ctx, n)
			wg.Done()
		}(i)
	}
	wg.Wait()

	log.Infof("All worker goroutines exited.")
}

// collector is used to collect results which all worker goroutines generated.
func (w *Work) runCollector() {
	log.V(2).Infof("Collector Start")
	// The result is written to the results channel after each
	// request is completed, and collector polls the results channel
	// until it is closed.
	for result := range w.results {
		err := w.Builder.PostRequest(result)
		if err != nil {
			log.Error("PostRequest got error : ", err)
		}
	}

	// If the result channel is closed then notify the main goroutine to collect results.
	w.done <- true
	log.V(2).Infof("Collector exited.")
}
