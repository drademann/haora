package list

import (
	"errors"
	"fmt"
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/command/internal/format"
	"github.com/spf13/cobra"
	"strings"
)

func printDefault(d data.Day, cmd *cobra.Command) error {
	headerStr := func(day data.Day) string {
		ds := day.Date.Format("02.01.2006 (Mon)")
		if day.IsToday() {
			return fmt.Sprintf("Tasks for today, %s\n", ds)
		}
		return fmt.Sprintf("Tasks for %s\n", ds)
	}
	cmd.Println(headerStr(d))

	if d.IsEmpty() {
		cmd.Println("no tasks recorded")
		return nil
	}

	tagsStr := func(tags []string) string {
		hashed := make([]string, len(tags))
		for i, tag := range tags {
			hashed[i] = "#" + tag
		}
		if len(hashed) == 0 {
			return ""
		}
		return " " + strings.Join(hashed, " ")
	}

	for _, task := range d.Tasks {
		start := task.Start.Format("15:04")
		var end string
		succ, err := d.Succ(task)
		if err == nil {
			end = succ.Start.Format("15:04")
		} else {
			if errors.Is(err, data.NoTaskSucc) && d.IsFinished() {
				end = d.Finished.Format("15:04")
			} else {
				end = " now "
			}
		}
		dur := format.Duration(d.TaskDuration(task))
		if task.IsPause {
			cmd.Printf("      |         %v   %v\n", dur, task.Text)
		} else {
			cmd.Printf("%v - %v   %v   %v%v\n", start, end, dur, task.Text, tagsStr(task.Tags))
		}
	}
	cmd.Println()
	f := "%23s\n"
	totalStr := fmt.Sprintf("total  %v", format.Duration(d.TotalDuration()))
	cmd.Printf(f, totalStr)
	totalBreakStr := fmt.Sprintf("paused  %v", format.Duration(d.TotalBreakDuration()))
	cmd.Printf(f, totalBreakStr)
	totalWorkStr := fmt.Sprintf("worked  %v", format.Duration(d.TotalWorkDuration()))
	cmd.Printf(f, totalWorkStr)
	return nil
}
