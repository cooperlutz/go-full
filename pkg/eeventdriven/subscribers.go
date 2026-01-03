package eeventdriven

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"go.opentelemetry.io/otel"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
	"go.opentelemetry.io/otel/trace"
)

const subscriberTracerName = "watermill/subscriber"

// Trace defines a middleware that will add tracing.
func Trace(options ...Option) message.HandlerMiddleware {
	return func(h message.HandlerFunc) message.HandlerFunc {
		return TraceHandler(h, options...)
	}
}

// TraceHandler decorates a watermill HandlerFunc to add tracing when a message is received.
func TraceHandler(h message.HandlerFunc, options ...Option) message.HandlerFunc {
	tracer := otel.Tracer(subscriberTracerName)
	config := &conf{}

	for _, opt := range options {
		opt(config)
	}

	spanOptions := []trace.SpanStartOption{
		trace.WithSpanKind(trace.SpanKindConsumer),
		trace.WithAttributes(config.spanAttributes...),
	}

	return func(msg *message.Message) ([]*message.Message, error) {
		spanName := message.HandlerNameFromCtx(msg.Context())
		ctx, span := tracer.Start(msg.Context(), spanName, spanOptions...)
		span.SetAttributes(
			semconv.MessagingDestinationNameKey.String(message.SubscribeTopicFromCtx(ctx)),
			semconv.MessagingSystemKey.String("watermill"),
			semconv.MessagingOperationReceive,
		)
		msg.SetContext(ctx)

		events, err := h(msg)
		if err != nil {
			span.RecordError(err)
		}

		span.End()

		return events, err
	}
}

// TraceConsumerHandler decorates a watermill NoPublishHandlerFunc to add tracing when a message is received.
func TraceConsumerHandler(h message.NoPublishHandlerFunc, options ...Option) message.NoPublishHandlerFunc {
	decoratedHandler := TraceHandler(func(msg *message.Message) ([]*message.Message, error) {
		return nil, h(msg)
	}, options...)

	return func(msg *message.Message) error {
		_, err := decoratedHandler(msg)

		return err
	}
}

func InitTracedRouter() (*message.Router, error) {
	router, err := message.NewRouter(message.RouterConfig{}, watermill.NopLogger{})
	if err != nil {
		return nil, err
	}

	router.AddMiddleware(Trace())

	return router, nil
}
