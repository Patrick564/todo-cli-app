package list

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/Patrick564/todo-cli-app/pkg/cmdutil"
	"github.com/spf13/cobra"
)

// Correct to uppercase
const tasksDir = "gtask_backup"

type ListOptions struct {
	All       bool
	Completed bool
	Pending   bool
	Name      string
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

		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.Completed {
				opts.Name = "completed"
			}

			if opts.Pending {
				opts.Name = "pending"
			}

			if opts.All || opts.Name == "" {
				opts.Name = "all"
			}

			return runList(opts.Name)
		},
	}

	cmd.PersistentFlags().BoolVarP(&opts.All, "all", "a", false, "List all tasks")
	cmd.PersistentFlags().BoolVarP(&opts.Completed, "completed", "c", false, "List completed tasks")
	cmd.PersistentFlags().BoolVarP(&opts.Pending, "pending", "p", false, "List pending tasks")

	return cmd
}

func runList(flag string) error {
	tasks, err := readTasksFromFS(os.DirFS(tasksDir), flag)
	if err != nil {
		if err == cmdutil.ErrFileEmpty {
			fmt.Printf("No %s tasks found.\n", flag)
			return nil
		}

		if err == cmdutil.ErrFileNotFound {
			fmt.Printf("File for %s tasks not found.\n", flag)
			return nil
		}

		return err
	}

	for idx, t := range tasks {
		if idx == 0 {
			fmt.Println()
		}
		fmt.Printf("%s: %s\n", t.Id, t.Content)
	}

	return nil
}

func readTasksFromFS(fileSystem fs.FS, flag string) ([]cmdutil.Task, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}

	entry, err := cmdutil.GetFile(dir, flag)
	if err != nil {
		return nil, err
	}

	task, err := getTasks(fileSystem, entry)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func getTasks(fileSystem fs.FS, f fs.DirEntry) ([]cmdutil.Task, error) {
	postFile, err := fileSystem.Open(f.Name())
	if err != nil {
		return nil, err
	}
	defer postFile.Close()

	return cmdutil.GetFileContent(postFile)
}
