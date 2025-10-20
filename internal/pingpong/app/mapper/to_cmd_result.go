package mapper

import (
	"github.com/cooperlutz/go-full/internal/pingpong/app/command"
	"github.com/cooperlutz/go-full/internal/pingpong/domain/entity"
)

func MapToCommandResult(entity entity.PingPongEntity) *command.PingPongCommandResult {
	command := &command.PingPongCommandResult{
		PingPongResult: MapToResult(entity),
	}

	return command
}
