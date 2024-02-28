package test

import (
	"bytes"
	"strings"
	"testing"
)

func AssertOutput(t *testing.T, out *bytes.Buffer, expected string) {
	if strings.HasPrefix(expected, "\n") {
		expected = expected[1:]
	}
	expected = strings.ReplaceAll(expected, "\t", "")
	if out.String() != expected {
		t.Errorf("expected output %q, but got %q", expected, out.String())
	}
}
