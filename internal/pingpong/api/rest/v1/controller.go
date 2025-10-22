package v1

import (
	"context"

	"go.opentelemetry.io/otel/trace"

	"github.com/cooperlutz/go-full/internal/pingpong/api/rest/v1/mapper"
	"github.com/cooperlutz/go-full/internal/pingpong/api/rest/v1/server"
	"github.com/cooperlutz/go-full/internal/pingpong/app/service"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ server.StrictServerInterface = (*PingPongRestAPIControllerV1)(nil)

// PingPongRestAPIControllerV1 is the controller for the PingPong API
type PingPongRestAPIControllerV1 struct {
	Service service.IPingPongService
}

// NewRestAPIController creates a new PingPongRestAPIControllerV1
func NewRestAPIController(svc service.IPingPongService) *PingPongRestAPIControllerV1 {
	return &PingPongRestAPIControllerV1{
		Service: svc,
	}
}

// PUT /ping-pongs
func (c *PingPongRestAPIControllerV1) PingPong(ctx context.Context, request server.PingPongRequestObject) (server.PingPongResponseObject, error) {
	spanCtx := trace.SpanContextFromContext(ctx)

	command := mapper.MapPingPongToCommand(request)

	cmdResponse, err := c.Service.PingPong(ctx, command)
	if err != nil {
		return server.PingPong400Response{}, err
	}

	httpPingPong := server.PingPong{Message: &cmdResponse.Message}

	response := server.PingPong200JSONResponse{
		Body:    httpPingPong,
		Headers: server.PingPong200ResponseHeaders{XRequestId: spanCtx.TraceID().String()},
	}
	return response, nil
}

// GET /ping-pongs
func (c *PingPongRestAPIControllerV1) GetFindAllPingPongs(ctx context.Context, request server.GetFindAllPingPongsRequestObject) (server.GetFindAllPingPongsResponseObject, error) {
	spanCtx := trace.SpanContextFromContext(ctx)

	queryResponse, err := c.Service.FindAll(ctx)
	if err != nil {
		return server.GetFindAllPingPongs400Response{}, err
	}

	pingpongs := mapper.MapFindAllToResponseRaw(queryResponse)

	response := server.GetFindAllPingPongs200JSONResponse{
		Body:    pingpongs,
		Headers: server.GetFindAllPingPongs200ResponseHeaders{XRequestId: spanCtx.TraceID().String()},
	}

	return response, nil
}

// GET /ping-pong/pings
func (c *PingPongRestAPIControllerV1) GetPings(ctx context.Context, request server.GetPingsRequestObject) (server.GetPingsResponseObject, error) {
	spanCtx := trace.SpanContextFromContext(ctx)

	allPings, err := c.Service.FindAllPings(ctx)
	if err != nil {
		return server.GetPings400Response{}, err
	}

	httpPings := mapper.MapFindAllToResponse(allPings)

	response := server.GetPings200JSONResponse{
		Body:    httpPings,
		Headers: server.GetPings200ResponseHeaders{XRequestId: spanCtx.TraceID().String()},
	}
	return response, nil
}

// GET /ping-pong/pongs
func (c *PingPongRestAPIControllerV1) GetPongs(ctx context.Context, request server.GetPongsRequestObject) (server.GetPongsResponseObject, error) {
	spanCtx := trace.SpanContextFromContext(ctx)

	allPongs, err := c.Service.FindAllPongs(ctx)
	if err != nil {
		return server.GetPongs400Response{}, err
	}

	httpPongs := mapper.MapFindAllToResponse(allPongs)

	response := server.GetPongs200JSONResponse{
		Body:    httpPongs,
		Headers: server.GetPongs200ResponseHeaders{XRequestId: spanCtx.TraceID().String()},
	}
	return response, nil
}

// GET /metrics/totalPingPongs
func (c *PingPongRestAPIControllerV1) GetTotalPingPongs(ctx context.Context, request server.GetTotalPingPongsRequestObject) (server.GetTotalPingPongsResponseObject, error) {
	spanCtx := trace.SpanContextFromContext(ctx)

	totalPingPongs, err := c.Service.TotalNumberOfPingPongs(ctx)
	if err != nil {
		return server.GetTotalPingPongs400Response{}, err
	}

	response := server.GetTotalPingPongs200JSONResponse{
		Body:    int(totalPingPongs),
		Headers: server.GetTotalPingPongs200ResponseHeaders{XRequestId: spanCtx.TraceID().String()},
	}
	return response, nil
}

// GET /metrics/totalPings
func (c *PingPongRestAPIControllerV1) GetTotalPings(ctx context.Context, request server.GetTotalPingsRequestObject) (server.GetTotalPingsResponseObject, error) {
	spanCtx := trace.SpanContextFromContext(ctx)

	totalPingPongs, err := c.Service.TotalNumberOfPings(ctx)
	if err != nil {
		return server.GetTotalPings400Response{}, err
	}

	response := server.GetTotalPings200JSONResponse{
		Body:    int(totalPingPongs),
		Headers: server.GetTotalPings200ResponseHeaders{XRequestId: spanCtx.TraceID().String()},
	}
	return response, nil
}

// GET /metrics/totalPongs
func (c *PingPongRestAPIControllerV1) GetTotalPongs(ctx context.Context, request server.GetTotalPongsRequestObject) (server.GetTotalPongsResponseObject, error) {
	spanCtx := trace.SpanContextFromContext(ctx)

	totalPingPongs, err := c.Service.TotalNumberOfPongs(ctx)
	if err != nil {
		return server.GetTotalPongs400Response{}, err
	}

	response := server.GetTotalPongs200JSONResponse{
		Body:    int(totalPingPongs),
		Headers: server.GetTotalPongs200ResponseHeaders{XRequestId: spanCtx.TraceID().String()},
	}
	return response, nil
}

// GET /metrics/dailyDistribution
func (c *PingPongRestAPIControllerV1) GetDailyDistribution(ctx context.Context, request server.GetDailyDistributionRequestObject) (server.GetDailyDistributionResponseObject, error) {
	spanCtx := trace.SpanContextFromContext(ctx)
	dailyDistribution, err := c.Service.TotalNumberOfPingPongsPerDay(ctx)
	if err != nil {
		return server.GetDailyDistribution400Response{}, err
	}

	parsed := mapper.MapMeasureCountByDateTimeToTrend(dailyDistribution)

	response := server.GetDailyDistribution200JSONResponse{
		Body:    parsed,
		Headers: server.GetDailyDistribution200ResponseHeaders{XRequestId: spanCtx.TraceID().String()},
	}
	return response, nil
}
