package root

import (
	"os"

	addCmd "github.com/Patrick564/todo-cli-app/pkg/cmd/add"
	listCmd "github.com/Patrick564/todo-cli-app/pkg/cmd/list"
	removeCmd "github.com/Patrick564/todo-cli-app/pkg/cmd/remove"
	versionCmd "github.com/Patrick564/todo-cli-app/pkg/cmd/version"

	"github.com/Patrick564/todo-cli-app/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "gtask <command> <subcommand>",
		Short:   "CLI task manager made with Go and Sqlite",
		Long:    "This app was create for practice how to create CLI apps with Go, you can create new tasks, mark as complete and delete.",
		Version: versionCmd.Version,

		Example: `  $ gtask list completed
  $ gtask add -f your_task
  $ gtask remove id_task
		`,

		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return cmdutil.CheckTasksDir()
		},
	}

	// Flags
	cmd.SetHelpFunc(func(c *cobra.Command, _ []string) {
		rootHelp(os.Stdout, c)
	})

	// Child commands
	cmd.AddCommand(versionCmd.NewCmdVersion())
	cmd.AddCommand(listCmd.NewCmdList())
	cmd.AddCommand(addCmd.NewCmdAdd())
	cmd.AddCommand(removeCmd.NewCmdRemove())

	return cmd
}
