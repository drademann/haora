package cmd

import (
	"testing"
	"time"
)

func TestDayDuration(t *testing.T) {
	t.Run("open day duration", func(t *testing.T) {
		ctx.workingDate = mockTime(0, 0)
		task1 := newTask(mockTime(9, 0), "task 1", nil)
		task2 := newTask(mockTime(10, 0), "task 2", nil)
		day := day{date: ctx.workingDate,
			tasks:    []Task{task1, task2},
			finished: time.Time{},
		}
		mockNowAt(t, mockTime(16, 0))

		d := day.totalDuration()
		assertDuration(t, "total", d, 7*time.Hour)
	})
	t.Run("finished day duration", func(t *testing.T) {
		ctx.workingDate = mockTime(0, 0)
		task1 := newTask(mockTime(9, 0), "task 1", nil)
		task2 := newTask(mockTime(10, 0), "task 2", nil)
		day := day{date: ctx.workingDate,
			tasks:    []Task{task1, task2},
			finished: mockTime(14, 0),
		}
		mockNowAt(t, mockTime(23, 59))

		d := day.totalDuration()
		assertDuration(t, "total", d, 5*time.Hour)
	})
}

func TestTotalWorkBreakDurations(t *testing.T) {
	t.Run("one break", func(t *testing.T) {
		ctx.workingDate = mockTime(0, 0)
		task1 := newTask(mockTime(9, 0), "task 1", nil)
		lunch := newTask(mockTime(12, 0), "lunch", nil)
		lunch.isPause = true
		task2 := newTask(mockTime(12, 45), "task 2", nil)
		day := day{date: ctx.workingDate,
			tasks:    []Task{task1, lunch, task2},
			finished: mockTime(14, 0),
		}
		mockNowAt(t, mockTime(23, 59))

		b := day.totalBreakDuration()
		assertDuration(t, "break", b, 45*time.Minute)

		w := day.totalWorkDuration()
		assertDuration(t, "work", w, 4*time.Hour+15*time.Minute)
	})
	t.Run("multiple breaks", func(t *testing.T) {
		ctx.workingDate = mockTime(0, 0)
		task1 := newTask(mockTime(10, 0), "task 1", nil)
		lunch := newTask(mockTime(12, 0), "lunch", nil)
		lunch.isPause = true
		task2 := newTask(mockTime(12, 45), "task 2", nil)
		tea := newTask(mockTime(16, 0), "tea", nil)
		tea.isPause = true
		task3 := newTask(mockTime(16, 15), "task 3", nil)
		day := day{date: ctx.workingDate,
			tasks:    []Task{task1, lunch, task2, tea, task3},
			finished: mockTime(17, 0),
		}
		mockNowAt(t, mockTime(23, 59))

		b := day.totalBreakDuration()
		assertDuration(t, "break", b, 1*time.Hour)

		w := day.totalWorkDuration()
		assertDuration(t, "work", w, 6*time.Hour)
	})
	t.Run("open break", func(t *testing.T) {
		ctx.workingDate = mockTime(0, 0)
		task1 := newTask(mockTime(9, 0), "task 1", nil)
		lunch := newTask(mockTime(12, 0), "break", nil)
		lunch.isPause = true
		day := day{date: ctx.workingDate,
			tasks:    []Task{task1, lunch},
			finished: time.Time{},
		}
		mockNowAt(t, mockTime(16, 0))

		b := day.totalBreakDuration()
		assertDuration(t, "break", b, 4*time.Hour)

		w := day.totalWorkDuration()
		assertDuration(t, "work", w, 3*time.Hour)
	})
}

func TestTagDuration(t *testing.T) {
	ctx.workingDate = mockTime(0, 0)
	task1 := newTask(mockTime(10, 0), "task 1", []string{"T1"})
	task2 := newTask(mockTime(12, 0), "task 2", []string{"T1", "T2"})
	task3 := newTask(mockTime(15, 0), "task 3", []string{"T3"})
	day := day{date: ctx.workingDate,
		tasks:    []Task{task1, task2, task3},
		finished: mockTime(16, 0),
	}

	b := day.totalTagDuration("T1")
	assertDuration(t, "T1", b, 5*time.Hour)
	b = day.totalTagDuration("T2")
	assertDuration(t, "T2", b, 3*time.Hour)
	b = day.totalTagDuration("T3")
	assertDuration(t, "T3", b, 1*time.Hour)
}

func TestTaskDuration(t *testing.T) {
	t.Run("task with a successor", func(t *testing.T) {
		ctx.workingDate = mockTime(0, 0)
		task1 := newTask(mockTime(9, 0), "task 1", nil)
		task2 := newTask(mockTime(10, 0), "task 2", nil)
		day := day{date: ctx.workingDate,
			tasks:    []Task{task1, task2},
			finished: time.Time{},
		}

		d := day.taskDuration(task1)
		assertDuration(t, "task", d, 1*time.Hour)
	})
	t.Run("task without a successor should return duration until now", func(t *testing.T) {
		ctx.workingDate = mockTime(0, 0)
		task1 := newTask(mockTime(9, 0), "task 1", nil)
		task2 := newTask(mockTime(10, 0), "task 2", nil)
		day := day{date: ctx.workingDate,
			tasks:    []Task{task1, task2},
			finished: time.Time{},
		}
		mockNowAt(t, mockTime(12, 0))

		d := day.taskDuration(task2)
		assertDuration(t, "task", d, 2*time.Hour)
	})
	t.Run("given day is finished should use this as end timestamp", func(t *testing.T) {
		ctx.workingDate = mockTime(0, 0)
		task1 := newTask(mockTime(9, 0), "task 1", nil)
		task2 := newTask(mockTime(10, 0), "task 2", nil)
		day := day{date: ctx.workingDate,
			tasks:    []Task{task1, task2},
			finished: mockTime(12, 0),
		}
		mockNowAt(t, mockTime(23, 59))

		d := day.taskDuration(task2)
		assertDuration(t, "task", d, 2*time.Hour)
	})
}
