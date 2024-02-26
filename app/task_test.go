package app

import (
	"github.com/google/uuid"
	"haora/test"
	"slices"
	"testing"
	"time"
)

func TestNewTask(t *testing.T) {
	t.Run("should set a random id", func(t *testing.T) {
		task := NewTask(time.Now(), "a task", nil)

		if err := uuid.Validate(task.Id.String()); err != nil {
			t.Errorf("expected task id to be a valid UUID, but got %q", task.Id)
		}
	})
	t.Run("should use working date with given hour and minute applied", func(t *testing.T) {
		WorkingDate = test.MockDate(2024, time.February, 25, 0, 0)

		task := NewTask(test.MockTime(10, 30), "a task", nil)

		got := task.Start.Format("02.01.2006 15:04")
		want := "25.02.2024 10:30"
		if got != want {
			t.Errorf("expected task start time to be %q, but got %q", want, got)
		}
	})
	t.Run("should ensure that tasks starting time has its seconds and nanoseconds truncated", func(t *testing.T) {
		task := NewTask(time.Now(), "a task", nil)

		if task.Start.Second() != 0 {
			t.Errorf("expected task start time seconds to be 0, but got %d", task.Start.Second())
		}
		if task.Start.Nanosecond() != 0 {
			t.Errorf("expected task start time nanoseconds to be 0, but got %d", task.Start.Nanosecond())
		}
	})
}

func TestTasksByStart(t *testing.T) {
	tasks := []Task{
		NewTask(test.MockTime(10, 0), "Y", nil),
		NewTask(test.MockTime(9, 0), "Z", nil),
		NewTask(test.MockTime(12, 0), "X", nil),
	}

	slices.SortFunc(tasks, tasksByStart)

	if tasks[0].Text != "Z" || tasks[1].Text != "Y" || tasks[2].Text != "X" {
		t.Errorf("expected sorted ordering to be task Z, Y, X, but got %s, %s, %s", tasks[0].Text, tasks[1].Text, tasks[2].Text)
	}
}
