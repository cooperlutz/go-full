package hteeteepee

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/trace"

	"github.com/cooperlutz/go-full/pkg/telemetree/metrics"
)

// Creates a new Chi multiplexer (mux) along with standard middleware and routes.
//
// The purpose of defining this function is to encapsulate the setup of the root router that all subsequent routers will be mounted on.
//
// Middleware includes:
//
// - RequestID: Injects a request ID into the context of each request.
//
// - RealIP: Sets the real IP address of the client.
//
// - OtelHTTP: Adds OpenTelemetry instrumentation for tracing HTTP requests.
//
// - Recoverer: Recovers from panics and returns a 500 error.
//
// Routes:
//
// - /metrics: Exposes Prometheus metrics for monitoring.
func NewRootRouterWithMiddleware() *chi.Mux {
	r := chi.NewRouter()
	r.Use(otelhttp.NewMiddleware("api-server"))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, //nolint:mnd // Maximum value not ignored by any of major browsers
	}))

	// Add Prometheus metrics endpoint
	r.Mount("/metrics", metrics.MetricsHandler())

	return r
}

// NewRouter creates a new Chi router with OpenTelemetry middleware for tracing.
func NewRouter(operationSpanName string) *chi.Mux {
	r := chi.NewRouter()

	defaultOtelOptions := []otelhttp.Option{
		otelhttp.WithSpanOptions(trace.WithSpanKind(trace.SpanKindServer)),
		otelhttp.WithSpanNameFormatter(func(operation string, r *http.Request) string {
			return operation
		}),
	}

	otelMW := otelhttp.NewMiddleware(operationSpanName, defaultOtelOptions...)

	r.Use(otelMW)

	return r
}
