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

package list

import (
	"github.com/drademann/haora/app/data"
	"github.com/spf13/cobra"
	"time"
)

var (
	tagsFlag bool
	weekFlag bool
)

func init() {
	Command.Flags().BoolVarP(&tagsFlag, "tags", "t", false, "shows durations per tag")
	Command.Flags().BoolVarP(&weekFlag, "week", "w", false, "shows week summary")
}

var Command = &cobra.Command{
	Use:   "list",
	Short: "List the recorded tasks of a day",
	Long:  `Provides a list of all tasks of a day, including their duration. A summary with total pause and working times is displayed at the end.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		d := data.State.WorkingDay()
		if tagsFlag {
			return printTags(*d, cmd)
		}
		if weekFlag {
			return printWeek(*d, cmd)
		}
		return printDefault(*d, cmd)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		tagsFlag = false
	},
}

func sign(d time.Duration) string {
	if d < 0 {
		return "-"
	}
	return "+"
}
