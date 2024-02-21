package app

import "time"

type Task struct {
	Start   time.Time
	Text    string
	IsPause bool
	Tags    []string
}
