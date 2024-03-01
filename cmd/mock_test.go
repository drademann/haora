package cmd

import (
	uuid2 "github.com/google/uuid"
	"testing"
	"time"
)

// mockNowAt allows pretending another timestamp for today (now).
// The returned function should be deferred to reestablish the original now() function.
func mockNowAt(t *testing.T, testNow time.Time) {
	now = func() time.Time {
		return testNow
	}
	t.Cleanup(func() { now = nowFunc })
}

func mockDate(year int, month time.Month, day, hour, minute int) time.Time {
	return time.Date(year, month, day, hour, minute, 0, 0, time.Local)
}

func mockTime(hour, minute int) time.Time {
	return mockDate(2024, time.June, 21, hour, minute)
}

func mockTask(t *testing.T, start time.Time, text string, tags ...string) Task {
	t.Helper()
	uuid, err := uuid2.NewRandom()
	if err != nil {
		t.Fatal(err)
	}
	return Task{
		id:      uuid,
		start:   start,
		text:    text,
		isPause: false,
		tags:    tags,
	}
}
