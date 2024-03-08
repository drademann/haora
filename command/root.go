package command

import (
	"github.com/drademann/haora/app/config"
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/command/add"
	"github.com/drademann/haora/command/finish"
	"github.com/drademann/haora/command/list"
	"github.com/drademann/haora/command/pause"
	"github.com/drademann/haora/command/version"
	"github.com/spf13/cobra"
	"os"
)

var (
	workingDateFlag string
)

var Root = &cobra.Command{
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
	Root.PersistentFlags().StringVarP(&workingDateFlag, "date", "d", "", "Date for the command to execute on (defaults to today)")
	Root.AddCommand(add.Command)
	Root.AddCommand(finish.Command)
	Root.AddCommand(list.Command)
	Root.AddCommand(pause.Command)
	Root.AddCommand(version.Command)
}

func Execute() {
	var err error
	if err = config.Load(); err != nil {
		Root.PrintErrf("failed to load app config: %v\n", err)
		os.Exit(1)
	}

	if err = data.Load(); err != nil {
		Root.PrintErrf("failed to load app data: %v\n", err)
		os.Exit(1)
	}

	if err = Root.Execute(); err != nil {
		Root.PrintErrf("error: %v\n\n", err)
		os.Exit(1)
	}

	if err = data.Save(); err != nil {
		Root.PrintErrf("failed to save app data: %v\n", err)
		os.Exit(1)
	}
}
