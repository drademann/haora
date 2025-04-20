//
// Copyright 2024-2025 The Haora Authors
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

package version

import (
	"github.com/drademann/haora/cmd/root"
	"github.com/spf13/cobra"
)

const version = "1.8.0"

func init() {
	root.Command.AddCommand(command)
}

var command = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v", "ve", "ver", "vers"},
	Short:   "Print the version",
	Long:    `Prints the version number of Haora.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("Haora v%s\n", version)
	},
}
