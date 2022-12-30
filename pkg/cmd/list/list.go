package list

import (
	"database/sql"
	"fmt"
	"time"

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
			db.Exec("insert into task (content) values ('example1')")
			db.Exec("insert into task (content) values ('example2')")

			rows, err := db.Query("select * from task")
			if err != nil {
				return err
			}
			defer rows.Close()

			var a struct {
				id      int
				content string
				date    time.Time
			}
			for rows.Next() {
				err = rows.Scan(&a.id, &a.content, &a.date)
				if err != nil {
					return err
				}
				fmt.Println(a)
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
