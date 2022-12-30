package database

import (
	"log"
	"os"
)

func init() {
	err := os.Setenv("PATH", "./tasks")
	if err != nil {
		log.Fatal(err)
	}
}
