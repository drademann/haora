package finish

import (
	"errors"
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/app/datetime"
	"github.com/drademann/haora/command"
	"github.com/spf13/cobra"
	"time"
)

var (
	endFlag string
)

func init() {
	finish.Flags().StringVarP(&endFlag, "end", "e", "", "The finish time, like 17:00, for the day")
	command.Root.AddCommand(finish)
}

var finish = &cobra.Command{
	Use:   "finish",
	Short: "Mark the day as done",
	RunE: func(cmd *cobra.Command, args []string) error {
		time, err := parseEndTime(args)
		if err != nil {
			return err
		}
		day := data.State.WorkingDay()
		day.Finish(time)
		cmd.Printf("Today finished at %v\n", day.Finished.Format("15:04"))
		return nil
	},
	PostRun: func(cmd *cobra.Command, args []string) { // reset flag so tests can rerun!
		endFlag = ""
	},
}

func parseEndTime(args []string) (time.Time, error) {
	if endFlag != "" {
		if endFlag == "now" {
			return datetime.Now(), nil
		}
		t, err := time.Parse("15:04", endFlag)
		if err != nil {
			return time.Time{}, err
		}
		return t, nil
	}
	if len(args) == 1 {
		if args[0] == "now" {
			return datetime.Now(), nil
		}
		t, err := time.Parse("15:04", args[0])
		if err != nil {
			return time.Time{}, err
		}
		return t, nil
	}
	return time.Time{}, errors.New("no day finish time found")
}
