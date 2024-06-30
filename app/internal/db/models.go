package db

import "gorm.io/gorm"

type KataFile struct {
	ScanId string `gorm:"unique;not null" json:"scan_id"`
	Status string `gorm:"not null" json:"status"`
	gorm.Model
}
