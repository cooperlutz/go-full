package command_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/internal/pingpong/app/command"
	"github.com/cooperlutz/go-full/internal/pingpong/app/common"
)

// TestNewPingPongCommandResult tests the NewPingPongCommandResult function.
// Test Cases:
//
// 1. Success - passing a `Ping!` returns a PingPongCommandResult containing a `Ping!`
func TestNewPingPongCommandResult(t *testing.T) {
	t.Parallel()

	unitTests := []struct {
		name           string
		input          string
		expectedReturn *command.PingPongCommandResult
	}{
		{
			name:  "`Ping!` param returns a PingPongCommandResult containing a `Ping!` message",
			input: "Ping!",
			expectedReturn: &command.PingPongCommandResult{
				PingPongResult: &common.PingPongResult{
					Message: "Ping!",
				},
			},
		},
	}
	for _, tt := range unitTests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			got := command.NewPingPongCommandResult(tt.input)

			// Assert
			assert.Equal(t, tt.input, got.Message)
		})
	}
}

// TestNewPingPongCommand tests the NewPingPongCommand function.
// Test Cases:
//
// 1. Success - passing a `ping` returns a PingPongCommand containing a `ping`
func TestNewPingPongCommand(t *testing.T) {
	t.Parallel()

	unitTests := []struct {
		name           string
		input          string
		expectedReturn *command.PingPongCommand
	}{
		{
			name:  "`ping` param returns a PingPongCommand containing a ping message",
			input: "ping",
			expectedReturn: &command.PingPongCommand{
				Message: "ping",
			},
		},
	}
	for _, tt := range unitTests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			got := command.NewPingPongCommand(tt.input)

			// Assert
			assert.Equal(t, tt.input, got.Message)
		})
	}
}
