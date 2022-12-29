package cmdutil

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func CheckDatabaseDir(db *sql.DB, fileName string) error {
	_, err := os.Stat(fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			_, err = db.Exec("CREATE TABLE IF NOT EXISTS task (id INTEGER PRIMARY KEY AUTOINCREMENT, content TEXT NOT NULL, created DATETIME DEFAULT CURRENT_TIMESTAMP)")
			if err != nil {
				return err
			}

			return nil
		}

		return err
	}

	return nil
}

func CheckTasksDir() error {
	if !exists(TasksDir) {
		err := os.Mkdir(TasksDir, 0755)
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func exists(file string) bool {
	_, err := os.Stat(file)

	return err == nil
}

func GetTaskList(fileSystem fs.FS, filename string) ([]*Task, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}

	fileName := strings.Join([]string{filename, "md"}, ".")
	var entry fs.DirEntry

	for _, d := range dir {
		if fileName == d.Name() {
			entry = d
		}
	}

	if entry == nil {
		return nil, ErrFileNotFound
	}

	postFile, err := fileSystem.Open(entry.Name())
	if err != nil {
		return nil, err
	}
	defer postFile.Close()

	return readFile(postFile)
}

func RemoveTask(fileSystem fs.FS, filename string, id string) error {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return err
	}

	fn := strings.Join([]string{filename, "md"}, ".")
	var entry fs.DirEntry

	for _, d := range dir {
		if fn == d.Name() {
			entry = d
		}
	}

	if entry == nil {
		return ErrFileNotFound
	}

	postFile, err := fileSystem.Open(entry.Name())
	if err != nil {
		return err
	}
	defer postFile.Close()

	return readFileById(postFile, id)
}

func readFile(postFile fs.File) ([]*Task, error) {
	scanner := bufio.NewScanner(postFile)

	var tasks []*Task

	for scanner.Scan() {
		t, err := NewTaskFromLine(scanner.Text())
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, t)
	}

	if tasks == nil {
		return nil, ErrFileEmpty
	}

	return tasks, nil
}

func readFileById(postFile fs.File, id string) error {
	temp, err := os.CreateTemp(TasksDir, TasksTempDir)
	if err != nil {
		return err
	}
	defer temp.Close()

	scanner := bufio.NewScanner(postFile)
	for scanner.Scan() {
		task, err := NewTaskFromLine(scanner.Text())
		if err != nil {
			return err
		}

		if task.Id != id {
			fmtTask := fmt.Sprintf("%s\n", task.ToString())
			_, err = temp.Write([]byte(fmtTask))
			if err != nil {
				return err
			}
		}

		continue
	}

	err = os.Rename(temp.Name(), TasksAddFile)
	if err != nil {
		return err
	}

	return nil
}
