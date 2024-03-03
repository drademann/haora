package version

import (
	"github.com/drademann/haora/command/root"
	"github.com/drademann/haora/test"
	"testing"
)

func TestVersionCmd(t *testing.T) {
	out := test.ExecuteCommand(t, root.Command, "version")

	expected := "Haora version 1.0.0\n"
	if out.String() != expected {
		t.Errorf("expected output %q but got %q", expected, out.String())
	}
}
