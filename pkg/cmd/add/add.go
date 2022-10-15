package add

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// TODO: create a raw editor for create many tasks at same time.
func NewCmdAdd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [flags] [<content>]",
		Short: "Add a new task",
		Long:  `Todo`,

		Args:    cobra.MinimumNArgs(1),
		Example: `  todo`,

		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				err := writeTask(args[0])
				if err != nil {
					return err
				}

				fmt.Printf("Added task: %s\n", args[0])
			}

			return nil
		},
	}

	return cmd
}

func writeTask(s string) error {
	task := fmt.Sprintf("\n%s", s)

	f, err := os.OpenFile("gtask_backup/pending.md", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write([]byte(task))
	if err != nil {
		return err
	}

	return nil
}
