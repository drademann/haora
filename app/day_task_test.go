package app

import (
	"errors"
	"testing"
	"time"
)

func TestTaskSuccPred(t *testing.T) {
	testDayDate := MockTime(0, 0)
	task1 := *NewTask(MockTime(9, 0), "task 1", false, nil)
	task2 := *NewTask(MockTime(10, 0), "task 2", false, nil)
	task3 := *NewTask(MockTime(12, 0), "task 3", false, nil)
	day := Day{Date: testDayDate,
		Tasks:    []Task{task1, task2, task3},
		Finished: time.Time{},
	}

	t.Run("find successor", func(t *testing.T) {
		s, err := day.succ(task2)

		if err != nil {
			t.Fatal(err)
		}
		if s.Id != task3.Id {
			t.Errorf("expected successor to be %q, but got %q", task3.Text, s.Text)
		}
	})
	t.Run("find no successor", func(t *testing.T) {
		_, err := day.succ(task3)

		if err == nil {
			t.Errorf("expected error, but got nil")
		}
		if !errors.Is(err, NoTaskSucc) {
			t.Errorf("expected error %q, but got %q", NoTaskSucc, err)
		}
	})
	t.Run("find predecessor", func(t *testing.T) {
		p, err := day.pred(task2)

		if err != nil {
			t.Fatal(err)
		}
		if p.Id != task1.Id {
			t.Errorf("expected predecessor to be %q, but got %q", task1.Text, p.Text)
		}
	})
	t.Run("find no predecessor", func(t *testing.T) {
		_, err := day.pred(task1)

		if err == nil {
			t.Errorf("expected error, but got nil")
		}
		if !errors.Is(err, NoTaskPred) {
			t.Errorf("expected error %q, but got %q", NoTaskPred, err)
		}
	})
}

func TestTaskDuration(t *testing.T) {
	testDayDate := MockTime(0, 0)
	task1 := *NewTask(MockTime(9, 0), "task 1", false, nil)
	task2 := *NewTask(MockTime(10, 0), "task 2", false, nil)
	day := Day{Date: testDayDate,
		Tasks:    []Task{task1, task2},
		Finished: time.Time{},
	}

	t.Run("task with a successor", func(t *testing.T) {
		d := day.Duration(task1)

		expected := time.Hour
		if d != expected {
			t.Errorf("expected duration to be %q, but got %q", expected, d)
		}
	})
	t.Run("task without a successor should return duration until now", func(t *testing.T) {
		Now = func() time.Time { return MockTime(12, 0) }

		d := day.Duration(task2)

		expected := 2 * time.Hour
		if d != expected {
			t.Errorf("expected duration to be %q, but got %q", expected, d)
		}
	})
}
