package main

import (
	"context"

	"aiisx.com/src/config"
	"aiisx.com/src/database"
	"aiisx.com/src/database/graphql"
	"aiisx.com/src/ent"
	"aiisx.com/src/gh"
	"aiisx.com/src/models"
	"aiisx.com/src/wakapi"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/apex/log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/sync/errgroup"
)

var (
	db     *ent.Client
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
	ctx := context.Background()
	logger = log.WithFields(log.Fields{
		"filename": "main.go",
	})

	db = database.Open(ctx, logger, config.Database)
	defer db.Close()
	c := ent.NewContext(log.NewContext(context.Background(), logger), db)
	database.Migrate(c, logger)

	gin.SetMode(config.GIN_MODE)
	r := gin.New()

	gh.NewChient(ctx, config.GITHUB_ACCESS_TOKEN)
	srv := graphql.New()
	r.GET("/graphql", playgroundHandler())
	r.POST("/query", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})
	g.Go(func() error {
		configWakAPI := models.ConfigWakAPI{
			URL:    config.WAKAPI_URL,
			APIKey: config.WAKAPI_KEY}
		return wakapi.NewRunner(logger, configWakAPI).Run(ctx)
	})
	g.Go(func() error {
		logger.Info("github UserRunner run!")
		return gh.UserRunner(ctx, config.GITHUB_USERNAME)
	})
	g.Go(func() error {
		logger.Info("server run!")
		return r.Run()
	})
	if e := g.Wait(); e != nil {
		log.Fatal(e.Error())
	}
}
