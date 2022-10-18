package root

import (
	"os"

	addCmd "github.com/Patrick564/todo-cli-app/pkg/cmd/add"
	listCmd "github.com/Patrick564/todo-cli-app/pkg/cmd/list"
	removeCmd "github.com/Patrick564/todo-cli-app/pkg/cmd/remove"
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

	// Flags
	cmd.Flags().Bool("version", false, "Show gtask version")
	cmd.SetHelpFunc(func(c *cobra.Command, args []string) {
		rootHelp(os.Stdout, c)
	})

	// Child commands
	cmd.AddCommand(versionCmd.NewCmdVersion(os.Stdout))
	cmd.AddCommand(listCmd.NewCmdList())
	cmd.AddCommand(addCmd.NewCmdAdd())
	cmd.AddCommand(removeCmd.NewCmdRemove())

	return cmd
}
