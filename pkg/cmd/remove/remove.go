package remove

import "github.com/spf13/cobra"

// TODO: create a raw editor for
// create many tasks at same time.
func NewCmdRemove() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "remove [flags]",
		Short:   "Remove an existent task",
		Long:    `Todo`,
		Example: `  todo`,
	}

	return cmd
}
