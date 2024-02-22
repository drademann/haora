package cmd_test

import (
	"bytes"
	"haora/app"
	"testing"
	"time"

	"haora/cmd"
)

func TestExecListCmd_givenNoDays(t *testing.T) {
	out := bytes.Buffer{}
	cmd.Config = cmd.Configuration{Out: &out, OutErr: &bytes.Buffer{}}

	err := cmd.ExecListCmd()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "no tasks recorded for today\n"
	if out.String() != expected {
		t.Errorf("expected output %q, but got %q", expected, out.String())
	}
}

func TestExecListCmd_givenTasksForToday(t *testing.T) {
	app.Now = func() time.Time {
		return time.Date(2024, time.February, 22, 16, 32, 0, 0, time.Local)
	}
	out := bytes.Buffer{}
	cmd.Config = cmd.Configuration{Out: &out, OutErr: &bytes.Buffer{}}

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

	err := cmd.ExecListCmd()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "09:00 - ... a task\n"
	if out.String() != expected {
		t.Errorf("expected output %q, but got %q", expected, out.String())
	}
}
