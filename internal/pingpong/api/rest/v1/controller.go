package v1

import (
	"context"

	"go.opentelemetry.io/otel/trace"

	"github.com/cooperlutz/go-full/internal/pingpong/api/rest/v1/mapper"
	"github.com/cooperlutz/go-full/internal/pingpong/api/rest/v1/server"
	"github.com/cooperlutz/go-full/internal/pingpong/app/usecase"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ server.StrictServerInterface = (*PingPongRestAPIControllerV1)(nil)

// PingPongRestAPIControllerV1 is the controller for the PingPong API
type PingPongRestAPIControllerV1 struct {
	UseCase usecase.IPingPongUseCase
}

// NewRestAPIController creates a new PingPongRestAPIControllerV1
func NewRestAPIController(svc usecase.IPingPongUseCase) *PingPongRestAPIControllerV1 {
	return &PingPongRestAPIControllerV1{
		UseCase: svc,
	}
}

// PUT /ping-pongs
func (c *PingPongRestAPIControllerV1) PingPong(ctx context.Context, request server.PingPongRequestObject) (server.PingPongResponseObject, error) {
	spanCtx := trace.SpanContextFromContext(ctx)

	command := mapper.MapPingPongToCommand(request)

	cmdResponse, err := c.UseCase.PingPong(ctx, command)
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

// STEP 3.3. Implement API Handlers & Mappers
// here we implement the handler for the endpoint defined in the openapi spec
func (c *PingPongRestAPIControllerV1) GetFindOneByID(ctx context.Context, request server.GetFindOneByIDRequestObject) (server.GetFindOneByIDResponseObject, error) {
	spanCtx := trace.SpanContextFromContext(ctx)

	// first, we map the request object from the API layer, to the query object in the Service Layer
	q := mapper.MapToQueryFindOneByID(request)

	// using the service Query object, we pass that to the Service method
	res, err := c.UseCase.FindOneByID(ctx, q)
	// if we encounter an error from the service layer, we pass back the 400 response and the resulting error
	if err != nil {
		return server.GetFindOneByID400Response{}, err
	}

	// we map the returned object from the service layer, to the object that we want to return within the API layer
	response := server.GetFindOneByID200JSONResponse{
		Body: server.PingPongRaw{
			CreatedAt: &res.CreatedAt,
			Deleted:   &res.Deleted,
			DeletedAt: res.DeletedAt,
			Id:        &res.ID,
			Message:   &res.Message,
			UpdatedAt: &res.UpdatedAt,
		},
		Headers: server.GetFindOneByID200ResponseHeaders{XRequestId: spanCtx.SpanID().String()},
	}

	// finally, we return the resulting object and `nil` for the error return
	return response, nil
}

// GET /ping-pongs
func (c *PingPongRestAPIControllerV1) GetFindAllPingPongs(ctx context.Context, request server.GetFindAllPingPongsRequestObject) (server.GetFindAllPingPongsResponseObject, error) {
	spanCtx := trace.SpanContextFromContext(ctx)

	queryResponse, err := c.UseCase.FindAll(ctx)
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

	allPings, err := c.UseCase.FindAllPings(ctx)
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

	allPongs, err := c.UseCase.FindAllPongs(ctx)
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

	totalPingPongs, err := c.UseCase.TotalNumberOfPingPongs(ctx)
	if err != nil {
		return server.GetTotalPingPongs400Response{}, err
	}

	response := server.GetTotalPingPongs200JSONResponse{
		Body:    int(totalPingPongs.Quantity),
		Headers: server.GetTotalPingPongs200ResponseHeaders{XRequestId: spanCtx.TraceID().String()},
	}
	return response, nil
}

// GET /metrics/totalPings
func (c *PingPongRestAPIControllerV1) GetTotalPings(ctx context.Context, request server.GetTotalPingsRequestObject) (server.GetTotalPingsResponseObject, error) {
	spanCtx := trace.SpanContextFromContext(ctx)

	totalPingPongs, err := c.UseCase.TotalNumberOfPings(ctx)
	if err != nil {
		return server.GetTotalPings400Response{}, err
	}

	response := server.GetTotalPings200JSONResponse{
		Body:    int(totalPingPongs.Quantity),
		Headers: server.GetTotalPings200ResponseHeaders{XRequestId: spanCtx.TraceID().String()},
	}
	return response, nil
}

// GET /metrics/totalPongs
func (c *PingPongRestAPIControllerV1) GetTotalPongs(ctx context.Context, request server.GetTotalPongsRequestObject) (server.GetTotalPongsResponseObject, error) {
	spanCtx := trace.SpanContextFromContext(ctx)

	totalPingPongs, err := c.UseCase.TotalNumberOfPongs(ctx)
	if err != nil {
		return server.GetTotalPongs400Response{}, err
	}

	response := server.GetTotalPongs200JSONResponse{
		Body:    int(totalPingPongs.Quantity),
		Headers: server.GetTotalPongs200ResponseHeaders{XRequestId: spanCtx.TraceID().String()},
	}
	return response, nil
}

// GET /metrics/dailyDistribution
func (c *PingPongRestAPIControllerV1) GetDailyDistribution(ctx context.Context, request server.GetDailyDistributionRequestObject) (server.GetDailyDistributionResponseObject, error) {
	spanCtx := trace.SpanContextFromContext(ctx)
	dailyDistribution, err := c.UseCase.TotalNumberOfPingPongsPerDay(ctx)
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
