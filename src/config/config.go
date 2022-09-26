package config

import (
	"os"
	"strconv"

	"aiisx.com/src/models"
	"github.com/apex/log"
	"github.com/joho/godotenv"
)

var (
	GIN_MODE            string
	GITHUB_USERNAME     string
	GITHUB_USERID       int
	GITHUB_ACCESS_TOKEN string
	WAKAPI_URL          string
	WAKAPI_KEY          string
	Database            models.ConfigDatabase
	Github              models.ConfigGithub
	HTTP                models.ConfigHTTP
)

func init() {
	err := godotenv.Load()
	logger := log.WithFields(log.Fields{"fileName": "config.go"})
	GIN_MODE = os.Getenv("GIN_MODE")
	GITHUB_ACCESS_TOKEN = os.Getenv("GITHUB_ACCESS_TOKEN")
	GITHUB_USERNAME = os.Getenv("GITHUB_USERNAME")
	GITHUB_USERID, _ = strconv.Atoi(os.Getenv("GITHUB_USERID"))
	Github.ClientID = os.Getenv("GITHUB_ClientID")
	Github.ClientSecret = os.Getenv("GITHUB_ClientSecret")
	HTTP.BaseURL = os.Getenv("BaseURL")
	WAKAPI_URL = os.Getenv("WAKAPI_URL")
	WAKAPI_KEY = os.Getenv("WAKAPI_KEY")
	Database.URL = os.Getenv("DATABASE_URL")
	if err != nil {
		logger.Fatal("Error loading .env file")
	}
}
