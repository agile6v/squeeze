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

package udp

import (
	"fmt"
	"net"
)

type UDPOptions struct {
	Addr        string
	Requests    int
	Concurrency int
	Timeout     int
	Duration    int
	MsgLength   int
	MaxResults  int
}

func NewUDPOptions() *UDPOptions {
	return &UDPOptions{}
}

func (udpOptions *UDPOptions) Validate(args []string) error {
	if udpOptions.Concurrency < 1 {
		return fmt.Errorf("option --concurrency must be greater than 0.")
	}

	// Check the validity of the target address
	_, _, err := net.SplitHostPort(args[0])
	if err != nil {
		return err
	}

	udpOptions.Addr = args[0]
	return nil
}
