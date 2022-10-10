package add

import (
	"fmt"

	"github.com/spf13/cobra"
)

// TODO: create a raw editor for
// create many tasks at same time.
func NewCmdAdd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [flags] [<content>]",
		Short: "Add a new task",
		Long:  `Todo`,

		Args:    cobra.ExactArgs(1),
		Example: `  todo`,

		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				fmt.Printf("Added task: %s\n", args[0])
				return
			}
		},
	}

	return cmd
}
