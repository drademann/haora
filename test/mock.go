package test

import (
	"github.com/drademann/haora/app/datetime"
	"testing"
	"time"
)

// MockNowAt allows pretending another timestamp for today (now).
// The returned function should be deferred to reestablish the original now() function.
func MockNowAt(t *testing.T, tm time.Time) time.Time {
	t.Helper()
	datetime.Now = func() time.Time {
		return tm
	}
	t.Cleanup(func() { datetime.Now = datetime.NowFunc })
	return tm
}

func MockDate(dateStr string) time.Time {
	tm, err := time.Parse("02.01.2006 15:04", dateStr)
	if err != nil {
		panic(err)
	}
	return time.Date(tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), 0, 0, time.Local)
}

func MockTime(timeStr string) time.Time {
	return MockDate("21.06.2024 " + timeStr)
}
