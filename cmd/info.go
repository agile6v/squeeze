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
	"bytes"
	"encoding/json"
	"github.com/agile6v/squeeze/pkg/config"
	"github.com/agile6v/squeeze/pkg/util"
	"github.com/spf13/cobra"
	log "github.com/golang/glog"
	"github.com/agile6v/squeeze/pkg/server"
)

func InfoCmd() *cobra.Command {
	configArgs := config.NewConfigArgs(nil)

	// infoCmd represents the info command
	infoCmd := &cobra.Command{
		Use:   "info",
		Short: "Show information about the squeeze cluster.",
		Long:  `Show information about the squeeze cluster.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := util.DoRequest("GET", configArgs.HttpAddr+"/info", "", 5)
			if err != nil {
				return err
			}

			info := &server.ClusterInfo{}
			ret, err := render(resp, &InfoTemplate, info)
			if err != nil {
				log.Errorf("failed to render response, %s", err)
				return err
			}

			fmt.Printf("%s", ret)

			return nil
		},
	}

	infoCmd.PersistentFlags().StringVar(&configArgs.HttpAddr, "httpAddr", "http://127.0.0.1:9998",
		"The address and port of the Squeeze master or slave.")
	return infoCmd
}

func render(data string, tmpl *string, info *server.ClusterInfo) (string, error) {
	err := json.Unmarshal([]byte(data), info)
	if err != nil {
		return "", err
	}

	buf := &bytes.Buffer{}
	if err := util.NewTemplate(*tmpl).Execute(buf, info); err != nil {
		return "", err
	}

	return buf.String(), nil
}

var (
	InfoTemplate = `
Cluster Information:
{{ if ne .Error "" }}
    {{.Error}}
{{ else }}
{{ range .Data }}   Agent: {{ .Addr }}, {{ .Status }}
{{ end }}
{{ end }}
`
)


