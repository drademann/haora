package data

import (
	"github.com/drademann/haora/app/datetime"
	"github.com/drademann/haora/test"
	"testing"
	"time"
)

func TestDayDuration(t *testing.T) {
	t.Run("open day duration", func(t *testing.T) {
		State.WorkingDate = test.MockTime("0:00")
		task1 := NewTask(test.MockTime("9:00"), "task 1")
		task2 := NewTask(test.MockTime("10:00"), "task 2")
		d := Day{Date: State.WorkingDate,
			Tasks:    []Task{task1, task2},
			Finished: time.Time{},
		}
		datetime.MockNowAt(t, test.MockTime("16:00"))

		dur := d.TotalDuration()
		test.AssertDuration(t, "total", dur, 7*time.Hour)
	})
	t.Run("finished day duration", func(t *testing.T) {
		State.WorkingDate = test.MockTime("0:00")
		task1 := NewTask(test.MockTime("9:00"), "task 1")
		task2 := NewTask(test.MockTime("10:00"), "task 2")
		d := Day{Date: State.WorkingDate,
			Tasks:    []Task{task1, task2},
			Finished: test.MockTime("14:00"),
		}
		datetime.MockNowAt(t, test.MockTime("23:59"))

		dur := d.TotalDuration()
		test.AssertDuration(t, "total", dur, 5*time.Hour)
	})
}

func TestTotalWorkBreakDurations(t *testing.T) {
	t.Run("one break", func(t *testing.T) {
		State.WorkingDate = test.MockTime("0:00")
		task1 := NewTask(test.MockTime("9:00"), "task 1")
		lunch := NewTask(test.MockTime("12:00"), "lunch")
		lunch.IsPause = true
		task2 := NewTask(test.MockTime("12:45"), "task 2")
		d := Day{Date: State.WorkingDate,
			Tasks:    []Task{task1, lunch, task2},
			Finished: test.MockTime("14:00"),
		}
		datetime.MockNowAt(t, test.MockTime("23:59"))

		b := d.TotalBreakDuration()
		test.AssertDuration(t, "break", b, 45*time.Minute)

		w := d.TotalWorkDuration()
		test.AssertDuration(t, "work", w, 4*time.Hour+15*time.Minute)
	})
	t.Run("multiple breaks", func(t *testing.T) {
		State.WorkingDate = test.MockTime("0:00")
		task1 := NewTask(test.MockTime("10:00"), "task 1")
		lunch := NewTask(test.MockTime("12:00"), "lunch")
		lunch.IsPause = true
		task2 := NewTask(test.MockTime("12:45"), "task 2")
		tea := NewTask(test.MockTime("16:00"), "tea")
		tea.IsPause = true
		task3 := NewTask(test.MockTime("16:15"), "task 3")
		d := Day{Date: State.WorkingDate,
			Tasks:    []Task{task1, lunch, task2, tea, task3},
			Finished: test.MockTime("17:00"),
		}
		datetime.MockNowAt(t, test.MockTime("23:59"))

		b := d.TotalBreakDuration()
		test.AssertDuration(t, "break", b, 1*time.Hour)

		w := d.TotalWorkDuration()
		test.AssertDuration(t, "work", w, 6*time.Hour)
	})
	t.Run("open break", func(t *testing.T) {
		State.WorkingDate = test.MockTime("0:00")
		task1 := NewTask(test.MockTime("9:00"), "task 1")
		lunch := NewTask(test.MockTime("12:00"), "break")
		lunch.IsPause = true
		d := Day{Date: State.WorkingDate,
			Tasks:    []Task{task1, lunch},
			Finished: time.Time{},
		}
		datetime.MockNowAt(t, test.MockTime("16:00"))

		b := d.TotalBreakDuration()
		test.AssertDuration(t, "break", b, 4*time.Hour)

		w := d.TotalWorkDuration()
		test.AssertDuration(t, "work", w, 3*time.Hour)
	})
}

func TestTagDuration(t *testing.T) {
	State.WorkingDate = test.MockTime("0:00")
	task1 := NewTask(test.MockTime("10:00"), "task 1", "T1")
	task2 := NewTask(test.MockTime("12:00"), "task 2", "T1", "T2")
	task3 := NewTask(test.MockTime("15:00"), "task 3", "T3")
	d := Day{Date: State.WorkingDate,
		Tasks:    []Task{task1, task2, task3},
		Finished: test.MockTime("16:00"),
	}

	b := d.TotalTagDuration("T1")
	test.AssertDuration(t, "T1", b, 5*time.Hour)
	b = d.TotalTagDuration("T2")
	test.AssertDuration(t, "T2", b, 3*time.Hour)
	b = d.TotalTagDuration("T3")
	test.AssertDuration(t, "T3", b, 1*time.Hour)
}

func TestTaskDuration(t *testing.T) {
	t.Run("task with a successor", func(t *testing.T) {
		State.WorkingDate = test.MockTime("0:00")
		task1 := NewTask(test.MockTime("9:00"), "task 1")
		task2 := NewTask(test.MockTime("10:00"), "task 2")
		d := Day{Date: State.WorkingDate,
			Tasks:    []Task{task1, task2},
			Finished: time.Time{},
		}

		dur := d.TaskDuration(task1)
		test.AssertDuration(t, "task", dur, 1*time.Hour)
	})
	t.Run("task without a successor should return duration until now", func(t *testing.T) {
		State.WorkingDate = test.MockTime("0:00")
		task1 := NewTask(test.MockTime("9:00"), "task 1")
		task2 := NewTask(test.MockTime("10:00"), "task 2")
		d := Day{Date: State.WorkingDate,
			Tasks:    []Task{task1, task2},
			Finished: time.Time{},
		}
		datetime.MockNowAt(t, test.MockTime("12:00"))

		dur := d.TaskDuration(task2)
		test.AssertDuration(t, "task", dur, 2*time.Hour)
	})
	t.Run("given day is finished should use this as end timestamp", func(t *testing.T) {
		State.WorkingDate = test.MockTime("0:00")
		task1 := NewTask(test.MockTime("9:00"), "task 1")
		task2 := NewTask(test.MockTime("10:00"), "task 2")
		d := Day{Date: State.WorkingDate,
			Tasks:    []Task{task1, task2},
			Finished: test.MockTime("12:00"),
		}
		datetime.MockNowAt(t, test.MockTime("23:59"))

		dur := d.TaskDuration(task2)
		test.AssertDuration(t, "task", dur, 2*time.Hour)
	})
}
