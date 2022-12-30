package list

import (
	"database/sql"
	"fmt"

	"github.com/Patrick564/todo-cli-app/pkg/database"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	All       bool
	Completed bool
	Pending   bool
	Verbose   bool
	File      string
}

func NewCmdList(db *sql.DB) *cobra.Command {
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
			tasks, err := database.AllTasks(db)
			if err != nil {
				return err
			}

			for _, t := range tasks {
				fmt.Println(t.Id, t.Content, t.Date)
			}

			return nil
		},
	}

	cmd.PersistentFlags().BoolVarP(&opts.All, "all", "a", false, "List all tasks")
	cmd.PersistentFlags().BoolVarP(&opts.Completed, "completed", "c", false, "List completed tasks")
	cmd.PersistentFlags().BoolVarP(&opts.Pending, "pending", "p", false, "List pending tasks")
	cmd.PersistentFlags().BoolVarP(&opts.Verbose, "verbose", "v", false, "List tasks with all information")

	return cmd
}
