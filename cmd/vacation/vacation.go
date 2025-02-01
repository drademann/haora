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

package vacation

import (
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/cmd/internal/parsing"
	"github.com/drademann/haora/cmd/root"
	"github.com/spf13/cobra"
)

func init() {
	command.Flags().Bool("remove", false, "removes the vacation flag instead")
	root.Command.AddCommand(command)
}

var command = &cobra.Command{
	Use:   "vacation",
	Short: "Marks a day as vacation",
	Long: `Marks a day as vacation so it won't be considered as working day anymore.
Any existing tasks will be removed.

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
		removeFlag, err := cmd.Flags().GetBool("remove")
		if err != nil {
			return err
		}

		day := dayList.Day(workingDate)
		day.Tasks = make([]*data.Task, 0)
		day.IsVacation = !removeFlag

		return data.Save(dayList)
	},
}
