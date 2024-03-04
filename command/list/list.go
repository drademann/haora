package list

import (
	"errors"
	"fmt"
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/command/internal/format"
	"github.com/spf13/cobra"
	"strings"
)

var Command = &cobra.Command{
	Use:   "list",
	Short: "List the recorded tasks of the selected day",
	RunE: func(cmd *cobra.Command, args []string) error {
		day := data.State.WorkingDay()

		cmd.Println(header(day))
		if day.IsEmpty() {
			cmd.Println(noTasks())
			return nil
		}

		for _, task := range day.Tasks {
			start := task.Start.Format("15:04")
			var end string
			succ, err := day.Succ(task)
			if err == nil {
				end = succ.Start.Format("15:04")
			} else {
				if errors.Is(err, data.NoTaskSucc) && day.IsFinished() {
					end = day.Finished.Format("15:04")
				} else {
					end = " now "
				}
			}
			dur := format.Duration(day.TaskDuration(task))
			if task.IsPause {
				cmd.Printf("      |         %v   %v\n", dur, task.Text)
			} else {
				cmd.Printf("%v - %v   %v   %v%v\n", start, end, dur, task.Text, tags(task.Tags))
			}
		}
		cmd.Println()
		f := "%23s\n"
		totalStr := fmt.Sprintf("total  %v", format.Duration(day.TotalDuration()))
		cmd.Printf(f, totalStr)
		totalBreakStr := fmt.Sprintf("paused  %v", format.Duration(day.TotalBreakDuration()))
		cmd.Printf(f, totalBreakStr)
		totalWorkStr := fmt.Sprintf("worked  %v", format.Duration(day.TotalWorkDuration()))
		cmd.Printf(f, totalWorkStr)
		for _, tag := range day.Tags() {
			tagStr := fmt.Sprintf("on %v  %v", tag, format.Duration(day.TotalTagDuration(tag)))
			cmd.Printf(f, tagStr)
		}
		return nil
	},
}

func header(day data.Day) string {
	ds := day.Date.Format("02.01.2006 (Mon)")
	if day.IsToday() {
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
