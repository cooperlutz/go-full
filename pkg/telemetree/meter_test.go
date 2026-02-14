package telemetree_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/attribute"

	"github.com/cooperlutz/go-full/app/config"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

func TestInitMeter(t *testing.T) {
	// t.Parallel() DO NOT PARALLELIZE

	// Set config values for resource definition
	config.ApplicationName = "test-service"
	config.ApplicationVersion = "2.0.0"
	config.ApplicationInstanceID = uuid.New().String()

	mp, err := telemetree.InitMeter(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, mp)

	// Check that the MeterProvider's resource has expected attributes
	// Since MeterProvider does not expose Resource(), test the resource via ResourceDefinition
	rd, err := telemetree.ResourceDefinition(context.Background())
	assert.NoError(t, err)
	attrs := rd.Attributes()

	expected := []attribute.KeyValue{
		attribute.String("service.name", "test-service"),
		attribute.String("service.version", "2.0.0"),
		attribute.String("service.instance.id", config.ApplicationInstanceID),
	}

	// Assert
	assert.Subset(t, attrs, expected)
}
