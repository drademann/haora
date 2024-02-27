package cmd

import (
	"fmt"
	"haora/app"
	"haora/test"
	"testing"
	"time"
)

func TestParseNoFlag(t *testing.T) {
	realNow := app.Now
	defer func() { app.Now = realNow }()
	testNow := test.MockDate(2024, time.February, 12, 10, 0)
	app.Now = func() time.Time { return testNow }

	*workingDateFlag = ""

	err := ParseDateFlag()
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	expected := testNow
	if app.WorkingDate != expected {
		t.Errorf("expected parsed working date to be %v, but got %v", expected, app.WorkingDate)
	}
}

func TestParseDateFlag(t *testing.T) {
	realNow := app.Now
	defer func() { app.Now = realNow }()
	testNow := test.MockDate(2024, time.February, 12, 10, 0)
	app.Now = func() time.Time { return testNow }

	testCases := []struct {
		name     string
		flag     string
		expected time.Time
	}{
		{"full date DD.MM.YYYY", "15.02.2024", test.MockDate(2024, time.February, 15, 0, 0)},
		{"full date D.M.YYYY with single digits", "1.2.2024", test.MockDate(2024, time.February, 1, 0, 0)},
		{"DD.MM. should assume current year", "15.02.", test.MockDate(2024, time.February, 15, 0, 0)},
		{"DD.MM should accept string without trailing point", "15.02", test.MockDate(2024, time.February, 15, 0, 0)},
		{"DD. should assume current month and current year", "15.", test.MockDate(2024, time.February, 15, 0, 0)},
		{"DD should accept string without trailing point", "15", test.MockDate(2024, time.February, 15, 0, 0)},
		{"D should accept single digits", "8", test.MockDate(2024, time.February, 8, 0, 0)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			app.WorkingDate = time.Time{}
			*workingDateFlag = tc.flag

			err := ParseDateFlag()

			if err != nil {
				t.Fatal(err)
			}
			if app.WorkingDate != tc.expected {
				t.Errorf("expected parsed working date to be %v, but got %v", tc.expected, app.WorkingDate)
			}
		})
	}
}

func TestWeekdayFlag(t *testing.T) {
	realNow := app.Now
	defer func() { app.Now = realNow }()
	testNow := test.MockDate(2024, time.February, 25, 10, 0) // sunday
	app.Now = func() time.Time { return testNow }

	testCases := []struct {
		flag     string
		expected time.Time
	}{
		{"mo", test.MockDate(2024, time.February, 19, 0, 0)},
		{"tu", test.MockDate(2024, time.February, 20, 0, 0)},
		{"we", test.MockDate(2024, time.February, 21, 0, 0)},
		{"th", test.MockDate(2024, time.February, 22, 0, 0)},
		{"fr", test.MockDate(2024, time.February, 23, 0, 0)},
		{"sa", test.MockDate(2024, time.February, 24, 0, 0)},
		// does not select today, instead it returned the sunday a week ago
		{"su", test.MockDate(2024, time.February, 18, 0, 0)},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("flag starting with %q", tc.flag), func(t *testing.T) {
			app.WorkingDate = time.Time{}
			*workingDateFlag = tc.flag

			err := ParseDateFlag()

			if err != nil {
				t.Fatal(err)
			}
			if app.WorkingDate != tc.expected {
				t.Errorf("expected parsed working date to be %v, but got %v", tc.expected, app.WorkingDate)
			}
		})
	}
}

func TestParseDayOnly(t *testing.T) {
	realNow := app.Now
	defer func() { app.Now = realNow }()
	testNow := test.MockDate(2024, time.February, 12, 10, 0)
	app.Now = func() time.Time { return testNow }

	*workingDateFlag = "35" // there is no "31st" feb

	err := ParseDateFlag()

	if err == nil {
		t.Errorf("expected error, but got nil")
	}
}
