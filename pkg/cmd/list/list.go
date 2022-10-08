package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCmdList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list <command>",
		Short: "List all tasks",
		Long:  `Todo list`,
		Example: `  $ gtask list all
  $ gtask list completed
  $ gtask list pending
		`,
	}

	cmd.AddCommand(NewCmdAll())
	cmd.AddCommand(NewCmdCompleted())
	cmd.AddCommand(NewCmdPending())

	return cmd
}

func NewCmdAll() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all",
		Short: "List all tasks",
		Long:  `todo`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("List of all tasks")
		},
	}

	return cmd
}

func NewCmdCompleted() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "completed",
		Short: "List all tasks pendings",
		Long:  `todo`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Tasks completed")
		},
	}

	return cmd
}

func NewCmdPending() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pending",
		Short: "List all tasks pendings",
		Long:  `todo`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Tasks pending")
		},
	}

	return cmd
}
