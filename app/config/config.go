package config

import (
	"encoding/json"
	"fmt"
	"strings"
)

var (
	ApplicationVersion string             //nolint:gochecknoglobals // set via build flags
	ApplicationName    string = "go-full" //nolint:gochecknoglobals // set via build flags
)

// Config.App settings.
type App struct {
	Name    string `env:"APP_NAME,required"`
	Version string `env:"APP_VERSION,required"`
}

// Config.DB settings.
type DB struct {
	Type     string `env:"DB_TYPE"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Host     string `env:"DB_HOST"`
	Port     int    `env:"DB_PORT"`
	DBName   string `env:"DB_DBNAME"`
	SSLMode  string `env:"DB_SSLMODE"`
}

// Config.HTTP settings.
type HTTP struct {
	Port string `env:"HTTP_PORT,required"`
}

type Telemetry struct {
	TraceEndpoint string
}

// Config holds the application configuration settings.
// Values are populated from environment variables.
//
// The Config struct and its nested structs are intentionally not pointer types to prevent mutation
// of configuration values at runtime.
type Config struct {
	App       App
	HTTP      HTTP
	Telemetry Telemetry
	DB        DB
}

func (c Config) String() string {
	c.DB.Password = "****" // hide password

	jsonData, _ := json.MarshalIndent(c, "", "  ") //nolint:errcheck // ignoring error

	return string(jsonData)
}

// GetDSN returns the Data Source Name for connecting to the database.
//
// Example: "user=youruser password=yourpassword host=localhost port=5432 dbname=yourdb sslmode=disable".
func (db DB) GetDSN() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s", db.User, db.Password, db.Host, db.Port, db.DBName, db.SSLMode))

	return b.String()
}

// GetURL returns the database connection URL.
//
// Example: "postgres://youruser:yourpassword@localhost:5432/yourdb?sslmode=disable"
func (db DB) GetURL() string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=%s", db.Type, db.User, db.Password, db.Host, db.Port, db.DBName, db.SSLMode)
}
