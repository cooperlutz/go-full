package command

import (
	"github.com/cooperlutz/go-full/internal/pingpong/app/common"
)

// PingPongCommand represents a command to perform a PingPong operation.
type PingPongCommand struct {
	Message string
}

// NewPingPongCommand creates a new PingPongCommand.
func NewPingPongCommand(msg string) *PingPongCommand {
	return &PingPongCommand{
		Message: msg,
	}
}

// PingPongCommandResult represents the result of a PingPongCommand.
type PingPongCommandResult struct {
	*common.PingPongResult
}

// NewPingPongCommandResult creates a new PingPongCommandResult.
func NewPingPongCommandResult(msg string) PingPongCommandResult {
	return PingPongCommandResult{
		PingPongResult: &common.PingPongResult{
			Message: msg,
		},
	}
}
