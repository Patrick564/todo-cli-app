package main

import (
	"os"

	"github.com/Patrick564/todo-cli-app/pkg/cmd/root"
)

// type exitCode int

// const (
// 	exitOK    exitCode = 0
// 	exitError exitCode = 1
// )

func main() {
	rootCmd := root.NewCmdRoot()
	code := 0

	if err := rootCmd.Execute(); err != nil {
		code = 0
	}

	os.Exit(code)
}
