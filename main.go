package main

import (
	"flag"
	"fmt"
	"haora/app"
	"os"

	"haora/cmd"
)

func main() {
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Fprintln(cmd.Config.OutErr, "expected at least one subcommand")
		os.Exit(1)
	}

	var err error
	if err = app.Load(); err != nil {
		fmt.Fprintf(cmd.Config.OutErr, "failed to load app data: %v\n", err)
		os.Exit(1)
	}

	switch os.Args[1] {
	case cmd.VersionCmd.Name():
		err = cmd.ExecVersion()
	default:
		flag.Usage()
	}
	if err != nil {
		fmt.Fprintf(cmd.Config.OutErr, "failed to execute command: %v\n", err)
		os.Exit(1)
	}

	if err = app.Save(); err != nil {
		fmt.Fprintf(cmd.Config.OutErr, "failed to save app data: %v\n", err)
	}
}
