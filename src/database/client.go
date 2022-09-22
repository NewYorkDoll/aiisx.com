package database

import (
	"context"
	"database/sql"
	"time"

	"aiisx.com/src/ent"
	"aiisx.com/src/ent/migrate"
	"aiisx.com/src/models"
	"ariga.io/entcache"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/apex/log"
	_ "github.com/go-sql-driver/mysql"
)

func Open(ctx context.Context, logger log.Interface, config models.ConfigDatabase) *ent.Client {
	var db *sql.DB
	var err error
	var attempt int
	for {
		attempt++
		db, err = sql.Open("mysql", config.URL)

		if err == nil {
			tctx, cancel := context.WithTimeout(ctx, 5*time.Second)

			err = db.PingContext(tctx)
			cancel()
		}

		if err != nil {

			logger.WithError(err).WithField("attempt", attempt).Error("failed to open database connection")

			if attempt > 2 {
				logger.Fatal("failed to open database connection after 3 attempts")
			}
			time.Sleep(time.Second * 5)
			continue
		}
		break
	}

	logger.Info("connected to database")

	// Create an ent.Driver from db.
	driver := entsql.OpenDB(dialect.MySQL, db)
	return ent.NewClient(ent.Driver(
		entcache.NewDriver(
			driver,
			entcache.TTL(120*time.Second),
			entcache.Levels(entcache.NewLRU(256)),
		),
	))
}

func Migrate(ctx context.Context, logger log.Interface) {
	logger.Info("initiating database schema migration")
	db := ent.FromContext(ctx)
	if db == nil {
		panic("database client is nil")
	}

	if err := db.Schema.Create(
		entcache.Skip(ctx),
		schema.WithDropColumn(true),
		schema.WithDropIndex(true),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		logger.WithError(err).Fatal("failed to create schema")
	}
	logger.Info("database schema migration complete")
}
