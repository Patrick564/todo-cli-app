package cmdutil

import (
	"errors"
	"os"
	"path/filepath"
)

// Create if not exists database file.
func NeedSchema(dir, fileName string) (bool, error) {
	filePath := filepath.Join(dir, fileName)

	_, err := os.Stat(filePath)
	if errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return false, err
		}

		file, err := os.Create(filePath)
		if err != nil {
			return false, err
		}
		file.Close()

		return true, nil
	}

	return false, nil
}
