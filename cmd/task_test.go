package cmd

import (
	"github.com/google/uuid"
	"slices"
	"testing"
	"time"
)

func TestNewTask(t *testing.T) {
	t.Run("should set a random id", func(t *testing.T) {
		task := newTask(time.Now(), "a task")

		if err := uuid.Validate(task.id.String()); err != nil {
			t.Errorf("expected task id to be a valid UUID, but got %q", task.id)
		}
	})
	t.Run("should ensure that tasks starting time has its seconds and nanoseconds truncated", func(t *testing.T) {
		task := newTask(time.Now(), "a task")

		if task.start.Second() != 0 {
			t.Errorf("expected task start time seconds to be 0, but got %d", task.start.Second())
		}
		if task.start.Nanosecond() != 0 {
			t.Errorf("expected task start time nanoseconds to be 0, but got %d", task.start.Nanosecond())
		}
	})
}

func TestTasksByStart(t *testing.T) {
	tasks := []task{
		newTask(mockTime("10:00"), "Y"),
		newTask(mockTime("9:00"), "Z"),
		newTask(mockTime("12:00"), "X"),
	}

	slices.SortFunc(tasks, tasksByStart)

	if tasks[0].text != "Z" || tasks[1].text != "Y" || tasks[2].text != "X" {
		t.Errorf("expected sorted ordering to be task Z, Y, X, but got %s, %s, %s", tasks[0].text, tasks[1].text, tasks[2].text)
	}
}
