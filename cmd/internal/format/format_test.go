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

package format

import (
	"testing"
	"time"
)

func TestFormatDuration(t *testing.T) {
	testCases := []struct {
		expected string
		duration time.Duration
		name     string
	}{
		{"     0m", 0 * time.Second, "no duration"},
		{"     0m", 59 * time.Second, "almost a minute"},
		{"     1m", 1 * time.Minute, "1 minute padded"},
		{"    59m", 59 * time.Minute, "59 minutes padded"},
		{" 1h  0m", 1 * time.Hour, "1 hour padded"},
		{"10h 42m", 10*time.Hour + 42*time.Minute, "a long time"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := Duration(tc.duration)
			if output != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, output)
			}
		})
	}
}

func TestFormatDurationShort(t *testing.T) {
	testCases := []struct {
		expected string
		duration time.Duration
		name     string
	}{
		{"", 0 * time.Second, "no duration"},
		{"", 59 * time.Second, "almost a minute"},
		{" 1m", 1 * time.Minute, "1 minute"},
		{"59m", 59 * time.Minute, "59 minutes"},
		{" 1h", 1 * time.Hour, "1 hour"},
		{"10h 42m", 10*time.Hour + 42*time.Minute, "a long time"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := DurationShort(tc.duration)
			if output != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, output)
			}
		})
	}
}

func TestFormatDurationDecimal(t *testing.T) {
	testCases := []struct {
		expected string
		duration time.Duration
	}{
		{" 0.00h", 0 * time.Second},
		{" 0.02h", 59 * time.Second},
		{" 0.17h", 10 * time.Minute},
		{" 0.50h", 30 * time.Minute},
		{" 1.00h", 1 * time.Hour},
		{"12.00h", 12 * time.Hour},
	}

	for _, tc := range testCases {
		t.Run(tc.expected, func(t *testing.T) {
			output := DurationDecimal(tc.duration)
			if output != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, output)
			}
		})
	}
}

func TestFormatDurationDecimalRounded(t *testing.T) {
	testCases := []struct {
		expected string
		duration time.Duration
	}{
		{" 0.00h", 0 * time.Second},
		{" 0.00h", 59 * time.Second},
		{" 0.25h", 10 * time.Minute},
		{" 0.50h", 30 * time.Minute},
		{" 0.75h", 1*time.Hour - 13*time.Minute},
		{"12.50h", 12*time.Hour + 24*time.Minute},
	}

	for _, tc := range testCases {
		t.Run(tc.expected, func(t *testing.T) {
			output := DurationDecimalRounded(tc.duration, 15*time.Minute)
			if output != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, output)
			}
		})
	}
}
