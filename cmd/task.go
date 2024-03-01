package cmd

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	id      uuid.UUID
	start   time.Time
	text    string
	isPause bool
	tags    []string
}

func newTask(start time.Time, text string, tags []string) Task {
	return Task{
		id:      uuid.New(),
		start:   ctx.atWorkingDateTime(start),
		text:    text,
		isPause: false,
		tags:    tags,
	}
}

func (t Task) with(start time.Time, text string, tags []string) Task {
	return Task{
		id:    t.id,
		start: ctx.atWorkingDateTime(start),
		text:  text,
		tags:  tags,
	}
}

var tasksByStart = func(t1, t2 Task) int {
	switch {
	case t1.start.Before(t2.start):
		return -1
	case t1.start.After(t2.start):
		return 1
	default:
		return 0
	}
}
