package add

import (
	"database/sql"
	"errors"

	"github.com/Patrick564/todo-cli-app/pkg/database"
	"github.com/spf13/cobra"
)

// TODO: create a raw editor for create many tasks at same time.
func NewCmdAdd(db *sql.DB) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [flags] [<content>]",
		Short: "Add a new task",
		Long:  `Todo`,

		Args:    cobra.MinimumNArgs(1),
		Example: `  todo`,

		RunE: func(_ *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("insufficient arguments")
			}

			return database.AddTask(db, args[0])
		},
	}

	return cmd
}
