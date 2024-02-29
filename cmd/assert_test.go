package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func assertOutput(t *testing.T, out *bytes.Buffer, expected string) {
	t.Helper()
	if strings.HasPrefix(expected, "\n") {
		expected = expected[1:]
	}
	expected = strings.ReplaceAll(expected, "\t", "")
	if out.String() != expected {
		t.Errorf("expected output %q, but got %q", expected, out.String())
	}
}
