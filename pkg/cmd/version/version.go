package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

const Version string = "gtask version 2.0.0 (2022-12-30)\nFeel free to fork this project https://github.com/Patrick564/todo-cli-app"

func NewCmdVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "version",
		Hidden: true,

		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println(Version)
		},
	}

	return cmd
}
