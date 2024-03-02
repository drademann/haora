package test

import (
	"bytes"
	"github.com/spf13/cobra"
	"strings"
	"testing"
)

func ExecuteCommand(t *testing.T, cmd *cobra.Command, argLine string) *bytes.Buffer {
	t.Helper()
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs(strings.Split(argLine, " "))
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}
	return buf
}
