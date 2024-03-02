package cmd

import (
	"reflect"
	"testing"
)

func TestAddCmd_simpleTask(t *testing.T) {
	ctx.data.days = nil

	_ = executeCommand(t, "--date 26.2.2024 add --start 12:15 --tags haora simple task")

	d := ctx.data.day(mockDate("26.02.2024 00:00"))
	if len(d.tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(d.tasks))
	}
	task := d.tasks[0]
	expectedStart := mockDate("26.02.2024 12:15")
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

func TestAddCmd_shouldAllowNowAsStartTime(t *testing.T) {
	ctx.data.days = nil
	mockNowAt(t, mockDate("26.02.2024 11:52"))

	_ = executeCommand(t, "--date 26.02.2024 add --start now --tags haora simple task")

	d := ctx.data.day(mockDate("26.02.2024 00:00"))
	if len(d.tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(d.tasks))
	}
	task := d.tasks[0]
	expectedStart := mockDate("26.02.2024 11:52")
	if task.start != expectedStart {
		t.Errorf("expected start time %v, got %v", expectedStart, task.start)
	}
}

func TestAddShouldUpdateExistingTaskAtSameTime(t *testing.T) {
	ctx.data = dayList{
		days: []day{
			{
				date: mockDate("26.02.2024 00:00"),
				tasks: []Task{
					{
						start: mockDate("26.02.2024 12:15"),
						text:  "existing task",
						tags:  []string{"beer"},
					},
				},
			},
		},
	}

	_ = executeCommand(t, "--date 26.02.2024 add --start 12:15 --tags haora simple task")

	d := ctx.data.day(mockDate("26.02.2024 00:00"))
	if len(d.tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(d.tasks))
	}
	task := d.tasks[0]
	if task.text != "simple task" {
		t.Errorf("expected updated task's text to be %q, but got %q", "simple task", task.text)
	}
	if !reflect.DeepEqual(task.tags, []string{"haora"}) {
		t.Errorf("expected updated task's tags to be %v, but got %v", []string{"haora"}, task.tags)
	}
}
