package hteeteepee_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/app/config"
	"github.com/cooperlutz/go-full/pkg/hteeteepee"
)

func TestNewHTTPServer(t *testing.T) {
	t.Parallel()
	mockConfig := config.Config{
		HTTP: config.HTTP{
			Port: ":8080",
		},
	}
	router := chi.NewRouter()

	server := hteeteepee.NewHTTPServer(mockConfig, router)

	assert.NotNil(t, server)
	assert.Equal(t, mockConfig, server.Config)
	assert.Equal(t, router, server.Router)
	assert.NotNil(t, server.Server)
	assert.Equal(t, ":8080", server.Server.Addr)
	assert.Equal(t, 1*time.Second, server.Server.ReadHeaderTimeout)
}

func TestRegisterController(t *testing.T) {
	t.Parallel()
	mockConfig := config.Config{
		HTTP: config.HTTP{
			Port: ":8080",
		},
	}
	router := chi.NewRouter()
	server := hteeteepee.NewHTTPServer(mockConfig, router)

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	server.RegisterController("/test", testHandler)

	// Use httptest to verify the route is mounted
	req, _ := http.NewRequest("GET", "/test", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestRegisterController_MountsHandler(t *testing.T) {
	t.Parallel()
	mockConfig := config.Config{
		HTTP: config.HTTP{
			Port: ":8080",
		},
	}
	router := chi.NewRouter()
	server := hteeteepee.NewHTTPServer(mockConfig, router)

	called := false
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	})

	server.RegisterController("/mounted", testHandler)

	req, _ := http.NewRequest("GET", "/mounted", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.True(t, called, "Handler should be called when route is mounted")
}

func TestRegisterController_DifferentEndpoints(t *testing.T) {
	t.Parallel()
	mockConfig := config.Config{
		HTTP: config.HTTP{
			Port: ":8080",
		},
	}
	router := chi.NewRouter()
	server := hteeteepee.NewHTTPServer(mockConfig, router)

	handler1 := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
	})
	handler2 := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
	})

	server.RegisterController("/one", handler1)
	server.RegisterController("/two", handler2)

	req1, _ := http.NewRequest("GET", "/one", nil)
	rr1 := httptest.NewRecorder()
	router.ServeHTTP(rr1, req1)
	assert.Equal(t, http.StatusCreated, rr1.Code)

	req2, _ := http.NewRequest("GET", "/two", nil)
	rr2 := httptest.NewRecorder()
	router.ServeHTTP(rr2, req2)
	assert.Equal(t, http.StatusAccepted, rr2.Code)
}

func TestHTTPServer_Run(t *testing.T) {
	t.Parallel()
	// Just ensure that Run() doesn't panic or return an error.
	// Full integration tests would be more complex and are not included here.
	server := hteeteepee.NewHTTPServer(config.Config{
		HTTP: config.HTTP{
			Port: ":0", // Use :0 to let the OS assign an available port
		},
	}, chi.NewRouter())

	go server.Run()
}
