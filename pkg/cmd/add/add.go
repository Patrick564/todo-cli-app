package add

import "github.com/spf13/cobra"

// TODO: create a raw editor for
// create many tasks at same time.
func NewCmdAdd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add [flags]",
		Short:   "Add a new task",
		Long:    `Todo`,
		Example: `  todo`,
	}

	return cmd
}
