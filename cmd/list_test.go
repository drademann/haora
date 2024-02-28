package cmd

import (
	"bytes"
	"haora/app"
	"haora/test"
	"testing"
	"time"
)

func TestExecListCmd_givenNoTasks(t *testing.T) {
	realNow := app.Now
	defer func() { app.Now = realNow }()
	app.Now = func() time.Time {
		return time.Date(2024, time.February, 22, 16, 32, 0, 0, time.Local)
	}

	t.Run("no days and thus no tasks for today", func(t *testing.T) {
		app.Data.Days = nil

		out := new(bytes.Buffer)
		rootCmd.SetOut(out)
		rootCmd.SetArgs([]string{"-d", "22.02.2024", "list"})

		if err := rootCmd.Execute(); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		test.AssertOutput(t, out,
			`
			Tasks for today, 22.02.2024 (Thu)

			no tasks recorded
			`)
	})
	t.Run("no tasks for other day than today", func(t *testing.T) {
		app.Data.Days = nil

		out := new(bytes.Buffer)
		rootCmd.SetOut(out)
		rootCmd.SetArgs([]string{"-d", "20.02.2024", "list"})

		if err := rootCmd.Execute(); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		test.AssertOutput(t, out,
			`
			Tasks for 20.02.2024 (Tue)
			
			no tasks recorded
			`)
	})
}

func TestExecListCmd_givenTasksForToday(t *testing.T) {
	out := new(bytes.Buffer)
	rootCmd.SetOut(out)
	rootCmd.SetArgs([]string{"list"})
	realNow := app.Now
	defer func() { app.Now = realNow }()
	*workingDateFlag = ""

	app.Now = func() time.Time {
		return time.Date(2024, time.February, 22, 16, 32, 0, 0, time.Local)
	}
	app.Data = app.DayList{
		Days: []app.Day{
			{Date: time.Date(2024, time.February, 22, 6, 24, 13, 0, time.Local),
				Tasks: []app.Task{
					{Start: time.Date(2024, time.February, 22, 9, 0, 0, 0, time.Local),
						Text:    "a task",
						IsPause: false,
						Tags:    []string{}},
				},
				Finished: time.Time{}}},
	}

	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	test.AssertOutput(t, out,
		`
		Tasks for today, 22.02.2024 (Thu)

		09:00 - ... a task
		`)
}
