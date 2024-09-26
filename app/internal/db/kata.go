package db

import "gorm.io/gorm"

type SensorId struct {
	gorm.Model
	SensorId         string `gorm:"unique;not null" json:"sensor_id"`
	SensorInstanceId string `gorm:"not null" json:"sensor_instance_id"`
	Active           bool   `gorm:"not null" json:"active"`
}

type KataFile struct {
	ScanId   string `gorm:"unique;not null" json:"scan_id"`
	State    string `gorm:"not null" json:"state"`
	SensorId string `gorm:"not null" json:"sensor_instance_id"`
	gorm.Model
}

type FileState struct {
	ScanId string `json:"scanId"`
	State  string `json:"state"`
}

const (
	KataProcessing  = "processing"
	KataNotDetected = "not detected"
	KataDetect      = "detect"
)

func (db *Database) KataCreate(file KataFile) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Model(KataFile{}).Create(file).Error; err != nil {
		return tx.Rollback().Error
	}
	return tx.Commit().Error
}

func (db *Database) KataUpdate(file KataFile) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Model(KataFile{}).Where("scan_id = ?", file.ScanId).Updates(file).Error; err != nil {
		return tx.Rollback().Error
	}
	return tx.Commit().Error
}

func (db *Database) KataGet(scanId string) (KataFile, error) {
	var file KataFile
	if err := db.Where("scan_id = ?", scanId).First(&file).Error; err != nil {
		return file, err
	}
	return file, nil
}
