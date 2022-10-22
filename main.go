package main

import (
	"os"

	"github.com/Patrick564/todo-cli-app/pkg/cmd/root"
)

func main() {
	rootCmd := root.NewCmdRoot()
	code := 0

	if err := rootCmd.Execute(); err != nil {
		code = 0
	}

	os.Exit(code)
}
