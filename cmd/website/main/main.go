package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"

	"website.friendsonly.me/pkg/repository/postgres"
	"website.friendsonly.me/pkg/tools/logging"
	"website.friendsonly.me/pkg/website/config"
	"website.friendsonly.me/pkg/website/handlers"
	"website.friendsonly.me/pkg/website/server"
)

func main() {
	cfg, err := config.NewWebsiteConfig()
	if err != nil {
		log.Fatal("Config error ", err)
	}

	logger, err := logging.InitLogger(cfg.LogPath, cfg.LogLevel)
	if err != nil {
		log.Fatal("logrus error ", err)
	}

	logger.Info("Connecting to DB...")
	ctx := context.Background()
	db, err := pgxpool.Connect(ctx, cfg.DBConnectionString)
	if err != nil {
		log.Fatal("DB Error:", cfg.DBConnectionString, err)
	}
	defer db.Close()
	logger.Info("Connect to DB success")

	logger.Info("Creating video repo...")
	videoRepo := postgres.NewVideoRepo(ctx, db, logger)
	logger.Info("Video repo was created")

	logger.Info("Creating public handlers...")
	publicHandlers := handlers.NewPublicHandler(videoRepo, logger)
	logger.Info("Public handlers was created")

	app := fiber.New(fiber.Config{
		// Prefork:       true,
		WriteTimeout:  time.Second * 15,
		ReadTimeout:   time.Second * 15,
		IdleTimeout:   time.Second * 20,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Apache 2.0",
		AppName:       "FriendsOnly.me website",
	})

	server.BindHandlers(app, publicHandlers, cfg)

	logger.Info("Service website listening ", cfg.HTTPAddr)

	err = app.Listen(cfg.HTTPAddr)
	if err != nil {
		logger.Error("listen error ", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	if err := app.Shutdown(); err != nil {
		logger.Info("Website service shutdown")
		return
	}

	logger.Info("Website service was stopped")
}
