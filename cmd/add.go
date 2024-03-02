package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"strings"
	"time"
)

var (
	startFlag string
	tagsFlag  string
)

func init() {
	addCmd.Flags().StringVarP(&startFlag, "start", "s", "", "The start time, like 10:00, of the task")
	addCmd.Flags().StringVarP(&tagsFlag, "tags", "t", "", "The comma separated tags of the task")
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to a day",
	RunE: func(cmd *cobra.Command, args []string) error {
		time, args, err := parseStartTime(args)
		if err != nil {
			return err
		}
		text := strings.Join(args, " ")
		tags, args := parseTags(args)
		d := ctx.workingDay()
		return d.addNewTask(time, text, tags)
	},
}

func parseStartTime(args []string) (time.Time, []string, error) {
	if startFlag != "" {
		if startFlag == "now" {
			return now(), args, nil
		}
		t, err := time.Parse("15:04", startFlag)
		if err != nil {
			return time.Time{}, args, err
		}
		return t, args, nil
	}
	if len(args) > 0 {
		if args[0] == "now" {
			return now(), args[1:], nil
		}
		t, err := time.Parse("15:04", args[0])
		if err != nil {
			return time.Time{}, args, err
		}
		return t, args[1:], nil
	}
	return time.Time{}, args, errors.New("no task starting time found")
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
