package telemetree_test

import (
	"testing"

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

	mp, err := telemetree.InitMeter()
	assert.NoError(t, err)
	assert.NotNil(t, mp)

	// Check that the MeterProvider's resource has expected attributes
	// Since MeterProvider does not expose Resource(), test the resource via ResourceDefinition
	rd, err := telemetree.ResourceDefinition()
	assert.NoError(t, err)
	attrs := rd.Attributes()

	expected := []attribute.KeyValue{
		attribute.String("service.name", "test-service"),
		attribute.String("service.version", "2.0.0"),
		attribute.String("service.instance.id", "unique-instance-id"),
	}

	// Assert
	assert.Subset(t, attrs, expected)
}
