package format

import (
	"testing"
	"time"
)

func TestFormatDuration(t *testing.T) {
	tests := []struct {
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
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := Duration(test.duration)
			if output != test.expected {
				t.Errorf("expected %q, but got %q", test.expected, output)
			}
		})
	}
}
