package event

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPingPongReceived(t *testing.T) {
	pingID := "ping123"
	message := "Hello, Ping!"

	event := NewPingPongReceived(pingID, message)

	assert.Equal(t, pingID, event.PingID)
	assert.Equal(t, message, event.Message)
}
