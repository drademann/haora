package cmd

import (
	"errors"
	"github.com/google/uuid"
	"slices"
	"time"
)

type Day struct {
	id       uuid.UUID
	date     time.Time
	tasks    []Task
	finished time.Time
}

func NewDay(date time.Time) Day {
	return Day{
		id:       uuid.New(),
		date:     date,
		tasks:    []Task{},
		finished: time.Time{},
	}
}

func (d *Day) IsEmpty() bool {
	return len(d.tasks) == 0
}

func (d *Day) totalDuration() time.Duration {
	if len(d.tasks) == 0 {
		return 0
	}
	slices.SortFunc(d.tasks, tasksByStart)
	start := d.tasks[0].start
	end := now()
	if !d.finished.IsZero() {
		end = d.finished
	}
	return end.Sub(start)
}

func (d *Day) totalBreakDuration() time.Duration {
	var sum time.Duration = 0
	for _, t := range d.tasks {
		if t.isPause {
			sum += d.taskDuration(t)
		}
	}
	return sum
}

func (d *Day) totalWorkDuration() time.Duration {
	var sum time.Duration = 0
	for _, t := range d.tasks {
		if !t.isPause {
			sum += d.taskDuration(t)
		}
	}
	return sum
}

func (d *Day) totalTagDuration(tag string) time.Duration {
	var sum time.Duration = 0
	for _, t := range d.tasks {
		if slices.Contains(t.tags, tag) {
			sum += d.taskDuration(t)
		}
	}
	return sum
}

func (d *Day) taskDuration(task Task) time.Duration {
	s, err := d.succ(task)
	if errors.Is(err, NoTaskSucc) {
		if d.finished.IsZero() {
			return now().Sub(task.start)
		}
		return d.finished.Sub(task.start)
	}
	return s.start.Sub(task.start)
}

var (
	NoTask     = errors.New("no task")
	NoTaskSucc = errors.New("no succeeding task")
	NoTaskPred = errors.New("no preceding task")
)

func (d *Day) succ(task Task) (Task, error) {
	slices.SortFunc(d.tasks, tasksByStart)
	for i, t := range d.tasks {
		if t.id == task.id {
			j := i + 1
			if j < len(d.tasks) {
				return d.tasks[j], nil
			}
		}
	}
	return task, NoTaskSucc
}

func (d *Day) pred(task Task) (Task, error) {
	slices.SortFunc(d.tasks, tasksByStart)
	for i, t := range d.tasks {
		if t.id == task.id {
			j := i - 1
			if j >= 0 {
				return d.tasks[j], nil
			}
		}
	}
	return task, NoTaskPred
}

func (d *Day) taskAt(start time.Time) (Task, error) {
	for _, t := range d.tasks {
		if t.start.Hour() == start.Hour() && t.start.Minute() == start.Minute() {
			return t, nil
		}
	}
	return Task{}, NoTask
}

func (d *Day) update(task Task) {
	for i, t := range d.tasks {
		if t.id == task.id {
			d.tasks[i] = task
		}
	}
}

func (d *Day) IsToday() bool {
	return isSameDay(d.date, now())
}

func (d *Day) tags() []string {
	set := make(map[string]struct{})
	for _, t := range d.tasks {
		for _, tag := range t.tags {
			set[tag] = struct{}{}
		}
	}
	var tags []string
	for tag := range set {
		tags = append(tags, tag)
	}
	slices.Sort(tags)
	return tags
}

func isSameDay(date1, date2 time.Time) bool {
	return date1.Location() == date2.Location() && date1.Day() == date2.Day() && date1.Month() == date2.Month() && date1.Year() == date2.Year()
}
