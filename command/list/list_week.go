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
	"github.com/drademann/haora/app/datetime"
	"github.com/drademann/haora/command/internal/format"
	"github.com/spf13/cobra"
	"time"
)

func printWeek(d data.Day, cmd *cobra.Command) error {
	date := datetime.FindWeekday(d.Date, datetime.Previous, time.Monday)
	week := data.CollectWeek(date)
	for _, day := range week.Days {
		dateStr := day.Date.Format("Mon 02.01.2006")
		if day.IsEmpty() {
			cmd.Printf("%s   -\n", dateStr)
		} else {
			startStr := day.Start().Format("15:04")
			endStr := day.End().Format("15:04")
			dur := day.TotalWorkDuration()
			durStr := format.Duration(dur)
			cmd.Printf("%s   %s - %s  worked %s\n", dateStr, startStr, endStr, durStr)
		}
	}
	overtime, exist := week.TotalOvertimeDuration()
	if !exist || overtime == 0 {
		cmd.Printf("\n                          total worked %s\n", format.Duration(week.TotalWorkDuration()))
	} else {
		cmd.Printf("\n                          total worked %s   (%s %v)\n", format.Duration(week.TotalWorkDuration()), sign(overtime), format.DurationShort(overtime))
	}
	return nil
}
