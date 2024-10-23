package main

import (
	"context"
	"database/sql"
	"os/signal"
	"runtime/debug"
	"strings"
	"syscall"

	"github.com/difmaj/ms-credit-score/internal/pkg/config"
	"github.com/difmaj/ms-credit-score/internal/pkg/logger"
	"github.com/difmaj/ms-credit-score/internal/pkg/migrations"
	"github.com/difmaj/ms-credit-score/internal/pkg/redis"
	"github.com/difmaj/ms-credit-score/internal/pkg/router/middleware"
	"github.com/difmaj/ms-credit-score/internal/repository"
	"github.com/difmaj/ms-credit-score/internal/subscriber"
	"github.com/difmaj/ms-credit-score/internal/usecase"
	"go.uber.org/zap"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()

	logger.Init()
	if err := config.Load(); err != nil {
		logger.Logger.Fatal("config.Load", zap.Error(err))
	}

	defer func() {
		if err := logger.Sync(); err != nil {
			logger.Logger.Error("logger.Sync", zap.Error(err))
		}
	}()

	if bi, ok := debug.ReadBuildInfo(); ok {
		logger.Logger.Info("Go: " + strings.TrimPrefix(bi.GoVersion, "go"))

		for _, bs := range bi.Settings {
			if strings.HasPrefix(bs.Key, "vcs") || bs.Key == "-tags" {
				logger.Logger.Info(bs.Key + ": " + bs.Value)
			}
		}
	}

	db, err := sql.Open("mysql", config.Env.DatabaseMYSQLDNS)
	if err != nil {
		logger.Logger.Fatal("sqlx.Open", zap.Error(err))
	}

	if err := db.Ping(); err != nil {
		logger.Logger.Fatal("db.Ping", zap.Error(err))
	}

	migrationResults, err := migrations.Run(db, &migrations.Config{
		Path:   "migrations/up",
		Schema: config.Env.DatabaseMYSQLSchema,
	})
	if err != nil {
		logger.Logger.Error("migrations.Run", zap.Error(err))
		return
	}

	if err := migrations.RunSeeds(ctx, db); err != nil {
		logger.Logger.Error("migrations.RunSeeds", zap.Error(err))
		return
	}

	logger.Logger.Info("migrations.Run", zap.Any("results", migrationResults))

	repo, err := repository.New(db)
	if err != nil {
		logger.Logger.Error("repository.New", zap.Error(err))
		return
	}

	redis, err := redis.NewClient()
	if err != nil {
		logger.Logger.Error("redis.NewClient", zap.Error(err))
		return
	}

	usecase := usecase.New(repo, redis)
	middle := middleware.NewMiddleware(usecase)

	subscriber, err := subscriber.New()
	if err != nil {
		logger.Logger.Error("subscriber.New", zap.Error(err))
		return
	}

	handler.NewHandler(subscriber, middle, usecase)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go subscriber.Process()

	<-ctx.Done()

	if err := subscriber.Close(); err != nil {
		logger.Logger.Error("server.Shutdown", zap.Error(err))
	}
}
