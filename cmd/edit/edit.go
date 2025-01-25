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

package edit

import (
	"errors"
	"fmt"
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/cmd/internal/parsing"
	"github.com/drademann/haora/cmd/root"
	"github.com/spf13/cobra"
	"strings"
	"time"
)

const IgnoreFlag = "╳"

func init() {
	command.Flags().StringP("update", "u", IgnoreFlag, "existing start timestamp to update, like 10:00, of the task")
	command.Flags().StringP("start", "s", IgnoreFlag, "new start timestamp, like 10:00, of the task")
	command.Flags().StringP("tags", "t", IgnoreFlag, "comma separated tags of the task")
	command.Flags().Bool("no-tags", false, "set if the new task should have no tags")
	command.Flags().StringP("text", "x", IgnoreFlag, "text to replace the existing text")
	root.Command.AddCommand(command)
}

var command = &cobra.Command{
	Use:     "edit",
	Aliases: []string{"e", "ed"},
	Short:   "Edit an existing task",
	Long: `Edit an existing task. 

The task to edit is chosen by its start time. Anything not set won't be changed.
Unlike the use of the add command, all flags to update must be set explicitly. 
All plain arguments are added to the text flag (-x).

Examples:

  $ haora edit -u 09:30 -s 10:00 -t programming -x "some more Go code"
  $ haora edit -u 10:00 -x "was Kotlin code"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		workingDateFlag, err := cmd.Flags().GetString("date")
		if err != nil {
			return err
		}
		updateFlag, err := cmd.Flags().GetString("update")
		if err != nil {
			return err
		}
		startFlag, err := cmd.Flags().GetString("start")
		if err != nil {
			return err
		}
		tagsFlag, err := cmd.Flags().GetString("tags")
		if err != nil {
			return err
		}
		noTagsFlag, err := cmd.Flags().GetBool("no-tags")
		if err != nil {
			return err
		}
		textFlag, err := cmd.Flags().GetString("text")
		if err != nil {
			return err
		}
		if len(args) > 0 {
			textFlag = fmt.Sprintf("%s %s", textFlag, strings.Join(args, " "))
		}

		dayList, err := data.Load()
		if err != nil {
			return err
		}
		workingDate, err := parsing.WorkingDate(workingDateFlag)
		if err != nil {
			return err
		}

		if err := editAction(workingDate, dayList, updateFlag, startFlag, tagsFlag, noTagsFlag, textFlag); err != nil {
			return err
		}
		return data.Save(dayList)
	},
	PostRun: func(cmd *cobra.Command, args []string) { // reset flag so tests can rerun!
		_ = cmd.Flags().Set("update", "")
		_ = cmd.Flags().Set("start", IgnoreFlag)
		_ = cmd.Flags().Set("tags", IgnoreFlag)
		_ = cmd.Flags().Set("no-tags", "false")
		_ = cmd.Flags().Set("text", IgnoreFlag)
	},
}

func editAction(workingDate time.Time, dayList *data.DayList, updateFlag, startFlag string, tagsFlag string, noTagsFlag bool, textFlag string) error {
	updateStartTime, err := parsing.Time(updateFlag)
	if err != nil {
		return err
	}
	var newStartTime *time.Time = nil
	if startFlag != IgnoreFlag {
		s, err := parsing.Time(startFlag)
		if err == nil {
			newStartTime = &s
		}
	}
	var tags []string = nil
	if tagsFlag != IgnoreFlag {
		tags = parsing.Tags(tagsFlag)
	}
	if noTagsFlag {
		tags = []string{}
	}
	var text *string = nil
	if textFlag != IgnoreFlag {
		text = &textFlag
	}

	day := dayList.Day(workingDate)
	err = day.EditTask(updateStartTime, newStartTime, text, tags)
	if errors.Is(err, data.NoTask) {
		return fmt.Errorf("no task found starting at %s", updateStartTime.Format("15:04"))
	}
	return nil
}
