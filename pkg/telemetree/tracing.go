package telemetree

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/cooperlutz/go-full/app/config"
)

func InitTracer(ctx context.Context, cfg config.Telemetry) (*sdktrace.TracerProvider, error) {
	// Create stdout exporter to be able to retrieve
	// the collected spans.
	stdOutExporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		return nil, err
	}

	httpExporter, err := otlptracehttp.New(ctx,
		otlptracehttp.WithInsecure(),
		otlptracehttp.WithEndpoint(cfg.TraceEndpoint),
	)
	if err != nil {
		panic(err)
	}

	// For the demonstration, use sdktrace.AlwaysSample sampler to sample all traces.
	// In a production application, use sdktrace.ProbabilitySampler with a desired probability.
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(stdOutExporter),
		sdktrace.WithBatcher(httpExporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(config.ApplicationName),
			semconv.ServiceVersion(config.ApplicationVersion),
		)),
	)

	otel.SetTracerProvider(tp)

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	return tp, err
}

// AddSpan adds an otel span to the existing trace.
func AddSpan(ctx context.Context, spanName string, keyValues ...attribute.KeyValue) (context.Context, trace.Span) {
	// tracer, ok := ctx.Value(tracerKey).(trace.Tracer)
	// if !ok || tracer == nil {
	// 	return ctx, trace.SpanFromContext(ctx)
	// }
	tracer := otel.Tracer(config.ApplicationName)

	ctx, span := tracer.Start(ctx, spanName)

	span.SetAttributes(keyValues...)

	return ctx, span
}

// RecordError records an error in the current span.
func RecordError(ctx context.Context, err error, msg string) {
	span := trace.SpanFromContext(ctx)
	if span == nil {
		return
	}

	span.SetStatus(codes.Error, msg)
}
