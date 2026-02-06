package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

type ErrEnvVarValueMissing struct {
	VarName string
}

func (e ErrEnvVarValueMissing) Error() string {
	return "Missing environment variable: " + e.VarName
}

// LoadEnvironmentVariables loads configuration settings from environment variables.
// if a variable is not set, the function returns an error.
func LoadEnvironmentVariables() (map[string]string, error) {
	keyVals := map[string]string{
		// Observability
		"OBSERVE_TRACE_ENDPOINT": os.Getenv("OBSERVE_TRACE_ENDPOINT"),
		// HTTP Server
		"HTTP_PORT": os.Getenv("HTTP_PORT"),
		// Database
		"DB_TYPE":     os.Getenv("DB_TYPE"),
		"DB_USER":     os.Getenv("DB_USER"),
		"DB_PASSWORD": os.Getenv("DB_PASSWORD"), // SENSITIVE
		"DB_HOST":     os.Getenv("DB_HOST"),
		"DB_PORT":     os.Getenv("DB_PORT"),
		"DB_DBNAME":   os.Getenv("DB_DBNAME"),
		"DB_SSLMODE":  os.Getenv("DB_SSLMODE"),
		// Security
		"SEC_JWT_SECRET": os.Getenv("SEC_JWT_SECRET"), // SENSITIVE
	}

	for k, v := range keyVals {
		if v == "" {
			return nil, ErrEnvVarValueMissing{VarName: k}
		}
	}

	return keyVals, nil
}

// LoadConfigFromEnvVars loads configuration settings from environment variables.
func LoadConfigFromEnvVars() (Config, error) {
	loadedEnvVars, err := LoadEnvironmentVariables()
	if err != nil {
		return Config{}, err
	}

	loadedCfg := Config{
		App: App{
			Name:    ApplicationName,
			Version: ApplicationVersion,
		},
		HTTP: HTTP{
			Port: loadedEnvVars["HTTP_PORT"],
		},
		Telemetry: Telemetry{
			TraceEndpoint: loadedEnvVars["OBSERVE_TRACE_ENDPOINT"],
		},
		DB: DB{
			Type:     loadedEnvVars["DB_TYPE"],
			User:     loadedEnvVars["DB_USER"],
			Password: loadedEnvVars["DB_PASSWORD"],
			Host:     loadedEnvVars["DB_HOST"],
			Port:     getEnvAsInt("DB_PORT", 5432), //nolint:mnd // default port for Postgres
			DBName:   loadedEnvVars["DB_DBNAME"],
			SSLMode:  loadedEnvVars["DB_SSLMODE"],
		},
		Security: Security{
			JWTSecret:      loadedEnvVars["SEC_JWT_SECRET"],
			AccessTokenTTL: time.Duration(15) * time.Minute, //nolint:mnd // default 15 minutes
		},
	}
	log.Printf("Loaded config: %+v\n", loadedCfg.String())

	return loadedCfg, nil
}

func getEnvAsInt(key string, defaultVal int) int {
	if valStr, ok := os.LookupEnv(key); ok {
		if val, err := strconv.Atoi(valStr); err == nil {
			return val
		}
	}

	return defaultVal
}
