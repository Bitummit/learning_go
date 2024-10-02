package main

import (
	"go_api/pkg/config"
	"go_api/pkg/logger"
	"go_api/pkg/storage/sqlite"
	"log/slog"
	"os"
)



func main() {
	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Env)
	log.Info("starting service ...", slog.String("env", cfg.Env))


	log.Info("Database connecting ...")
	_, err := sqlite.New(cfg.StoragePath)
	
	if err != nil {
		log.Error("Faled to connect to storage", logger.Err(err))
		os.Exit(1)
	}

	log.Info("Database connected")


}


