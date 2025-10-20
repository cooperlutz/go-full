package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func MetricsHandler() http.Handler {
	// router := chi.NewRouter()
	reg := prometheus.NewRegistry()
	reg.MustRegister(
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		collectors.NewGoCollector(),
	)

	// Expose /metrics HTTP endpoint using the created custom registry.
	mw := NewPromMiddleware(
		reg, nil).
		WrapHandler("/metrics", promhttp.HandlerFor(
			reg,
			promhttp.HandlerOpts{}),
		)

	return mw
}
