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

// SlaveCmd represents the slave command
var SlaveCmd = &cobra.Command{
	Use:   "slave",
	Short: "Squeeze slave node.",
	Long:  `Slave initiates stress testing to the target.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("run squeeze with slave mode.")

		stopChan := make(chan struct{})

		// Create the server
		srv := server.NewServer(server.Slave)

		// Initialize the server
		err := srv.Initialize(server.SrvArgs)
		if err != nil {
			return fmt.Errorf("failed to initialize slave server: %v", err)
		}

		// Start the server
		err = srv.Start(stopChan)
		if err != nil {
			return fmt.Errorf("failed to start slave server: %v", err)
		}

		util.WaitSignal(stopChan)

		return nil
	},
}

func init() {
	SlaveCmd.PersistentFlags().StringVar(&server.SrvArgs.HTTPAddr, "httpAddr", ":9998",
		"Squeeze service HTTP address")
	SlaveCmd.PersistentFlags().StringVar(&server.SrvArgs.GRPCAddr, "grpcAddr", ":9997",
		"Squeeze service grpc address")
	SlaveCmd.PersistentFlags().StringVar(&server.SrvArgs.MasterAddr, "masterAddr", "",
		"The address of the master server")
	SlaveCmd.PersistentFlags().StringVar(&server.SrvArgs.GrpcMasterAddr, "grpcMasterAddr", "",
		"The address of the grpc master server")
	SlaveCmd.PersistentFlags().DurationVar(&server.SrvArgs.ReportInterval, "reportInterval", 5,
		"Task reporting interval to the master")
	SlaveCmd.PersistentFlags().IntVar(&server.SrvArgs.ResultCapacity, "resultCapacity", 2000,
		"The capacity of the results channel for aggregating.")
	SlaveCmd.MarkPersistentFlagRequired("masterAddr")
	SlaveCmd.MarkPersistentFlagRequired("grpcMasterAddr")
}
