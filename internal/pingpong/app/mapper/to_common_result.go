package mapper

import (
	"github.com/cooperlutz/go-full/internal/pingpong/app/common"
	"github.com/cooperlutz/go-full/internal/pingpong/domain/entity"
)

func MapToResult(entity entity.PingPongEntity) common.PingPongResult {
	result := common.PingPongResult{
		Message: entity.GetMessage(),
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
