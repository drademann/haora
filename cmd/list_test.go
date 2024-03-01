package cmd

import (
	"testing"
	"time"
)

func TestExecListCmd_givenNoTasks(t *testing.T) {
	mockNowAt(t, time.Date(2024, time.February, 22, 16, 32, 0, 0, time.Local))

	t.Run("no days and thus no tasks for today", func(t *testing.T) {
		ctx.data.days = nil

		out := executeCommand(t, "-d 22.02.2024 list")

		assertOutput(t, out,
			`
			Tasks for today, 22.02.2024 (Thu)

			no tasks recorded
			`)
	})
	t.Run("no tasks for other day than today", func(t *testing.T) {
		ctx.data.days = nil

		out := executeCommand(t, "-d 20.02.2024 list")

		assertOutput(t, out,
			`
			Tasks for 20.02.2024 (Tue)
			
			no tasks recorded
			`)
	})
}

func TestExecListCmd_oneOpenTaskForToday(t *testing.T) {
	*workingDateFlag = ""
	mockNowAt(t, time.Date(2024, time.February, 22, 16, 32, 0, 0, time.Local))

	ctx.data = dayList{
		days: []Day{
			{date: time.Date(2024, time.February, 22, 6, 24, 13, 0, time.Local),
				tasks: []Task{
					{start: time.Date(2024, time.February, 22, 9, 0, 0, 0, time.Local),
						text:    "a task",
						isPause: false,
						tags:    []string{"Haora"}},
				},
				finished: time.Time{}}},
	}

	out := executeCommand(t, "list")

	assertOutput(t, out,
		`
		Tasks for today, 22.02.2024 (Thu)

		09:00 - now      7h 32m   Haora   a task
		          total  7h 32m
		         breaks      0m
		         worked  7h 32m
		       on Haora  7h 32m
		`)
}
