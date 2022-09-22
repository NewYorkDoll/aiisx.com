package models

type ConfigWakAPI struct {
	URL    string `env:"URL"     long:"url"     required:"true" description:"wakapi connection url"`
	APIKey string `env:"API_KEY" long:"api-key" required:"true" description:"WakAPI API key"`
}

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
