package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "list <command>",
	Short: "List of all tasks",
	Long:  "todo: ---",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list executed")
	},
}

func init() {
	// root.AddCommand(versionCmd)
}
