package main

import (
	"log"
	"os"

	"github.com/Patrick564/todo-cli-app/pkg/cmd/root"
	"github.com/Patrick564/todo-cli-app/pkg/database"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	rootCmd := root.NewCmdRoot(db)
	code := 0

	if err := rootCmd.Execute(); err != nil {
		code = 0
	}

	os.Exit(code)
}
