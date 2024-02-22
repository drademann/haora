package cmd

import (
	"flag"
	"fmt"
	"haora/app"
)

var ListCmd = flag.NewFlagSet("list", flag.ExitOnError)

func ExecListCmd() error {
	now := app.Now()
	day := app.Data.Day(now)

	if day.HasNoTasks() {
		fmt.Fprintln(Config.Out, "no tasks recorded for today")
		return nil
	}

	for _, task := range day.Tasks {
		fmt.Fprintf(Config.Out, "%v - ... %v\n", task.Start.Format("15:04"), task.Text)
	}
	return nil
}
