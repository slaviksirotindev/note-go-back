package infrastructure

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/labstack/gommon/log"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigration(handler *SqlHandler) error {
	dbConn, err := handler.db.DB()
	if err != nil {
		log.Fatal("Could not connect to database: %v", err)
	}

	driver, err := postgres.WithInstance(dbConn, &postgres.Config{})
	if err != nil {
		log.Fatal("failed to create database driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)

	if err != nil {
		log.Fatal("failed to create migration instance: %v", err)
	}
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Info("nothing to migrate")
		} else {
			log.Fatal("failed to migrate: %v", err)
		}
	} else {
		log.Info("migrated")
	}
	return nil
}
