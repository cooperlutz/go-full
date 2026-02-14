package telemetree

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"

	"github.com/cooperlutz/go-full/app/config"
)

// InitTracer initializes an OpenTelemetry TracerProvider with a stdout exporter and an OTLP HTTP exporter.
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

	rd, err := ResourceDefinition(ctx)
	if err != nil {
		return nil, err
	}
	// For the demonstration, use sdktrace.AlwaysSample sampler to sample all traces.
	// In a production application, use sdktrace.ProbabilitySampler with a desired probability.
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(stdOutExporter),
		sdktrace.WithBatcher(httpExporter),
		sdktrace.WithResource(rd),
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
	tracer := otel.Tracer(config.ApplicationName)

	ctx, span := tracer.Start(ctx, spanName)

	span.SetAttributes(keyValues...)

	return ctx, span
}

// RecordError records an error in the current span.
func RecordError(ctx context.Context, err error, messages ...string) {
	span := trace.SpanFromContext(ctx)
	if span == nil {
		return
	}

	if len(messages) > 0 {
		for _, msg := range messages {
			span.AddEvent(msg)
		}
	}

	span.SetStatus(codes.Error, err.Error())
}
