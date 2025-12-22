package metrics_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/pkg/telemetree/metrics"
)

func TestMetricsHandler(t *testing.T) {
	t.Parallel()

	// Arrange
	metricHandler := metrics.MetricsHandler()
	rr := httptest.NewRecorder()
	// Act
	req, err := http.NewRequest("GET", "/metrics", nil)
	if err != nil {
		t.Fatal(err)
	}
	metricHandler.ServeHTTP(rr, req)
	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)
}
