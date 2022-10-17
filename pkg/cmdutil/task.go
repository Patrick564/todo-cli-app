package cmdutil

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	TasksDir     = "gtask_backup"
	TasksAddFile = "gtask_backup/all.md"
)

type Task struct {
	Id      string
	Content string
}

func (t *Task) ToString() string {
	return fmt.Sprintf("%s: %s", t.Id, t.Content)
}

func New(content string) *Task {
	id := uuid.New()

	return &Task{Id: id.String(), Content: content}
}
