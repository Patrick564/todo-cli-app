package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

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
				fmt.Println("List of all tasks")
			}

			if Completed {
				fmt.Println("List of completed tasks")
			}

			if Pending {
				fmt.Println("List of pending tasks")
			}

			cmd.Help()
		},
	}

	cmd.PersistentFlags().BoolVarP(&All, "all", "a", false, "List all tasks")
	cmd.PersistentFlags().BoolVarP(&Completed, "completed", "c", false, "List completed tasks")
	cmd.PersistentFlags().BoolVarP(&Pending, "pending", "p", false, "List pending tasks")

	return cmd
}
