package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"haora/app"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the recorded tasks of the selected day",
	Run: func(cmd *cobra.Command, args []string) {
		now := app.Now()
		day := app.Data.Day(now)

		if day.HasNoTasks() {
			fmt.Fprintln(cmd.OutOrStdout(), "no tasks recorded for today")
			return
		}

		for _, task := range day.Tasks {
			fmt.Fprintf(cmd.OutOrStdout(), "%v - ... %v\n", task.Start.Format("15:04"), task.Text)
		}
	},
}
