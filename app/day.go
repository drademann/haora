package app

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

func NewDay(date time.Time) *Day {
	return &Day{
		Id:       uuid.New(),
		Date:     date,
		Tasks:    []Task{},
		Finished: time.Time{},
	}
}

func (d Day) HasNoTasks() bool {
	return len(d.Tasks) == 0
}

var (
	NoTaskSucc = errors.New("no succeeding task")
	NoTaskPred = errors.New("no preceding task")
)

func (d Day) succ(task Task) (Task, error) {
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

func (d Day) pred(task Task) (Task, error) {
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

func isSameDay(date1, date2 time.Time) bool {
	return date1.Location() == date2.Location() && date1.Day() == date2.Day() && date1.Month() == date2.Month() && date1.Year() == date2.Year()
}
