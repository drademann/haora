package cmd

import (
	"fmt"
	"testing"
	"time"
)

func TestParseNoFlag(t *testing.T) {
	testNow := mockDate(2024, time.February, 12, 10, 0)
	mockNowAt(t, testNow)

	date, err := parseDateFlag("")
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	expected := testNow
	if date != expected {
		t.Errorf("expected parsed working date to be %v, but got %v", expected, date)
	}
}

func TestParseDateFlag(t *testing.T) {
	mockNowAt(t, mockDate(2024, time.February, 12, 10, 0))

	testCases := []struct {
		name     string
		flag     string
		expected time.Time
	}{
		{"full date DD.MM.YYYY", "15.02.2024", mockDate(2024, time.February, 15, 0, 0)},
		{"full date D.M.YYYY with single digits", "1.2.2024", mockDate(2024, time.February, 1, 0, 0)},
		{"DD.MM. should assume current year", "15.02.", mockDate(2024, time.February, 15, 0, 0)},
		{"DD.MM should accept string without trailing point", "15.02", mockDate(2024, time.February, 15, 0, 0)},
		{"DD. should assume current month and current year", "15.", mockDate(2024, time.February, 15, 0, 0)},
		{"DD should accept string without trailing point", "15", mockDate(2024, time.February, 15, 0, 0)},
		{"D should accept single digits", "8", mockDate(2024, time.February, 8, 0, 0)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx.workingDate = time.Time{}

			date, err := parseDateFlag(tc.flag)
			if err != nil {
				t.Fatal(err)
			}

			if date != tc.expected {
				t.Errorf("expected parsed date to be %v, but got %v", tc.expected, date)
			}
		})
	}
}

func TestWeekdayFlag(t *testing.T) {
	mockNowAt(t, mockDate(2024, time.February, 25, 10, 0)) // sunday

	testCases := []struct {
		flag     string
		expected time.Time
	}{
		{"mo", mockDate(2024, time.February, 19, 0, 0)},
		{"tu", mockDate(2024, time.February, 20, 0, 0)},
		{"we", mockDate(2024, time.February, 21, 0, 0)},
		{"th", mockDate(2024, time.February, 22, 0, 0)},
		{"fr", mockDate(2024, time.February, 23, 0, 0)},
		{"sa", mockDate(2024, time.February, 24, 0, 0)},
		// does not select today, instead it returned the sunday a week ago
		{"su", mockDate(2024, time.February, 18, 0, 0)},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("flag starting with %q", tc.flag), func(t *testing.T) {
			ctx.workingDate = time.Time{}

			date, err := parseDateFlag(tc.flag)
			if err != nil {
				t.Fatal(err)
			}

			if date != tc.expected {
				t.Errorf("expected parsed working date to be %v, but got %v", tc.expected, date)
			}
		})
	}
}

func TestParseDayOnly(t *testing.T) {
	mockNowAt(t, mockDate(2024, time.February, 12, 10, 0))

	_, err := parseDateFlag("35")

	if err == nil {
		t.Errorf("expected error, but got nil")
	}
}
