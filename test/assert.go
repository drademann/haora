package test

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func AssertOutput(t *testing.T, out *bytes.Buffer, expected string) {
	t.Helper()
	if strings.HasPrefix(expected, "\n") {
		expected = expected[1:]
	}
	expected = strings.ReplaceAll(expected, "\t", "")
	if out.String() != expected {
		t.Errorf("expected output \n%q, but got \n%q", expected, out.String())
	}
}

func AssertDuration(t *testing.T, name string, d, expected time.Duration) {
	t.Helper()
	if d != expected {
		t.Errorf("expected %s duration of %v, but got %v", name, expected, d)
	}
}
