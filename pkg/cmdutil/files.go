package cmdutil

import (
	"bufio"
	"io/fs"
	"os"
	"strings"
)

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
	file, err := ReadFromFS(fileSystem, filename)
	if err != nil {
		return nil, err
	}

	return ReadFile(file)
}

func GetTaskById(fileSystem fs.FS, name string, id string) (*Task, error) {
	_, err := ReadFromFS(fileSystem, name)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func ReadFromFS(fileSystem fs.FS, filename string) (fs.File, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}

	return getFile(fileSystem, dir, filename)
}

func getFile(fileSystem fs.FS, dir []fs.DirEntry, filename string) (fs.File, error) {
	fileName := strings.Join([]string{filename, "md"}, ".")

	for _, d := range dir {
		if fileName == d.Name() {
			return openFile(fileSystem, d)
		}
	}

	return nil, ErrFileNotFound
}

func openFile(fileSystem fs.FS, f fs.DirEntry) (fs.File, error) {
	postFile, err := fileSystem.Open(f.Name())
	if err != nil {
		return nil, err
	}
	defer postFile.Close()

	return postFile, nil
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
