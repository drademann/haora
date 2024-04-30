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

package remove

import (
	"fmt"
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/cmd/internal/parsing"
	"github.com/spf13/cobra"
	"time"
)

func init() {
	Command.Flags().StringP("start", "s", "", "starting timestamp of the task to delete, like 10:00")
}

var Command = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"r", "re", "rm", "rem"},
	Short:   "Remove a task",
	Long: `Remove a task of a day.

The specific task is identified by its starting timestamp, like so

$ haora delete 10:00
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		workingDateFlag, err := cmd.Flags().GetString("date")
		if err != nil {
			return err
		}
		startFlag, err := cmd.Flags().GetString("start")
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

		if err := removeAction(workingDate, dayList, startFlag, args); err != nil {
			return err
		}
		return data.Save(dayList)
	},
	PostRun: func(cmd *cobra.Command, args []string) { // reset flag so tests can rerun!
		_ = cmd.Flags().Set("start", "")
	},
}

func removeAction(workingDate time.Time, dayList *data.DayList, startFlag string, args []string) error {
	startTimeToDelete, args, err := parsing.Time(startFlag, args)
	if err != nil {
		return err
	}
	day := dayList.Day(workingDate)
	if removed := day.RemoveTask(startTimeToDelete); !removed {
		return fmt.Errorf("no task found at %s", startTimeToDelete.Format("15:04"))
	}
	if day.IsEmpty() {
		day.Unfinished()
	}
	return nil
}
