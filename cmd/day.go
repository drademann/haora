package cmd

import (
	"errors"
	"github.com/google/uuid"
	"slices"
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

func (d *Day) Duration(task Task) time.Duration {
	s, err := d.succ(task)
	if errors.Is(err, NoTaskSucc) {
		return now().Sub(task.Start)
	}
	return s.Start.Sub(task.Start)
}

var (
	NoTask     = errors.New("no task")
	NoTaskSucc = errors.New("no succeeding task")
	NoTaskPred = errors.New("no preceding task")
)

func (d *Day) succ(task Task) (Task, error) {
	slices.SortFunc(d.Tasks, tasksByStart)
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

func (d *Day) pred(task Task) (Task, error) {
	slices.SortFunc(d.Tasks, tasksByStart)
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

func (d *Day) update(task Task) {
	for i, t := range d.Tasks {
		if t.Id == task.Id {
			d.Tasks[i] = task
		}
	}
}

func (d *Day) IsToday() bool {
	return isSameDay(d.Date, now())
}

func isSameDay(date1, date2 time.Time) bool {
	return date1.Location() == date2.Location() && date1.Day() == date2.Day() && date1.Month() == date2.Month() && date1.Year() == date2.Year()
}
