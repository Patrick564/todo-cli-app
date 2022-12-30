package database

import (
	"log"
	"os"
)

func init() {
	err := os.Setenv("FILE_PATH", "./tasks")
	if err != nil {
		log.Fatal(err)
	}
}
