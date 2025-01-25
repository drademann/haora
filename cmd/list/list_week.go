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

package list

import (
	"fmt"
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/app/datetime"
	"github.com/drademann/haora/cmd/config"
	"github.com/drademann/haora/cmd/internal/format"
	"github.com/spf13/cobra"
	"time"
)

func printWeek(cmd *cobra.Command, workingDate time.Time, dayList *data.DayList) error {
	date := datetime.FindWeekday(workingDate, datetime.Previous, time.Monday)
	week := dayList.Week(date)
	for _, day := range week.Days {
		dateStr := day.Date.Format("Mon 02.01.2006")
		if config.IsHidden(day.Date.Weekday()) && day.IsEmpty() {
			continue
		} else if day.IsEmpty() {
			cmd.Printf("%s   -\n", dateStr)
		} else {
			startStr := day.Start().Format("15:04")
			var endStr string
			if day.IsFinished() {
				endStr = day.End().Format("15:04")
			} else {
				endStr = " now "
			}
			dur := day.TotalWorkDuration()
			durStr := format.Duration(dur)
			durDecStr := format.DurationDecimal(dur)
			durDecRoundedStr := format.DurationDecimalRounded(dur, 15*time.Minute)
			overtime, exist := day.OvertimeDuration()
			if !exist || overtime == 0 {
				//goland:noinspection GrazieInspection
				cmd.Printf("%s   %s - %s  worked %s  %s  %s\n", dateStr, startStr, endStr, durStr, durDecStr, durDecRoundedStr)
			} else {
				//goland:noinspection GrazieInspection
				cmd.Printf("%s   %s - %s  worked %s  %s  %s   (%s %v)\n", dateStr, startStr, endStr, durStr, durDecStr, durDecRoundedStr, sign(overtime), format.DurationShort(overtime))
			}
		}
	}
	totalDur := week.TotalWorkDuration()
	totalDurStr := format.Duration(totalDur)
	totalDurDecStr := format.DurationDecimal(totalDur)
	totalDurDecRoundedStr := format.DurationDecimalRounded(totalDur, 15*time.Minute)
	totalOvertimeDur, exist := week.TotalOvertimeDuration()
	if !exist || totalOvertimeDur == 0 {
		cmd.Printf("\n                          total worked %s  %s  %s\n", totalDurStr, totalDurDecStr, totalDurDecRoundedStr)
	} else {
		cmd.Printf("\n                          total worked %s  %s  %s   (%s)\n", totalDurStr, totalDurDecStr, totalDurDecRoundedStr, formattedWeekOvertime(totalOvertimeDur))
	}
	return nil
}

func formattedWeekOvertime(overtime time.Duration) string {
	s := sign(overtime)
	d := format.DurationShort(overtime)
	now := datetime.Now()
	weekFinish := now.Add(-overtime)
	if overtime >= 0 || !data.IsSameDay(now, weekFinish) {
		return fmt.Sprintf("%s %v", s, d)
	}

	return fmt.Sprintf("%s %v to %v", s, d, weekFinish.Format("15:04"))
}
