package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"oop/internal/config"
	"oop/internal/http/middleware/logger"
	"oop/internal/storage"
	"os"
)

const (
	envDev  = "dev"
	envProd = "prod"
)

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
	log = log.With(slog.String("env", cfg.Env))

	log.Info("init server", slog.String("address", cfg.Address))
	log.Debug("logger debug mode enabled")

	db, err := storage.OpenDB(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	))
	if err != nil {
		log.Error(err.Error(), slog.String("error", err.Error()))
	}
	defer db.Close()
	//userModel, err := db.CreateUserTable()
	//if err != nil {
	//	log.Fatal(err)

	router := chi.NewRouter()

	router.Use(middleware.RequestID) // adds a unique request ID to the context
	router.Use(logger.New(log))
	router.Use(middleware.Recoverer) // recover after panic
	router.Use(middleware.URLFormat) // parse URL format

}
