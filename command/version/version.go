package version

import (
	"github.com/drademann/haora/command"
	"github.com/spf13/cobra"
)

const version = "1.0.0"

func init() {
	command.Root.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Haora",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("Haora version %s\n", version)
	},
}
