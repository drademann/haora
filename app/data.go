package app

import (
	"time"
)

// Data represents the so far added Days.
//
// The app will load the list before executing a command,
// and will save the (changed) list after the command finishes without an error.
var Data DayList

// WorkingDate represents the global set date to apply commands on.
var WorkingDate time.Time

func atWorkingDateTime(t time.Time) time.Time {
	return time.Date(WorkingDate.Year(), WorkingDate.Month(), WorkingDate.Day(), t.Hour(), t.Minute(), 0, 0, WorkingDate.Location())
}
