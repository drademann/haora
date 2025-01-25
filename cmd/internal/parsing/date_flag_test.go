//
// Copyright 2024-2024 The Haora Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package parsing

import (
	"fmt"
	"github.com/drademann/fugo/test"
	"github.com/drademann/haora/app/datetime"
	"testing"
	"time"
)

func TestParseNoFlag(t *testing.T) {
	testNow := datetime.AssumeForTestNowAt(t, test.Date("12.02.2024 10:00"))

	date, err := WorkingDate("")
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	expected := testNow
	if date != expected {
		t.Errorf("expected parsed working date to be %v, but got %v", expected, date)
	}
}

func TestParseDateFlag(t *testing.T) {
	datetime.AssumeForTestNowAt(t, test.Date("12.02.2024 10:00"))

	testCases := []struct {
		name     string
		flag     string
		expected time.Time
	}{
		{"full date DD.MM.YYYY", "15.02.2024", test.Date("15.02.2024 00:00")},
		{"full date D.M.YYYY with single digits", "1.2.2024", test.Date("01.02.2024 00:00")},
		{"full date D.M.YY with short year", "1.2.24", test.Date("01.02.2024 00:00")},
		{"DD.MM. should assume current year", "15.02.", test.Date("15.02.2024 00:00")},
		{"DD.MM should accept string without trailing point", "15.02", test.Date("15.02.2024 00:00")},
		{"DD. should assume current month and current year", "15.", test.Date("15.02.2024 00:00")},
		{"DD should accept string without trailing point", "15", test.Date("15.02.2024 00:00")},
		{"D should accept single digits", "8", test.Date("08.02.2024 00:00")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			date, err := WorkingDate(tc.flag)
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
	datetime.AssumeForTestNowAt(t, test.Date("25.02.2024 10:00")) // sunday

	testCases := []struct {
		flag     string
		expected time.Time
	}{
		{"mo", test.Date("19.02.2024 00:00")},
		{"tu", test.Date("20.02.2024 00:00")},
		{"we", test.Date("21.02.2024 00:00")},
		{"th", test.Date("22.02.2024 00:00")},
		{"fr", test.Date("23.02.2024 00:00")},
		{"sa", test.Date("24.02.2024 00:00")},
		// does not select today, instead it returned the sunday a week ago
		{"su", test.Date("18.02.2024 00:00")},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("flag starting with %q", tc.flag), func(t *testing.T) {
			date, err := WorkingDate(tc.flag)
			if err != nil {
				t.Fatal(err)
			}

			if date != tc.expected {
				t.Errorf("expected parsed working date to be %v, but got %v", tc.expected, date)
			}
		})
	}
}

func TestYesterdayFlag(t *testing.T) {
	datetime.AssumeForTestNowAt(t, test.Date("18.05.2024 12:00"))

	testCases := []string{
		"yesterday",
		"yes",
		"ye",
		"yd",
		"y",
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("flag for yesterday: %q", tc), func(t *testing.T) {
			date, err := WorkingDate(tc)
			if err != nil {
				t.Fatal(err)
			}

			expectedYesterday := test.Date("17.05.2024 00:00")
			if date != expectedYesterday {
				t.Errorf("expected parsed working date to be yesterday, the %v, but got %v", expectedYesterday, date)
			}
		})
	}
}

func TestParseDayOnly(t *testing.T) {
	datetime.AssumeForTestNowAt(t, test.Date("12.02.2024 10:00"))

	_, err := WorkingDate("35")

	if err == nil {
		t.Errorf("expected error, but got nil")
	}
}
