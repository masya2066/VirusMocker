package consumer

import (
	"mime/multipart"
	"virus_mocker/app/pkg/files"
	"virus_mocker/app/pkg/logger"
)

func PtmsChecker(file multipart.File, header *multipart.FileHeader, err error) bool{
	log := logger.Init()
	size, err := files.FileSizeChecker(file, header, err)
	if err != nil {
		log.Error(err.Error())
	}
	if size == 0 {
		log.Error("File is empty")
	}

	exist, err := files.CheckFileContent(file, "virus_exist")
	if err != nil {
		log.Error(err.Error())
	}

	return exist
}
