package cmd

import (
	"github.com/google/uuid"
	"slices"
	"testing"
	"time"
)

func TestNewTask(t *testing.T) {
	t.Run("should set a random id", func(t *testing.T) {
		task := newTask(time.Now(), "a task", nil)

		if err := uuid.Validate(task.id.String()); err != nil {
			t.Errorf("expected task id to be a valid UUID, but got %q", task.id)
		}
	})
	t.Run("should use working date with given hour and minute applied", func(t *testing.T) {
		ctx.workingDate = mockDate(2024, time.February, 25, 0, 0)

		task := newTask(mockTime(10, 30), "a task", nil)

		got := task.start.Format("02.01.2006 15:04")
		want := "25.02.2024 10:30"
		if got != want {
			t.Errorf("expected task start time to be %q, but got %q", want, got)
		}
	})
	t.Run("should ensure that tasks starting time has its seconds and nanoseconds truncated", func(t *testing.T) {
		task := newTask(time.Now(), "a task", nil)

		if task.start.Second() != 0 {
			t.Errorf("expected task start time seconds to be 0, but got %d", task.start.Second())
		}
		if task.start.Nanosecond() != 0 {
			t.Errorf("expected task start time nanoseconds to be 0, but got %d", task.start.Nanosecond())
		}
	})
}

func TestTasksByStart(t *testing.T) {
	tasks := []Task{
		newTask(mockTime(10, 0), "Y", nil),
		newTask(mockTime(9, 0), "Z", nil),
		newTask(mockTime(12, 0), "X", nil),
	}

	slices.SortFunc(tasks, tasksByStart)

	if tasks[0].text != "Z" || tasks[1].text != "Y" || tasks[2].text != "X" {
		t.Errorf("expected sorted ordering to be task Z, Y, X, but got %s, %s, %s", tasks[0].text, tasks[1].text, tasks[2].text)
	}
}
