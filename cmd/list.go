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
		if day.isEmpty() {
			cmd.Println(noTasks())
			return
		}

		for _, task := range day.tasks {
			start := task.start.Format("15:04")
			end := " now "
			succ, err := day.succ(task)
			if err == nil {
				end = succ.start.Format("15:04")
			}
			dur := formatDuration(day.taskDuration(task))
			cmd.Printf("%v - %v   %v   %v%v\n", start, end, dur, task.text, tags(task.tags))
		}
		cmd.Println()
		f := "%23s\n"
		totalStr := fmt.Sprintf("total %v", formatDuration(day.totalDuration()))
		cmd.Printf(f, totalStr)
		totalBreakStr := fmt.Sprintf("breaks %v", formatDuration(day.totalBreakDuration()))
		cmd.Printf(f, totalBreakStr)
		totalWorkStr := fmt.Sprintf("worked %v", formatDuration(day.totalWorkDuration()))
		cmd.Printf(f, totalWorkStr)
		for _, tag := range day.tags() {
			tagStr := fmt.Sprintf("on %v %v", tag, formatDuration(day.totalTagDuration(tag)))
			cmd.Printf(f, tagStr)
		}
	},
}

func header(day day) string {
	ds := day.date.Format("02.01.2006 (Mon)")
	if day.isToday() {
		return fmt.Sprintf("Tasks for today, %s\n", ds)
	}
	return fmt.Sprintf("Tasks for %s\n", ds)
}

func noTasks() string {
	return "no tasks recorded"
}

func tags(tags []string) string {
	hashed := make([]string, len(tags))
	for i, tag := range tags {
		hashed[i] = "#" + tag
	}
	if len(hashed) == 0 {
		return ""
	}
	return " " + strings.Join(hashed, " ")
}
