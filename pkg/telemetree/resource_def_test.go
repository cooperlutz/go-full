package telemetree_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/attribute"

	"github.com/cooperlutz/go-full/app/config"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

func TestResourceDefinition(t *testing.T) {
	config.ApplicationName = "my-service-name"
	config.ApplicationVersion = "1.0.0"

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
			Value: attribute.StringValue("unique-instance-id"),
		},
	}
	rd, err := telemetree.ResourceDefinition()
	rdAttributes := rd.Attributes()
	assert.NoError(t, err)
	assert.ElementsMatch(t, expectedAttributes, rdAttributes)
}
