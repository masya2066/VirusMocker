package files

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

func Create(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("")
	if err != nil {
		return err
	}

	fmt.Println("File created and data written successfully")
	return nil
}

func FileSizeChecker(file multipart.File, header *multipart.FileHeader, err error) (int, error) {
	if err != nil {
		return 0, err
	}
	defer file.Close()

	return int(header.Size / (1024 * 1024)), nil
}

func CheckFileContent(file multipart.File, data string) (bool, error) {
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return false, err
	}
	content, err := io.ReadAll(file)
	if err != nil {
		return false, err
	}

	if bytes.Contains(content, []byte(data)) {
		return true, nil
	}
	return false, nil
}
