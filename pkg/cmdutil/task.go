package cmdutil

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

const (
	TasksDir           string = "gtask_backup"
	TasksTempDir       string = "gtask_temp_"
	TasksAddFile       string = "gtask_backup/all.md"
	TasksPendingFile   string = "gtask_backup/pending.md"
	TasksCompletedFile string = "gtask_backup/completed.md"
)

const taskRawSep string = ": "

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

func NewTaskFromLine(line string) (*Task, error) {
	rawTask := strings.Split(line, taskRawSep)
	if len(rawTask) == 0 {
		return nil, ErrEmptyLineFound
	}

	return &Task{Id: rawTask[0], Content: rawTask[1]}, nil
}
