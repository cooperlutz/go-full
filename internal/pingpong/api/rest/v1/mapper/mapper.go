package mapper

import (
	"github.com/cooperlutz/go-full/internal/pingpong/api/rest/v1/server"
	"github.com/cooperlutz/go-full/internal/pingpong/app/command"
	"github.com/cooperlutz/go-full/internal/pingpong/app/query"
	"github.com/cooperlutz/go-full/pkg/types"
	"github.com/cooperlutz/go-full/pkg/utilitee"
)

func MapPingPongToCommand(req server.PingPongRequestObject) command.PingPongCommand {
	return command.PingPongCommand{
		Message: *req.JSONBody.Message,
	}
}

// STEP 3.3. Implement API Handlers & Mappers
// here, we write code that implements our logic for mapping objects from Service Layer to API Layer
func MapToQueryFindOneByID(req server.GetFindOneByIDRequestObject) query.FindOneByID {
	return query.FindOneByID{
		ID: req.PingPongID,
	}
}

func MapFindAllToResponse(res query.FindAllQueryResponse) server.PingPongs {
	var httpPings []server.PingPong
	for _, p := range res.PingPongs {
		httpPing := server.PingPong{
			Message: &p.Message,
		}
		httpPings = append(httpPings, httpPing)
	}
	response := server.PingPongs{Pingpongs: &httpPings}

	return response
}

func MapFindAllToResponseRaw(res query.FindAllQueryResponseRaw) server.PingPongsRaw {
	var httpPings []server.PingPongRaw

	for _, p := range res.Entities {
		httpPing := server.PingPongRaw{
			Id:        utilitee.StrPtr(p.GetIdString()),
			Message:   utilitee.StrPtr(p.GetMessage()),
			CreatedAt: utilitee.TimePtr(p.GetCreatedAtTime()),
			UpdatedAt: utilitee.TimePtr(p.GetUpdatedAtTime()),
			DeletedAt: p.GetDeletedAtTime(),
			Deleted:   utilitee.BoolPtr(p.IsDeleted()),
		}
		httpPings = append(httpPings, httpPing)
	}

	response := server.PingPongsRaw{Pingpongs: &httpPings}

	return response
}

func MapMeasureCountByDateTimeToTrend(msr []types.MeasureCountbyDateTimeMetric) server.Trend {
	var keys []server.TrendKey
	var values []server.TrendValue

	for _, item := range msr {
		keys = append(keys, server.TrendKey(item.DateTime.String()))
		values = append(values, server.TrendValue(item.Count))
	}

	response := server.Trend{
		DimensionKeys:   &keys,
		DimensionValues: &values,
	}

	return response
}
