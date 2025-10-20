package config_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/app/config"
)

// TestConfig_ToString tests the String method of the Config struct.

// Test Cases:

// 1. Config to string returns a string type.
func TestConfig_ToString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		cfg      config.Config
		wantType string
	}{
		{
			name: "config to string returns a string type",
			cfg: config.Config{
				App: config.App{
					Name:    "mygoapp",
					Version: "1.2.3",
				},
				HTTP: config.HTTP{
					Port: "8080",
				},
				DB: config.DB{
					Type:     "postgres",
					Host:     "db",
					Port:     5432,
					User:     "user",
					Password: "THIS_IS_NOT_A_REAL_PASSWORD",
					DBName:   "db",
					SSLMode:  "disable",
				},
			},
			wantType: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.cfg.String()
			assert.IsType(t, tt.wantType, got)
		})
	}
}

// TestDB_GetURL tests the GetURL method of the DB struct.
//
// Test Cases:
//
// 1. Method returns the correct Postgres connection URL.
func TestDB_GetURL(t *testing.T) {
	t.Parallel()

	// Arrange
	unitTests := []struct {
		name     string
		input    config.DB
		expected string
	}{
		{
			name: "should return the correct Postgres connection URL",
			input: config.DB{
				Type:     "postgres",
				Host:     "db",
				Port:     5432,
				User:     "user",
				Password: "THIS_IS_NOT_A_REAL_PASSWORD",
				DBName:   "db",
				SSLMode:  "disable",
			},
			expected: "postgres://user:THIS_IS_NOT_A_REAL_PASSWORD@db:5432/db?sslmode=disable",
		},
	}
	for _, tt := range unitTests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			got := tt.input.GetURL()
			// Assert
			assert.Equal(t, tt.expected, got)
		})
	}
}

// TestDB_GetDSN tests the GetDSN method of the DB struct.
//
// Test Cases:
//
// 1. Method returns the correct Postgres DSN.
func TestDB_GetDSN(t *testing.T) {
	t.Parallel()

	// Arrange
	unitTests := []struct {
		name     string
		input    config.DB
		expected string
	}{
		{
			name: "should return the correct Postgres DSN",
			input: config.DB{
				Type:     "postgres",
				Host:     "db",
				Port:     5432,
				User:     "user",
				Password: "THIS_IS_NOT_A_REAL_PASSWORD",
				DBName:   "db",
				SSLMode:  "disable",
			},
			expected: "host=db port=5432 user=user password=THIS_IS_NOT_A_REAL_PASSWORD dbname=db sslmode=disable",
		},
	}
	for _, tt := range unitTests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			got := tt.input.GetDSN()

			// Assert

			// The strings contain the same elements, ignoring order
			gotSlice := strings.Split(got, " ")
			expectedSlice := strings.Split(tt.expected, " ")
			assert.ElementsMatch(t, expectedSlice, gotSlice)
		})
	}
}
