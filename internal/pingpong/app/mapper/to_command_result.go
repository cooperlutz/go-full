package mapper

import (
	"github.com/cooperlutz/go-full/internal/pingpong/app/command"
	"github.com/cooperlutz/go-full/internal/pingpong/domain/entity"
)

func MapToCommandResult(e entity.PingPongEntity) command.PingPongCommandResult {
	res := MapToResult(e)

	command := command.PingPongCommandResult{
		PingPongResult: &res,
	}

	return command
}
