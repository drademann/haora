package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"haora/app"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "haora",
	Short: "Time tracking with Haora",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return ParseDateFlag()
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintln(cmd.OutOrStdout(), "The Haora CLI. Choose your command ...")
	},
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	workingDateFlag = rootCmd.PersistentFlags().StringP("date", "d", "", "Date for the command to execute on (defaults to today)")
}

func Execute() {
	var err error
	if err = app.Load(); err != nil {
		fmt.Fprintf(rootCmd.OutOrStderr(), "failed to load app data: %v\n", err)
		os.Exit(1)
	}

	if err = rootCmd.Execute(); err != nil {
		fmt.Fprintf(rootCmd.OutOrStderr(), "error: %v\n\n", err)
		os.Exit(1)
	}

	if err = app.Save(); err != nil {
		fmt.Fprintf(rootCmd.OutOrStderr(), "failed to save app data: %v\n", err)
		os.Exit(1)
	}
}
