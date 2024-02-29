package cmd

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

func newTask(start time.Time, text string, tags []string) Task {
	return Task{
		Id:      uuid.New(),
		Start:   ctx.atWorkingDateTime(start),
		Text:    text,
		IsPause: false,
		Tags:    tags,
	}
}

func (t Task) with(start time.Time, text string, tags []string) Task {
	return Task{
		Id:    t.Id,
		Start: ctx.atWorkingDateTime(start),
		Text:  text,
		Tags:  tags,
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
