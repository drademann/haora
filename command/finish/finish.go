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

package finish

import (
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/command/internal/parsing"
	"github.com/spf13/cobra"
)

func init() {
	Command.Flags().StringP("end", "e", "", "finish timestamp, like 17:00, for the day")
	Command.Flags().Bool("remove", false, "removes the set finish timestamp")
}

var Command = &cobra.Command{
	Use:     "finish",
	Aliases: []string{"f", "fi", "fin", "fini"},
	Short:   "Mark the day as done",
	Long: `Marks the day as done by setting its final end timestamp. 
The command accepts the first arg as timestamp:

$ haora finish 17:00`,
	RunE: func(cmd *cobra.Command, args []string) error {
		workingDateFlag, err := cmd.Flags().GetString("date")
		if err != nil {
			return err
		}
		endFlag, err := cmd.Flags().GetString("end")
		if err != nil {
			return err
		}
		removeFlag, err := cmd.Flags().GetBool("remove")
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
		day := dayList.Day(workingDate)

		if removeFlag {
			day.Unfinished()
		} else if err := finishAction(day, endFlag, args); err != nil {
			return err
		}
		return data.Save(dayList)
	},
	PostRun: func(cmd *cobra.Command, args []string) { // reset flag so tests can rerun!
		_ = cmd.Flags().Set("end", "")
	},
}

func finishAction(day *data.Day, endFlag string, args []string) error {
	endTime, _, err := parsing.Time(endFlag, args)
	if err != nil {
		return err
	}
	if err = day.Finish(endTime); err != nil {
		return err
	}
	return nil
}
