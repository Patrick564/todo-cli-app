package add

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Patrick564/todo-cli-app/pkg/cmdutil"
	"github.com/spf13/cobra"
)

// TODO: create a raw editor for create many tasks at same time.
func NewCmdAdd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [flags] [<content>]",
		Short: "Add a new task",
		Long:  `Todo`,

		Args:    cobra.MinimumNArgs(1),
		Example: `  todo`,

		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("insufficient arguments")
			}

			return runAdd(args)
		},
	}

	return cmd
}

func runAdd(args []string) error {
	content := strings.Join(args, " ")
	task := cmdutil.NewTask(content)

	err := writeTask(task)
	if err != nil {
		return err
	}

	fmt.Printf("Added task: %s\n", task.ToString())

	return nil
}

func writeTask(t cmdutil.Task) error {
	f, err := os.OpenFile(cmdutil.TasksAddFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	task := fmt.Sprintf("%s\n", t.ToString())

	_, err = f.Write([]byte(task))
	if err != nil {
		return err
	}

	return nil
}
