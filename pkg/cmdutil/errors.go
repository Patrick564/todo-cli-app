package cmdutil

import "errors"

var (
	ErrFileEmpty      = errors.New("not found any task in file")
	ErrFileNotFound   = errors.New("not found file with tasks")
	ErrEmptyLineFound = errors.New("found an empty line instead a task")
)
