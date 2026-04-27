package migration

import (
	"errors"
	"fmt"
	"log/slog"
	"slices"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	DefaultMigrationAttempts = 20
	DefaultMigrationTimeout  = 1 * time.Second
	DefaultMigrationSource   = "file://db/migrations"
)

// Migrate runs the database migrations.
func Migrate(databaseDriver, databaseURL string) {
	slog.Info("Migrate: starting migrations")

	supportedDatabaseDrivers := []string{
		"postgres",
	}

	if !slices.Contains(supportedDatabaseDrivers, databaseDriver) {
		slog.Error("Migrate: unsupported database driver " + databaseDriver + ", supported drivers are: " + fmt.Sprint(supportedDatabaseDrivers))
	}

	m, err := NewPostgresMigration(databaseURL)
	if err != nil {
		slog.Error("Migrate: postgres connect error: " + err.Error())
	}

	// Execute the migrations.
	err = m.Up()
	defer m.Close()

	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		slog.Error("Migrate: up error: " + err.Error())
	}

	if !errors.Is(err, migrate.ErrNoChange) {
		slog.Info("Migrate: migrations applied successfully for driver " + databaseDriver)
	}

	slog.Info("Migrate: migrations applied successfully for driver " + databaseDriver)
}

func NewPostgresMigration(databaseURL string) (*migrate.Migrate, error) {
	var (
		attempts = DefaultMigrationAttempts
		err      error
		m        *migrate.Migrate
	)

	for attempts > 0 {
		m, err = migrate.New(DefaultMigrationSource, databaseURL)
		if err == nil {
			break
		}

		slog.Info("Migrate: postgres is trying to connect, attempts left: " + fmt.Sprint(attempts))

		time.Sleep(DefaultMigrationTimeout)

		attempts--
	}

	return m, err
}
