package telemetree_test

import (
	"context"
	"testing"

	"go.opentelemetry.io/otel/attribute"

	"github.com/cooperlutz/go-full/app/config"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

func TestInitTracer_Success(t *testing.T) {
	// t.Parallel() DO NOT PARALLELIZE

	cfg := config.Telemetry{
		TraceEndpoint: "localhost:4318",
	}
	ctx := context.Background()

	tp, err := telemetree.InitTracer(ctx, cfg)
	if err != nil {
		t.Fatalf("InitTracer returned error: %v", err)
	}
	if tp == nil {
		t.Fatal("InitTracer returned nil TracerProvider")
	}
}

func TestAddSpan_Basic(t *testing.T) {
	ctx := context.Background()
	spanName := "test-span"

	newCtx, span := telemetree.AddSpan(ctx, spanName)
	if newCtx == nil {
		t.Error("AddSpan returned nil context")
	}
	if span == nil {
		t.Error("AddSpan returned nil span")
	}
	// The Span type does not have a Name() method, so we cannot check the span name directly.
	span.End()
}

func TestAddSpan_WithAttributes(t *testing.T) {
	ctx := context.Background()
	spanName := "span-with-attrs"
	attr1 := attribute.String("key1", "value1")
	attr2 := attribute.Int("key2", 42)

	newCtx, span := telemetree.AddSpan(ctx, spanName, attr1, attr2)
	if newCtx == nil {
		t.Error("AddSpan returned nil context")
	}
	if span == nil {
		t.Error("AddSpan returned nil span")
	}
	span.End()
}
