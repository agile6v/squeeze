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
	"github.com/agile6v/squeeze/pkg/server"
	"github.com/agile6v/squeeze/pkg/util"
	"github.com/spf13/cobra"
)

func MasterCmd() *cobra.Command {
	serverArgs := server.NewServerArgs()

	// masterCmd represents the master command
	masterCmd := &cobra.Command{
		Use:   "master",
		Short: "Squeeze master node",
		Long:  `Master node is responsible for managing all slave nodes.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("run squeeze with master mode.")

			stopChan := make(chan struct{})

			// Create the server
			srv := server.NewServer(server.Master)

			// Initialize the server
			err := srv.Initialize(serverArgs)
			if err != nil {
				return fmt.Errorf("failed to initialize master server: %v", err)
			}

			// Start the server
			err = srv.Start(stopChan)
			if err != nil {
				return fmt.Errorf("failed to start master server: %v", err)
			}

			util.WaitSignal(stopChan)

			return nil
		},
	}

	masterCmd.PersistentFlags().StringVar(&serverArgs.HTTPAddr, "httpAddr", ":9998",
		"Squeeze service HTTP address")
	masterCmd.PersistentFlags().StringVar(&serverArgs.GRPCAddr, "grpcAddr", ":9997",
		"Squeeze service grpc address")
	return masterCmd
}
