package app

import "time"

type Task struct {
	start   time.Time
	text    string
	isPause bool
	tags    []string
}
