package cmdutil

import (
	"time"
)

type Task struct {
	Id       int
	Content  string
	Complete int
	Date     time.Time
}
