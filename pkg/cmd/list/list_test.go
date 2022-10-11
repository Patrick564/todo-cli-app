package list

import (
	"reflect"
	"testing"
	"testing/fstest"
)

func TestList(t *testing.T) {
	fs := fstest.MapFS{
		"all.md": {Data: []byte(`1. example first task
2. example second task`)},
	}

	task, err := readTasksFromFile(fs)
	if err != nil {
		t.Fatal(err)
	}

	got := task
	want := []Task{
		{Id: "1", Content: "example first task"},
		{Id: "2", Content: "example second task"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %s but got %s", want, got)
	}
}
