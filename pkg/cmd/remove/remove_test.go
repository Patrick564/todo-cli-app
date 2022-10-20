package remove

import (
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/Patrick564/todo-cli-app/pkg/cmdutil"
)

func TestRemove(t *testing.T) {
	cases := []struct {
		name string
		file fstest.MapFS
		want []*cmdutil.Task
	}{
		{
			name: "remove a first task from file all.md",
			file: fstest.MapFS{"all.md": {Data: []byte("123: example\n125: example 2")}},
			want: []*cmdutil.Task{
				{Id: "125", Content: "example 2"},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := cmdutil.ReadFromFS(c.file, "all")
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("want error '%+v' but got '%+v'", c.want, got)
			}
		})
	}
}
