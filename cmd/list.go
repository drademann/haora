package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the recorded tasks of the selected day",
	Run: func(cmd *cobra.Command, args []string) {
		day := ctx.workingDay()

		cmd.Println(header(day))
		if day.IsEmpty() {
			cmd.Println(noTasks())
			return
		}

		for _, task := range day.Tasks {
			start := task.Start.Format("15:04")
			end := "now  "
			succ, err := day.succ(task)
			if err == nil {
				end = succ.Start.Format("15:04")
			}
			dur := formatDuration(day.duration(task))
			cmd.Printf("%v - %v   %v   %v   %v\n", start, end, dur, strings.Join(task.Tags, ","), task.Text)
		}
	},
}

func header(day Day) string {
	ds := day.Date.Format("02.01.2006 (Mon)")
	if day.IsToday() {
		return fmt.Sprintf("Tasks for today, %s\n", ds)
	}
	return fmt.Sprintf("Tasks for %s\n", ds)
}

func noTasks() string {
	return "no tasks recorded"
}
