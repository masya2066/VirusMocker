package files

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
	"os"
)

type customFile struct {
	*bytes.Reader
}

func (cf *customFile) Close() error {
	// No operation needed for in-memory data, but method is required.
	return nil
}

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
	// Возвращаем вес в МБ
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

	// Проверяем наличие фразы в содержимом
	if bytes.Contains(content, []byte(data)) {
		return true, nil
	}
	return false, nil
}

func ByteSliceToMultipartFile(data []byte, fileName string) (multipart.File, *multipart.FileHeader, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	header := make(textproto.MIMEHeader)
	header.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, fileName))
	header.Set("Content-Type", "application/octet-stream")

	part, err := writer.CreatePart(header)
	if err != nil {
		return nil, nil, err
	}

	if _, err = part.Write(data); err != nil {
		return nil, nil, err
	}

	writer.Close()

	file := &customFile{Reader: bytes.NewReader(data)}

	fileHeader := &multipart.FileHeader{
		Filename: fileName,
		Header:   header,
		Size:     int64(len(data)),
	}

	return file, fileHeader, nil
}
