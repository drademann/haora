package cmd

import (
	"flag"
	"fmt"
)

const version = "1.0.0"

var VersionCmd = flag.NewFlagSet("version", flag.ExitOnError)

func ExecVersion() error {
	_, err := fmt.Fprintf(Config.Out, "haora version %s\n", version)
	if err != nil {
		return err
	}
	return nil
}
