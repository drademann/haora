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
	"fmt"
	"github.com/drademann/haora/app/datetime"
	"github.com/drademann/haora/cmd/config"
	"slices"
	"sort"
	"strings"
	"time"
)

type Day struct {
	Date     time.Time
	Tasks    []*Task
	Finished time.Time
}

func NewDay(date time.Time) *Day {
	return &Day{
		Date:     date,
		Tasks:    make([]*Task, 0),
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

func (d *Day) AddNewTask(s time.Time, txt string, tgs []string) error {
	s = datetime.Combine(d.Date, s)
	t, err := d.taskAt(s)
	if err != nil {
		if errors.Is(err, NoTask) {
			t = NewTask(s, txt, tgs...)
			d.AddTask(t)
		} else {
			return err
		}
	} else {
		t.Start = s
		t.Text = txt
		t.Tags = tgs
	}
	return nil
}

func (d *Day) AddNewPause(s time.Time, txt string) error {
	s = datetime.Combine(d.Date, s)
	p, err := d.taskAt(s)
	if err != nil {
		if errors.Is(err, NoTask) {
			p = NewPause(s, txt)
			d.AddTask(p)
		} else {
			return err
		}
	} else {
		p.Start = s
		p.Text = txt
	}
	return nil
}

func (d *Day) AddTask(task *Task) {
	d.Tasks = append(d.Tasks, task)
	slices.SortFunc(d.Tasks, tasksByStart)
}

// RemoveTask removes a task from the Day’s list of tasks based on the specified time.
// After removing a found task it returns true.
func (d *Day) RemoveTask(timeToDelete time.Time) bool {
	timeToDelete = datetime.Combine(d.Date, timeToDelete)
	i, found := sort.Find(len(d.Tasks), func(i int) int {
		return timeToDelete.Compare(d.Tasks[i].Start)
	})
	if found {
		d.Tasks = append(d.Tasks[:i], d.Tasks[i+1:]...)
		return true
	}
	return false
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

func (d *Day) Finish(f time.Time) error {
	if len(d.Tasks) == 0 {
		return errors.New("no tasks to finish")
	}
	f = datetime.Combine(d.Date, f)
	last := d.Tasks[len(d.Tasks)-1]
	if f.Before(last.Start) {
		return fmt.Errorf("can't finish before last task's start timestamp (%s)", last.Start.Format("15:04"))
	}
	d.Finished = f
	return nil
}

func (d *Day) Unfinished() {
	d.Finished = time.Time{}
}

// SuggestedFinish returns a suggested finish time. Empty days or already finished days return false.
func (d *Day) SuggestedFinish() (time.Time, bool) {
	durationPerDay, exist := config.DurationPerDay()
	if !exist || d.IsFinished() || d.IsEmpty() {
		return time.Time{}, false
	}
	return d.Tasks[0].Start.Add(durationPerDay).Add(d.TotalPauseDuration()), true
}

func (d *Day) TotalDuration() time.Duration {
	var total time.Duration
	for _, task := range d.Tasks {
		total += d.TaskDuration(*task)
	}
	return total
}

func (d *Day) TotalPauseDuration() time.Duration {
	var sum time.Duration = 0
	for _, t := range d.Tasks {
		if t.IsPause {
			sum += d.TaskDuration(*t)
		}
	}
	return sum
}

func (d *Day) TotalWorkDuration() time.Duration {
	var sum time.Duration = 0
	for _, t := range d.Tasks {
		if !t.IsPause {
			sum += d.TaskDuration(*t)
		}
	}
	return sum
}

func (d *Day) OvertimeDuration() (time.Duration, bool) {
	durationPerDay, exist := config.DurationPerDay()
	if !exist {
		return 0, false
	}
	return d.TotalWorkDuration() - durationPerDay, true
}

func (d *Day) TotalTagDuration(tag string) time.Duration {
	var sum time.Duration = 0
	for _, t := range d.Tasks {
		if slices.Contains(t.Tags, tag) {
			sum += d.TaskDuration(*t)
		}
	}
	return sum
}

func (d *Day) TaskDuration(task Task) time.Duration {
	s, err := d.Succ(task)
	if errors.Is(err, NoTaskSucc) {
		switch {
		case !d.IsFinished() && task.Start.After(datetime.Now()):
			return 0
		case !d.IsFinished():
			return datetime.Now().Sub(task.Start)
		default:
			return d.Finished.Sub(task.Start)
		}
	}
	return s.Start.Sub(task.Start)
}

var (
	NoTask     = errors.New("no task")
	NoTaskSucc = errors.New("no succeeding task")
	NoTaskPred = errors.New("no preceding task")
)

func (d *Day) Succ(task Task) (*Task, error) {
	for i, t := range d.Tasks {
		if t.Start == task.Start {
			j := i + 1
			if j < len(d.Tasks) {
				return d.Tasks[j], nil
			}
		}
	}
	return nil, NoTaskSucc
}

func (d *Day) Pred(task Task) (*Task, error) {
	for i, t := range d.Tasks {
		if t.Start == task.Start {
			j := i - 1
			if j >= 0 {
				return d.Tasks[j], nil
			}
		}
	}
	return nil, NoTaskPred
}

func (d *Day) taskAt(start time.Time) (*Task, error) {
	for _, t := range d.Tasks {
		if t.Start.Hour() == start.Hour() && t.Start.Minute() == start.Minute() {
			return t, nil
		}
	}
	return nil, NoTask
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

func (d *Day) sanitize() *Day {
	for _, t := range d.Tasks {
		for j := range t.Tags {
			t.Tags[j] = strings.ToLower(strings.TrimSpace(t.Tags[j]))
		}
	}
	return d
}

func isSameDay(date1, date2 time.Time) bool {
	return date1.Location() == date2.Location() && date1.Day() == date2.Day() && date1.Month() == date2.Month() && date1.Year() == date2.Year()
}
