package mapper

import (
	"github.com/cooperlutz/go-full/internal/pingpong/api/rest/v1/server"
	"github.com/cooperlutz/go-full/internal/pingpong/app/command"
	"github.com/cooperlutz/go-full/internal/pingpong/app/query"
	"github.com/cooperlutz/go-full/pkg/types"
)

func MapPingPongToCommand(req server.PingPongRequestObject) command.PingPongCommand {
	return command.PingPongCommand{
		Message: *req.JSONBody.Message,
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
		pingpongId := p.PingPongID.String()

		httpPing := server.PingPongRaw{
			Id:        &pingpongId,
			Message:   &p.Message,
			CreatedAt: &p.CreatedAt,
			UpdatedAt: &p.UpdatedAt,
			DeletedAt: p.DeletedAt,
			Deleted:   &p.Deleted,
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
