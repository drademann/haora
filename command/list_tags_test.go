package command

import (
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/app/datetime"
	"github.com/drademann/haora/test"
	"testing"
	"time"
)

func TestListTagsCmd_givenNoTasks(t *testing.T) {
	datetime.MockNowAt(t, time.Date(2024, time.February, 22, 16, 32, 0, 0, time.Local))

	t.Run("no days and thus no tasks for today", func(t *testing.T) {
		data.State.DayList.Days = nil

		out := test.ExecuteCommand(t, Root, "-d 22.02.2024 list --tags")

		test.AssertOutput(t, out,
			`
			Tag summary for today, 22.02.2024 (Thu)

			no tags found
			`)
	})
	t.Run("no tasks for other day than today", func(t *testing.T) {
		data.State.DayList.Days = nil

		out := test.ExecuteCommand(t, Root, "-d 20.02.2024 list --tags")

		test.AssertOutput(t, out,
			`
			Tag summary for 20.02.2024 (Tue)
			
			no tags found
			`)
	})
}

func TestListTagsCmd(t *testing.T) {
	d := data.Day{Date: test.MockDate("22.02.2024 00:00")}
	d.AddTasks(
		data.NewTask(test.MockDate("22.02.2024 9:00"), "a task", "haora"),
		data.NewTask(test.MockDate("22.02.2024 12:00"), "a task", "learning"),
		data.NewTask(test.MockDate("22.02.2024 15:00"), "a task", "go", "learning"),
	)
	data.State.DayList = data.DayListType{Days: []data.Day{d}}

	datetime.MockNowAt(t, test.MockDate("22.02.2024 16:32"))

	out := test.ExecuteCommand(t, Root, "list --tags")

	test.AssertOutput(t, out,
		`
		Tag summary for today, 22.02.2024 (Thu)

		 1h 32m   1.53h   1.50h   #go
		 3h  0m   3.00h   3.00h   #haora
		 4h 32m   4.53h   4.50h   #learning
		`)
}
