package root

import (
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"
)

const bodyPad int = 13

type helpEntry struct {
	Title string
	Body  string
}

// Add spacing at start and between name and short description
func formatCommand(command *cobra.Command) string {
	pad := fmt.Sprintf("  %%-%ds", bodyPad)
	template := fmt.Sprintf(pad, command.Name()+":")

	return template + command.Short
}

// list, version, add, remove
func rootHelp(w io.Writer, command *cobra.Command) {
	helpEntries := []helpEntry{}
	commandEntries := []string{}

	for _, c := range command.Commands() {
		commandEntries = append(commandEntries, formatCommand(c))
	}

	helpEntries = append(helpEntries, helpEntry{"USAGE", "  " + command.UseLine()})
	helpEntries = append(helpEntries, helpEntry{"COMMANDS", strings.Join(commandEntries, "\n")})
	helpEntries = append(helpEntries, helpEntry{"EXAMPLES", command.Example})

	fmt.Fprintln(w, command.Short)

	for _, e := range helpEntries {
		fmt.Fprint(w, "\n")
		fmt.Fprintln(w, e.Title)
		fmt.Fprintln(w, e.Body)
	}
}
