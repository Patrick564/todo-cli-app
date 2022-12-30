package database

import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/Patrick564/todo-cli-app/pkg/cmdutil"
	_ "github.com/mattn/go-sqlite3"
)

const (
	taskDir  string = "tasks"
	taskFile string = "tasks.db"
)

func createSchema(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS task (id INTEGER PRIMARY KEY AUTOINCREMENT, content TEXT NOT NULL, created DATETIME DEFAULT CURRENT_TIMESTAMP)")
	if err != nil {
		return err
	}

	_, err = db.Exec("insert into task (content) values ('Example task')")
	if err != nil {
		return err
	}

	return nil
}

func Connect() (*sql.DB, error) {
	path := os.Getenv("PATH")

	// Comprove if database file needSchm
	needSchm, err := cmdutil.NeedSchema(path, taskFile)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", filepath.Join(path, taskFile))
	if err != nil {
		return nil, err
	}

	if needSchm {
		createSchema(db)
	}

	return db, nil
}
