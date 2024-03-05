package datetime

import (
	"testing"
	"time"
)

// Now timestamp as variable to allow tests to override it.
//
// Seconds and nanoseconds are truncated and set to zero because all calculations in Haora are based on minutes.
var Now = NowFunc

func NowFunc() time.Time {
	return time.Now().Truncate(time.Minute)
}

// MockNowAt allows pretending another timestamp for today (now).
// The returned function should be deferred to reestablish the original now() function.
// Should not be called from production code!
func MockNowAt(t *testing.T, tm time.Time) time.Time {
	t.Helper()
	Now = func() time.Time {
		return tm
	}
	t.Cleanup(func() { Now = NowFunc })
	return tm
}

func Combine(d, t time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), t.Hour(), t.Minute(), 0, 0, d.Location())
}

type Direction int

const (
	Previous Direction = -1
	// Next Direction = 1
)

func FindWeekday(date time.Time, dir Direction, weekday time.Weekday) time.Time {
	step := time.Duration(dir * 24)
	date = date.Add(step * time.Hour)
	if date.Weekday() == weekday {
		return date
	}
	return FindWeekday(date, dir, weekday)
}
