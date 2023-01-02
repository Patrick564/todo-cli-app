package cmdutil

import (
	"time"
)

const (
	DriverName   string = "sqlite3"
	FormatLayout string = "2006-01-02 03:04:05"
)

type Task struct {
	Id       int
	Content  string
	Complete int
	Date     time.Time
}
