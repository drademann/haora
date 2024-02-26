package app

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	Id      uuid.UUID
	Start   time.Time
	Text    string
	IsPause bool
	Tags    []string
}

func NewTask(start time.Time, text string, tags []string) Task {
	s := time.Date(WorkingDate.Year(), WorkingDate.Month(), WorkingDate.Day(), start.Hour(), start.Minute(), 0, 0, WorkingDate.Location())
	return Task{
		Id:      uuid.New(),
		Start:   s,
		Text:    text,
		IsPause: false,
		Tags:    tags,
	}
}

var tasksByStart = func(t1, t2 Task) int {
	switch {
	case t1.Start.Before(t2.Start):
		return -1
	case t1.Start.After(t2.Start):
		return 1
	default:
		return 0
	}
}
