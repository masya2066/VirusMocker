package response

type CreateFileMS struct {
	Errors []string         `json:"errors"`
	Data   CreateFileMSData `json:"data"`
}
type CreateFileMSData struct {
	FileUri string `json:"file_uri"`
	Ttl     int    `json:"ttl"`
}

type CreateScanTaskMS struct {
	Errors []string             `json:"errors"`
	Data   CreateScanTaskMSData `json:"data"`
}

type CreateScanTaskMSData struct {
	ScanId string `json:"scan_id"`
}
