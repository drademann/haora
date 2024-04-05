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

package pause

import (
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/command/internal/parsing"
	"github.com/spf13/cobra"
	"strings"
	"time"
)

func init() {
	Command.Flags().StringP("start", "s", "", "starting timestamp, like 12:00, of the pause")
}

var Command = &cobra.Command{
	Use:     "pause",
	Aliases: []string{"p", "pa", "pau", "break", "bre", "br"},
	Short:   "Adds a pause to a day",
	Long: `Adds a new pause to a day.

The command accepts the first arg as timestamp, and any following as text (optional), like

$ haora pause 12:00 Lunch`,
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

		if err := pauseAction(workingDate, dayList, startFlag, args); err != nil {
			return err
		}
		return data.Save(dayList)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		_ = cmd.Flags().Set("start", "")
	},
}

func pauseAction(workingDate time.Time, dayList *data.DayList, startFlag string, args []string) error {
	pauseTime, args, err := parsing.Time(startFlag, args)
	if err != nil {
		return err
	}
	text := strings.Join(args, " ")
	day := dayList.Day(workingDate)
	return day.AddNewPause(pauseTime, text)
}
