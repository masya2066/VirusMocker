package req

type CreateScanTaskMS struct {
	FileUri       string `json:"file_uri"`
	FileName      string `json:"file_name"`
	AsyncResult   string `json:"async_result"`
	AnalysisDepth int    `json:"analysis_depth"`
}
