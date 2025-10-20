package command

import (
	"github.com/cooperlutz/go-full/internal/pingpong/app/common"
)

type PingPongCommand struct {
	Message string
}

func NewPingPongCommand(msg string) *PingPongCommand {
	return &PingPongCommand{
		Message: msg,
	}
}

type PingPongCommandResult struct {
	*common.PingPongResult
}

func NewPingPongCommandResult(msg string) *PingPongCommandResult {
	return &PingPongCommandResult{
		PingPongResult: &common.PingPongResult{
			Message: msg,
		},
	}
}
