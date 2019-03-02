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
	goflag "flag"

	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
)

// RootCmd is the main command of the Squeeze
var RootCmd = &cobra.Command{
	Use:          "squeeze",
	Short:        "A Load Testing Tool.",
	Long:         "Squeeze provides scalable and easy-to-use load testing tool for performance testing.",
	SilenceUsage: true,
}

func init() {
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	// For https://github.com/kubernetes/kubernetes/issues/17162#issuecomment-225596212
	flag.CommandLine.Parse([]string{})

	RootCmd.AddCommand(MasterCmd())
	RootCmd.AddCommand(SlaveCmd())
	RootCmd.AddCommand(ClientCmd())
	RootCmd.AddCommand(WebCmd())
	RootCmd.AddCommand(InfoCmd())
	RootCmd.AddCommand(VersionCmd)
}
