package list

import (
	"database/sql"
	"fmt"

	"github.com/Patrick564/todo-cli-app/pkg/database"
	"github.com/spf13/cobra"
)

const formatLayout string = "2006-01-02 03:04:05"

func runList(db *sql.DB) error {
	tasks, err := database.AllTasks(db)
	if err != nil {
		return err
	}

	fmt.Print("Pending tasks:\n\n")
	for _, t := range tasks {
		if t.Complete == 0 {
			fmt.Printf(
				"     %-3d: %-30s %s\n",
				t.Id,
				t.Content,
				t.Date.Format(formatLayout),
			)
		}
	}

	return nil
}

func NewCmdList(db *sql.DB) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list <command>",
		Short:   "List all tasks",
		Long:    `Todo list`,
		Aliases: []string{"l"},

		Example: "  $ gtask list all\n  $ gtask list completed\n  $ gtask list pending",

		RunE: func(_ *cobra.Command, _ []string) error {
			return runList(db)
		},
	}

	return cmd
}
