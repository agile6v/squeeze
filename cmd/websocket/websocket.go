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
	"os/signal"
	"github.com/agile6v/squeeze/pkg/config"
	"github.com/agile6v/squeeze/pkg/pb"
	log "github.com/golang/glog"
	"github.com/spf13/cobra"
	"github.com/agile6v/squeeze/pkg/proto/builder"
)

func WsCmd(configArgs *config.ProtoConfigArgs) *cobra.Command {
	wsOptions := config.NewWsOptions()
	wsCmd := &cobra.Command{
		Use:   "websocket",
		Short: "websocket protocol benchmark",
		Long:  `websocket protocol benchmark`,
		Args:  cobra.ExactArgs(1),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return wsOptions.Validate(args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			configArgs.Options = wsOptions
			builder := builder.NewBuilder(pb.Protocol_WEBSOCKET)

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

			if configArgs.Callback != "" {
				fmt.Printf("%s", resp)
				return nil
			}

			ret, err := builder.Render(resp, configArgs.Callback)
			if err != nil {
				log.Errorf("failed to render response, %s", err)
				return err
			}

			fmt.Printf("%s", ret)
			return nil
		},
	}

	wsCmd.PersistentFlags().IntVarP(&wsOptions.Requests, "requests", "n",
		math.MaxInt32, "Number of requests to perform")
	wsCmd.PersistentFlags().IntVarP(&wsOptions.Concurrency, "concurrency", "c",
		1, "Number of multiple requests to make at a time")
	wsCmd.PersistentFlags().IntVarP(&wsOptions.Timeout, "timeout", "s",
		30, "Websocket handshake timeout in seconds (Default is 30 seconds)")
	wsCmd.PersistentFlags().StringVarP(&wsOptions.Body, "body", "b",
		"", "Request body string")
	wsCmd.PersistentFlags().IntVarP(&wsOptions.Duration, "duration", "z",
		0, "Duration of application to send requests. if duration is specified, n is ignored.")
	wsCmd.PersistentFlags().IntVar(&wsOptions.MaxResults, "maxResults", 1000000,
		"The maximum number of response results that can be used")

	return wsCmd
}

