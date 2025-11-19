package migration_test

import (
	"testing"
	"time"

	"github.com/cooperlutz/go-full/pkg/deebee/migration"
)

// TestNewPostgresMigration_InvalidURL tests that NewPostgresMigration returns an error for an invalid URL.
func TestNewPostgresMigration_InvalidURL(t *testing.T) {
	t.Parallel()

	invalidURL := "invalid://localhost:5432/db"
	m, err := migration.NewPostgresMigration(invalidURL)
	if err == nil {
		t.Errorf("expected error for invalid URL, got nil")
	}
	if m != nil {
		t.Errorf("expected nil migrate instance, got %v", m)
	}
}

// TestNewPostgresMigration_Timeout tests that newPostgresMigration retries the correct number of times.
func TestNewPostgresMigration_Timeout(t *testing.T) {
	t.Parallel()

	start := time.Now()
	invalidURL := "invalid://localhost:5432/db"
	_, _ = migration.NewPostgresMigration(invalidURL)
	elapsed := time.Since(start)
	minExpected := migration.DefaultMigrationAttempts * migration.DefaultMigrationTimeout
	if elapsed < minExpected {
		t.Errorf("expected at least %v of retries, got %v", minExpected, elapsed)
	}
}
