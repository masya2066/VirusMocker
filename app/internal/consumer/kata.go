package consumer

import (
	"mime/multipart"
	"time"
	"virus_mocker/app/internal/db"
	"virus_mocker/app/pkg/files"
	"virus_mocker/app/pkg/logger"
)

func KataChecker(scanid string, file multipart.File, header *multipart.FileHeader, err error) {
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
		time.Sleep(time.Duration(5) * time.Second)
		if updateError := db.DB.Model(db.KataFile{}).Where("scan_id = ?", scanid).Update("state", db.KataDetect); updateError.Error != nil {
			log.Error(updateError.Error.Error())
		}
		return
	}

	time.Sleep(time.Duration(size*3) * time.Second)
	if updateError := db.DB.Model(db.KataFile{}).Where("scan_id = ?", scanid).Update("state", db.KataNotDetected); updateError.Error != nil {
		log.Error(updateError.Error.Error())
	}
}
