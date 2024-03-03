package pause

import (
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/command/internal/parsing"
	"github.com/spf13/cobra"
	"strings"
)

var (
	startFlag string
)

func init() {
	Command.Flags().StringVarP(&startFlag, "start", "s", "", "The start time, like 12:00, of the pause")
}

var Command = &cobra.Command{
	Use:   "pause",
	Short: "Add a pause to a day",
	RunE: func(cmd *cobra.Command, args []string) error {
		time, args, err := parsing.Time(startFlag, args)
		if err != nil {
			return err
		}
		text := strings.Join(args, " ")
		d := data.State.WorkingDay()
		return d.AddNewPause(time, text)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		startFlag = ""
	},
}
