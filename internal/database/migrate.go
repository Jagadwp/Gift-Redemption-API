package database

import (
	"log"

	"github.com/gift-redemption/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(cfg *config.Config) {
	dsn := cfg.Database.MigrationURL()

	m, err := migrate.New("file://migrations", dsn)
	if err != nil {
		log.Fatalf("failed to initialize migrations: %v", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migration failed: %v", err)
	}

	log.Println("migrations applied successfully")
}
