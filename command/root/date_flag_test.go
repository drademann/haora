package root

import (
	"fmt"
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/test"
	"testing"
	"time"
)

func TestParseNoFlag(t *testing.T) {
	testNow := test.MockNowAt(t, test.MockDate("12.02.2024 10:00"))

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
	test.MockNowAt(t, test.MockDate("12.02.2024 10:00"))

	testCases := []struct {
		name     string
		flag     string
		expected time.Time
	}{
		{"full date DD.MM.YYYY", "15.02.2024", test.MockDate("15.02.2024 00:00")},
		{"full date D.M.YYYY with single digits", "1.2.2024", test.MockDate("01.02.2024 00:00")},
		{"DD.MM. should assume current year", "15.02.", test.MockDate("15.02.2024 00:00")},
		{"DD.MM should accept string without trailing point", "15.02", test.MockDate("15.02.2024 00:00")},
		{"DD. should assume current month and current year", "15.", test.MockDate("15.02.2024 00:00")},
		{"DD should accept string without trailing point", "15", test.MockDate("15.02.2024 00:00")},
		{"D should accept single digits", "8", test.MockDate("08.02.2024 00:00")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data.State.WorkingDate = time.Time{}

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
	test.MockNowAt(t, test.MockDate("25.02.2024 10:00")) // sunday

	testCases := []struct {
		flag     string
		expected time.Time
	}{
		{"mo", test.MockDate("19.02.2024 00:00")},
		{"tu", test.MockDate("20.02.2024 00:00")},
		{"we", test.MockDate("21.02.2024 00:00")},
		{"th", test.MockDate("22.02.2024 00:00")},
		{"fr", test.MockDate("23.02.2024 00:00")},
		{"sa", test.MockDate("24.02.2024 00:00")},
		// does not select today, instead it returned the sunday a week ago
		{"su", test.MockDate("18.02.2024 00:00")},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("flag starting with %q", tc.flag), func(t *testing.T) {
			data.State.WorkingDate = time.Time{}

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
	test.MockNowAt(t, test.MockDate("12.02.2024 10:00"))

	_, err := parseDateFlag("35")

	if err == nil {
		t.Errorf("expected error, but got nil")
	}
}
