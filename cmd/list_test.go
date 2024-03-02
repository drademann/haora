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
	d := day{date: mockDate("22.02.2024 00:00")}
	d.addTasks(newTask(mockDate("22.02.2024 9:00"), "a task", "Haora"))
	ctx.data = dayList{days: []day{d}}

	workingDateFlag = ""
	mockNowAt(t, mockDate("22.02.2024 16:32"))

	out := executeCommand(t, "list")

	assertOutput(t, out,
		`
		Tasks for today, 22.02.2024 (Thu)

		09:00 -  now     7h 32m   a task #Haora
		
		          total  7h 32m
		         breaks      0m
		         worked  7h 32m
		       on Haora  7h 32m
		`)
}

func TestExecListCmd_multipleTasksLastOpen(t *testing.T) {
	d := day{date: mockDate("22.02.2024 00:00")}
	d.addTasks(
		newTask(mockDate("22.02.2024 9:00"), "some programming", "Haora"),
		newTask(mockDate("22.02.2024 10:00"), "fixing bugs"),
	)
	ctx.data = dayList{days: []day{d}}

	workingDateFlag = ""
	mockNowAt(t, mockDate("22.02.2024 16:32"))

	out := executeCommand(t, "list")

	assertOutput(t, out,
		`
		Tasks for today, 22.02.2024 (Thu)

		09:00 - 10:00    1h  0m   some programming #Haora
		10:00 -  now     6h 32m   fixing bugs

		          total  7h 32m
		         breaks      0m
		         worked  7h 32m
		       on Haora  1h  0m
		`)
}
