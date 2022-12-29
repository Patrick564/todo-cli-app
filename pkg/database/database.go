package database

import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/Patrick564/todo-cli-app/pkg/cmdutil"
	_ "github.com/mattn/go-sqlite3"
)

const (
	dirName  string = "/tasks"
	fileName string = "/tasks.db"
)

func Connect() (*sql.DB, error) {
	path, err := os.Executable()
	if err != nil {
		return nil, err
	}
	path = filepath.Dir(path)

	err = os.MkdirAll(path+dirName, 0755)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", path+dirName+fileName)
	if err != nil {
		return nil, err
	}

	err = cmdutil.CheckDatabaseDir(db, fileName)
	if err != nil {
		return nil, err
	}

	return db, nil
}
