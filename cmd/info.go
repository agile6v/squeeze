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
	"github.com/spf13/cobra"
)

// InfoCmd represents the info command
var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show information about the squeeze cluster.",
	Long:  `Show information about the squeeze cluster.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := util.DoRequest("GET", config.ConfigArgs.HttpAddr+"/info", "", 5)
		if err != nil {
			return err
		}

		fmt.Printf("Info: %s\n", resp)
		return nil
	},
}

func init() {
	InfoCmd.PersistentFlags().StringVar(&config.ConfigArgs.HttpAddr, "httpAddr", "http://127.0.0.1:9998",
		"The address and port of the Squeeze master or slave.")
}
