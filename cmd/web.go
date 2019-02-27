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
	"github.com/agile6v/squeeze/pkg/config"
	"github.com/agile6v/squeeze/pkg/util"
	"github.com/agile6v/squeeze/pkg/server"
	"github.com/spf13/cobra"
)

// WebCmd represents the web command
var WebCmd = &cobra.Command{
	Use:   "web",
	Short: "Backend server that supports the Squeeze UI.",
	Long:  `Backend server that supports the Squeeze UI.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return validate(args)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("run squeeze with web mode.")

		stopChan := make(chan struct{})

		// Create the server
		srv := server.NewServer(server.Web)

		// Initialize the server
		err := srv.Initialize(server.SrvArgs)
		if err != nil {
			return fmt.Errorf("failed to initialize web server: %v", err)
		}

		// Start the server
		err = srv.Start(stopChan)
		if err != nil {
			return fmt.Errorf("failed to start web server: %v", err)
		}

		util.WaitSignal(stopChan)
		return nil
	},
}

func init() {
	WebCmd.PersistentFlags().StringVar(&server.SrvArgs.HTTPAddr, "httpAddr", "http://127.0.0.1:9998",
		"The address and port of the Squeeze master or slave.")
	WebCmd.PersistentFlags().StringVar(&config.ConfigArgs.WebOpts.DSN, "dsn", "",
		`Data Source Name. If you specify --type=mysql, need to set this option.
Format: username:password@protocol(address)/dbname?param=value`)
	WebCmd.PersistentFlags().StringVar(&config.ConfigArgs.WebOpts.File, "file", "/tmp/sqlite.db",
		"SQLite database files. If you specify --type=sqlite, need to set this option.")
	WebCmd.PersistentFlags().StringVar(&config.ConfigArgs.WebOpts.Type, "type", "sqlite",
		"The type of the database, one of the mysql and sqlite.")
	WebCmd.MarkPersistentFlagRequired("type")
}

func validate(args []string) error {
	if config.ConfigArgs.WebOpts.Type != "mysql" && config.ConfigArgs.WebOpts.Type != "sqlite" {
		return fmt.Errorf("option --type must be one of the mysql and sqlite.")
	}

	if config.ConfigArgs.WebOpts.Type == "mysql" && config.ConfigArgs.WebOpts.DSN == "" {
		return fmt.Errorf("option --dsn cannot be empty.")
	}

	return nil
}
