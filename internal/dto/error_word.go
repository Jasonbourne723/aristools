package dto

type ErrorWordDto struct {
	Id    int64    `json:"id"`
	En    string   `json:"en"`
	Cn    []string `json:"cn"`
	Times int      `json:"times"`
}

type WordAnalysisDto struct {
	Today WordAnalysisByDayDto
	Items []WordAnalysisByDayDto
}

type WordAnalysisByDayDto struct {
	Date     string `json:"date"`
	Count    int    `json:"count"`
	ErrCount int    `json:"err_count"`
}
