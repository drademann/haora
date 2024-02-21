package app

import "time"

type Day struct {
	date     time.Time
	tasks    []Task
	finished time.Time
}
