package version

import (
	"github.com/drademann/haora/command"
	"github.com/drademann/haora/test"
	"testing"
)

func TestVersionCmd(t *testing.T) {
	out := test.ExecuteCommand(t, command.Root, "version")

	expected := "Haora version 1.0.0\n"
	if out.String() != expected {
		t.Errorf("expected output %q but got %q", expected, out.String())
	}
}
