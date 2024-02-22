package cmd

import (
	"flag"
	"fmt"
)

var ListCmd = flag.NewFlagSet("list", flag.ExitOnError)

func ExecListCmd() error {
	fmt.Fprintln(Config.Out, "no times recorded for today")
	return nil
}
