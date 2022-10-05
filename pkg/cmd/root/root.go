package root

import (
	"os"

	versionCmd "github.com/Patrick564/todo-cli-app/pkg/cmd/version"
	"github.com/spf13/cobra"
)

// var version = "0.1.0"

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gtask",
		Short: "A simple CLI task manager",
		Long:  "Todo: under construction",

		Example: `
			$ gtask add -f your_task
		`,
	}

	cmd.Flags().Bool("version", false, "Show gtask version")
	cmd.SetHelpFunc(func(c *cobra.Command, args []string) {
		rootHelp(os.Stdout)
	})

	// Child commands
	cmd.AddCommand(versionCmd.NewCmdVersion(os.Stdout))

	return cmd
}
