package cmdutil

import "errors"

var (
	ErrFileEmpty      error = errors.New("not found any task in file")
	ErrFileNotFound   error = errors.New("not found file with tasks")
	ErrEmptyLineFound error = errors.New("found an empty line instead a task")
)
