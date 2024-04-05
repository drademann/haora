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

package command

import (
	"github.com/drademann/haora/app/config"
	"github.com/drademann/haora/command/add"
	"github.com/drademann/haora/command/finish"
	"github.com/drademann/haora/command/list"
	"github.com/drademann/haora/command/pause"
	"github.com/drademann/haora/command/version"
	"github.com/spf13/cobra"
	"os"
)

var Root = &cobra.Command{
	Use:   "haora",
	Short: "Time tracking with Haora",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("The Haora CLI. Choose your command ...")
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) { // reset flag so tests can rerun!
		_ = cmd.Flags().Set("date", "")
	},
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	cobra.OnInitialize(config.InitViper)
	Root.PersistentFlags().StringP("date", "d", "", "date for the command to execute on (defaults to today)")
	Root.AddCommand(add.Command)
	Root.AddCommand(finish.Command)
	Root.AddCommand(list.Command)
	Root.AddCommand(pause.Command)
	Root.AddCommand(version.Command)
}

func Execute() {
	var err error
	if err = Root.Execute(); err != nil {
		Root.PrintErrf("error: %v\n\n", err)
		os.Exit(1)
	}
}
