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

package add

import (
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/cmd/internal/parsing"
	"github.com/spf13/cobra"
	"strings"
	"time"
)

func init() {
	Command.Flags().StringP("start", "s", "", "starting timestamp, like 10:00, of the task")
	Command.Flags().StringP("tags", "t", "", "comma separated tags of the task")
	Command.Flags().Bool("no-tags", false, "set if the new task shall have no tags")
}

var Command = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a", "ad"},
	Short:   "Adds a task to a day",
	Long: `Adds a new task to a day. 

The default and simplest to use format for the add command is 

$ haora add [time] [single tag] [text...]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		workingDateFlag, err := cmd.Flags().GetString("date")
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

		dayList, err := data.Load()
		if err != nil {
			return err
		}
		workingDate, err := parsing.WorkingDate(workingDateFlag)
		if err != nil {
			return err
		}

		if err := addAction(workingDate, dayList, startFlag, tagsFlag, noTagsFlag, args); err != nil {
			return err
		}
		return data.Save(dayList)
	},
	PostRun: func(cmd *cobra.Command, args []string) { // reset flag so tests can rerun!
		_ = cmd.Flags().Set("start", "")
		_ = cmd.Flags().Set("tags", "")
		_ = cmd.Flags().Set("no-tags", "")
	},
}

func addAction(workingDate time.Time, dayList *data.DayList, startFlag, tagsFlag string, noTagsFlag bool, args []string) error {
	startTime, args, err := parsing.Time(startFlag, args)
	if err != nil {
		return err
	}
	var tags []string
	if !noTagsFlag {
		tags, args, err = parseTags(tagsFlag, args)
		if err != nil {
			return err
		}
	}
	text := strings.Join(args, " ")
	day := dayList.Day(workingDate)
	return day.AddNewTask(startTime, text, tags)
}

func parseTags(tagsFlag string, args []string) ([]string, []string, error) {
	if tagsFlag != "" {
		tags := strings.Split(tagsFlag, ",")
		return tags, args, nil
	}
	if len(args) > 0 {
		tags := strings.Split(args[0], ",")
		return tags, args[1:], nil
	}
	return []string{}, args, nil
}
