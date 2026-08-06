package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/azicussdu/GoProjG2/internal/app"
	"github.com/azicussdu/GoProjG2/internal/config"
)

// @title           GoProjG2 API
// @version         1.0
// @description     REST API for an online course platform: courses, lessons, enrollments and authentication.

// @contact.name   API Support
// @contact.email  azicus.sdu@gmail.com

// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.apikey  BearerAuth
// @in                          header
// @name                        Authorization
// @description                 Type "Bearer" followed by a space and the JWT access token.

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	logger := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}),
	)

	slog.SetDefault(logger)

	if err := app.Run(cfg); err != nil {
		slog.Error("server failed")
	}
}
