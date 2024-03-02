package cmd

import (
	"testing"
	"time"
)

func TestDayDuration(t *testing.T) {
	t.Run("open day duration", func(t *testing.T) {
		ctx.workingDate = mockTime("0:00")
		task1 := newTask(mockTime("9:00"), "task 1")
		task2 := newTask(mockTime("10:00"), "task 2")
		d := day{date: ctx.workingDate,
			tasks:    []Task{task1, task2},
			finished: time.Time{},
		}
		mockNowAt(t, mockTime("16:00"))

		dur := d.totalDuration()
		assertDuration(t, "total", dur, 7*time.Hour)
	})
	t.Run("finished day duration", func(t *testing.T) {
		ctx.workingDate = mockTime("0:00")
		task1 := newTask(mockTime("9:00"), "task 1")
		task2 := newTask(mockTime("10:00"), "task 2")
		d := day{date: ctx.workingDate,
			tasks:    []Task{task1, task2},
			finished: mockTime("14:00"),
		}
		mockNowAt(t, mockTime("23:59"))

		dur := d.totalDuration()
		assertDuration(t, "total", dur, 5*time.Hour)
	})
}

func TestTotalWorkBreakDurations(t *testing.T) {
	t.Run("one break", func(t *testing.T) {
		ctx.workingDate = mockTime("0:00")
		task1 := newTask(mockTime("9:00"), "task 1")
		lunch := newTask(mockTime("12:00"), "lunch")
		lunch.isPause = true
		task2 := newTask(mockTime("12:45"), "task 2")
		d := day{date: ctx.workingDate,
			tasks:    []Task{task1, lunch, task2},
			finished: mockTime("14:00"),
		}
		mockNowAt(t, mockTime("23:59"))

		b := d.totalBreakDuration()
		assertDuration(t, "break", b, 45*time.Minute)

		w := d.totalWorkDuration()
		assertDuration(t, "work", w, 4*time.Hour+15*time.Minute)
	})
	t.Run("multiple breaks", func(t *testing.T) {
		ctx.workingDate = mockTime("0:00")
		task1 := newTask(mockTime("10:00"), "task 1")
		lunch := newTask(mockTime("12:00"), "lunch")
		lunch.isPause = true
		task2 := newTask(mockTime("12:45"), "task 2")
		tea := newTask(mockTime("16:00"), "tea")
		tea.isPause = true
		task3 := newTask(mockTime("16:15"), "task 3")
		d := day{date: ctx.workingDate,
			tasks:    []Task{task1, lunch, task2, tea, task3},
			finished: mockTime("17:00"),
		}
		mockNowAt(t, mockTime("23:59"))

		b := d.totalBreakDuration()
		assertDuration(t, "break", b, 1*time.Hour)

		w := d.totalWorkDuration()
		assertDuration(t, "work", w, 6*time.Hour)
	})
	t.Run("open break", func(t *testing.T) {
		ctx.workingDate = mockTime("0:00")
		task1 := newTask(mockTime("9:00"), "task 1")
		lunch := newTask(mockTime("12:00"), "break")
		lunch.isPause = true
		d := day{date: ctx.workingDate,
			tasks:    []Task{task1, lunch},
			finished: time.Time{},
		}
		mockNowAt(t, mockTime("16:00"))

		b := d.totalBreakDuration()
		assertDuration(t, "break", b, 4*time.Hour)

		w := d.totalWorkDuration()
		assertDuration(t, "work", w, 3*time.Hour)
	})
}

func TestTagDuration(t *testing.T) {
	ctx.workingDate = mockTime("0:00")
	task1 := newTask(mockTime("10:00"), "task 1", "T1")
	task2 := newTask(mockTime("12:00"), "task 2", "T1", "T2")
	task3 := newTask(mockTime("15:00"), "task 3", "T3")
	d := day{date: ctx.workingDate,
		tasks:    []Task{task1, task2, task3},
		finished: mockTime("16:00"),
	}

	b := d.totalTagDuration("T1")
	assertDuration(t, "T1", b, 5*time.Hour)
	b = d.totalTagDuration("T2")
	assertDuration(t, "T2", b, 3*time.Hour)
	b = d.totalTagDuration("T3")
	assertDuration(t, "T3", b, 1*time.Hour)
}

func TestTaskDuration(t *testing.T) {
	t.Run("task with a successor", func(t *testing.T) {
		ctx.workingDate = mockTime("0:00")
		task1 := newTask(mockTime("9:00"), "task 1")
		task2 := newTask(mockTime("10:00"), "task 2")
		d := day{date: ctx.workingDate,
			tasks:    []Task{task1, task2},
			finished: time.Time{},
		}

		dur := d.taskDuration(task1)
		assertDuration(t, "task", dur, 1*time.Hour)
	})
	t.Run("task without a successor should return duration until now", func(t *testing.T) {
		ctx.workingDate = mockTime("0:00")
		task1 := newTask(mockTime("9:00"), "task 1")
		task2 := newTask(mockTime("10:00"), "task 2")
		d := day{date: ctx.workingDate,
			tasks:    []Task{task1, task2},
			finished: time.Time{},
		}
		mockNowAt(t, mockTime("12:00"))

		dur := d.taskDuration(task2)
		assertDuration(t, "task", dur, 2*time.Hour)
	})
	t.Run("given day is finished should use this as end timestamp", func(t *testing.T) {
		ctx.workingDate = mockTime("0:00")
		task1 := newTask(mockTime("9:00"), "task 1")
		task2 := newTask(mockTime("10:00"), "task 2")
		d := day{date: ctx.workingDate,
			tasks:    []Task{task1, task2},
			finished: mockTime("12:00"),
		}
		mockNowAt(t, mockTime("23:59"))

		dur := d.taskDuration(task2)
		assertDuration(t, "task", dur, 2*time.Hour)
	})
}
