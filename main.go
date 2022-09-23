package main

import (
	"context"

	"aiisx.com/src/config"
	"aiisx.com/src/database"
	"aiisx.com/src/database/graphql"
	"aiisx.com/src/ent"
	_ "aiisx.com/src/ent/runtime"
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
	// 初始化环境变量
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// gh.SyncOnStart = true
	ctx := context.Background()
	logger = log.WithFields(log.Fields{
		"filename": "main.go",
	})
	/// 链接数据库
	db = database.Open(ctx, logger, config.Database)
	defer db.Close()
	ctx = ent.NewContext(log.NewContext(context.Background(), logger), db)
	database.Migrate(ctx, logger)

	// 模式设置
	gin.SetMode(config.GIN_MODE)
	r := gin.New()

	// github服务
	gh.NewChient(ctx, config.GITHUB_ACCESS_TOKEN)

	// graphql服务
	srv := graphql.New(db)
	r.GET("/graphql", playgroundHandler())
	r.POST("/query", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})

	// 运行wakapi服务
	g.Go(func() error {
		configWakAPI := models.ConfigWakAPI{
			URL:    config.WAKAPI_URL,
			APIKey: config.WAKAPI_KEY}
		return wakapi.NewRunner(logger, configWakAPI).Run(ctx)
	})

	// github获取用户服务
	g.Go(func() error {
		logger.Info("github UserRunner run!")
		return gh.UserRunner(ctx, config.GITHUB_USERNAME)
	})

	// github event
	g.Go(func() error {
		logger.Info("EventsRunner run!")
		return gh.EventsRunner(ctx)
	})

	// server主服务
	g.Go(func() error {
		logger.Info("server run!")
		return r.Run()
	})
	if e := g.Wait(); e != nil {
		log.Fatal(e.Error())
	}
}
