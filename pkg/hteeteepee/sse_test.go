package hteeteepee_test

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cooperlutz/go-full/pkg/hteeteepee"
	"github.com/stretchr/testify/assert"
)

func TestNewSSEBroker(t *testing.T) {
	sseBroker := hteeteepee.NewSSEBroker()
	assert.NotNil(t, sseBroker, "Expected NewSSEBroker to return a non-nil broker")
	sseBroker.Start()
	assert.NotNil(t, sseBroker, "Expected SSEBroker to be running after Start()")

}

func TestSSEBroker_ServerHTTP(t *testing.T) {
	// arrange
	sseBroker := hteeteepee.NewSSEBroker()
	sseBroker.Start()
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/events", nil)

	// Notify the broker in a separate goroutine so that the ServeHTTP call can receive the event
	done := make(chan struct{})
	go func() {
		time.Sleep(100 * time.Millisecond)
		sseBroker.NotifyString("Hello, SSE!")
		time.Sleep(100 * time.Millisecond)
		close(done)
	}()

	// Call ServeHTTP in a separate goroutine so that it can block waiting for events while the broker is notified
	go sseBroker.ServeHTTP(rr, req)

	// Wait for the notification goroutine to finish so that the SSE message has been sent
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("Test timeout")
	}

	// assert
	assert.Equal(t, 200, rr.Code, "Expected status code 200")
	assert.Equal(t, "text/event-stream", rr.Header().Get("Content-Type"), "Expected Content-Type 'text/event-stream'")
	assert.Contains(t, rr.Body.String(), "data: Hello, SSE!", "Expected SSE message 'Hello, SSE!' in response body")
}
