package add

import (
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/command"
	"github.com/drademann/haora/test"
	"reflect"
	"testing"
)

func TestAddCmd_simpleTask(t *testing.T) {
	data.State.DayList.Days = nil

	_ = test.ExecuteCommand(t, command.Root, "--date 26.2.2024 add --start 12:15 --tags haora simple task")

	d := data.State.DayList.Day(test.MockDate("26.02.2024 00:00"))
	if len(d.Tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(d.Tasks))
	}
	task := d.Tasks[0]
	expectedStart := test.MockDate("26.02.2024 12:15")
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

func TestAddCmd_shouldAllowNowAsStartTime(t *testing.T) {
	data.State.DayList.Days = nil
	test.MockNowAt(t, test.MockDate("26.02.2024 11:52"))

	_ = test.ExecuteCommand(t, command.Root, "--date 26.02.2024 add --start now --tags haora simple task")

	d := data.State.DayList.Day(test.MockDate("26.02.2024 00:00"))
	if len(d.Tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(d.Tasks))
	}
	task := d.Tasks[0]
	expectedStart := test.MockDate("26.02.2024 11:52")
	if task.Start != expectedStart {
		t.Errorf("expected start time %v, got %v", expectedStart, task.Start)
	}
}

func TestAddShouldUpdateExistingTaskAtSameTime(t *testing.T) {
	data.State.DayList = data.DayListType{
		Days: []data.Day{
			{
				Date: test.MockDate("26.02.2024 00:00"),
				Tasks: []data.Task{
					{
						Start: test.MockDate("26.02.2024 12:15"),
						Text:  "existing task",
						Tags:  []string{"beer"},
					},
				},
			},
		},
	}

	_ = test.ExecuteCommand(t, command.Root, "--date 26.02.2024 add --start 12:15 --tags haora simple task")

	d := data.State.DayList.Day(test.MockDate("26.02.2024 00:00"))
	if len(d.Tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(d.Tasks))
	}
	task := d.Tasks[0]
	if task.Text != "simple task" {
		t.Errorf("expected updated task's text to be %q, but got %q", "simple task", task.Text)
	}
	if !reflect.DeepEqual(task.Tags, []string{"haora"}) {
		t.Errorf("expected updated task's tags to be %v, but got %v", []string{"haora"}, task.Tags)
	}
}
