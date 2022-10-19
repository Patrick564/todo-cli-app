package cmdutil

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

// type File interface {
// 	Close()
// 	Open()
// }

type fileReader func(f fs.File) ([]*Task, error)

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

func ReadFromFS(fileSystem fs.FS, name string, reader fileReader) ([]*Task, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}

	f, err := getFile(dir, name)
	if err != nil {
		return nil, err
	}

	t, err := openFile(fileSystem, f, reader)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func getFile(dir []fs.DirEntry, name string) (fs.DirEntry, error) {
	fileName := fmt.Sprintf("%s.md", name)

	for _, d := range dir {
		if fileName == d.Name() {
			return d, nil
		}
	}

	return nil, ErrFileNotFound
}

func openFile(fileSystem fs.FS, f fs.DirEntry, reader fileReader) ([]*Task, error) {
	postFile, err := fileSystem.Open(f.Name())
	if err != nil {
		return nil, err
	}
	defer postFile.Close()

	return reader(postFile)
}

func ReadFile(postFile fs.File) ([]*Task, error) {
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

// Just a test for filter lines by ID
func ReadFileById(postFile fs.File, filter ...string) ([]*Task, error) {
	scanner := bufio.NewScanner(postFile)

	var tasks []*Task

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ": ")

		t, err := NewTaskFromArray(line)
		if err != nil {
			return nil, err
		}

		if len(filter) != 0 {
			if t.Id == filter[0] {
				tasks = append(tasks, t)
				break
			}
		}

		tasks = append(tasks, t)
	}

	if tasks == nil {
		return nil, ErrFileEmpty
	}

	return tasks, nil
}
