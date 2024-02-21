package app_test

import (
	"fmt"
	"testing"
	"time"

	"haora/app"
)

var location *time.Location

func init() {
	var err error
	location, err = time.LoadLocation("Europe/Berlin")
	if err != nil {
		panic(fmt.Sprintf("unable to load location for tests: %v", err))
	}
}

func TestNewDay(t *testing.T) {
	t.Run("should initialise the days date", func(t *testing.T) {
		date := time.Date(2024, time.February, 21, 14, 58, 42, 12, location)
		day := app.NewDay(date)

		if len(day.Tasks) != 0 {
			t.Errorf("expected new day to have no tasks, but found %d", len(day.Tasks))
		}
		if !day.Finished.IsZero() {
			t.Errorf("didn't expect new day to be finished from the beginning")
		}
	})
}

func TestSameDay(t *testing.T) {
	testCases := []struct {
		name     string
		date1    time.Time
		date2    time.Time
		expected bool
	}{
		{"dates at exact same time should return true",
			time.Date(2024, time.February, 21, 10, 0, 0, 0, location),
			time.Date(2024, time.February, 21, 10, 0, 0, 0, location),
			true},
		{"dates at same day should return true",
			time.Date(2024, time.February, 21, 10, 0, 0, 0, location),
			time.Date(2024, time.February, 21, 15, 22, 31, 0, location),
			true},
		{"dates at different days should return false",
			time.Date(2024, time.February, 21, 10, 0, 0, 0, location),
			time.Date(2024, time.February, 12, 10, 0, 0, 0, location),
			false},
		{"dates at different month should return false",
			time.Date(2024, time.February, 21, 10, 0, 0, 0, location),
			time.Date(2024, time.December, 21, 10, 0, 0, 0, location),
			false},
		{"dates at different years should return false",
			time.Date(2024, time.February, 21, 10, 0, 0, 0, location),
			time.Date(2025, time.February, 21, 10, 0, 0, 0, location),
			false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := app.IsSameDay(tc.date1, tc.date2)
			if result != tc.expected {
				t.Errorf("expected %t, but got %t", tc.expected, result)
			}
		})
	}
}
