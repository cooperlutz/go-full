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
	return server.PingPongs{Pingpongs: &httpPings}
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
			// TODO: fix this
			// DeletedAt: p.DeletedAt,
			Deleted: &p.Deleted,
		}
		httpPings = append(httpPings, httpPing)
	}
	return server.PingPongsRaw{Pingpongs: &httpPings}
}

func MapMeasureCountByDateTimeToTrend(p []types.MeasureCountbyDateTime) server.Trend {
	var keys []server.TrendKey
	var values []server.TrendValue

	for _, item := range p {
		keys = append(keys, server.TrendKey(item.DateTime.String()))
		values = append(values, server.TrendValue(item.Count))
	}

	outPutresponse := server.Trend{
		DimensionKeys:   &keys,
		DimensionValues: &values,
	}

	return outPutresponse
}
