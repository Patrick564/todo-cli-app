package cmdutil

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"strings"
)

var (
	ErrFileEmpty    = errors.New("not found any task in file")
	ErrFileNotFound = errors.New("not found file with tasks")
)

type Task struct {
	Id      string
	Content string
}

func GetFile(dir []fs.DirEntry, flag string) (fs.DirEntry, error) {
	cmdFlag := fmt.Sprintf("%s.md", flag)

	for _, d := range dir {
		if cmdFlag == d.Name() {
			return d, nil
		}
	}

	return nil, ErrFileNotFound
}

func GetFileContent(postFile fs.File) ([]Task, error) {
	scanner := bufio.NewScanner(postFile)

	var tasks []Task

	for scanner.Scan() {
		task := strings.Split(scanner.Text(), ". ")
		tasks = append(tasks, Task{Id: task[0], Content: task[1]})
	}

	if tasks == nil {
		return nil, ErrFileEmpty
	}

	return tasks, nil
}
