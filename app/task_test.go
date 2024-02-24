package app

import (
	"github.com/google/uuid"
	"slices"
	"testing"
	"time"
)

func TestNewTask(t *testing.T) {
	t.Run("should set a random id", func(t *testing.T) {
		task := NewTask(time.Now(), "a task", false, nil)

		if err := uuid.Validate(task.Id.String()); err != nil {
			t.Errorf("expected task id to be a valid UUID, but got %q", task.Id)
		}
	})
	t.Run("should ensure that tasks starting time has its seconds and nanoseconds truncated", func(t *testing.T) {
		task := NewTask(time.Now(), "a task", false, nil)

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
		*NewTask(MockTime(10, 0), "Y", false, nil),
		*NewTask(MockTime(9, 0), "Z", false, nil),
		*NewTask(MockTime(12, 0), "X", false, nil),
	}

	slices.SortFunc(tasks, tasksByStart)

	if tasks[0].Text != "Z" || tasks[1].Text != "Y" || tasks[2].Text != "X" {
		t.Errorf("expected sorted ordering to be task Z, Y, X, but got %s, %s, %s", tasks[0].Text, tasks[1].Text, tasks[2].Text)
	}
}
