package main

import (
	"Kulibyka/internal/config"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfg := config.MustLoad()

	// Валидация параметров
	if cfg.StoragePath == "" {
		panic("storage-path is required")
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.PostgreSQL.User, cfg.PostgreSQL.Password, cfg.PostgreSQL.Host,
		cfg.PostgreSQL.Port, cfg.PostgreSQL.DBName, cfg.PostgreSQL.SSLMode)

	m, err := migrate.New("file://migrations", connStr)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migration to change")
			return
		}
		panic(err)
	}

	fmt.Println("migrations applied successfully")
}
