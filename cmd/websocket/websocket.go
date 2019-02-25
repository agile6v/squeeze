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
	"os"
	"fmt"
	"math"
	"errors"
	"net/url"
	"os/signal"
	"github.com/agile6v/squeeze/pkg/config"
	"github.com/agile6v/squeeze/pkg/pb"
	"github.com/agile6v/squeeze/pkg/proto/websocket"
	log "github.com/golang/glog"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "websocket",
	Short: "websocket protocol benchmark",
	Long:  `websocket protocol benchmark`,
	Args:  cobra.ExactArgs(1),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return validate(args)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		builder := websocket.NewBuilder()

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			<-c
			fmt.Printf("\nCanceling...\n")
			_, err := builder.CancelTask(&config.ConfigArgs, pb.Protocol_WEBSOCKET)
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

		if config.ConfigArgs.Callback != "" {
			fmt.Printf("%s", resp)
			return nil
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
	Command.PersistentFlags().IntVarP(&config.ConfigArgs.WsOpts.Requests, "requests", "n",
		math.MaxInt32, "Number of requests to perform")
	Command.PersistentFlags().IntVarP(&config.ConfigArgs.WsOpts.Concurrency, "concurrency", "c",
		1, "Number of multiple requests to make at a time")
	Command.PersistentFlags().IntVarP(&config.ConfigArgs.WsOpts.Timeout, "timeout", "s",
		30, "Websocket handshake timeout in seconds (Default is 30 seconds)")
	Command.PersistentFlags().StringVarP(&config.ConfigArgs.WsOpts.Body, "body", "b",
		"", "Request body string")
	Command.PersistentFlags().IntVarP(&config.ConfigArgs.WsOpts.Duration, "duration", "z",
		0, "Duration of application to send requests. if duration is specified, n is ignored.")
	Command.PersistentFlags().IntVar(&config.ConfigArgs.WsOpts.MaxResults, "maxResults", 1000000,
		"The maximum number of response results that can be used")
}

func validate(args []string) error {
	if config.ConfigArgs.HttpOpts.Concurrency < 1 {
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

	config.ConfigArgs.WsOpts.Scheme = u.Scheme
	config.ConfigArgs.WsOpts.Host = u.Host
	config.ConfigArgs.WsOpts.Path = u.Path

	return nil
}
