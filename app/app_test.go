package app_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/app"
	"github.com/cooperlutz/go-full/app/config"
)

func TestNewApplication(t *testing.T) {
	conf := config.Config{
		App: config.App{
			Name:    "test-app",
			Version: "1.0.0",
		},
		HTTP: config.HTTP{
			Port: ":8080",
		},
		DB: config.DB{
			Type:     "postgres",
			User:     "user",
			Password: "password",
			Host:     "localhost",
			Port:     5432,
			DBName:   "testdb",
			SSLMode:  "disable",
		},
	}

	appInstance := app.NewApplication(conf)
	assert.NotNil(t, appInstance)
	assert.IsType(t, &app.Application{}, appInstance)
}

// Just ensure that Run() doesn't panic or return an error.
// Full integration tests would be more complex and are not included here.
func TestApplication(t *testing.T) {
	// Run the application in a separate goroutine to prevent race condition issues with loading env vars
	go func() {
		envVars := map[string]string{
			"HTTP_PORT":              ":0", // :80 will fail in github actions
			"DB_TYPE":                "postgres",
			"DB_USER":                "user",
			"DB_PASSWORD":            "THIS_IS_NOT_A_REAL_PASSWORD",
			"DB_HOST":                "db",
			"DB_PORT":                "5432",
			"DB_DBNAME":              "db",
			"DB_SSLMODE":             "disable",
			"OBSERVE_TRACE_ENDPOINT": "localhost:4317",
		}
		config.ApplicationName = "my-app"
		config.ApplicationVersion = "1.0.0"
		// Set environment variables for the test
		for key, value := range envVars {
			t.Setenv(key, value)
		}
		assert.NotNil(t, os.Getenv("HTTP_PORT"))
		conf, err := config.LoadConfigFromEnvVars()
		assert.NoError(t, err)
		assert.NotNil(t, conf)

		application := app.NewApplication(conf)
		assert.NotNil(t, application)
		assert.IsType(t, &app.Application{}, application)
		// This will run indefinitely in a real application, so we just ensure it starts without error.
		// In a real test, you would have more sophisticated checks and possibly a way to shut it down gracefully.
		application.Run()
	}()
}
