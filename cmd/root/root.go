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

package root

import (
	"github.com/drademann/haora/app/config"
	"github.com/spf13/cobra"
	"os"
)

var Command = &cobra.Command{
	Use:   "haora",
	Short: "Time tracking with Haora",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("The Haora CLI.")
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) { // reset flag so tests can rerun!
		_ = cmd.Flags().Set("date", "")
	},
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	cobra.OnInitialize(config.InitViper)
	Command.PersistentFlags().StringP("date", "d", "", "date for the command to execute on (defaults to today)")
}

func Execute() {
	var err error
	if err = Command.Execute(); err != nil {
		Command.PrintErrf("error: %v\n", err)
		os.Exit(1)
	}
}
