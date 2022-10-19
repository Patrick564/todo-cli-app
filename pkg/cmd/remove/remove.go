package remove

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCmdRemove() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove [flags]",
		Short: "Remove an existent task",
		Long:  `Todo`,

		Args:    cobra.ExactArgs(1),
		Example: `  todo`,

		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				fmt.Printf("Arg pass: %s\n", args[0])
				return
			}
		},
	}

	return cmd
}
