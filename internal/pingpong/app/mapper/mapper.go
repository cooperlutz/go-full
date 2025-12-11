package mapper

import (
	"github.com/cooperlutz/go-full/internal/pingpong/app/command"
	"github.com/cooperlutz/go-full/internal/pingpong/app/common"
	"github.com/cooperlutz/go-full/internal/pingpong/app/query"
	"github.com/cooperlutz/go-full/internal/pingpong/domain/entity"
)

// MapFromCommandPingPong maps a PingPongCommand to a VALIDATED PingPongEntity, returning a validation error if invalid.
func MapFromCommandPingPong(c command.PingPongCommand) (entity.PingPongEntity, error) {
	pingpong, err := entity.New(
		c.Message,
	)
	if err != nil {
		return entity.PingPongEntity{}, err
	}

	return pingpong, nil
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
		ID:        e.GetIdString(),
		Message:   e.GetMessage(),
		CreatedAt: e.GetCreatedAtTime(),
		UpdatedAt: e.GetUpdatedAtTime(),
		Deleted:   e.IsDeleted(),
		DeletedAt: e.GetDeletedAtTime(),
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
	var resultingPingPongs []common.PingPongRawResult

	for _, pp := range l.PingPongs {
		resultingPingPongs = append(resultingPingPongs, MapToRawResult(pp))
	}

	return query.FindAllQueryResponseRaw{
		Entities: resultingPingPongs,
	}
}
