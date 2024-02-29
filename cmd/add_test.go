package cmd

import (
	"bytes"
	"reflect"
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

	day := ctx.data.day(mockDate(2024, time.February, 26, 0, 0))
	if len(day.Tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(day.Tasks))
	}
	task := day.Tasks[0]
	expectedStart := mockDate(2024, time.February, 26, 12, 15)
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

func TestAddShouldUpdateExistingTaskAtSameTime(t *testing.T) {
	ctx.data = dayList{
		days: []Day{
			{
				Date: mockDate(2024, time.February, 26, 0, 0),
				Tasks: []Task{
					{
						Start: mockDate(2024, time.February, 26, 12, 15),
						Text:  "existing task",
						Tags:  []string{"beer"},
					},
				},
			},
		},
	}

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetArgs([]string{"--date", "26.2.2024", "add", "--start", "12:15", "--tags", "haora", "simple", "task"})

	if err := rootCmd.Execute(); err != nil {
		t.Fatal(err)
	}

	day := ctx.data.day(mockDate(2024, time.February, 26, 0, 0))
	if len(day.Tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(day.Tasks))
	}
	task := day.Tasks[0]
	if task.Text != "simple task" {
		t.Errorf("expected updated task's text to be %q, but got %q", "simple task", task.Text)
	}
	if !reflect.DeepEqual(task.Tags, []string{"haora"}) {
		t.Errorf("expected updated task's tags to be %v, but got %v", []string{"haora"}, task.Tags)
	}
}
