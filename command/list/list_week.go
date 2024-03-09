package list

import (
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/app/datetime"
	"github.com/drademann/haora/command/internal/format"
	"github.com/spf13/cobra"
	"time"
)

func printWeek(d data.Day, cmd *cobra.Command) error {
	date := datetime.FindWeekday(d.Date, datetime.Previous, time.Monday)
	week := data.CollectWeek(date)
	for _, day := range week.Days {
		dateStr := day.Date.Format("Mon 02.01.2006")
		if day.IsEmpty() {
			cmd.Printf("%s   -\n", dateStr)
		} else {
			startStr := day.Start().Format("15:04")
			endStr := day.End().Format("15:04")
			dur := day.TotalWorkDuration()
			durStr := format.Duration(dur)
			cmd.Printf("%s   %s - %s  worked %s\n", dateStr, startStr, endStr, durStr)
		}
	}
	overtime, err := week.TotalOvertimeDuration()
	if err != nil || overtime == 0 {
		cmd.Printf("\n                          total worked %s\n", format.Duration(week.TotalWorkDuration()))
	} else {
		cmd.Printf("\n                          total worked %s   (%s %v)\n", format.Duration(week.TotalWorkDuration()), sign(overtime), format.DurationShort(overtime))
	}
	return nil
}
