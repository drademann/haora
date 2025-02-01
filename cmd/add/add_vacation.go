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

package add

import (
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/cmd/internal/parsing"
	"github.com/spf13/cobra"
	"time"
)

func init() {
	command.AddCommand(vacationCommand)
}

var vacationCommand = &cobra.Command{
	Use:   "vacation",
	Short: "Marks the day as vacation.",
	Long: `Adds a new task to a day.

$ haora add vacation`,
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

		if err := addVacationAction(workingDate, dayList); err != nil {
			return err
		}
		return data.Save(dayList)
	},
}

func addVacationAction(workingDate time.Time, dayList *data.DayList) error {
	day := dayList.Day(workingDate)
	day.IsVacation = true
	return nil
}
