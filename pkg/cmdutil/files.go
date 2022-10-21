package cmdutil

import (
	"bufio"
	"io/fs"
	"os"
	"strings"
)

const taskRawSep string = ": "

// type fileReader interface {
// 	func(fs.FS) ([]*Task, error)
// 	readFileById(postFile fs.File, id string) error
// }

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
		line := strings.Split(scanner.Text(), ": ")

		t, err := NewTaskFromArray(line)
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
	temp, err := os.CreateTemp("gtask_backup", "gtask_temp")
	if err != nil {
		return err
	}
	defer temp.Close()

	scanner := bufio.NewScanner(postFile)
	writer := bufio.NewWriter(temp)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), taskRawSep)
		if line[0] == id {
			continue
		}
		writer.Write([]byte(scanner.Text()))
	}

	err = os.Rename(temp.Name(), "gtask_backup/all.md")
	if err != nil {
		return err
	}

	return nil
}
