package add

import (
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/command/internal/parsing"
	"github.com/drademann/haora/command/root"
	"github.com/spf13/cobra"
	"strings"
)

var (
	startFlag  string
	tagsFlag   string
	noTagsFlag bool
)

func init() {
	Command.Flags().StringVarP(&startFlag, "start", "s", "", "The start time, like 10:00, of the task")
	Command.Flags().StringVarP(&tagsFlag, "tags", "t", "", "The comma separated tags of the task")
	Command.Flags().BoolVar(&noTagsFlag, "no-tags", false, "Set when the new task has no tags")
	root.Command.AddCommand(Command)
}

var Command = &cobra.Command{
	Use:   "add",
	Short: "Add a task to a day",
	RunE: func(cmd *cobra.Command, args []string) error {
		time, args, err := parsing.Time(startFlag, args)
		if err != nil {
			return err
		}
		var tags []string
		if !noTagsFlag {
			tags, args = parseTags(args)
		}
		text := strings.Join(args, " ")
		d := data.State.WorkingDay()
		return d.AddNewTask(time, text, tags)
	},
	PostRun: func(cmd *cobra.Command, args []string) { // reset flag so tests can rerun!
		startFlag = ""
		tagsFlag = ""
		noTagsFlag = false
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
