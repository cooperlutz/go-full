package mapper

import (
	"github.com/cooperlutz/go-full/internal/pingpong/api/rest/v1/server"
	"github.com/cooperlutz/go-full/internal/pingpong/app/command"
	"github.com/cooperlutz/go-full/internal/pingpong/app/query"
	"github.com/cooperlutz/go-full/pkg/types"
	"github.com/cooperlutz/go-full/pkg/utilitee"
)

// MapPingPongToCommand maps an API request object to a PingPongCommand.
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

// MapFindAllToResponse maps a FindAllQueryResponse to an API response object.
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

// MapFindAllToResponseRaw maps a FindAllQueryResponseRaw to an API response object.
func MapFindAllToResponseRaw(res query.FindAllQueryResponseRaw) server.PingPongsRaw {
	var httpPings []server.PingPongRaw

	for _, p := range res.Entities {
		httpPing := server.PingPongRaw{
			Id:        utilitee.StrPtr(p.ID),
			Message:   utilitee.StrPtr(p.Message),
			CreatedAt: utilitee.TimePtr(p.CreatedAt),
			UpdatedAt: utilitee.TimePtr(p.UpdatedAt),
			DeletedAt: p.DeletedAt,
			Deleted:   utilitee.BoolPtr(p.Deleted),
		}
		httpPings = append(httpPings, httpPing)
	}

	response := server.PingPongsRaw{Pingpongs: &httpPings}

	return response
}

// MapMeasureCountByDateTimeToTrend maps a slice of MeasureCountbyDateTimeMetric to a Trend response.
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
