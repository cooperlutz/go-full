package config_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/app/config"
)

// TestLoadConfigFromEnvVars tests the LoadConfigFromEnvVars function.
func TestLoadConfigFromEnvVars(t *testing.T) {
	// Arrange
	tests := []struct {
		name        string
		envVars     map[string]string
		expectedCfg config.Config
	}{
		{
			name: "success",
			envVars: map[string]string{
				"HTTP_PORT":              "8080",
				"DB_TYPE":                "postgres",
				"DB_USER":                "user",
				"DB_PASSWORD":            "THIS_IS_NOT_A_REAL_PASSWORD",
				"DB_HOST":                "db",
				"DB_PORT":                "5432",
				"DB_DBNAME":              "db",
				"DB_SSLMODE":             "disable",
				"OBSERVE_TRACE_ENDPOINT": "localhost:4317",
				"SEC_JWT_SECRET":         "THIS_IS_NOT_A_REAL_PASSWORD",
			},
			expectedCfg: config.Config{
				App: config.App{
					Name:    "go-full",
					Version: "",
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
				Telemetry: config.Telemetry{
					TraceEndpoint: "localhost:4317",
				},
				Security: config.Security{
					JWTSecret:      "THIS_IS_NOT_A_REAL_PASSWORD",
					AccessTokenTTL: time.Duration(15) * time.Minute,
				},
			},
		},
		{
			name: "test type conversion default values: invalid DB_PORT & METRICS_ENABLED is overriden with default",
			envVars: map[string]string{
				"HTTP_PORT":              "8080",
				"DB_TYPE":                "postgres",
				"DB_USER":                "user",
				"DB_PASSWORD":            "THIS_IS_NOT_A_REAL_PASSWORD",
				"DB_HOST":                "db",
				"DB_PORT":                "this-is-not-an-int",
				"DB_DBNAME":              "db",
				"DB_SSLMODE":             "disable",
				"OBSERVE_TRACE_ENDPOINT": "localhost:4317",
				"SEC_JWT_SECRET":         "THIS_IS_NOT_A_REAL_PASSWORD",
			},
			expectedCfg: config.Config{
				App: config.App{
					Name:    "go-full",
					Version: "",
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
				Telemetry: config.Telemetry{
					TraceEndpoint: "localhost:4317",
				},
				Security: config.Security{
					JWTSecret:      "THIS_IS_NOT_A_REAL_PASSWORD",
					AccessTokenTTL: time.Duration(15) * time.Minute,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			for key, value := range tt.envVars {
				t.Setenv(key, value)
			}
			// Act
			cfg, err := config.LoadConfigFromEnvVars()
			// Assert
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedCfg, cfg)
		})
	}
}

// TestLoadConfigFromEnvVars_MissingVars tests the LoadConfigFromEnvVars function.
func TestLoadConfigFromEnvVars_MissingVars(t *testing.T) {
	// Arrange
	tests := []struct {
		name        string
		envVars     map[string]string
		expectedCfg config.Config
		expectedErr error
	}{
		{
			name: "success",
			envVars: map[string]string{
				"HTTP_PORT":              "",
				"DB_TYPE":                "postgres",
				"DB_USER":                "user",
				"DB_PASSWORD":            "THIS_IS_NOT_A_REAL_PASSWORD",
				"DB_HOST":                "db",
				"DB_PORT":                "5432",
				"DB_DBNAME":              "db",
				"DB_SSLMODE":             "disable",
				"OBSERVE_TRACE_ENDPOINT": "localhost:4317",
				"SEC_JWT_SECRET":         "THIS_IS_NOT_A_REAL_PASSWORD",
			},
			expectedCfg: config.Config{},
			expectedErr: config.ErrEnvVarValueMissing{VarName: "HTTP_PORT"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			for key, value := range tt.envVars {
				t.Setenv(key, value)
			}
			// Act
			cfg, err := config.LoadConfigFromEnvVars()
			// Assert
			assert.Error(t, err)
			assert.Equal(t, tt.expectedErr, err)
			assert.Equal(t, tt.expectedCfg, cfg)
			assert.Equal(t, "Missing environment variable: HTTP_PORT", err.Error())
		})
	}
}
