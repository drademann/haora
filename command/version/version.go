package version

import (
	"github.com/spf13/cobra"
)

const version = "0.2.0-alpha"

var Command = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Haora",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("Haora v%s\n", version)
	},
}
