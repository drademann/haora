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

package app

import (
	"testing"
	"time"
)

func TestDurationPerWeek(t *testing.T) {
	testCases := []struct {
		input    string
		expected time.Duration
	}{
		{"1h 30m", 90 * time.Minute},
		{"1h", 1 * time.Hour},
		{"30m", 30 * time.Minute},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			dur, err := Duration(tc.input)
			if err != nil {
				t.Fatal(err)
			}
			if dur != tc.expected {
				t.Errorf("expected duration %v, got %v", tc.expected, dur)
			}
		})
	}
}

func TestDurationPerDay(t *testing.T) {
	testCases := []struct {
		DurationPerWeek string
		DaysPerWeek     int
		expected        time.Duration
	}{
		{"40h", 5, 8 * time.Hour},
		{"36h 15m", 5, 7*time.Hour + 15*time.Minute},
	}

	for _, tc := range testCases {
		t.Run(tc.DurationPerWeek, func(t *testing.T) {
			Config.Times.DurationPerWeek = tc.DurationPerWeek
			Config.Times.DaysPerWeek = tc.DaysPerWeek

			dur, err := DurationPerDay()
			if err != nil {
				t.Fatal(err)
			}

			if dur != tc.expected {
				t.Errorf("expected duration per day %v, got %v", tc.expected, dur)
			}
		})
	}
}
