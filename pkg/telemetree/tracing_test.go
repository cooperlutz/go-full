package telemetree_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/attribute"

	"github.com/cooperlutz/go-full/app/config"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

func TestInitTracer_Success(t *testing.T) {
	t.Parallel()

	// Arrange
	cfg := config.Telemetry{
		TraceEndpoint: "localhost:4318",
	}
	ctx := context.Background()

	// Act
	tp, err := telemetree.InitTracer(ctx, cfg)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, tp)

	// Cleanup
	defer func() {
		_ = tp.Shutdown(ctx)
	}()
}

func TestAddSpan_Basic(t *testing.T) {
	t.Parallel()
	// Arrange
	ctx := context.Background()
	spanName := "test-span"

	newCtx, span := telemetree.AddSpan(ctx, spanName)
	assert.NotNil(t, span)
	assert.NotNil(t, newCtx)

	// Cleanup
	span.End()
}

func TestAddSpan_WithAttributes(t *testing.T) {
	t.Parallel()
	// Arrange
	ctx := context.Background()
	spanName := "span-with-attrs"
	attr1 := attribute.String("key1", "value1")
	attr2 := attribute.Int("key2", 42)

	// Act
	newCtx, span := telemetree.AddSpan(ctx, spanName, attr1, attr2)

	// Assert
	assert.NotNil(t, span)
	assert.NotNil(t, newCtx)

	// Cleanup
	span.End()
}
