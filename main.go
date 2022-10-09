package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strconv"
	"time"

	"aiisx.com/src/auth"
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
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	adapter "github.com/gwatts/gin-adapter"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
	"golang.org/x/sync/errgroup"
)

var (
	db     *ent.Client
	logger log.Interface
	g      errgroup.Group
	r      *gin.Engine
)

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
func GetSessionUser(r *http.Request, session *sessions.Session) *ent.User {
	s, ok := session.Values["gothUser"]
	if !ok {
		return nil
	}
	u := &ent.User{}
	json.Unmarshal([]byte(s.(string)), u)
	return u
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
	database.RegisterHooks(ctx)
	auth := auth.NewAuthHandler[ent.User, int](
		database.NewAuthService(db, config.GITHUB_USERID),
		securecookie.GenerateRandomKey(16),
		securecookie.GenerateRandomKey(16),
	)

	// 模式设置
	gin.SetMode(config.GIN_MODE)
	r = gin.New()

	nextHandler, wrapper := adapter.New()
	ns := database.UseContextIP(nextHandler)
	newAuthMiddleware := auth.NewAuthMiddleware(nextHandler)
	r.Use(cors.Default())
	r.Use(wrapper(ns))
	r.Use(RequestInfo)
	r.Use(wrapper(newAuthMiddleware))
	goth.UseProviders(
		github.New(
			config.Github.ClientID,
			config.Github.ClientSecret,
			config.HTTP.BaseURL+"/-/auth/providers/github/callback",
		),
	)

	r.GET("/-/auth/providers/github/callback", func(ctx *gin.Context) {
		ctx.Header("Cache-Control", "no-store")
		provider := "github"
		ctx.Request = contextWithProviderName(ctx, provider)
		if gothUser, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request); err == nil {
			id, err := auth.Auth.Set(ctx, &gothUser)

			if err != nil {
				logger.Error(err.Error())
				return
			}
			gothic.StoreInSession("_auth", fmt.Sprintf("%v", id), ctx.Request, ctx.Writer)
			ctx.Redirect(http.StatusMovedPermanently, "/admin")
		} else {
			ctx.Redirect(http.StatusMovedPermanently, "/")

		}
	})

	gh.NewChient(ctx, config.GITHUB_ACCESS_TOKEN)
	// graphql服务
	srv := graphql.New(db)

	r.GET("/-/auth", func(ctx *gin.Context) {
		ctx.Header("Cache-Control", "no-store")
		gothic.BeginAuthHandler(ctx.Writer, gothic.GetContextWithProvider(ctx.Request, "github"))
	})
	r.GET("/-/auth/logout", func(ctx *gin.Context) {
		ctx.Header("Cache-Control", "no-store")
		err := gothic.Logout(ctx.Writer, gothic.GetContextWithProvider(ctx.Request, "github"))
		auth.Logout()
		if err != nil {
			ctx.JSON(http.StatusFound, gin.H{
				"message": err,
			})
		}
		ctx.Redirect(http.StatusMovedPermanently, "/")
	})
	r.GET("/", ReverseProxy())
	r.GET("/p/*id", ReverseProxy())
	r.GET("/error/*id", ReverseProxy())
	r.GET("/posts", ReverseProxy())
	r.GET("/repost", ReverseProxy())
	r.GET("/admin/*id", ReverseProxy())
	r.GET("/_nuxt/*id", ReverseProxy())
	r.GET("/graphql", func(c *gin.Context) {
		playground.Handler("GraphQL playground", "/query").ServeHTTP(c.Writer, c.Request)
	})

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
		err := gh.UserRunner(ctx, config.GITHUB_USERNAME)
		return err
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

	r.NoRoute(func(c *gin.Context) {
		code := c.Writer.Status()
		fmt.Println(code)
		url := "/error/" + strconv.FormatInt(int64(code), 10)
		// 实现内部重定向
		c.Redirect(http.StatusMovedPermanently, url)

	})

	if e := g.Wait(); e != nil {
		log.Fatal(e.Error())
	}
}

func ReverseProxy() gin.HandlerFunc {

	target := "localhost:3000"

	return func(c *gin.Context) {

		director := func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = target
			req.Host = target
		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func contextWithProviderName(c *gin.Context, provider string) *http.Request {
	return c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", provider))
}

func RequestInfo(ctx *gin.Context) {
	now := time.Now()
	timeInfo := fmt.Sprintf("%v-%v-%v %v-%v-%v", now.Year(), int(now.Month()), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	fmt.Printf("[%v] %v - %v\n", timeInfo, ctx.Request.Method, ctx.FullPath())
}
