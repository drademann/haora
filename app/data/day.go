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
	"errors"
	"github.com/drademann/haora/app"
	"github.com/drademann/haora/app/datetime"
	"github.com/google/uuid"
	"slices"
	"strings"
	"time"
)

type Day struct {
	Id       uuid.UUID
	Date     time.Time
	Tasks    []Task
	Finished time.Time
}

func NewDay(date time.Time) Day {
	return Day{
		Id:       uuid.New(),
		Date:     date,
		Tasks:    []Task{},
		Finished: time.Time{},
	}
}

func (d *Day) IsEmpty() bool {
	return len(d.Tasks) == 0
}

func (d *Day) IsToday() bool {
	return isSameDay(d.Date, datetime.Now())
}

func (d *Day) IsFinished() bool {
	return !d.Finished.IsZero()
}

// AddNewTask creates a new task.
//
// The new task starts at the start timestamp with given text and tags.
// The date part of the start timestamp is not used, instead the day's date is applied.
// If a task at the specific timestamp already exists, it will be updated instead of added.
func (d *Day) AddNewTask(s time.Time, tx string, tgs []string) error {
	s = datetime.Combine(d.Date, s)
	t, err := d.taskAt(s)
	if err != nil {
		if errors.Is(err, NoTask) {
			t = NewTask(s, tx, tgs...)
			d.AddTasks(t)
		} else {
			return err
		}
	} else {
		t = t.with(s, tx, tgs...)
		d.updateTask(t)
	}
	State.DayList.update(*d)
	return nil
}

func (d *Day) AddNewPause(s time.Time, tx string) error {
	s = datetime.Combine(d.Date, s)
	p, err := d.taskAt(s)
	if err != nil {
		if errors.Is(err, NoTask) {
			p = NewPause(s, tx)
			d.AddTasks(p)
		} else {
			return err
		}
	} else {
		p = p.with(s, tx)
		d.updateTask(p)
	}
	State.DayList.update(*d)
	return nil
}

func (d *Day) AddTasks(tasks ...Task) {
	d.Tasks = append(d.Tasks, tasks...)
	slices.SortFunc(d.Tasks, tasksByStart)
}

func (d *Day) Start() time.Time {
	if d.IsEmpty() {
		panic("check Day.IsEmpty before trying to get start time")
	}
	return d.Tasks[0].Start
}

func (d *Day) End() time.Time {
	if d.Finished.IsZero() {
		return datetime.Now()
	}
	return d.Finished
}

func (d *Day) Finish(f time.Time) {
	f = datetime.Combine(d.Date, f)
	d.Finished = f
	State.DayList.update(*d)
}

func (d *Day) TotalDuration() time.Duration {
	if d.IsEmpty() {
		return 0
	}
	return d.End().Sub(d.Start())
}

func (d *Day) TotalBreakDuration() time.Duration {
	var sum time.Duration = 0
	for _, t := range d.Tasks {
		if t.IsPause {
			sum += d.TaskDuration(t)
		}
	}
	return sum
}

func (d *Day) TotalWorkDuration() time.Duration {
	var sum time.Duration = 0
	for _, t := range d.Tasks {
		if !t.IsPause {
			sum += d.TaskDuration(t)
		}
	}
	return sum
}

func (d *Day) OvertimeDuration() (time.Duration, error) {
	durationPerDay, err := app.DurationPerDay()
	if err != nil {
		return 0, err
	}
	return d.TotalWorkDuration() - durationPerDay, nil
}

func (d *Day) TotalTagDuration(tag string) time.Duration {
	var sum time.Duration = 0
	for _, t := range d.Tasks {
		if slices.Contains(t.Tags, tag) {
			sum += d.TaskDuration(t)
		}
	}
	return sum
}

func (d *Day) TaskDuration(task Task) time.Duration {
	s, err := d.Succ(task)
	if errors.Is(err, NoTaskSucc) {
		if d.Finished.IsZero() {
			return datetime.Now().Sub(task.Start)
		}
		return d.Finished.Sub(task.Start)
	}
	return s.Start.Sub(task.Start)
}

var (
	NoTask     = errors.New("no task")
	NoTaskSucc = errors.New("no succeeding task")
	NoTaskPred = errors.New("no preceding task")
)

func (d *Day) Succ(task Task) (Task, error) {
	for i, t := range d.Tasks {
		if t.Id == task.Id {
			j := i + 1
			if j < len(d.Tasks) {
				return d.Tasks[j], nil
			}
		}
	}
	return task, NoTaskSucc
}

func (d *Day) Pred(task Task) (Task, error) {
	for i, t := range d.Tasks {
		if t.Id == task.Id {
			j := i - 1
			if j >= 0 {
				return d.Tasks[j], nil
			}
		}
	}
	return task, NoTaskPred
}

func (d *Day) taskAt(start time.Time) (Task, error) {
	for _, t := range d.Tasks {
		if t.Start.Hour() == start.Hour() && t.Start.Minute() == start.Minute() {
			return t, nil
		}
	}
	return Task{}, NoTask
}

func (d *Day) Tags() []string {
	set := make(map[string]struct{})
	for _, t := range d.Tasks {
		for _, tag := range t.Tags {
			set[tag] = struct{}{}
		}
	}
	var res []string
	for tag := range set {
		res = append(res, tag)
	}
	slices.Sort(res)
	return res
}

func (d *Day) updateTask(task Task) {
	for i, t := range d.Tasks {
		if t.Id == task.Id {
			d.Tasks[i] = task
		}
	}
}

func (d *Day) sanitize() Day {
	for _, t := range d.Tasks {
		for j := range t.Tags {
			t.Tags[j] = strings.ToLower(strings.TrimSpace(t.Tags[j]))
		}
	}
	return *d
}

func isSameDay(date1, date2 time.Time) bool {
	return date1.Location() == date2.Location() && date1.Day() == date2.Day() && date1.Month() == date2.Month() && date1.Year() == date2.Year()
}
