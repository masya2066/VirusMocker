package folders

import (
	"os"
)

func Create(path string) error {
	if err := os.Mkdir(path, 0755); err != nil {
		return err
	}
	return nil
}

func CheckFolderExists(folderPath string) error {
	if _, err := os.Stat(folderPath); err != nil {
		return err
	}
	return nil
}
