package list

import (
	"reflect"
	"testing"
	"testing/fstest"
)

// TODO: add table test for list all, completed, and pending
func TestList(t *testing.T) {
	cases := []struct {
		name    string
		flag    string
		file    fstest.MapFS
		want    []Task
		wantErr error
	}{
		{
			name: "empty tasks file",
			flag: "all.md",
			file: fstest.MapFS{
				"all.md": {Data: []byte("")},
			},
			want:    nil,
			wantErr: ErrFileEmpty,
		},
		{
			name: "all.md file with 3 tasks",
			flag: "all.md",
			file: fstest.MapFS{
				"all.md": {Data: []byte("1. fake task example\n5. second fake task\n13. two digits id fake task")},
			},
			want: []Task{
				{Id: "1", Content: "fake task example"},
				{Id: "5", Content: "second fake task"},
				{Id: "13", Content: "two digits id fake task"},
			},
		},
		{
			name: "all.md and completed.md file with 2 tasks",
			flag: "completed.md",
			file: fstest.MapFS{
				"all.md":       {Data: []byte("15. fake task")},
				"completed.md": {Data: []byte("12. completed fake task\n28. second completed fake task")},
			},
			want: []Task{
				{Id: "12", Content: "completed fake task"},
				{Id: "28", Content: "second completed fake task"},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := readTasksFromFS(c.file, c.flag)
			if err != nil {
				assertError(t, err, c.wantErr)
			}

			assertDeepEqual(t, got, c.want)
		})
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("want error '%+v' but got '%+v'", want, got)
	}
}

func assertDeepEqual(t testing.TB, got, want []Task) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %+v but got %+v", want, got)
	}
}
