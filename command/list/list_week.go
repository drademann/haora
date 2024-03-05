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
	var total time.Duration
	for i := 0; i < 7; i++ {
		d = data.State.DayList.Day(date)
		dateStr := d.Date.Format("Mon 02.01.2006")
		if d.IsEmpty() {
			cmd.Printf("%s   -\n", dateStr)
		} else {
			startStr := d.Start().Format("15:04")
			endStr := d.End().Format("15:04")
			dur := d.TotalWorkDuration()
			durStr := format.Duration(dur)
			cmd.Printf("%s   %s - %s  worked %s\n", dateStr, startStr, endStr, durStr)
			total += dur
		}
		date = date.Add(24 * time.Hour)
	}
	cmd.Printf("\n                          total worked %s\n", format.Duration(total))
	return nil
}
