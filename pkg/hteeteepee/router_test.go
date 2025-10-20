package hteeteepee_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/pkg/hteeteepee"
)

func TestNewRootRouterWithMiddleware(t *testing.T) {
	t.Parallel()
	router := hteeteepee.NewRootRouterWithMiddleware()

	// Check that the router is not nil
	assert.NotNil(t, router, "Router should not be nil")

	// Check that the router is of type *chi.Mux

	assert.IsType(t, &chi.Mux{}, router, "Router should be of type *chi.Mux")

	// Check that middleware is applied by inspecting the routes
	routes := []string{}
	err := chi.Walk(router, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		routes = append(routes, route)
		return nil
	})
	assert.NoError(t, err, "Error walking through routes")

	// Ensure that the /metrics endpoint is present
	assert.Contains(t, routes, "/metrics/*", "Router should contain /metrics endpoint")
}

func TestNewRouter(t *testing.T) {
	t.Parallel()
	operationSpanName := "test-operation"
	router := hteeteepee.NewRouter(operationSpanName)

	// Check that the router is not nil
	assert.NotNil(t, router, "Router should not be nil")

	// Check that the router is of type *chi.Mux
	assert.IsType(t, &chi.Mux{}, router, "Router should be of type *chi.Mux")

	// Check that OpenTelemetry middleware is applied by making a test request
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	router.Method("GET", "/test", testHandler)

	req, err := http.NewRequest("GET", "/test", nil)
	assert.NoError(t, err, "Error creating test request")

	// Use a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Check that the response status code is 200 OK
	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code 200 OK")
}
