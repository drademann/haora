package cmd

import (
	"bytes"
	"haora/app"
	"haora/test"
	"testing"
	"time"
)

func TestAddSimpleTask(t *testing.T) {
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetArgs([]string{"--date", "26.2.2024", "add", "--start", "12:15", "--tags", "haora", "simple", "task"})

	if err := rootCmd.Execute(); err != nil {
		t.Fatal(err)
	}

	day := app.Data.Day(test.MockDate(2024, time.February, 26, 0, 0))
	if len(day.Tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(day.Tasks))
	}
	task := day.Tasks[0]
	expectedStart := test.MockDate(2024, time.February, 26, 12, 15)
	if task.Start != expectedStart {
		t.Errorf("expected start time %v, got %v", expectedStart, task.Start)
	}
	if task.Text != "simple task" {
		t.Errorf("expected text 'simple task', got %s", task.Text)
	}
	if len(task.Tags) != 1 || task.Tags[0] != "haora" {
		t.Errorf("expected tags ['haora'], got %v", task.Tags)
	}
}
