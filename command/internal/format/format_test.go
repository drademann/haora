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
