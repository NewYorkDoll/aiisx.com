package config

import (
	"os"

	"github.com/apex/log"
	"github.com/joho/godotenv"
)

var (
	GIN_MODE            string
	GITHUB_USERNAME     string
	GITHUB_ACCESS_TOKEN string
	WAKAPI_URL          string
	WAKAPI_KEY          string
)

func init() {
	err := godotenv.Load()
	logger := log.WithFields(log.Fields{"fileName": "config.go"})
	GIN_MODE = os.Getenv("GIN_MODE")
	GITHUB_ACCESS_TOKEN = os.Getenv("GITHUB_ACCESS_TOKEN")
	GITHUB_USERNAME = os.Getenv("GITHUB_USERNAME")
	WAKAPI_URL = os.Getenv("WAKAPI_URL")
	WAKAPI_KEY = os.Getenv("WAKAPI_KEY")
	if err != nil {
		logger.Fatal("Error loading .env file")
	}
}
