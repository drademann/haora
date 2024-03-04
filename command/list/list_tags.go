package list

import (
	"fmt"
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/command/internal/format"
	"github.com/spf13/cobra"
	"time"
)

func printTags(d data.Day, cmd *cobra.Command) error {
	headerStr := func(day data.Day) string {
		ds := day.Date.Format("02.01.2006 (Mon)")
		if day.IsToday() {
			return fmt.Sprintf("Tag summary for today, %s\n", ds)
		}
		return fmt.Sprintf("Tag summary for %s\n", ds)
	}
	cmd.Println(headerStr(d))

	if d.IsEmpty() {
		cmd.Println("no tags found")
		return nil
	}

	for _, tg := range d.Tags() {
		tagDur := d.TotalTagDuration(tg)
		cmd.Printf("%v  %v  %v   #%v\n",
			format.Duration(tagDur),
			format.DurationDecimal(tagDur),
			format.DurationDecimalRounded(tagDur, 15*time.Minute),
			tg,
		)
	}

	return nil
}
