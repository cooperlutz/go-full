package query

import (
	"github.com/cooperlutz/go-full/internal/pingpong/app/common"
	"github.com/cooperlutz/go-full/internal/pingpong/domain/entity"
)

type FindAllQueryResponse struct {
	PingPongs []*common.PingPongResult
}

type FindAllQueryResponseRaw struct {
	Entities []entity.PingPongEntity
}
