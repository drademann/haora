//
// Copyright 2024-2025 The Haora Authors
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

package config

import (
	"github.com/drademann/fugo/test/assert"
	"testing"
	"time"
)

func TestDurationPerDay(t *testing.T) {
	testCases := []struct {
		name            string
		durationPerWeek time.Duration
		daysPerWeek     int
		expected        time.Duration
	}{
		{"40h per week", 40 * time.Hour, 5, 8 * time.Hour},
		{"36h 15m per week", 36*time.Hour + 15*time.Minute, 5, 7*time.Hour + 15*time.Minute},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			durationPerWeek = &tc.durationPerWeek
			daysPerWeek = &tc.daysPerWeek

			dur, exist := DurationPerDay()
			if !exist {
				t.Fatal("duration per day should exist")
			}

			if dur != tc.expected {
				t.Errorf("expected duration per day %v, got %v", tc.expected, dur)
			}
		})
	}
}

func TestHiddenWeekdays(t *testing.T) {
	SetHiddenWeekdays(t, "sa sun")
	InitViper()

	assert.True(t, IsHidden(time.Saturday))
	assert.True(t, IsHidden(time.Sunday))
	assert.False(t, IsHidden(time.Wednesday))
}
