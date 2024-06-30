package db

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
