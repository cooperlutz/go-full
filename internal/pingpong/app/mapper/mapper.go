package mapper

import (
	"github.com/cooperlutz/go-full/internal/pingpong/app/command"
	"github.com/cooperlutz/go-full/internal/pingpong/app/common"
	"github.com/cooperlutz/go-full/internal/pingpong/app/query"
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

func MapToCommandResult(e entity.PingPongEntity) command.PingPongCommandResult {
	res := MapToResult(e)

	command := command.PingPongCommandResult{
		PingPongResult: &res,
	}

	return command
}

func MapToResult(e entity.PingPongEntity) common.PingPongResult {
	result := common.PingPongResult{
		Message: e.GetMessage(),
	}

	return result
}

func MapToRawResult(e entity.PingPongEntity) common.PingPongRawResult {
	return common.PingPongRawResult{
		ID:        e.PingPongID.String(),
		Message:   e.Message,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
		Deleted:   e.Deleted,
		DeletedAt: e.DeletedAt,
	}
}

func MapListToQueryResponse(l entity.ListOfPingPongs) query.FindAllQueryResponse {
	var resultingPings []common.PingPongResult

	for _, pp := range l.PingPongs {
		resultingPings = append(resultingPings, MapToResult(pp))
	}

	return query.FindAllQueryResponse{
		PingPongs: resultingPings,
	}
}

func MapListToQueryResponseRaw(l entity.ListOfPingPongs) query.FindAllQueryResponseRaw {
	return query.FindAllQueryResponseRaw{
		Entities: l.PingPongs,
	}
}
