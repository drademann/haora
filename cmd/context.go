package cmd

import (
	"time"
)

var ctx context

// Context in which a command executes.
type context struct {

	// DayList with all days recorded so far.
	data dayList

	// WorkingDate represents the global set date to apply commands on.
	workingDate time.Time
}

func initContext(workingDate time.Time) {
	ctx.workingDate = workingDate
}

func (c context) workingDay() day {
	return c.data.day(ctx.workingDate)
}

func (c context) atWorkingDateTime(t time.Time) time.Time {
	return time.Date(c.workingDate.Year(), c.workingDate.Month(), c.workingDate.Day(), t.Hour(), t.Minute(), 0, 0, c.workingDate.Location())
}
