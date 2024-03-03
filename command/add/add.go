package add

import (
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/command"
	"github.com/drademann/haora/command/internal/parsing"
	"github.com/spf13/cobra"
	"strings"
)

var (
	startFlag string
	tagsFlag  string
)

func init() {
	add.Flags().StringVarP(&startFlag, "start", "s", "", "The start time, like 10:00, of the task")
	add.Flags().StringVarP(&tagsFlag, "tags", "t", "", "The comma separated tags of the task")
	command.Root.AddCommand(add)
}

var add = &cobra.Command{
	Use:   "add",
	Short: "Add a task to a day",
	RunE: func(cmd *cobra.Command, args []string) error {
		time, args, err := parsing.Time(startFlag, args)
		if err != nil {
			return err
		}
		text := strings.Join(args, " ")
		tags, args := parseTags(args)
		d := data.State.WorkingDay()
		return d.AddNewTask(time, text, tags)
	},
	PostRun: func(cmd *cobra.Command, args []string) { // reset flag so tests can rerun!
		startFlag = ""
		tagsFlag = ""
	},
}

func parseTags(args []string) ([]string, []string) {
	if tagsFlag != "" {
		tags := strings.Split(tagsFlag, ",")
		return tags, args
	}
	if len(args) > 0 {
		tags := strings.Split(args[0], ",")
		return tags, args[1:]
	}
	return []string{}, args
}
