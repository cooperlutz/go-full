package metrics_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cooperlutz/go-full/pkg/telemetree/metrics"
)

func TestMetricsHandler(t *testing.T) {
	t.Parallel()

	metricHandler := metrics.MetricsHandler()
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/metrics", nil)
	if err != nil {
		t.Fatal(err)
	}
	metricHandler.ServeHTTP(rr, req)
}
