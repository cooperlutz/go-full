package deebee

import (
	"testing"
	"time"
)

// TestNewPostgresMigration_InvalidURL tests that newPostgresMigration returns an error for an invalid URL.
func TestNewPostgresMigration_InvalidURL(t *testing.T) {
	t.Parallel()

	invalidURL := "invalid://localhost:5432/db"
	m, err := newPostgresMigration(invalidURL)
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
	_, _ = newPostgresMigration(invalidURL)
	elapsed := time.Since(start)
	minExpected := _defaultAttempts * _defaultTimeout
	if elapsed < minExpected {
		t.Errorf("expected at least %v of retries, got %v", minExpected, elapsed)
	}
}
