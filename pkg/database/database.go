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

	return nil
}

func AddTask(db *sql.DB, content string) error {
	db.Exec("INSERT INTO task (content) VALUES ($1)", content)

	return nil
}

func RemoveTask(db *sql.DB, id int) error {
	db.Exec("DELETE FROM task WHERE id = $1", id)

	return nil
}

func AllTasks(db *sql.DB) ([]cmdutil.TaskSQL, error) {
	rows, err := db.Query("SELECT * FROM task")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]cmdutil.TaskSQL, 0)

	for rows.Next() {
		task := cmdutil.TaskSQL{}

		err = rows.Scan(&task.Id, &task.Content, &task.Date)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func Connect() (*sql.DB, error) {
	path := os.Getenv("FILE_PATH")

	// Comprove if database file need schema creation
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