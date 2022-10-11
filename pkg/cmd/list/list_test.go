package list

import (
	"reflect"
	"testing"
	"testing/fstest"
)

// TODO: add table test for list all, completed, and pending
func TestList(t *testing.T) {
	fs := fstest.MapFS{
		"all.md": {Data: []byte(`1. fake task example
5. second fake task
13. two digits id fake task
`)},
	}

	task, err := readTasksFromFile(fs)
	if err != nil {
		t.Fatal(err)
	}

	got := task
	want := []Task{
		{Id: "1", Content: "fake task example"},
		{Id: "5", Content: "second fake task"},
		{Id: "13", Content: "two digits id fake task"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %s but got %s", want, got)
	}
}
