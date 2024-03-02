package data

import (
	"errors"
	"github.com/drademann/haora/test"
	"testing"
	"time"
)

func TestTaskSuccPred(t *testing.T) {
	testDayDate := test.MockTime("0:00")
	task1 := NewTask(test.MockTime("9:00"), "task 1")
	task2 := NewTask(test.MockTime("10:00"), "task 2")
	task3 := NewTask(test.MockTime("12:00"), "task 3")
	d := Day{Date: testDayDate,
		Tasks:    []Task{task1, task2, task3},
		Finished: time.Time{},
	}

	t.Run("find successor", func(t *testing.T) {
		s, err := d.Succ(task2)

		if err != nil {
			t.Fatal(err)
		}
		if s.Id != task3.Id {
			t.Errorf("expected successor to be %q, but got %q", task3.Text, s.Text)
		}
	})
	t.Run("find no successor", func(t *testing.T) {
		_, err := d.Succ(task3)

		if err == nil {
			t.Errorf("expected error, but got nil")
		}
		if !errors.Is(err, NoTaskSucc) {
			t.Errorf("expected error %q, but got %q", NoTaskSucc, err)
		}
	})
	t.Run("find predecessor", func(t *testing.T) {
		p, err := d.pred(task2)

		if err != nil {
			t.Fatal(err)
		}
		if p.Id != task1.Id {
			t.Errorf("expected predecessor to be %q, but got %q", task1.Text, p.Text)
		}
	})
	t.Run("find no predecessor", func(t *testing.T) {
		_, err := d.pred(task1)

		if err == nil {
			t.Errorf("expected error, but got nil")
		}
		if !errors.Is(err, NoTaskPred) {
			t.Errorf("expected error %q, but got %q", NoTaskPred, err)
		}
	})
}
