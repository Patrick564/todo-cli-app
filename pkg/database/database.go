package database

import (
	"database/sql"

	"github.com/Patrick564/todo-cli-app/pkg/cmdutil"
	_ "github.com/mattn/go-sqlite3"
)

const fileName string = "tasks.db"

func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		return nil, err
	}

	err = cmdutil.CheckDatabaseDir(db, fileName)
	if err != nil {
		return nil, err
	}

	return db, nil
}
