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

// Close implements io.Closer, which is required by multipart.File.
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

func ByteSliceToMultipartFile(data []byte, fileName string) (multipart.File, *multipart.FileHeader, error) {
	// Create a buffer to write our multipart form data
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create a form-data header with a file part
	header := make(textproto.MIMEHeader)
	header.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, fileName))
	header.Set("Content-Type", "application/octet-stream")

	// Create the file part and write the byte slice to it
	part, err := writer.CreatePart(header)
	if err != nil {
		return nil, nil, err
	}

	if _, err = part.Write(data); err != nil {
		return nil, nil, err
	}

	// Close the writer to finalize the form data
	writer.Close()

	// Create a customFile that implements the multipart.File interface
	file := &customFile{Reader: bytes.NewReader(data)}

	// Create a multipart.FileHeader to describe the file
	fileHeader := &multipart.FileHeader{
		Filename: fileName,
		Header:   header,
		Size:     int64(len(data)),
	}

	// Return the custom file as a multipart.File
	return file, fileHeader, nil
}
