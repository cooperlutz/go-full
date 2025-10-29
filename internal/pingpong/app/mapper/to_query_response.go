package mapper

import (
	"github.com/cooperlutz/go-full/internal/pingpong/app/common"
	"github.com/cooperlutz/go-full/internal/pingpong/app/query"
	"github.com/cooperlutz/go-full/internal/pingpong/domain/entity"
)

func MapListToQueryResponse(pps entity.ListOfPingPongs) query.FindAllQueryResponse {
	var resultingPings []common.PingPongResult

	for _, pp := range pps.PingPongs {
		resultingPings = append(resultingPings, MapToResult(pp))
	}

	return query.FindAllQueryResponse{
		PingPongs: resultingPings,
	}
}

func MapListToQueryResponseRaw(pps entity.ListOfPingPongs) query.FindAllQueryResponseRaw {
	return query.FindAllQueryResponseRaw{
		Entities: pps.PingPongs,
	}
}
