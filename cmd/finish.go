package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"time"
)

var (
	endFlag string
)

func init() {
	finishCmd.Flags().StringVarP(&endFlag, "end", "e", "", "The finish time, like 17:00, for the day")
	rootCmd.AddCommand(finishCmd)
}

var finishCmd = &cobra.Command{
	Use:   "finish",
	Short: "Mark the day as done",
	RunE: func(cmd *cobra.Command, args []string) error {
		time, err := parseEndTime(args)
		if err != nil {
			return err
		}
		day := ctx.workingDay()
		day.finish(time)
		cmd.Printf("Today finished at %v\n", day.finished.Format("15:04"))
		return nil
	},
}

func parseEndTime(args []string) (time.Time, error) {
	if endFlag != "" {
		if endFlag == "now" {
			return now(), nil
		}
		t, err := time.Parse("15:04", endFlag)
		if err != nil {
			return time.Time{}, err
		}
		return t, nil
	}
	if len(args) == 1 {
		if args[0] == "now" {
			return now(), nil
		}
		t, err := time.Parse("15:04", args[0])
		if err != nil {
			return time.Time{}, err
		}
		return t, nil
	}
	return time.Time{}, errors.New("no day finish time found")
}
