package main

import (
	"log/slog"
	"os"
	"url-shortener/src/internal/config"
	"url-shortener/src/internal/helpers/sl"
	"url-shortener/src/internal/storage/sqlite"
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

	storage, err := sqlite.NewStorage(cfg.StoragePath)
	if err != nil {
		log.Error("db init is failed", sl.Err(err))
		os.Exit(1)
	}

	_ = storage

	// todo init storage

	// todo init router

	// todo run server
}
