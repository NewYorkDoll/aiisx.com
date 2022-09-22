package main

import (
	"context"

	"aiisx.com/src/database/graphql"
	"aiisx.com/src/gh"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/apex/log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/sync/errgroup"
)

var (
	logger log.Interface
	g      errgroup.Group
)

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	gin.SetMode(GIN_MODE)
	r := gin.New()

	logger = log.WithFields(log.Fields{
		"filename": "main.go",
	})

	ctx := context.Background()

	gh.NewChient(ctx, GITHUB_ACCESS_TOKEN)
	srv := graphql.New()
	r.GET("/graphql", playgroundHandler())
	r.POST("/query", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})

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
