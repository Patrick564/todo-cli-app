package list

import (
	"reflect"
	"testing"
	"testing/fstest"
)

// TODO: add table test for list all, completed, and pending
func TestList(t *testing.T) {
	cases := []struct {
		name string
		file fstest.MapFS
		want []Task
	}{
		{
			name: "all.md file with 3 tasks",
			file: fstest.MapFS{
				"all.md": {Data: []byte("1. fake task example\n5. second fake task\n13. two digits id fake task")},
			},
			want: []Task{
				{Id: "1", Content: "fake task example"},
				{Id: "5", Content: "second fake task"},
				{Id: "13", Content: "two digits id fake task"},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := readTasksFromFile(c.file)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("want %s but got %s", c.want, got)
			}
		})
	}
}
