//
// Copyright 2024-2024 The Haora Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package data

import (
	"github.com/drademann/haora/app/datetime"
	"github.com/drademann/haora/test"
	"github.com/drademann/haora/test/assert"
	"testing"
	"time"
)

func TestDayDuration(t *testing.T) {
	t.Run("open day duration", func(t *testing.T) {
		State.WorkingDate = test.Time("0:00")
		task1 := NewTask(test.Time("9:00"), "task 1")
		task2 := NewTask(test.Time("10:00"), "task 2")
		d := Day{Date: State.WorkingDate,
			Tasks:    []Task{task1, task2},
			Finished: time.Time{},
		}
		datetime.AssumeForTestNowAt(t, test.Time("16:00"))

		dur := d.TotalDuration()
		assert.Duration(t, "total", dur, 7*time.Hour)
	})
	t.Run("finished day duration", func(t *testing.T) {
		State.WorkingDate = test.Time("0:00")
		task1 := NewTask(test.Time("9:00"), "task 1")
		task2 := NewTask(test.Time("10:00"), "task 2")
		d := Day{Date: State.WorkingDate,
			Tasks:    []Task{task1, task2},
			Finished: test.Time("14:00"),
		}
		datetime.AssumeForTestNowAt(t, test.Time("23:59"))

		dur := d.TotalDuration()
		assert.Duration(t, "total", dur, 5*time.Hour)
	})
}

func TestTotalWorkBreakDurations(t *testing.T) {
	t.Run("one break", func(t *testing.T) {
		State.WorkingDate = test.Time("0:00")
		task1 := NewTask(test.Time("9:00"), "task 1")
		lunch := NewTask(test.Time("12:00"), "lunch")
		lunch.IsPause = true
		task2 := NewTask(test.Time("12:45"), "task 2")
		d := Day{Date: State.WorkingDate,
			Tasks:    []Task{task1, lunch, task2},
			Finished: test.Time("14:00"),
		}
		datetime.AssumeForTestNowAt(t, test.Time("23:59"))

		b := d.TotalBreakDuration()
		assert.Duration(t, "break", b, 45*time.Minute)

		w := d.TotalWorkDuration()
		assert.Duration(t, "work", w, 4*time.Hour+15*time.Minute)
	})
	t.Run("multiple breaks", func(t *testing.T) {
		State.WorkingDate = test.Time("0:00")
		task1 := NewTask(test.Time("10:00"), "task 1")
		lunch := NewTask(test.Time("12:00"), "lunch")
		lunch.IsPause = true
		task2 := NewTask(test.Time("12:45"), "task 2")
		tea := NewTask(test.Time("16:00"), "tea")
		tea.IsPause = true
		task3 := NewTask(test.Time("16:15"), "task 3")
		d := Day{Date: State.WorkingDate,
			Tasks:    []Task{task1, lunch, task2, tea, task3},
			Finished: test.Time("17:00"),
		}
		datetime.AssumeForTestNowAt(t, test.Time("23:59"))

		b := d.TotalBreakDuration()
		assert.Duration(t, "break", b, 1*time.Hour)

		w := d.TotalWorkDuration()
		assert.Duration(t, "work", w, 6*time.Hour)
	})
	t.Run("open break", func(t *testing.T) {
		State.WorkingDate = test.Time("0:00")
		task1 := NewTask(test.Time("9:00"), "task 1")
		lunch := NewTask(test.Time("12:00"), "break")
		lunch.IsPause = true
		d := Day{Date: State.WorkingDate,
			Tasks:    []Task{task1, lunch},
			Finished: time.Time{},
		}
		datetime.AssumeForTestNowAt(t, test.Time("16:00"))

		b := d.TotalBreakDuration()
		assert.Duration(t, "break", b, 4*time.Hour)

		w := d.TotalWorkDuration()
		assert.Duration(t, "work", w, 3*time.Hour)
	})
}

func TestTagDuration(t *testing.T) {
	State.WorkingDate = test.Time("0:00")
	task1 := NewTask(test.Time("10:00"), "task 1", "T1")
	task2 := NewTask(test.Time("12:00"), "task 2", "T1", "T2")
	task3 := NewTask(test.Time("15:00"), "task 3", "T3")
	d := Day{Date: State.WorkingDate,
		Tasks:    []Task{task1, task2, task3},
		Finished: test.Time("16:00"),
	}

	b := d.TotalTagDuration("T1")
	assert.Duration(t, "T1", b, 5*time.Hour)
	b = d.TotalTagDuration("T2")
	assert.Duration(t, "T2", b, 3*time.Hour)
	b = d.TotalTagDuration("T3")
	assert.Duration(t, "T3", b, 1*time.Hour)
}

func TestTaskDuration(t *testing.T) {
	t.Run("task with a successor", func(t *testing.T) {
		State.WorkingDate = test.Time("0:00")
		task1 := NewTask(test.Time("9:00"), "task 1")
		task2 := NewTask(test.Time("10:00"), "task 2")
		d := Day{Date: State.WorkingDate,
			Tasks:    []Task{task1, task2},
			Finished: time.Time{},
		}

		dur := d.TaskDuration(task1)
		assert.Duration(t, "task", dur, 1*time.Hour)
	})
	t.Run("task without a successor should return duration until now", func(t *testing.T) {
		State.WorkingDate = test.Time("0:00")
		task1 := NewTask(test.Time("9:00"), "task 1")
		task2 := NewTask(test.Time("10:00"), "task 2")
		d := Day{Date: State.WorkingDate,
			Tasks:    []Task{task1, task2},
			Finished: time.Time{},
		}
		datetime.AssumeForTestNowAt(t, test.Time("12:00"))

		dur := d.TaskDuration(task2)
		assert.Duration(t, "task", dur, 2*time.Hour)
	})
	t.Run("given day is finished should use this as end timestamp", func(t *testing.T) {
		State.WorkingDate = test.Time("0:00")
		task1 := NewTask(test.Time("9:00"), "task 1")
		task2 := NewTask(test.Time("10:00"), "task 2")
		d := Day{Date: State.WorkingDate,
			Tasks:    []Task{task1, task2},
			Finished: test.Time("12:00"),
		}
		datetime.AssumeForTestNowAt(t, test.Time("23:59"))

		dur := d.TaskDuration(task2)
		assert.Duration(t, "task", dur, 2*time.Hour)
	})
}
