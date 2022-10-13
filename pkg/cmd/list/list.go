package list

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// Correct to uppercase
const tasksDir = "gtask_backup"

var (
	ErrFileNotFound = errors.New("not found file with tasks")
	ErrFileEmpty    = errors.New("not found any task in file")
)

type Task struct {
	Id      string
	Content string
}

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
				opts.Name = "completed.md"
			}

			if opts.Pending {
				opts.Name = "pending.md"
			}

			if opts.All || opts.Name == "" {
				opts.Name = "all.md"
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

func readTasksFromFS(fileSystem fs.FS, flag string) ([]Task, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}

	entry, err := getTasksFile(dir, flag)
	if err != nil {
		return nil, err
	}

	task, err := getTasks(fileSystem, entry)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func getTasksFile(dir []fs.DirEntry, flag string) (fs.DirEntry, error) {
	for _, d := range dir {
		if flag == d.Name() {
			return d, nil
		}
	}

	return nil, ErrFileNotFound
}

func getTasks(fileSystem fs.FS, f fs.DirEntry) ([]Task, error) {
	postFile, err := fileSystem.Open(f.Name())
	if err != nil {
		return nil, err
	}
	defer postFile.Close()

	return newTasks(postFile)
}

func newTasks(postFile fs.File) ([]Task, error) {
	scanner := bufio.NewScanner(postFile)

	var tasks []Task

	for scanner.Scan() {
		task := strings.Split(scanner.Text(), ". ")
		tasks = append(tasks, Task{Id: task[0], Content: task[1]})
	}

	if tasks == nil {
		return nil, ErrFileEmpty
	}

	return tasks, nil
}
