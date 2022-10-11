package list

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const tasksDir = "gtask_backup"

type Task struct {
	Id      string
	Content string
}

func NewCmdList() *cobra.Command {
	var (
		All       bool
		Completed bool
		Pending   bool
	)

	cmd := &cobra.Command{
		Use:   "list <command>",
		Short: "List all tasks",
		Long:  `Todo list`,

		Example: `  $ gtask list all
  $ gtask list completed
  $ gtask list pending
		`,

		Run: func(cmd *cobra.Command, args []string) {
			if All {
				tasks, _ := readTasksFromFile(os.DirFS(tasksDir))

				fmt.Println()
				for _, t := range tasks {
					fmt.Printf("%s: %s\n", t.Id, t.Content)
				}

				return
			}

			if Completed {
				fmt.Println("List of completed tasks")
				return
			}

			if Pending {
				fmt.Println("List of pending tasks")
				return
			}

			cmd.Help()
		},
	}

	cmd.PersistentFlags().BoolVarP(&All, "all", "a", false, "List all tasks")
	cmd.PersistentFlags().BoolVarP(&Completed, "completed", "c", false, "List completed tasks")
	cmd.PersistentFlags().BoolVarP(&Pending, "pending", "p", false, "List pending tasks")

	return cmd
}

func readTasksFromFile(fileSystem fs.FS) ([]Task, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}

	task, err := getTask(fileSystem, dir[0])
	if err != nil {
		return nil, err
	}

	return task, nil
}

// TODO: just open a file, change name to something more descriptive
func getTask(fileSystem fs.FS, f fs.DirEntry) ([]Task, error) {
	postFile, err := fileSystem.Open(f.Name())
	if err != nil {
		return nil, err
	}
	defer postFile.Close()

	return newTask(postFile)
}

// TODO: same, change name to some more descriptive
func newTask(postBody io.Reader) ([]Task, error) {
	scanner := bufio.NewScanner(postBody)

	var tasks []Task

	for scanner.Scan() {
		task := strings.Split(scanner.Text(), ". ")
		tasks = append(tasks, Task{Id: task[0], Content: task[1]})
	}

	return tasks, nil
}
