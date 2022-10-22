package remove

import (
	"fmt"
	"os"

	"github.com/Patrick564/todo-cli-app/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdRemove() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove [flags]",
		Short: "Remove an existent task",
		Long:  `Todo`,

		Args:    cobra.ExactArgs(1),
		Example: `  todo`,

		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				return runRemove(args[0])
			}

			return nil
		},
	}

	return cmd
}

func runRemove(id string) error {
	err := cmdutil.RemoveTask(os.DirFS(cmdutil.TasksDir), "all", id)
	if err != nil {
		return err
	}

	fmt.Println("Task remove correctly.")

	return nil
}
