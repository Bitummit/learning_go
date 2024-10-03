package main

import (
	"go_api/internal/storage/sqlite"
	urltoshort "go_api/internal/url_to_short"
	"go_api/pkg/config"
	"go_api/pkg/logger"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)



func main() {
	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Env)
	log.Info("starting service ...", slog.String("env", cfg.Env))


	log.Info("Database connecting ...")
	storage, err := sqlite.New(cfg.StoragePath)
	
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

	router.Post("/url", urltoshort.NewAlias(log, storage))

	log.Info("Starting server", slog.String("address", cfg.Address))

	srv := &http.Server{
		Addr: cfg.Address,
		Handler: router,
		// ReadTimeout: cfg.HTTPServer.Timeout,
		// WriteTimeout: cfg.HTTPServer.Timeout,
		// IdleTimeout: cfg.HTTPServer.IdleTimeout,
	}

	err = srv.ListenAndServe(); if err != nil {
		log.Error("failed to start server")
	}

}


