package models

type CodingStats struct {
	Languages      []LanguageStat `json:"languages"`
	TotalSeconds   int            `json:"total_seconds"`
	TotalDuration  string         `json:"total_duration"`
	CalculatedDays int            `json:"calculated_days"`
}

type LanguageStat struct {
	Language      string `json:"key"`
	TotalSeconds  int    `json:"total"`
	TotalDuration string `json:"total_duration"`
}
