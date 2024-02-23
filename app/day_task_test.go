package app

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestSuccPred(t *testing.T) {
	testDayDate := MockTime(0, 0)
	task1 := *NewTask(MockTime(9, 0), "task 1", false, nil)
	task2 := *NewTask(MockTime(10, 0), "task 2", false, nil)
	task3 := *NewTask(MockTime(12, 0), "task 3", false, nil)
	fmt.Println(task1.Id)
	fmt.Println(task2.Id)
	fmt.Println(task3.Id)
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
