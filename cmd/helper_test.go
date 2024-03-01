package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func executeCommand(t *testing.T, argLine string) *bytes.Buffer {
	t.Helper()
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs(strings.Split(argLine, " "))
	if err := rootCmd.Execute(); err != nil {
		t.Fatal(err)
	}
	return buf
}
