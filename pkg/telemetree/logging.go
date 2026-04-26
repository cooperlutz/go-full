package telemetree

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/sdk/log"

	"github.com/cooperlutz/go-full/app/config"
)

func InitLogger(ctx context.Context, cfg config.Telemetry) error {
	// Initialize the OpenTelemetry logger and set it as the default logger for the application
	logger := otelslog.NewLogger("go-full-logger")

	// Initialize the OTLP log exporter with the provided configuration
	logExporter, err := otlploghttp.New(
		ctx,
		otlploghttp.WithInsecure(),
		otlploghttp.WithEndpoint(cfg.OTLPHttpEndpoint),
	)
	if err != nil {
		return err
	}

	// Define resource attributes for the logger (e.g., service name, environment)
	res, err := ResourceDefinition(ctx)
	if err != nil {
		return err
	}

	// Create a new OpenTelemetry logger provider with the OTLP log exporter and resource attributes
	lp := log.NewLoggerProvider(
		log.WithResource(res),
		log.WithProcessor(
			log.NewBatchProcessor(logExporter),
		),
	)

	// Set the global logger provider to our OpenTelemetry logger
	global.SetLoggerProvider(lp)

	// Set the default slog logger to our OpenTelemetry logger
	slog.SetDefault(logger)

	return nil
}
