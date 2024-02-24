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

func NewTask(start time.Time, text string, isPause bool, tags []string) *Task {
	return &Task{
		Id:      uuid.New(),
		Start:   start.Truncate(time.Minute),
		Text:    text,
		IsPause: isPause,
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
