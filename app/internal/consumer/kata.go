package consumer

import (
	"mime/multipart"
	"time"
	"virus_mocker/app/internal/db"
	"virus_mocker/app/pkg/files"
	"virus_mocker/app/pkg/logger"
)

func KataChecker(scanid string, file multipart.File, header *multipart.FileHeader, DB db.Database, err error) {
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
	if exist {
		time.Sleep(time.Duration(size*3) * time.Second)
		if err := DB.KataUpdateState(scanid, db.KataDetect); err != nil {
			log.Error(err.Error())
		}
		return
	}

	existError, err := files.CheckFileContent(file, "error_exist")
	if err != nil {
		log.Error(err.Error())
	}
	if existError {
		time.Sleep(time.Duration(size*3) * time.Second)
		if err := DB.KataUpdateState(scanid, db.KataError); err != nil {
			log.Error(err.Error())
		}
		return
	}

	time.Sleep(time.Duration(size*3) * time.Second)
	if err := DB.KataUpdateState(scanid, db.KataNotDetected); err != nil {
		log.Error(err.Error())
	}
	return
}
