package root

import (
	"fmt"
	"io"
)

type helpEntry struct {
	Title string
	Body  string
}

func rootHelp(w io.Writer) {
	helpEntries := []helpEntry{}

	helpEntries = append(helpEntries, helpEntry{"USAGE", "  gtask <command>"})

	for _, e := range helpEntries {
		fmt.Fprintln(w, e.Title)
		fmt.Fprintln(w, e.Body)
	}
}
