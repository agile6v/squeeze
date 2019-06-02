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

package tcp

import (
	"math"
	"github.com/spf13/cobra"
	"github.com/agile6v/squeeze/pkg/config"
	"github.com/agile6v/squeeze/pkg/pb"
	"github.com/agile6v/squeeze/pkg/proto/builder"
	"github.com/agile6v/squeeze/pkg/proto/tcp"
)

func Command(configArgs *config.ProtoConfigArgs) *cobra.Command {
	tcpOptions := tcp.NewTCPOptions()
	tcpCmd := &cobra.Command{
		Use:   "tcp",
		Short: "tcp protocol benchmark",
		Long:  `tcp protocol benchmark`,
		Args:  cobra.ExactArgs(1),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return tcpOptions.Validate(args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			configArgs.Options = tcpOptions
			builder := builder.NewBuilder(pb.Protocol_TCP)
            return builder.RunTask(configArgs)
		},
	}

	tcpCmd.PersistentFlags().IntVarP(&tcpOptions.Requests, "requests", "n",
		math.MaxInt32, "Number of requests to perform")
	tcpCmd.PersistentFlags().IntVarP(&tcpOptions.Concurrency, "concurrency", "c",
		1, "Number of multiple requests to make at a time")
	tcpCmd.PersistentFlags().IntVarP(&tcpOptions.MsgLength, "message-length", "l",
		1, "The length of the message to send")
	tcpCmd.PersistentFlags().IntVarP(&tcpOptions.Timeout, "timeout", "s",
		30, "Timeout in seconds (Default is 30 seconds)")
	tcpCmd.PersistentFlags().IntVarP(&tcpOptions.Duration, "duration", "z",
		0, "Duration of application to send requests. if duration is specified, n is ignored.")
	tcpCmd.PersistentFlags().IntVar(&tcpOptions.MaxResults, "maxResults", 1000000,
		"The maximum number of response results that can be used")

	return tcpCmd
}

