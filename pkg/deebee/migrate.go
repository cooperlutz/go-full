package deebee

import (
	"errors"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	_defaultAttempts = 20
	_defaultTimeout  = 1 * time.Second
)

func newPostgresMigration(databaseURL string) (*migrate.Migrate, error) {
	var (
		attempts = _defaultAttempts
		err      error
		m        *migrate.Migrate
	)

	for attempts > 0 {
		m, err = migrate.New("file://db/migrations", databaseURL)
		if err == nil {
			break
		}

		log.Printf("Migrate: postgres is trying to connect, attempts left: %d", attempts)
		time.Sleep(_defaultTimeout)

		attempts--
	}

	return m, err
}

// Migrate runs the database migrations.
func Migrate(databaseDriver, databaseURL string) {
	log.Printf("Migrate: starting migrations for driver %s, url: %s", databaseDriver, databaseURL)

	supportedDatabaseDrivers := []string{
		"postgres",
	}

	if databaseDriver != "postgres" {
		log.Fatalf("Migrate: unsupported database driver %s, supported drivers are: %v", databaseDriver, supportedDatabaseDrivers)
	}

	m, err := newPostgresMigration(databaseURL)
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
