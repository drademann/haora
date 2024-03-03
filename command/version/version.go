package version

import (
	"github.com/drademann/haora/command/root"
	"github.com/spf13/cobra"
)

const version = "1.0.0"

func init() {
	root.Command.AddCommand(Command)
}

var Command = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Haora",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("Haora version %s\n", version)
	},
}
