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
	"github.com/drademann/haora/cmd/internal/parsing"
	"github.com/spf13/cobra"
	"time"
)

func init() {
	Command.Flags().BoolP("tags", "t", false, "shows durations per tag")
	Command.Flags().BoolP("week", "w", false, "shows week summary")
}

var Command = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l", "li", "lis"},
	Short:   "List the recorded tasks of a day",
	Long:    `Provides a list of all tasks of a day, including their duration. A summary with total pause and working times is displayed at the end.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		workingDateFlag, err := cmd.Flags().GetString("date")
		if err != nil {
			return err
		}

		dayList, err := data.Load()
		if err != nil {
			return err
		}
		workingDate, err := parsing.WorkingDate(workingDateFlag)
		if err != nil {
			return err
		}

		tagsFlag, err := cmd.Flags().GetBool("tags")
		if err != nil {
			return err
		}
		if tagsFlag {
			return printTags(cmd, workingDate, dayList)
		}

		weekFlag, err := cmd.Flags().GetBool("week")
		if err != nil {
			return err
		}
		if weekFlag {
			return printWeek(cmd, workingDate, dayList)
		}

		return printDefault(cmd, workingDate, dayList)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		_ = cmd.Flags().Set("tags", "")
		_ = cmd.Flags().Set("week", "")
	},
}

func sign(d time.Duration) string {
	if d < 0 {
		return "-"
	}
	return "+"
}
