package cmd

import (
	"bytes"
	"testing"
)

func TestVersionCmd(t *testing.T) {
	out := bytes.Buffer{}
	rootCmd.SetOut(&out)
	rootCmd.SetArgs([]string{"version"})

	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "Haora version 1.0.0\n"
	if out.String() != expected {
		t.Errorf("expected output %q but got %q", expected, out.String())
	}
}
