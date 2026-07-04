package main

import (
	"log"

	"github.com/azicussdu/GoProjG2/internal/app"
	"github.com/azicussdu/GoProjG2/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	if err := app.Run(cfg); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
