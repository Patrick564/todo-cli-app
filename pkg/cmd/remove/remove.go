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
				return runRemove()
			}

			return nil
		},
	}

	return cmd
}

func runRemove() error {
	dirFS := os.DirFS(cmdutil.TasksDir)
	err := cmdutil.RemoveTask(dirFS, "all", "df6b13dd-e7d8-4d5e-80e6-9edaf5b0c64c")
	if err != nil {
		return err
	}

	fmt.Println("correct delete")

	return nil
}
