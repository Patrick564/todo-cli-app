package list

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/Patrick564/todo-cli-app/pkg/cmdutil"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	All       bool
	Completed bool
	Pending   bool
	Verbose   bool
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
	tasks, err := readTasksFromFS(os.DirFS(cmdutil.TasksDir), opts.Name)
	if err != nil {
		if err == cmdutil.ErrFileEmpty {
			fmt.Println("No tasks found.")
			return nil
		}

		if err == cmdutil.ErrFileNotFound {
			fmt.Printf("File for %s tasks not found.\n", opts.Name)
			return nil
		}

		return err
	}

	for idx, t := range tasks {
		if idx == 0 {
			fmt.Println()
		}

		if opts.Verbose {
			fmt.Printf("%s: %s\n", t.Id, t.Content)
		} else {
			fmt.Printf("%s: %s\n", t.Id, t.Content)
		}
	}

	return nil
}

// func listTasks(tasks []cmdutil.Task, verbose bool) {
// 	fmt.Println()

// 	format := func(t cmdutil.Task, v bool) string {
// 		if v {
// 			return fmt.Sprintf("%s: %s\n", t.Id, t.Content)
// 		}

// 		return fmt.Sprintf("%s\n", t.Content)
// 	}

// 	for _, t := range tasks {
// 		if v {
// 			fmt.Printf("%s: %s\n", t.Id, t.Content)
// 		} else {
// 			fmt.Printf("%s: %s\n", t.Id, t.Content)
// 		}
// 	}
// }

func readTasksFromFS(fileSystem fs.FS, flag string) ([]cmdutil.Task, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}

	f, err := getFile(dir, flag)
	if err != nil {
		return nil, err
	}

	t, err := getTasks(fileSystem, f)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func getFile(dir []fs.DirEntry, flag string) (fs.DirEntry, error) {
	cmdFlag := fmt.Sprintf("%s.md", flag)

	for _, d := range dir {
		if cmdFlag == d.Name() {
			return d, nil
		}
	}

	return nil, cmdutil.ErrFileNotFound
}

func getTasks(fileSystem fs.FS, f fs.DirEntry) ([]cmdutil.Task, error) {
	postFile, err := fileSystem.Open(f.Name())
	if err != nil {
		return nil, err
	}
	defer postFile.Close()

	return getFileContent(postFile)
}

func getFileContent(postFile fs.File) ([]cmdutil.Task, error) {
	scanner := bufio.NewScanner(postFile)

	var tasks []cmdutil.Task

	for scanner.Scan() {
		task := strings.Split(scanner.Text(), ": ")

		if len(task) == 0 {
			return nil, cmdutil.ErrEmptyLineFound
		}

		tasks = append(tasks, cmdutil.Task{Id: task[0], Content: task[1]})
	}

	if tasks == nil {
		return nil, cmdutil.ErrFileEmpty
	}

	return tasks, nil
}
