package command

import (
	"github.com/drademann/haora/app"
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/app/datetime"
	"github.com/drademann/haora/test"
	"github.com/drademann/haora/test/assert"
	"testing"
	"time"
)

func TestListWeekCmd_givenNoTasks(t *testing.T) {
	datetime.AssumeForTestNowAt(t, time.Date(2024, time.February, 22, 16, 32, 0, 0, time.Local))
	app.Config.Times.DurationPerWeek = "40h"
	app.Config.Times.DaysPerWeek = 5

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
	d.AddTasks(
		data.NewTask(test.Date("22.02.2024 09:00"), "task 1"),
		data.NewPause(test.Date("22.02.2024 12:00"), "lunch"),
		data.NewTask(test.Date("22.02.2024 12:45"), "task 2"),
	)
	d.Finished = test.Date("22.02.2024 17:00")
	data.State.DayList = data.DayListType{Days: []data.Day{d}}

	out := test.ExecuteCommand(t, Root, "list --week")

	assert.Output(t, out,
		`
		Mon 19.02.2024   -
		Tue 20.02.2024   -
		Wed 21.02.2024   -
		Thu 22.02.2024   09:00 - 17:00  worked  7h 15m
		Fri 23.02.2024   -
		Sat 24.02.2024   -
		Sun 25.02.2024   -
		
		                          total worked  7h 15m   (- 32h 45m)
		`)
}

func TestListWeekCmd_withTotalDuration(t *testing.T) {
	datetime.AssumeForTestNowAt(t, test.Date("22.02.2024 16:32"))

	d1 := data.Day{Date: test.Date("22.02.2024 00:00")}
	d1.AddTasks(
		data.NewTask(test.Date("22.02.2024 09:00"), "task 1"),
		data.NewPause(test.Date("22.02.2024 12:00"), "lunch"),
		data.NewTask(test.Date("22.02.2024 12:45"), "task 2"),
	)
	d1.Finished = test.Date("22.02.2024 17:00")
	d2 := data.Day{Date: test.Date("23.02.2024 00:00")}
	d2.AddTasks(
		data.NewTask(test.Date("23.02.2024 10:30"), "task a"),
		data.NewPause(test.Date("23.02.2024 12:15"), "some bread"),
		data.NewTask(test.Date("23.02.2024 12:30"), "task b"),
	)
	d2.Finished = test.Date("23.02.2024 15:00")

	data.State.DayList = data.DayListType{Days: []data.Day{d1, d2}}

	out := test.ExecuteCommand(t, Root, "list --week")

	assert.Output(t, out,
		`
		Mon 19.02.2024   -
		Tue 20.02.2024   -
		Wed 21.02.2024   -
		Thu 22.02.2024   09:00 - 17:00  worked  7h 15m
		Fri 23.02.2024   10:30 - 15:00  worked  4h 15m
		Sat 24.02.2024   -
		Sun 25.02.2024   -
		
		                          total worked 11h 30m   (- 28h 30m)
		`)
}
