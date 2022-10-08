package version

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

const version = "0.1.0"

func NewCmdVersion(w io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "version",
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(w, "gtaks version %s\n", version)
		},
	}

	return cmd
}
