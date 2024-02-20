package cmd_test

import (
	"bytes"
	"testing"

	"haora/cmd"
)

func TestExecVersion(t *testing.T) {
	out := bytes.Buffer{}
	cmd.Config = cmd.Configuration{Out: &out, OutErr: &bytes.Buffer{}}

	err := cmd.ExecVersion()

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	expected := "haora version 1.0.0\n"
	if out.String() != expected {
		t.Errorf("expected output %q but got %q", expected, out.String())
	}
}
