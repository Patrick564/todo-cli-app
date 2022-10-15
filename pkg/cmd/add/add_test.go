package add

import (
	"reflect"
	"testing"
	"testing/fstest"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		name string
		task string
		file fstest.MapFile
		want fstest.MapFS
	}{
		{
			name: "create new task in an empty file",
			task: "1. task example",
			file: fstest.MapFile{Data: []byte("")},
			want: fstest.MapFS{
				"all.md": {Data: []byte("1. task example")},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if !reflect.DeepEqual(c.file, c.want) {
				t.Errorf("want '%+v' but got '%+v'", string(c.want["all.md"].Data), string(c.file.Data))
			}
		})
	}
}
