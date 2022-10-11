package list

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"os"

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
				content, _ := readTasksFromFile(os.DirFS(tasksDir))
				fmt.Printf("content: \n%s", content)

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

	var tasks []Task

	for _, f := range dir {
		task, err := getTask(fileSystem, f)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func getTask(fileSystem fs.FS, f fs.DirEntry) (Task, error) {
	postFile, err := fileSystem.Open(f.Name())
	if err != nil {
		return Task{}, err
	}
	defer postFile.Close()

	return newTask(postFile)
}

func newTask(postBody io.Reader) (Task, error) {
	scanner := bufio.NewScanner(postBody)

	scanner.Scan()
	line := scanner.Text()

	return Task{
		Id:      line[:1],
		Content: line[3:],
	}, nil
}
