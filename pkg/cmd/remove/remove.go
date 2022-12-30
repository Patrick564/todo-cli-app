package remove

import (
	"database/sql"
	"strconv"

	"github.com/Patrick564/todo-cli-app/pkg/database"
	"github.com/spf13/cobra"
)

func NewCmdRemove(db *sql.DB) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove [flags]",
		Short: "Remove an existent task",
		Long:  `Todo`,

		Args:    cobra.ExactArgs(1),
		Example: `  todo`,

		RunE: func(_ *cobra.Command, args []string) error {
			if len(args) > 0 {
				i, err := strconv.Atoi(args[0])
				if err != nil {
					return err
				}

				return database.RemoveTask(db, i)
			}

			return nil
		},
	}

	return cmd
}
