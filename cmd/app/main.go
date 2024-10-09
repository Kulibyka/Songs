package main

import (
	"Kulibyka/internal/config"
	api "Kulibyka/internal/http"
	"Kulibyka/internal/storage/postgresql"
	"fmt"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"os"
)

const (
	envLocal = "local"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	log := setupLogger(cfg.Env)
	log.Info("app started")
	log.Debug("error")

	db, err := postgresql.New(cfg.PostgreSQL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := mux.NewRouter()

	api.RegisterRoutes(router, db)

	addr := ":8080"
	log.Info("Starting server", slog.String("addr", addr))
	if err := http.ListenAndServe(addr, router); err != nil {
		panic(err)
	}

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(
			os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	}
	return log
}
