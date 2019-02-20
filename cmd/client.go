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
	"github.com/agile6v/squeeze/cmd/http"
	"github.com/agile6v/squeeze/cmd/websocket"
	"github.com/agile6v/squeeze/pkg/config"
	"github.com/spf13/cobra"
)

// ClientCmd represents the client command
var ClientCmd = &cobra.Command{
	Use:   "client",
	Short: "A handy tool that can call the Squeeze's API.",
	Long: `This command allows you to interact with Squeeze and stress targets with multiple protocols.
Currently supported protocol is only http, other protocols are under development. Look forward
to your contribution.
	`,
}

// StopCmd represents the stop command
var StopCmd = &cobra.Command{
	Use:   "stop",
	Short: "",
	Long:  ``,
}

func init() {
	ClientCmd.PersistentFlags().StringVar(&config.ConfigArgs.Callback, "callback", "",
		"If this call is asynchronous then stress result will be sent to the address.")
	ClientCmd.PersistentFlags().StringVar(&config.ConfigArgs.HttpAddr, "httpAddr", "http://127.0.0.1:9998",
		"The address and port of the Squeeze master or slave.")

	ClientCmd.AddCommand(http.Command)
	ClientCmd.AddCommand(websocket.Command)
	ClientCmd.AddCommand(StopCmd)
}
