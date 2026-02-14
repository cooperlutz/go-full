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

func TestResourceDefinition(t *testing.T) {
	// Arrange
	config.ApplicationName = "my-service-name"
	config.ApplicationVersion = "1.0.0"
	config.ApplicationInstanceID = uuid.New().String()
	expectedAttributes := []attribute.KeyValue{
		{
			Key:   "service.name",
			Value: attribute.StringValue("my-service-name"),
		},
		{
			Key:   "service.version",
			Value: attribute.StringValue("1.0.0"),
		},
		{
			Key:   "service.instance.id",
			Value: attribute.StringValue(config.ApplicationInstanceID),
		},
	}
	rd, err := telemetree.ResourceDefinition(context.Background())
	rdAttributes := rd.Attributes()
	// Assert
	assert.NoError(t, err)
	assert.ElementsMatch(t, expectedAttributes, rdAttributes)
}
