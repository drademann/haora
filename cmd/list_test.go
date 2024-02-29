package cmd

import (
	"bytes"
	"testing"
	"time"
)

func TestExecListCmd_givenNoTasks(t *testing.T) {
	now = func() time.Time {
		return time.Date(2024, time.February, 22, 16, 32, 0, 0, time.Local)
	}
	defer func() { now = time.Now }()

	t.Run("no days and thus no tasks for today", func(t *testing.T) {
		ctx.data.days = nil

		out := new(bytes.Buffer)
		rootCmd.SetOut(out)
		rootCmd.SetArgs([]string{"-d", "22.02.2024", "list"})

		if err := rootCmd.Execute(); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		assertOutput(t, out,
			`
			Tasks for today, 22.02.2024 (Thu)

			no tasks recorded
			`)
	})
	t.Run("no tasks for other day than today", func(t *testing.T) {
		ctx.data.days = nil

		out := new(bytes.Buffer)
		rootCmd.SetOut(out)
		rootCmd.SetArgs([]string{"-d", "20.02.2024", "list"})

		if err := rootCmd.Execute(); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		assertOutput(t, out,
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
	*workingDateFlag = ""

	mockNowAt(t, time.Date(2024, time.February, 22, 16, 32, 0, 0, time.Local))

	ctx.data = dayList{
		days: []Day{
			{Date: time.Date(2024, time.February, 22, 6, 24, 13, 0, time.Local),
				Tasks: []Task{
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

	assertOutput(t, out,
		`
		Tasks for today, 22.02.2024 (Thu)

		09:00 - ... a task
		`)
}
