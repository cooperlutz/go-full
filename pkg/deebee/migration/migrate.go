package migration

import (
	"errors"
	"log"
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
	log.Printf("Migrate: starting migrations")

	supportedDatabaseDrivers := []string{
		"postgres",
	}

	if !slices.Contains(supportedDatabaseDrivers, databaseDriver) {
		log.Fatalf("Migrate: unsupported database driver %s, supported drivers are: %v", databaseDriver, supportedDatabaseDrivers)
	}

	m, err := NewPostgresMigration(databaseURL)
	if err != nil {
		log.Fatalf("Migrate: postgres connect error: %s", err)
	}

	// Execute the migrations.
	err = m.Up()
	defer m.Close()

	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Migrate: up error: %s", err)
	}

	if !errors.Is(err, migrate.ErrNoChange) {
		log.Printf("Migrate: migrations applied successfully for driver %s", databaseDriver)
	}

	log.Printf("Migrate: migrations applied successfully for driver %s", databaseDriver)
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

		log.Printf("Migrate: postgres is trying to connect, attempts left: %d", attempts)

		time.Sleep(DefaultMigrationTimeout)

		attempts--
	}

	return m, err
}
