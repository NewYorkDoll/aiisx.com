package main

import (
	"os"

	"github.com/apex/log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	logger log.Interface
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	GIN_MODE := os.Getenv("GIN_MODE")
	gin.SetMode(GIN_MODE)
	r := gin.Default()
	logger := log.WithFields(log.Fields{})
	logger.Info("test info message")
	if err := r.Run(); err != nil {
		logger.WithError(err).Fatal("shutting down")
	} else {
		logger.Info("Listening and serving HTTP on http://localhost:8080")
	} // 监听并在 0.0.0.0:8080 上启动服务
}
