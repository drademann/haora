package list

import (
	"github.com/drademann/haora/app/data"
	"github.com/spf13/cobra"
	"time"
)

var (
	tagsFlag bool
	weekFlag bool
)

func init() {
	Command.Flags().BoolVar(&tagsFlag, "tags", false, "Show durations per tag")
	Command.Flags().BoolVar(&weekFlag, "week", false, "Show week summary")
}

var Command = &cobra.Command{
	Use:   "list",
	Short: "List the recorded tasks of the selected day",
	RunE: func(cmd *cobra.Command, args []string) error {
		d := data.State.WorkingDay()
		if tagsFlag {
			return printTags(d, cmd)
		}
		if weekFlag {
			return printWeek(d, cmd)
		}
		return printDefault(d, cmd)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		tagsFlag = false
	},
}

func sign(d time.Duration) string {
	if d < 0 {
		return "-"
	}
	return "+"
}
