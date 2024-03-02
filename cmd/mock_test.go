package cmd

import (
	"testing"
	"time"
)

// mockNowAt allows pretending another timestamp for today (now).
// The returned function should be deferred to reestablish the original now() function.
func mockNowAt(t *testing.T, tm time.Time) time.Time {
	t.Helper()
	now = func() time.Time {
		return tm
	}
	t.Cleanup(func() { now = nowFunc })
	return tm
}

func mockDate(dateStr string) time.Time {
	tm, err := time.Parse("02.01.2006 15:04", dateStr)
	if err != nil {
		panic(err)
	}
	return time.Date(tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), 0, 0, time.Local)
}

func mockTime(timeStr string) time.Time {
	return mockDate("21.06.2024 " + timeStr)
}
