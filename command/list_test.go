package command

import (
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/test"
	"testing"
	"time"
)

func TestListCmd_givenNoTasks(t *testing.T) {
	test.MockNowAt(t, time.Date(2024, time.February, 22, 16, 32, 0, 0, time.Local))

	t.Run("no days and thus no tasks for today", func(t *testing.T) {
		data.State.DayList.Days = nil

		out := test.ExecuteCommand(t, Root, "-d 22.02.2024 list")

		test.AssertOutput(t, out,
			`
			Tasks for today, 22.02.2024 (Thu)

			no tasks recorded
			`)
	})
	t.Run("no tasks for other day than today", func(t *testing.T) {
		data.State.DayList.Days = nil

		out := test.ExecuteCommand(t, Root, "-d 20.02.2024 list")

		test.AssertOutput(t, out,
			`
			Tasks for 20.02.2024 (Tue)
			
			no tasks recorded
			`)
	})
}

func TestListCmd_oneOpenTaskForToday(t *testing.T) {
	d := data.Day{Date: test.MockDate("22.02.2024 00:00")}
	d.AddTasks(data.NewTask(test.MockDate("22.02.2024 9:00"), "a task", "Haora"))
	data.State.DayList = data.DayListType{Days: []data.Day{d}}

	test.MockNowAt(t, test.MockDate("22.02.2024 16:32"))

	out := test.ExecuteCommand(t, Root, "list")

	test.AssertOutput(t, out,
		`
		Tasks for today, 22.02.2024 (Thu)

		09:00 -  now     7h 32m   a task #Haora
		
		         total   7h 32m
		        paused       0m
		        worked   7h 32m
		      on Haora   7h 32m
		`)
}

func TestListCmd_multipleTasksLastOpen(t *testing.T) {
	d := data.Day{Date: test.MockDate("22.02.2024 00:00")}
	d.AddTasks(
		data.NewTask(test.MockDate("22.02.2024 9:00"), "some programming", "Haora"),
		data.NewTask(test.MockDate("22.02.2024 10:00"), "fixing bugs"),
	)
	data.State.DayList = data.DayListType{Days: []data.Day{d}}

	test.MockNowAt(t, test.MockDate("22.02.2024 16:32"))

	out := test.ExecuteCommand(t, Root, "list")

	test.AssertOutput(t, out,
		`
		Tasks for today, 22.02.2024 (Thu)

		09:00 - 10:00    1h  0m   some programming #Haora
		10:00 -  now     6h 32m   fixing bugs

		         total   7h 32m
		        paused       0m
		        worked   7h 32m
		      on Haora   1h  0m
		`)
}

func TestListCmd_withPause(t *testing.T) {
	test.MockNowAt(t, test.MockDate("22.02.2024 16:32"))

	d := data.Day{Date: test.MockDate("22.02.2024 00:00")}
	d.AddTasks(
		data.NewTask(test.MockDate("22.02.2024 9:00"), "some programming", "Haora"),
		data.NewPause(test.MockDate("22.02.2024 12:15"), ""),
		data.NewTask(test.MockDate("22.02.2024 13:00"), "fixing bugs"),
	)
	data.State.DayList = data.DayListType{Days: []data.Day{d}}

	out := test.ExecuteCommand(t, Root, "list")

	test.AssertOutput(t, out,
		`
		Tasks for today, 22.02.2024 (Thu)

		09:00 - 12:15    3h 15m   some programming #Haora
		      |             45m   
		13:00 -  now     3h 32m   fixing bugs

		         total   7h 32m
		        paused      45m
		        worked   6h 47m
		      on Haora   3h 15m
		`)
}

func TestListCmd_withFinished(t *testing.T) {
	test.MockNowAt(t, test.MockDate("22.02.2024 16:32"))

	d := data.Day{Date: test.MockDate("22.02.2024 00:00")}
	d.AddTasks(
		data.NewTask(test.MockDate("22.02.2024 9:00"), "some programming", "Haora"),
		data.NewPause(test.MockDate("22.02.2024 12:15"), "lunch"),
		data.NewTask(test.MockDate("22.02.2024 13:00"), "fixing bugs"),
	)
	d.Finished = test.MockDate("22.02.2024 17:00")
	data.State.DayList = data.DayListType{Days: []data.Day{d}}

	out := test.ExecuteCommand(t, Root, "list")

	test.AssertOutput(t, out,
		`
		Tasks for today, 22.02.2024 (Thu)

		09:00 - 12:15    3h 15m   some programming #Haora
		      |             45m   lunch
		13:00 - 17:00    4h  0m   fixing bugs

		         total   8h  0m
		        paused      45m
		        worked   7h 15m
		      on Haora   3h 15m
		`)
}
