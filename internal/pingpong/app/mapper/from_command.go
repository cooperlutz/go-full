package mapper

import (
	"github.com/cooperlutz/go-full/internal/pingpong/app/command"
	"github.com/cooperlutz/go-full/internal/pingpong/domain/entity"
)

// MapFromCommandPingPong maps a PingPongCommand to a VALIDATED PingPongEntity, returning a validation error if invalid.
func MapFromCommandPingPong(c command.PingPongCommand) (*entity.PingPongEntity, error) {
	entity, err := entity.New(
		c.Message,
	)
	if err != nil {
		return nil, err
	}

	return entity, nil
}
