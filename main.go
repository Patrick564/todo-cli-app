package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Patrick564/todo-cli-app/pkg/cmd/root"
	"github.com/Patrick564/todo-cli-app/pkg/cmdutil"
	"github.com/Patrick564/todo-cli-app/pkg/database"
)

const dbFile string = "tasks.db"

func main() {
	path := os.Getenv("FILE_PATH")

	conn, err := database.New(filepath.Join(path, dbFile))
	if err != nil {
		log.Fatal(err)
	}

	needSchm, err := cmdutil.NeedsRestoreSchema(path, dbFile)
	if err != nil {
		log.Fatal(err)
	}

	if needSchm {
		conn.RestoreSchema()
	}

	rootCmd := root.NewCmdRoot(conn.DB)
	code := 0

	if err := rootCmd.Execute(); err != nil {
		code = 0
	}

	os.Exit(code)
}
