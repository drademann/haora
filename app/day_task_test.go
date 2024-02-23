package app

import (
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
		s := day.succ(task2)

		if s.Id != task3.Id {
			t.Errorf("expected successor to be %q, but got %q", task3.Text, s.Text)
		}
	})
}
