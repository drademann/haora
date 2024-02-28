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
		day := app.Data.Day(app.WorkingDate)

		cmd.Println(header(day))
		if day.IsEmpty() {
			cmd.Println(noTasks())
			return
		}

		for _, task := range day.Tasks {
			cmd.Printf("%v - ... %v\n", task.Start.Format("15:04"), task.Text)
		}
	},
}

func header(day app.Day) string {
	ds := day.Date.Format("02.01.2006 (Mon)")
	if day.IsToday() {
		return fmt.Sprintf("Tasks for today, %s\n", ds)
	}
	return fmt.Sprintf("Tasks for %s\n", ds)
}

func noTasks() string {
	return "no tasks recorded"
}
