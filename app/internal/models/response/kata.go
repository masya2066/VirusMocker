package response

import "virus_mocker/app/internal/db"

type ScanResult struct {
	ScanId           string `json:"scan_id"`
	State            string `json:"state"`
	SensorInstanceId string `json:"sensor_instance_id"`
}

type ScanFilesResult struct {
	Scans []db.FileState `json:"scans"`
}
