package root

import (
	"github.com/drademann/haora/app"
	"github.com/drademann/haora/app/data"
	"github.com/spf13/cobra"
	"os"
)

var (
	workingDateFlag string
)

var Command = &cobra.Command{
	Use:   "haora",
	Short: "Time tracking with Haora",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		workingDate, err := parseDateFlag(workingDateFlag)
		if err != nil {
			return err
		}
		data.InitState(workingDate)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("The Haora CLI. Choose your command ...")
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) { // reset flag so tests can rerun!
		workingDateFlag = ""
	},
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	Command.PersistentFlags().StringVarP(&workingDateFlag, "date", "d", "", "Date for the command to execute on (defaults to today)")
}

func Execute() {
	var err error
	if err = app.Load(); err != nil {
		Command.PrintErrf("failed to load app data: %v\n", err)
		os.Exit(1)
	}

	if err = Command.Execute(); err != nil {
		Command.PrintErrf("error: %v\n\n", err)
		os.Exit(1)
	}

	if err = app.Save(); err != nil {
		Command.PrintErrf("failed to save app data: %v\n", err)
		os.Exit(1)
	}
}
