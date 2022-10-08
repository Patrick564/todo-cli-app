package root

import (
	"os"

	listCmd "github.com/Patrick564/todo-cli-app/pkg/cmd/list"
	versionCmd "github.com/Patrick564/todo-cli-app/pkg/cmd/version"
	"github.com/spf13/cobra"
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gtask <command> <subcommand>",
		Short: "A simple CLI task manager",
		Long:  "Todo: under construction",

		Example: `  $ gtask list completed
  $ gtask add -f your_task
  $ gtask remove id_task
		`,
	}

	cmd.Flags().Bool("version", false, "Show gtask version")
	cmd.SetHelpFunc(func(c *cobra.Command, args []string) {
		rootHelp(os.Stdout, c)
	})

	// Child commands
	cmd.AddCommand(versionCmd.NewCmdVersion(os.Stdout))
	cmd.AddCommand(listCmd.NewCmdList())

	return cmd
}
