package eeventdriven

import (
	"go.opentelemetry.io/otel/attribute"
)

// conf represents the configuration options available for subscriber
// middlewares and publisher decorators.
type conf struct {
	spanAttributes []attribute.KeyValue
}

// Option provides a convenience wrapper for simple options that can be
// represented as functions.
type Option func(*conf)

// WithSpanAttributes includes the given attributes to the generated Spans.
func WithSpanAttributes(attributes ...attribute.KeyValue) Option {
	return func(c *conf) {
		c.spanAttributes = attributes
	}
}
