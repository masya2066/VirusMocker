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
