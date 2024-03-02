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

func newTask(s time.Time, tx string, tgs ...string) Task {
	return Task{
		id:      uuid.New(),
		start:   s.Truncate(time.Minute),
		text:    tx,
		isPause: false,
		tags:    tgs,
	}
}

func (t Task) with(s time.Time, tx string, tgs ...string) Task {
	return Task{
		id:    t.id,
		start: s,
		text:  tx,
		tags:  tgs,
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
