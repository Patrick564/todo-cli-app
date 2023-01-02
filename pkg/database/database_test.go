package database

import (
	"os"
	"testing"
)

var db Conn

func mockValues() {
	db.DB.Exec("INSERT INTO task (id, content) VALUES (1, 'Mock task 1')")
	db.DB.Exec("INSERT INTO task (id, content) VALUES (2, 'Mock task 2')")
	db.DB.Exec("INSERT INTO task (id, content) VALUES (3, 'Mock task 3')")
}

func TestAllTasks(t *testing.T) {
	tasks, err := AllTasks(db.DB)
	if err != nil {
		t.Fatal(err)
	}

	if len(tasks) != 3 {
		t.Fatal("incorrect number of tasks")
	}
}

func TestAddTask(t *testing.T) {
	err := AddTask(db.DB, []string{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveTask(t *testing.T) {
	err := RemoveTask(db.DB, 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMain(m *testing.M) {
	db, _ = New(":memory:")
	db.RestoreSchema()

	mockValues()

	exitCode := m.Run()

	os.Exit(exitCode)
}
