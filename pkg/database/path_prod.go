//go:build prod
// +build prod

package database

import (
	"log"
	"os"
	"path/filepath"
)

const taskDir string = "tasks"

func init() {
	path, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	err = os.Unsetenv("FILE_PATH")
	if err != nil {
		log.Fatal(err)
	}

	err = os.Setenv("FILE_PATH", filepath.Join(filepath.Dir(path), taskDir))
	if err != nil {
		log.Fatal(err)
	}
}
