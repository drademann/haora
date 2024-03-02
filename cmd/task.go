package cmd

import (
	"github.com/google/uuid"
	"time"
)

type task struct {
	id      uuid.UUID
	start   time.Time
	text    string
	isBreak bool
	tags    []string
}

func newTask(s time.Time, tx string, tgs ...string) task {
	return task{
		id:      uuid.New(),
		start:   s.Truncate(time.Minute),
		text:    tx,
		isBreak: false,
		tags:    tgs,
	}
}

func (t task) with(s time.Time, tx string, tgs ...string) task {
	return task{
		id:    t.id,
		start: s,
		text:  tx,
		tags:  tgs,
	}
}

var tasksByStart = func(t1, t2 task) int {
	switch {
	case t1.start.Before(t2.start):
		return -1
	case t1.start.After(t2.start):
		return 1
	default:
		return 0
	}
}
