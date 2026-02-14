package telemetree

import (
	"context"

	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"

	"github.com/cooperlutz/go-full/app/config"
)

func ResourceDefinition(ctx context.Context) (*resource.Resource, error) {
	res, err := resource.New(
		ctx,
		resource.WithAttributes(
			semconv.ServiceInstanceIDKey.String(config.ApplicationInstanceID),
			semconv.ServiceName(config.ApplicationName),
			semconv.ServiceVersion(config.ApplicationVersion),
		),
	)

	return res, err
}
