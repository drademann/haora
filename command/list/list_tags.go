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
	"fmt"
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/command/internal/format"
	"github.com/spf13/cobra"
	"time"
)

func printTags(cmd *cobra.Command, workingDate time.Time, dayList *data.DayList) error {
	d := dayList.Day(workingDate)

	headerStr := func(day data.Day) string {
		ds := day.Date.Format("02.01.2006 (Mon)")
		if day.IsToday() {
			return fmt.Sprintf("Tag summary for today, %s\n", ds)
		}
		return fmt.Sprintf("Tag summary for %s\n", ds)
	}
	cmd.Println(headerStr(*d))

	if d.IsEmpty() {
		cmd.Println("no tags found")
		return nil
	}

	for _, tg := range d.Tags() {
		tagDur := d.TotalTagDuration(tg)
		cmd.Printf("%v  %v  %v   #%v\n",
			format.Duration(tagDur),
			format.DurationDecimal(tagDur),
			format.DurationDecimalRounded(tagDur, 15*time.Minute),
			tg,
		)
	}

	return nil
}
