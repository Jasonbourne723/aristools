package dto

type AddWordDto struct {
	En string   `json:"en"`
	Cn []string `json:"cn"`
}

type WordDto struct {
	Id    int64    `json:"id"`
	En    string   `json:"en"`
	Cn    []string `json:"cn"`
	Times int      `json:"times"`
}
