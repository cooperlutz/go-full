package hteeteepee_test

import (
	"context"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/pkg/hteeteepee"
)

func TestNewSSEBroker(t *testing.T) {
	sseBroker := hteeteepee.NewSSEBroker()
	assert.NotNil(t, sseBroker, "Expected NewSSEBroker to return a non-nil broker")
	sseBroker.Start()
	assert.NotNil(t, sseBroker, "Expected SSEBroker to be running after Start()")
}

func TestSSEBroker_ServeHTTP(t *testing.T) {
	// arrange
	sseBroker := hteeteepee.NewSSEBroker()
	sseBroker.Start()
	rr := httptest.NewRecorder()
	ctx, cancel := context.WithCancel(context.Background())
	req := httptest.NewRequest("GET", "/events", nil).WithContext(ctx)

	// Call ServeHTTP in a separate goroutine so that it can block waiting for events while the broker is notified
	serveHTTPDone := make(chan struct{})
	go func() {
		defer close(serveHTTPDone)
		sseBroker.ServeHTTP(rr, req)
	}()

	// Notify the broker, then cancel the context to unblock ServeHTTP
	select {
	case <-time.After(5 * time.Second):
		t.Fatal("Test timeout waiting to notify")
	default:
		time.Sleep(100 * time.Millisecond)
		sseBroker.NotifyString("Hello, SSE!")
		time.Sleep(100 * time.Millisecond)
		cancel()
	}

	// Wait for ServeHTTP to finish before reading rr to avoid a data race
	select {
	case <-serveHTTPDone:
	case <-time.After(5 * time.Second):
		t.Fatal("Test timeout waiting for ServeHTTP to finish")
	}

	// assert
	assert.Equal(t, 200, rr.Code, "Expected status code 200")
	assert.Equal(t, "text/event-stream", rr.Header().Get("Content-Type"), "Expected Content-Type 'text/event-stream'")
	assert.Contains(t, rr.Body.String(), "data: Hello, SSE!", "Expected SSE message 'Hello, SSE!' in response body")
}
