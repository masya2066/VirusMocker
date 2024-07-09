package db

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
