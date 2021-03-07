package util

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type DatabaseDriverType string

const (
	PGDriver DatabaseDriverType = "postgres"
)

type MigrationConfig struct {
	Driver      DatabaseDriverType
	DatabaseURL string
	FileDir     string
}

func RunMigrations(config MigrationConfig) error {
	switch config.Driver {
	case PGDriver:
		return runPGMigrations(config)
	default:
		return errors.New("not supported driver type")
	}
}

func runPGMigrations(config MigrationConfig) error {
	m, err := migrate.New(
		config.FileDir,
		config.DatabaseURL)
	if err != nil {
		return err
	}
	return m.Up()
}
