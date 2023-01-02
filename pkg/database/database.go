package database

import (
	"database/sql"
	"strings"

	"github.com/Patrick564/todo-cli-app/pkg/cmdutil"
	_ "github.com/mattn/go-sqlite3"
)

type Conn struct {
	DB *sql.DB
}

func (c *Conn) RestoreSchema() error {
	_, err := c.DB.Exec(`
		CREATE TABLE IF NOT EXISTS task (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			content TEXT NOT NULL,
			complete INTEGER DEFAULT 0,
			created DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return err
	}

	return nil
}

func New(source string) (Conn, error) {
	db, err := sql.Open(cmdutil.DriverName, source)
	if err != nil {
		return Conn{}, err
	}

	return Conn{DB: db}, nil
}

func AddTask(db *sql.DB, content []string) error {
	task := strings.Join(content, " ")

	_, err := db.Exec("INSERT INTO task (content) VALUES ($1)", task)
	if err != nil {
		return err
	}

	return nil
}

func RemoveTask(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM task WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func AllTasks(db *sql.DB) ([]cmdutil.Task, error) {
	rows, err := db.Query("SELECT * FROM task")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]cmdutil.Task, 0)

	for rows.Next() {
		task := cmdutil.Task{}

		err = rows.Scan(&task.Id, &task.Content, &task.Complete, &task.Date)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}
