package cmd

import (
	"bytes"
	"haora/app"
	"testing"
	"time"
)

func TestExecListCmd_givenNoDays(t *testing.T) {
	out := bytes.Buffer{}
	rootCmd.SetOut(&out)
	rootCmd.SetArgs([]string{"list"})
	*workingDateFlag = ""

	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "no tasks recorded for today\n"
	if out.String() != expected {
		t.Errorf("expected output %q, but got %q", expected, out.String())
	}
}

func TestExecListCmd_givenTasksForToday(t *testing.T) {
	out := bytes.Buffer{}
	rootCmd.SetOut(&out)
	rootCmd.SetArgs([]string{"list"})
	realNow := app.Now
	defer func() { app.Now = realNow }()
	*workingDateFlag = ""

	app.Now = func() time.Time {
		return time.Date(2024, time.February, 22, 16, 32, 0, 0, time.Local)
	}
	app.Data = app.DayList{
		{Date: time.Date(2024, time.February, 22, 6, 24, 13, 0, time.Local),
			Tasks: []app.Task{
				{Start: time.Date(2024, time.February, 22, 9, 0, 0, 0, time.Local),
					Text:    "a task",
					IsPause: false,
					Tags:    []string{}},
			},
			Finished: time.Time{}},
	}

	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "09:00 - ... a task\n"
	if out.String() != expected {
		t.Errorf("expected output %q, but got %q", expected, out.String())
	}
}
