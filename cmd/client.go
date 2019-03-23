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

package cmd

import (
	"fmt"
	log "github.com/golang/glog"
	"github.com/spf13/cobra"
	"github.com/agile6v/squeeze/cmd/http"
	"github.com/agile6v/squeeze/cmd/websocket"
	"github.com/agile6v/squeeze/pkg/config"
	"github.com/agile6v/squeeze/pkg/proto"
)

func ClientCmd() *cobra.Command {
	configArgs := config.NewConfigArgs(nil)
	// clientCmd represents the client command
	clientCmd := &cobra.Command{
		Use:   "client",
		Short: "A handy tool that can call the Squeeze's API.",
		Long: `This command allows you to interact with Squeeze and stress targets with multiple protocols.
Currently supported protocol is only http, other protocols are under development. Look forward
to your contribution.
	`,
	}

	clientCmd.PersistentFlags().StringVar(&configArgs.Callback, "callback", "",
		"If this call is asynchronous then stress result will be sent to the address.")
	clientCmd.PersistentFlags().StringVar(&configArgs.HttpAddr, "httpAddr", "http://127.0.0.1:9998",
		"The address and port of the Squeeze master or slave.")

	clientCmd.AddCommand(http.HttpCmd(configArgs))
	clientCmd.AddCommand(websocket.WsCmd(configArgs))
	clientCmd.AddCommand(StopCmd(configArgs))

	return clientCmd
}

// stopCmd represents the stop command
func StopCmd(configArgs *config.ProtoConfigArgs) *cobra.Command {
	stopCmd := &cobra.Command{
		Use:   "stop",
		Short: "Stop the running task",
		Long:  `Stop the running task`,
		RunE: func(cmd *cobra.Command, args []string) error {
			builder := &proto.ProtoBuilderBase{}
			_, err := builder.CancelTask(configArgs)
			if err != nil {
				log.Errorf("failed to cancel task %s", err)
			}
			fmt.Printf("\nCancelled\n")
			return nil
		},
	}
	return stopCmd
}