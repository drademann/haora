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
	"errors"
	"fmt"
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/cmd/config"
	"github.com/drademann/haora/cmd/internal/format"
	"github.com/spf13/cobra"
	"strings"
	"time"
)

func printDefault(cmd *cobra.Command, workingDate time.Time, dayList *data.DayList) error {
	d := dayList.Day(workingDate)

	headerStr := func(day data.Day) string {
		ds := day.Date.Format("02.01.2006 (Mon)")
		if day.IsToday() {
			return fmt.Sprintf("Tasks for today, %s\n", ds)
		}
		return fmt.Sprintf("Tasks for %s\n", ds)
	}
	cmd.Println(headerStr(*d))

	if d.IsVacation {
		label, isSet := config.LabelVacation()
		if !isSet {
			label = "vacation"
		}
		cmd.Println(label)
		return nil
	}
	if d.IsEmpty() {
		cmd.Println("no tasks recorded")
		return nil
	}

	tagsStr := func(tags []string) string {
		hashed := make([]string, len(tags))
		for i, tag := range tags {
			hashed[i] = "#" + tag
		}
		if len(hashed) == 0 {
			return ""
		}
		return " " + strings.Join(hashed, " ")
	}

	for _, task := range d.Tasks {
		start := task.Start.Format("15:04")
		var end string
		succ, err := d.Succ(*task)
		if err == nil {
			end = succ.Start.Format("15:04")
		} else {
			if errors.Is(err, data.NoTaskSucc) && d.IsFinished() {
				end = d.Finished.Format("15:04")
			} else {
				end = " now "
			}
		}
		dur := format.Duration(d.TaskDuration(*task))
		if task.IsPause {
			cmd.Printf("      |         %v   %v\n", dur, task.Text)
		} else {
			//goland:noinspection GrazieInspection
			cmd.Printf("%v - %v   %v   %v%v\n", start, end, dur, task.Text, tagsStr(task.Tags))
		}
	}
	cmd.Println()
	cmd.Printf("         total  %v\n", format.Duration(d.TotalDuration()))
	totalPause := d.TotalPauseDuration()
	//goland:noinspection GrazieInspection
	cmd.Printf("        paused  %v\n", format.Duration(totalPause))
	printWorkedOvertime(cmd, d)
	return nil
}

func printWorkedOvertime(cmd *cobra.Command, d *data.Day) {
	output := fmt.Sprintf("        worked  %v", format.Duration(d.TotalWorkDuration()))
	overtime, exist := d.OvertimeDuration()
	if exist && overtime != 0 {
		suggestion, ok := d.SuggestedFinish()
		if ok && overtime < 0 {
			pauseMarker := ""
			if d.UsesDefaultPause() {
				pauseMarker = "*"
			}
			output += fmt.Sprintf("   (%s %v to %v)%s", sign(overtime), format.DurationShort(overtime), suggestion.Format("15:04"), pauseMarker)
		} else {
			output += fmt.Sprintf("   (%s %v)", sign(overtime), format.DurationShort(overtime))
		}
	}
	cmd.Println(output)
}
