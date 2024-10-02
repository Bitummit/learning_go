package main

import (
	"go_api/internal/storage/sqlite"
	"go_api/pkg/config"
	"go_api/pkg/logger"
	"log/slog"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
}


