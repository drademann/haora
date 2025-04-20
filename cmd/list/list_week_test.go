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
	"github.com/drademann/fugo/test"
	"github.com/drademann/fugo/test/assert"
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/app/datetime"
	"github.com/drademann/haora/cmd"
	"github.com/drademann/haora/cmd/config"
	"github.com/drademann/haora/cmd/root"
	"testing"
	"time"
)

func TestListWeekCmd_givenNoTasks(t *testing.T) {
	datetime.AssumeForTestNowAt(t, time.Date(2024, time.February, 22, 16, 32, 0, 0, time.Local))
	config.SetDurationPerWeek(t, 40*time.Hour)
	config.SetDaysPerWeek(t, 5)
	config.ApplyConfigOptions(t)

	data.MockLoadSave(t, &data.DayList{})

	out := cmd.TestExecute(t, root.Command, "-d 22.02.2024 list --week")

	//goland:noinspection GrazieInspection
	assert.Output(t, out,
		`
		Mon 19.02.2024   -
		Tue 20.02.2024   -
		Wed 21.02.2024   -
		Thu 22.02.2024   -
		Fri 23.02.2024   -
		Sat 24.02.2024   -
		Sun 25.02.2024   -

		                          total worked      0m   0.00h   0.00h   (- 40h)
		`)
}

func TestListWeekCmd(t *testing.T) {
	datetime.AssumeForTestNowAt(t, test.Date("22.02.2024 16:32"))
	config.SetDurationPerWeek(t, 40*time.Hour)
	config.SetDaysPerWeek(t, 5)
	config.ApplyConfigOptions(t)

	d := data.Day{Date: test.Date("22.02.2024 00:00")}
	d.AddTask(data.NewTask(test.Date("22.02.2024 09:00"), "task 1"))
	d.AddTask(data.NewPause(test.Date("22.02.2024 12:00"), "lunch"))
	d.AddTask(data.NewTask(test.Date("22.02.2024 12:45"), "task 2"))
	d.Finished = test.Date("22.02.2024 17:00")
	data.MockLoadSave(t, &data.DayList{Days: []*data.Day{&d}})

	out := cmd.TestExecute(t, root.Command, "list --week")

	//goland:noinspection GrazieInspection
	assert.Output(t, out,
		`
		Mon 19.02.2024   -
		Tue 20.02.2024   -
		Wed 21.02.2024   -
		Thu 22.02.2024   09:00 - 17:00  worked  7h 15m   7.25h   7.25h   (- 45m)
		Fri 23.02.2024   -
		Sat 24.02.2024   -
		Sun 25.02.2024   -
		
		                          total worked  7h 15m   7.25h   7.25h   (- 32h 45m)
		`)
}

func TestListWeekCmd_shouldMarkVacationDays(t *testing.T) {
	datetime.AssumeForTestNowAt(t, test.Date("22.02.2024 16:32"))
	config.SetDurationPerWeek(t, 40*time.Hour)
	config.SetDaysPerWeek(t, 5)
	config.ApplyConfigOptions(t)

	d1 := data.Day{Date: test.Date("22.02.2024 00:00")}
	d1.AddTask(data.NewTask(test.Date("22.02.2024 09:00"), "task 1"))
	d1.AddTask(data.NewPause(test.Date("22.02.2024 12:00"), "lunch"))
	d1.AddTask(data.NewTask(test.Date("22.02.2024 12:45"), "task 2"))
	d1.Finished = test.Date("22.02.2024 17:00")

	d2 := data.Day{Date: test.Date("23.02.2024 00:00")}
	d2.IsVacation = true

	data.MockLoadSave(t, &data.DayList{Days: []*data.Day{&d1, &d2}})

	out := cmd.TestExecute(t, root.Command, "list --week")

	//goland:noinspection GrazieInspection
	assert.Output(t, out,
		`
		Mon 19.02.2024   -
		Tue 20.02.2024   -
		Wed 21.02.2024   -
		Thu 22.02.2024   09:00 - 17:00  worked  7h 15m   7.25h   7.25h   (- 45m)
		Fri 23.02.2024   vacation
		Sat 24.02.2024   -
		Sun 25.02.2024   -
		
		                          total worked  7h 15m   7.25h   7.25h   (- 24h 45m)
		`)
}

func TestListWeekCmd_givenDayIsOpen_shouldDisplayNowAsEndTime(t *testing.T) {
	datetime.AssumeForTestNowAt(t, test.Date("22.02.2024 16:32"))
	config.SetDurationPerWeek(t, 40*time.Hour)
	config.SetDaysPerWeek(t, 5)
	config.ApplyConfigOptions(t)

	d := data.Day{Date: test.Date("22.02.2024 00:00")}
	d.AddTask(data.NewTask(test.Date("22.02.2024 09:00"), "task 1"))
	d.AddTask(data.NewPause(test.Date("22.02.2024 12:00"), "lunch"))
	d.AddTask(data.NewTask(test.Date("22.02.2024 12:45"), "task 2"))
	data.MockLoadSave(t, &data.DayList{Days: []*data.Day{&d}})

	out := cmd.TestExecute(t, root.Command, "list --week")

	//goland:noinspection GrazieInspection
	assert.Output(t, out,
		`
		Mon 19.02.2024   -
		Tue 20.02.2024   -
		Wed 21.02.2024   -
		Thu 22.02.2024   09:00 -  now   worked  6h 47m   6.78h   6.75h   (-  1h 13m)
		Fri 23.02.2024   -
		Sat 24.02.2024   -
		Sun 25.02.2024   -
		
		                          total worked  6h 47m   6.78h   6.75h   (- 33h 13m)
		`)
}

func TestListWeekCmd_givenTodayIsMonday_shouldStartOneWeekBack(t *testing.T) {
	datetime.AssumeForTestNowAt(t, test.Date("18.03.2024 16:32"))
	config.SetDurationPerWeek(t, 40*time.Hour)
	config.SetDaysPerWeek(t, 5)
	config.ApplyConfigOptions(t)

	d := data.Day{Date: test.Date("22.02.2024 00:00")}
	data.MockLoadSave(t, &data.DayList{Days: []*data.Day{&d}})

	out := cmd.TestExecute(t, root.Command, "-d mo list --week")

	//goland:noinspection GrazieInspection
	assert.Output(t, out,
		`
		Mon 11.03.2024   -
		Tue 12.03.2024   -
		Wed 13.03.2024   -
		Thu 14.03.2024   -
		Fri 15.03.2024   -
		Sat 16.03.2024   -
		Sun 17.03.2024   -
		
		                          total worked      0m   0.00h   0.00h   (- 40h)
		`)
}

func TestListWeekCmd_withTotalDuration(t *testing.T) {
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

	data.MockLoadSave(t, &data.DayList{Days: []*data.Day{&d1, &d2}})

	flagCases := []string{"--week", "-w"}
	for _, fc := range flagCases {
		command := fmt.Sprintf("list %s", fc)
		t.Run(command, func(t *testing.T) {
			datetime.AssumeForTestNowAt(t, test.Date("22.02.2024 16:32"))
			config.SetDurationPerWeek(t, 40*time.Hour)
			config.SetDaysPerWeek(t, 5)
			config.ApplyConfigOptions(t)

			out := cmd.TestExecute(t, root.Command, command)

			//goland:noinspection GrazieInspection
			assert.Output(t, out,
				`
				Mon 19.02.2024   -
				Tue 20.02.2024   -
				Wed 21.02.2024   -
				Thu 22.02.2024   09:00 - 17:00  worked  7h 15m   7.25h   7.25h   (- 45m)
				Fri 23.02.2024   10:30 - 15:00  worked  4h 15m   4.25h   4.25h   (-  3h 45m)
				Sat 24.02.2024   -
				Sun 25.02.2024   -
				
				                          total worked 11h 30m  11.50h  11.50h   (- 28h 30m)
				`)
		})
	}
}

func TestListWeekCmd_withTotalDuration_andWeeksFinishTime(t *testing.T) {
	d1 := data.Day{Date: test.Date("21.02.2024 00:00")}
	d1.AddTask(data.NewTask(test.Date("21.02.2024 09:00"), "task 1"))
	d1.AddTask(data.NewPause(test.Date("21.02.2024 12:00"), "lunch"))
	d1.AddTask(data.NewTask(test.Date("21.02.2024 12:45"), "task 2"))
	d1.Finished = test.Date("21.02.2024 17:00")

	d2 := data.Day{Date: test.Date("22.02.2024 00:00")}
	d2.AddTask(data.NewTask(test.Date("22.02.2024 10:30"), "task a"))
	d2.AddTask(data.NewPause(test.Date("22.02.2024 12:15"), "some bread"))
	d2.AddTask(data.NewTask(test.Date("22.02.2024 12:30"), "task b"))
	d2.Finished = test.Date("22.02.2024 15:00")

	d3 := data.Day{Date: test.Date("23.02.2024 00:00")}
	d3.AddTask(data.NewTask(test.Date("23.02.2024 09:00"), "task c"))

	data.MockLoadSave(t, &data.DayList{Days: []*data.Day{&d1, &d2, &d3}})

	flagCases := []string{"--week", "-w"}
	for _, fc := range flagCases {
		command := fmt.Sprintf("list %s", fc)
		t.Run(command, func(t *testing.T) {
			datetime.AssumeForTestNowAt(t, test.Date("23.02.2024 10:32"))
			config.SetDurationPerWeek(t, 15*time.Hour)
			config.SetDaysPerWeek(t, 3)
			config.SetDefaultPause(t, 10*time.Minute)
			config.ApplyConfigOptions(t)

			out := cmd.TestExecute(t, root.Command, command)

			//goland:noinspection GrazieInspection
			assert.Output(t, out,
				`
				Mon 19.02.2024   -
				Tue 20.02.2024   -
				Wed 21.02.2024   09:00 - 17:00  worked  7h 15m   7.25h   7.25h   (+  2h 15m)
				Thu 22.02.2024   10:30 - 15:00  worked  4h 15m   4.25h   4.25h   (- 45m)
				Fri 23.02.2024   09:00 -  now   worked  1h 32m   1.53h   1.50h   (-  3h 38m)
				Sat 24.02.2024   -
				Sun 25.02.2024   -
				
				                          total worked 13h  2m  13.03h  13.00h   (-  2h  8m to 12:40)
				`)
		})
	}
}

func TestListWeekCmd_withHiddenWeekdays(t *testing.T) {
	datetime.AssumeForTestNowAt(t, test.Date("20.03.2024 16:32"))
	config.SetDurationPerWeek(t, 40*time.Hour)
	config.SetDaysPerWeek(t, 5)
	config.SetHiddenWeekdays(t, "sat sun")
	config.ApplyConfigOptions(t)

	d := data.Day{Date: test.Date("22.03.2024 00:00")}
	data.MockLoadSave(t, &data.DayList{Days: []*data.Day{&d}})

	out := cmd.TestExecute(t, root.Command, "-d mo list --week")

	//goland:noinspection GrazieInspection
	assert.Output(t, out,
		`
		Mon 18.03.2024   -
		Tue 19.03.2024   -
		Wed 20.03.2024   -
		Thu 21.03.2024   -
		Fri 22.03.2024   -
		
		                          total worked      0m   0.00h   0.00h   (- 40h)
		`)
}

func TestListWeekCmd_withHiddenWeekdays_showsHiddenWeekdaysWhenNotEmpty(t *testing.T) {
	datetime.AssumeForTestNowAt(t, test.Date("20.03.2024 16:32"))
	config.SetDurationPerWeek(t, 40*time.Hour)
	config.SetDaysPerWeek(t, 5)
	config.SetHiddenWeekdays(t, "mon sat sun")
	config.ApplyConfigOptions(t)

	d := data.Day{Date: test.Date("18.03.2024 00:00")}
	d.AddTask(data.NewTask(test.Date("18.03.2024 09:00"), "task 1"))
	d.AddTask(data.NewPause(test.Date("18.03.2024 12:00"), "lunch"))
	d.AddTask(data.NewTask(test.Date("18.03.2024 12:45"), "task 2"))
	d.Finished = test.Date("18.03.2024 17:00")
	data.MockLoadSave(t, &data.DayList{Days: []*data.Day{&d}})

	out := cmd.TestExecute(t, root.Command, "-d mo list --week")

	//goland:noinspection GrazieInspection
	assert.Output(t, out,
		`
		Mon 18.03.2024   09:00 - 17:00  worked  7h 15m   7.25h   7.25h   (- 45m)
		Tue 19.03.2024   -
		Wed 20.03.2024   -
		Thu 21.03.2024   -
		Fri 22.03.2024   -
		
		                          total worked  7h 15m   7.25h   7.25h   (- 32h 45m)
		`)
}
