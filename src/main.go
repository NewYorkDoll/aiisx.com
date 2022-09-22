package main

import (
	"context"

	"aiisx.com/src/gh"
	"github.com/apex/log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/sync/errgroup"
)

var (
	logger log.Interface
	g      errgroup.Group
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	gin.SetMode(GIN_MODE)
	r := gin.Default()
	logger = log.WithFields(log.Fields{
		"filename": "main.go",
	})

	ctx := context.Background()
	gh.NewChient(ctx, GITHUB_ACCESS_TOKEN)
	g.Go(func() error {
		logger.Info("github UserRunner run!")
		return gh.UserRunner(ctx, GITHUB_USERNAME)
	})
	g.Go(func() error {
		logger.Info("server run!")
		return r.Run()
	})
	if e := g.Wait(); e != nil {
		log.Fatal(e.Error())
	}
}
