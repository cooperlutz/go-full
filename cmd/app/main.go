package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/cooperlutz/go-full/app"
	"github.com/cooperlutz/go-full/app/config"
	"github.com/cooperlutz/go-full/pkg/deebee/migration"
)

// main is the entry point of the application.
func main() {
	// Command-line flags
	flagMigrate := flag.Bool("migrate", false, "run migrations before starting the app")
	flag.Parse()

	// Load configuration
	conf, err := config.LoadConfigFromEnvVars()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	// Run migrations if the flag is set
	if *flagMigrate {
		slog.Info("Running migrations before starting the app...")
		migration.Migrate(conf.DB.Type, conf.DB.GetURL())
	}

	// Create and start the application
	application := app.NewApplication(conf)
	// Start the application
	application.Run()
}
