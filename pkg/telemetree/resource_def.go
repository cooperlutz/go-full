package telemetree

import (
	"context"

	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"

	"github.com/cooperlutz/go-full/app/config"
)

func ResourceDefinition() (*resource.Resource, error) {
	res, err := resource.New(
		context.Background(), // Use a background context
		resource.WithAttributes(
			semconv.ServiceNameKey.String(config.ApplicationName),
			semconv.ServiceVersionKey.String(config.ApplicationVersion),
			semconv.ServiceInstanceIDKey.String("unique-instance-id"),
		),
	)

	return res, err
}
