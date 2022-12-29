package list

import (
	"errors"
	"fmt"
	"os"

	"github.com/Patrick564/todo-cli-app/pkg/cmdutil"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	All       bool
	Completed bool
	Pending   bool
	Verbose   bool
	File      string
}

func NewCmdList() *cobra.Command {
	opts := ListOptions{}

	cmd := &cobra.Command{
		Use:   "list <command>",
		Short: "List all tasks",
		Long:  `Todo list`,

		Example: `  $ gtask list all
  $ gtask list completed
  $ gtask list pending
		`,

		RunE: func(_ *cobra.Command, _ []string) error {
			if opts.Completed {
				opts.File = "completed"
			}

			if opts.Pending {
				opts.File = "pending"
			}

			if opts.All || opts.File == "" {
				opts.File = "all"
			}

			return runList(opts)
		},
	}

	cmd.PersistentFlags().BoolVarP(&opts.All, "all", "a", false, "List all tasks")
	cmd.PersistentFlags().BoolVarP(&opts.Completed, "completed", "c", false, "List completed tasks")
	cmd.PersistentFlags().BoolVarP(&opts.Pending, "pending", "p", false, "List pending tasks")
	cmd.PersistentFlags().BoolVarP(&opts.Verbose, "verbose", "v", false, "List tasks with all information")

	return cmd
}

func runList(opts ListOptions) error {
	tasks, err := cmdutil.GetTaskList(os.DirFS(cmdutil.TasksDir), opts.File)
	if err != nil {
		if errors.Is(err, cmdutil.ErrFileEmpty) || errors.Is(err, cmdutil.ErrFileNotFound) {
			fmt.Println("No tasks found, create new with 'gtask add <...>'.")
			return nil
		}

		return err
	}

	return listTasks(tasks, opts.Verbose)
}

func listTasks(tasks []*cmdutil.Task, v bool) error {
	format := func(t cmdutil.Task, v bool) string {
		if v {
			return t.ToString()
		}

		return t.Content
	}

	fmt.Println("All tasks: ")
	for idx, t := range tasks {
		if idx == 0 {
			fmt.Println()
		}
		fmt.Println(format(*t, v))
	}

	return nil
}
