package cmd_test

import (
	"bytes"
	"testing"

	"haora/cmd"
)

func TestExecListCmd_givenNoDays(t *testing.T) {
	out := bytes.Buffer{}
	cmd.Config = cmd.Configuration{Out: &out, OutErr: &bytes.Buffer{}}

	err := cmd.ExecListCmd()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "no times recorded for today\n"
	if out.String() != expected {
		t.Errorf("expected output %q, but got %q", expected, out.String())
	}
}
