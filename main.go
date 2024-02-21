package main

import (
	"flag"
	"fmt"
	"os"

	"haora/cmd"
)

func main() {
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("expected at least one subcommand")
		os.Exit(1)
	}

	var err error
	switch os.Args[1] {
	case cmd.VersionCmd.Name():
		err = cmd.ExecVersion()
	default:
		flag.Usage()
	}
	if err != nil {
		_, _ = fmt.Fprintln(cmd.Config.OutErr, err)
		os.Exit(1)
	}
}
