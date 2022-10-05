package main

import (
	"fmt"
	"os"

	"github.com/Patrick564/todo-cli-app/pkg/cmd/root"
)

func main() {
	rootCmd := root.NewCmdRoot()

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
