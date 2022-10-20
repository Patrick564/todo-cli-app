package list

import (
	"fmt"
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/Patrick564/todo-cli-app/pkg/cmdutil"
)

// TODO: add table test for list all, completed, and pending
func TestList(t *testing.T) {
	cases := []struct {
		name string
		flag string
		file fstest.MapFS
		want []*cmdutil.Task
	}{
		{
			name: "all.md file with 3 tasks",
			flag: "all",
			file: fstest.MapFS{
				"all.md": {Data: []byte("1: fake task example\n5: second fake task\n13: two digits id fake task")},
			},
			want: []*cmdutil.Task{
				{Id: "1", Content: "fake task example"},
				{Id: "5", Content: "second fake task"},
				{Id: "13", Content: "two digits id fake task"},
			},
		},
		{
			name: "all.md and completed.md file with 2 tasks",
			flag: "completed",
			file: fstest.MapFS{
				"all.md":       {Data: []byte("15: fake task")},
				"completed.md": {Data: []byte("12: completed fake task\n28: second completed fake task")},
			},
			want: []*cmdutil.Task{
				{Id: "12", Content: "completed fake task"},
				{Id: "28", Content: "second completed fake task"},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := cmdutil.GetTaskList(c.file, c.flag)

			assertNoError(t, err)
			assertDeepEqual(t, got, c.want)
		})
	}
}

func TestListWithError(t *testing.T) {
	cases := []struct {
		name    string
		flag    string
		file    fstest.MapFS
		want    []*cmdutil.Task
		wantErr error
	}{
		{
			name:    "empty files dir",
			flag:    "all",
			file:    fstest.MapFS{},
			wantErr: cmdutil.ErrFileNotFound,
		},
		{
			name:    "empty tasks file",
			flag:    "all",
			file:    fstest.MapFS{"all.md": {Data: []byte("")}},
			wantErr: cmdutil.ErrFileEmpty,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := cmdutil.GetTaskList(c.file, c.flag)

			assertError(t, err, c.wantErr)
		})
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	fmt.Println("aaaaaaa assert error")

	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}
	fmt.Printf("%+v, %+v assert", got, want)
	if got != want {
		t.Errorf("want error '%+v' but got '%+v'", want, got)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertDeepEqual(t testing.TB, got, want []*cmdutil.Task) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %+v but got %+v", want, got)
	}
}
