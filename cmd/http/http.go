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
	"os"
	"fmt"
	"math"
	"errors"
	"os/signal"
	"github.com/agile6v/squeeze/pkg/config"
	"github.com/agile6v/squeeze/pkg/pb"
	log "github.com/golang/glog"
	"github.com/spf13/cobra"
	"github.com/agile6v/squeeze/pkg/proto/builder"
)

func HttpCmd(configArgs *config.ProtoConfigArgs) *cobra.Command {
	httpOptions := config.NewHttpOptions()
	httpCmd := &cobra.Command{
		Use:   "http",
		Short: "http protocol benchmark",
		Long:  `http protocol benchmark`,
		Args:  cobra.ExactArgs(1),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return httpOptions.Validate(args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			configArgs.Options = httpOptions
			httpOptions.URL = args[0]
			builder := builder.NewBuilder(pb.Protocol_HTTP)

			c := make(chan os.Signal, 1)
			signal.Notify(c, os.Interrupt)
			go func() {
				<-c
				fmt.Printf("\nCanceling...\n")
				_, err := builder.CancelTask(configArgs)
				if err != nil {
					log.Errorf("failed to cancel task %s", err)
				}
			}()

			resp, err := builder.CreateTask(configArgs)
			if err != nil {
				log.Errorf("failed to create task %s", err)
				if resp != "" {
					return errors.New(resp)
				}
				return err
			}

			ret, err := builder.Render(resp)
			if err != nil {
				log.Errorf("failed to render response, %s", err)
				return err
			}

			fmt.Printf("%s", ret)
			return nil
		},
	}

	httpCmd.PersistentFlags().IntVarP(&httpOptions.Requests, "requests", "n",
		math.MaxInt32, "Number of requests to perform")
	httpCmd.PersistentFlags().StringVarP(&httpOptions.Method, "method", "m",
		"GET", "Method name")
	httpCmd.PersistentFlags().IntVarP(&httpOptions.Concurrency, "concurrency", "c",
		1, "Number of multiple requests to make at a time")
	httpCmd.PersistentFlags().IntVarP(&httpOptions.Timeout, "timeout", "s",
		30, "Seconds to max. wait for each response(Default is 30 seconds)")
	httpCmd.PersistentFlags().IntVarP(&httpOptions.RateLimit, "rateLimit", "q",
		0, "Rate limit, in queries per second (QPS). Default is no rate limit")
	httpCmd.PersistentFlags().IntVarP(&httpOptions.Duration, "duration", "z",
		0, "Duration of application to send requests. if duration is specified, n is ignored.")
	httpCmd.PersistentFlags().BoolVar(&httpOptions.DisableKeepAlive, "disable-keepalive",
		false, "Disable keepalive, connection will use keepalive by default.")
	httpCmd.PersistentFlags().BoolVar(&httpOptions.DisableCompression, "disable-compression",
		false, "Disable compression of body received from the server.")
	httpCmd.PersistentFlags().StringVarP(&httpOptions.ProxyAddr, "proxy", "x",
		"", "HTTP Proxy address as host:port")
	httpCmd.PersistentFlags().StringSliceVar(&httpOptions.Headers, "header", nil,
		"Custom HTTP header.(Repeatable)")
	httpCmd.PersistentFlags().StringVarP(&httpOptions.Body, "body", "d",
		"", "Request body string")
	httpCmd.PersistentFlags().StringVarP(&httpOptions.BodyFile, "bodyfile", "D",
		"", "Request body from file")
	httpCmd.PersistentFlags().StringVarP(&httpOptions.ContentType, "content-type", "T",
		"text/plain", "Content-type header to use for POST/PUT data")
	httpCmd.PersistentFlags().IntVar(&httpOptions.MaxResults, "maxResults", 1000000,
		"The maximum number of response results that can be used")
	httpCmd.PersistentFlags().BoolVar(&httpOptions.HTTP2, "http2",
		false, "Enable http2")

	return httpCmd
}