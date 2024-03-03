package data

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

func NewTask(s time.Time, tx string, tgs ...string) Task {
	return Task{
		Id:      uuid.New(),
		Start:   s.Truncate(time.Minute),
		Text:    tx,
		IsPause: false,
		Tags:    tgs,
	}
}

func NewPause(s time.Time, tx string) Task {
	return Task{
		Id:      uuid.New(),
		Start:   s.Truncate(time.Minute),
		Text:    tx,
		IsPause: true,
		Tags:    nil,
	}
}

func (t Task) with(s time.Time, tx string, tgs ...string) Task {
	return Task{
		Id:    t.Id,
		Start: s,
		Text:  tx,
		Tags:  tgs,
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
