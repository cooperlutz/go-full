package ai

import (
	"testing"
	"go.opentelemetry.io/otel"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	t.Parallel()
	client := NewClient("dummy_api_key")
	assert.NotNil(t, client, "Client should not be nil")
	assert.NotNil(t, client.openAIClient, "OpenAI wrapped client should not be nil")
}

func TestTracerRegistered(t *testing.T) {
	t.Parallel()
	trc := otel.Tracer("pkg/ai")
	assert.NotNil(t, trc, "Tracer should be registered globally")
}
