package cmd

import (
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

		if day.IsEmpty() {
			cmd.Println("no tasks recorded for today")
			return
		}

		for _, task := range day.Tasks {
			cmd.Printf("%v - ... %v\n", task.Start.Format("15:04"), task.Text)
		}
	},
}
