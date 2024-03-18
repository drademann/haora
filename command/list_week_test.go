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

package command

import (
	"github.com/drademann/haora/app/config"
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/app/datetime"
	"github.com/drademann/haora/test"
	"github.com/drademann/haora/test/assert"
	"testing"
	"time"
)

func TestListWeekCmd_givenNoTasks(t *testing.T) {
	datetime.AssumeForTestNowAt(t, time.Date(2024, time.February, 22, 16, 32, 0, 0, time.Local))
	config.SetDurationPerWeek(t, 40*time.Hour)
	config.SetDaysPerWeek(t, 5)

	data.State.DayList.Days = nil

	out := test.ExecuteCommand(t, Root, "-d 22.02.2024 list --week")

	assert.Output(t, out,
		`
		Mon 19.02.2024   -
		Tue 20.02.2024   -
		Wed 21.02.2024   -
		Thu 22.02.2024   -
		Fri 23.02.2024   -
		Sat 24.02.2024   -
		Sun 25.02.2024   -

		                          total worked      0m   (- 40h)
		`)
}

func TestListWeekCmd(t *testing.T) {
	datetime.AssumeForTestNowAt(t, test.Date("22.02.2024 16:32"))

	d := data.Day{Date: test.Date("22.02.2024 00:00")}
	d.AddTask(data.NewTask(test.Date("22.02.2024 09:00"), "task 1"))
	d.AddTask(data.NewPause(test.Date("22.02.2024 12:00"), "lunch"))
	d.AddTask(data.NewTask(test.Date("22.02.2024 12:45"), "task 2"))
	d.Finished = test.Date("22.02.2024 17:00")
	data.State.DayList = &data.DayListType{Days: []*data.Day{&d}}

	out := test.ExecuteCommand(t, Root, "list --week")

	assert.Output(t, out,
		`
		Mon 19.02.2024   -
		Tue 20.02.2024   -
		Wed 21.02.2024   -
		Thu 22.02.2024   09:00 - 17:00  worked  7h 15m   (- 45m)
		Fri 23.02.2024   -
		Sat 24.02.2024   -
		Sun 25.02.2024   -
		
		                          total worked  7h 15m   (- 32h 45m)
		`)
}

func TestListWeekCmd_givenTodayIsMonday_shouldStartOneWeekBack(t *testing.T) {
	datetime.AssumeForTestNowAt(t, test.Date("18.03.2024 16:32"))

	d := data.Day{Date: test.Date("22.02.2024 00:00")}
	data.State.DayList = &data.DayListType{Days: []*data.Day{&d}}

	out := test.ExecuteCommand(t, Root, "-d mo list --week")

	assert.Output(t, out,
		`
		Mon 11.03.2024   -
		Tue 12.03.2024   -
		Wed 13.03.2024   -
		Thu 14.03.2024   -
		Fri 15.03.2024   -
		Sat 16.03.2024   -
		Sun 17.03.2024   -
		
		                          total worked      0m   (- 40h)
		`)
}

func TestListWeekCmd_withTotalDuration(t *testing.T) {
	datetime.AssumeForTestNowAt(t, test.Date("22.02.2024 16:32"))

	d1 := data.Day{Date: test.Date("22.02.2024 00:00")}
	d1.AddTask(data.NewTask(test.Date("22.02.2024 09:00"), "task 1"))
	d1.AddTask(data.NewPause(test.Date("22.02.2024 12:00"), "lunch"))
	d1.AddTask(data.NewTask(test.Date("22.02.2024 12:45"), "task 2"))
	d1.Finished = test.Date("22.02.2024 17:00")

	d2 := data.Day{Date: test.Date("23.02.2024 00:00")}
	d2.AddTask(data.NewTask(test.Date("23.02.2024 10:30"), "task a"))
	d2.AddTask(data.NewPause(test.Date("23.02.2024 12:15"), "some bread"))
	d2.AddTask(data.NewTask(test.Date("23.02.2024 12:30"), "task b"))
	d2.Finished = test.Date("23.02.2024 15:00")

	data.State.DayList = &data.DayListType{Days: []*data.Day{&d1, &d2}}

	out := test.ExecuteCommand(t, Root, "list --week")

	assert.Output(t, out,
		`
		Mon 19.02.2024   -
		Tue 20.02.2024   -
		Wed 21.02.2024   -
		Thu 22.02.2024   09:00 - 17:00  worked  7h 15m   (- 45m)
		Fri 23.02.2024   10:30 - 15:00  worked  4h 15m   (-  3h 45m)
		Sat 24.02.2024   -
		Sun 25.02.2024   -
		
		                          total worked 11h 30m   (- 28h 30m)
		`)
}
