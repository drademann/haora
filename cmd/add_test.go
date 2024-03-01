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
	if len(day.tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(day.tasks))
	}
	task := day.tasks[0]
	expectedStart := mockDate(2024, time.February, 26, 12, 15)
	if task.start != expectedStart {
		t.Errorf("expected start time %v, got %v", expectedStart, task.start)
	}
	if task.text != "simple task" {
		t.Errorf("expected text 'simple task', got %s", task.text)
	}
	if len(task.tags) != 1 || task.tags[0] != "haora" {
		t.Errorf("expected tags ['haora'], got %v", task.tags)
	}
}

func TestAddShouldUpdateExistingTaskAtSameTime(t *testing.T) {
	ctx.data = dayList{
		days: []Day{
			{
				date: mockDate(2024, time.February, 26, 0, 0),
				tasks: []Task{
					{
						start: mockDate(2024, time.February, 26, 12, 15),
						text:  "existing task",
						tags:  []string{"beer"},
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
	if len(day.tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(day.tasks))
	}
	task := day.tasks[0]
	if task.text != "simple task" {
		t.Errorf("expected updated task's text to be %q, but got %q", "simple task", task.text)
	}
	if !reflect.DeepEqual(task.tags, []string{"haora"}) {
		t.Errorf("expected updated task's tags to be %v, but got %v", []string{"haora"}, task.tags)
	}
}
