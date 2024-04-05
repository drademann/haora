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
}

var Command = &cobra.Command{
	Use:   "finish",
	Short: "Mark the day as done",
	Long: `Marks the day as done by setting its final end timestamp. 
The command accepts the first arg as timestamp:

$ haora finish 17:00`,
	RunE: func(cmd *cobra.Command, args []string) error {
		endFlag, err := cmd.Flags().GetString("end")
		if err != nil {
			return err
		}
		return finishAction(endFlag, args)
	},
	PostRun: func(cmd *cobra.Command, args []string) { // reset flag so tests can rerun!
		_ = cmd.Flags().Set("end", "")
	},
}

func finishAction(endFlag string, args []string) error {
	time, _, err := parsing.Time(endFlag, args)
	if err != nil {
		return err
	}
	day := data.State.WorkingDay()
	day.Finish(time)
	return nil
}
