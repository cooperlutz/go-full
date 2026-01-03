package query

import (
	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/internal/pingpong/app/common"
)

type FindOneByID struct {
	ID uuid.UUID
}

type FindOneByIDResponse struct {
	common.PingPongRawResult
}

type FindAllQueryResponse struct {
	PingPongs []common.PingPongResult
}

type FindAllQueryResponseRaw struct {
	Entities []common.PingPongRawResult
}
