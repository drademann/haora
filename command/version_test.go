package command

import (
	"github.com/drademann/haora/test"
	"testing"
)

func TestVersionCmd(t *testing.T) {
	out := test.ExecuteCommand(t, Root, "version")

	expected := "Haora v0.1.0-alpha\n"
	if out.String() != expected {
		t.Errorf("expected output %q but got %q", expected, out.String())
	}
}
