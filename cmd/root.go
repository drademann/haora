package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var (
	workingDateFlag string
)

var rootCmd = &cobra.Command{
	Use:   "haora",
	Short: "Time tracking with Haora",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		workingDate, err := parseDateFlag(workingDateFlag)
		if err != nil {
			return err
		}
		initContext(workingDate)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("The Haora CLI. Choose your command ...")
	},
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&workingDateFlag, "date", "d", "", "Date for the command to execute on (defaults to today)")
}

func Execute() {
	var err error
	if err = Load(); err != nil {
		rootCmd.PrintErrf("failed to load app data: %v\n", err)
		os.Exit(1)
	}

	if err = rootCmd.Execute(); err != nil {
		rootCmd.PrintErrf("error: %v\n\n", err)
		os.Exit(1)
	}

	if err = Save(); err != nil {
		rootCmd.PrintErrf("failed to save app data: %v\n", err)
		os.Exit(1)
	}
}
