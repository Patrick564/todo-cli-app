package cmdutil

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	TasksDir           string = "gtask_backup"
	TasksAddFile       string = "gtask_backup/all.md"
	TasksPendingFile   string = "gtask_backup/pending.md"
	TasksCompletedFile string = "gtask_backup/completed.md"
)

type Task struct {
	Id      string
	Content string
}

func (t *Task) ToString() string {
	return fmt.Sprintf("%s: %s", t.Id, t.Content)
}

func NewTask(content string) Task {
	id := uuid.New()

	return Task{Id: id.String(), Content: content}
}

func NewTaskFromArray(line []string) (*Task, error) {
	if len(line) == 0 {
		return nil, ErrEmptyLineFound
	}

	return &Task{Id: line[0], Content: line[1]}, nil
}
