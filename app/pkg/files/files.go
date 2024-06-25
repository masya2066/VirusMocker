package files

import (
	"fmt"
	"os"
)

func Create(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("Hello, World!\nThis is a test file.")
	if err != nil {
		return err
	}

	fmt.Println("File created and data written successfully")
	return nil
}
