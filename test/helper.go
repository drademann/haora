//
// Copyright 2024-2024 The Haora Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package test

import (
	"bytes"
	"github.com/spf13/cobra"
	"strings"
	"testing"
)

func ExecuteCommand(t *testing.T, cmd *cobra.Command, argLine string) *bytes.Buffer {
	t.Helper()
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs(strings.Split(argLine, " "))
	if err := cmd.Execute(); err != nil {
		cmd.PrintErrf("error: %v\n\n", err)
	}
	return buf
}
