package main

import (
	"log/slog"
	"url-shortener/src/internal/config"
	"url-shortener/src/pkg/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetUpLogger(cfg.Env)

	log.Info(
		"url shortener is starting",
		slog.String("address", cfg.Server.Address),
		slog.String("env", cfg.Env),
	)

	// todo init storage

	// todo init router

	// todo run server
}
