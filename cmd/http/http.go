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
	"net/url"
	"os/signal"
	"github.com/agile6v/squeeze/pkg/config"
	"github.com/agile6v/squeeze/pkg/pb"
	"github.com/agile6v/squeeze/pkg/proto/http"
	"github.com/agile6v/squeeze/pkg/util"
	log "github.com/golang/glog"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "http",
	Short: "http protocol benchmark",
	Long:  `http protocol benchmark`,
	Args:  cobra.ExactArgs(1),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return validate(args)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		config.ConfigArgs.HttpOpts.URL = args[0]
		builder := http.NewBuilder()

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			<-c
			fmt.Printf("\nCanceling...\n")
			_, err := builder.CancelTask(&config.ConfigArgs, pb.Protocol_HTTP)
			if err != nil {
				log.Errorf("failed to cancel task %s", err)
			}
		}()

		resp, err := builder.CreateTask(&config.ConfigArgs)
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

func init() {
	Command.PersistentFlags().IntVarP(&config.ConfigArgs.HttpOpts.Requests, "requests", "n",
		math.MaxInt32, "Number of requests to perform")
	Command.PersistentFlags().StringVarP(&config.ConfigArgs.HttpOpts.Method, "method", "m",
		"GET", "Method name")
	Command.PersistentFlags().IntVarP(&config.ConfigArgs.HttpOpts.Concurrency, "concurrency", "c",
		1, "Number of multiple requests to make at a time")
	Command.PersistentFlags().IntVarP(&config.ConfigArgs.HttpOpts.Timeout, "timeout", "s",
		30, "Seconds to max. wait for each response(Default is 30 seconds)")
	Command.PersistentFlags().IntVarP(&config.ConfigArgs.HttpOpts.RateLimit, "rateLimit", "q",
		0, "Rate limit, in queries per second (QPS). Default is no rate limit")
	Command.PersistentFlags().IntVarP(&config.ConfigArgs.HttpOpts.Duration, "duration", "z",
		0, "Duration of application to send requests. if duration is specified, n is ignored.")
	Command.PersistentFlags().BoolVar(&config.ConfigArgs.HttpOpts.DisableKeepAlive, "disable-keepalive",
		false, "Disable keepalive, connection will use keepalive by default.")
	Command.PersistentFlags().StringVarP(&config.ConfigArgs.HttpOpts.ProxyAddr, "proxy", "x",
		"", "HTTP Proxy address as host:port")
	Command.PersistentFlags().StringSliceVar(&config.ConfigArgs.HttpOpts.Headers, "header", nil,
		"Custom HTTP header.(Repeatable)")
	Command.PersistentFlags().StringVarP(&config.ConfigArgs.HttpOpts.Body, "body", "b",
		"", "Request body string")
	Command.PersistentFlags().StringVarP(&config.ConfigArgs.HttpOpts.ContentType, "content-type", "T",
		"text/plain", "Content-type header to use for POST/PUT data")
	Command.PersistentFlags().IntVar(&config.ConfigArgs.HttpOpts.MaxResults, "maxResults", 1000000,
		"The maximum number of response results that can be used")
}

func validate(args []string) error {
	// Check the validity of the concurrency
	if config.ConfigArgs.HttpOpts.Concurrency < 1 {
		return fmt.Errorf("option --concurrency must be greater than 0.")
	}

	// Check if the options are missing
	if config.ConfigArgs.HttpOpts.Requests == 0 && config.ConfigArgs.HttpOpts.Duration == 0 {
		return fmt.Errorf("option --requests or --duration must be specified one of them.")
	}

	//
	if config.ConfigArgs.HttpOpts.Duration == 0 {
		if config.ConfigArgs.HttpOpts.Requests < config.ConfigArgs.HttpOpts.Concurrency {
			return fmt.Errorf("option --concurrecny must be greater than --requests.")
		}
	}

	// Check if the format of http headers' is vaild
	if len(config.ConfigArgs.HttpOpts.Headers) > 0 {
		for _, h := range config.ConfigArgs.HttpOpts.Headers {
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

	if config.ConfigArgs.HttpOpts.ProxyAddr != "" {
		_, err := url.Parse(config.ConfigArgs.HttpOpts.ProxyAddr)
		if err != nil {
			return fmt.Errorf("invalid argument %s: %s", config.ConfigArgs.HttpOpts.ProxyAddr, err.Error())
		}
	}

	return nil
}
