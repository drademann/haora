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
	if workingDate != expected {
		t.Errorf("expected parsed working date to be %v, but got %v", expected, workingDate)
	}
}

func TestParseDateFlag(t *testing.T) {
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
			realNow := app.Now
			defer func() { app.Now = realNow }()
			testNow := test.MockDate(2024, time.February, 12, 10, 0)
			app.Now = func() time.Time { return testNow }

			workingDate = time.Time{}
			*workingDateFlag = tc.flag

			err := ParseDateFlag()

			if err != nil {
				t.Fatal(err)
			}
			if workingDate != tc.expected {
				t.Errorf("expected parsed working date to be %v, but got %v", tc.expected, workingDate)
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

	fmt.Println(workingDate)
	if err == nil {
		t.Errorf("expected error, but got nil")
	}
}
