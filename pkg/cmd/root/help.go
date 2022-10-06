package root

import (
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"
)

type helpEntry struct {
	Title string
	Body  string
}

// list, version, add, remove
func rootHelp(w io.Writer, command *cobra.Command) {
	// coreCommands := []string{}
	// actionsCommands := []string{}
	helpEntries := []helpEntry{}
	commandEntries := []string{}

	for _, c := range command.Commands() {
		commandEntries = append(commandEntries, c.Name()+c.Short)
	}

	helpEntries = append(helpEntries, helpEntry{"USAGE", command.Use})
	helpEntries = append(helpEntries, helpEntry{"COMMANDS", strings.Join(commandEntries, "\n")})
	helpEntries = append(helpEntries, helpEntry{"EXAMPLES", command.Example})

	fmt.Fprintln(w, command.Short)

	for _, e := range helpEntries {
		fmt.Fprintln(w, e.Title)
		fmt.Fprintln(w, e.Body)
	}
}

// func commandPad() {
// 	template := fmt.Sprintf("%%-%ds ")
// }
