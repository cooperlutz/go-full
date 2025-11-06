package service

import (
	"context"

	"github.com/cooperlutz/go-full/internal/pingpong/app/command"
	"github.com/cooperlutz/go-full/internal/pingpong/app/mapper"
	"github.com/cooperlutz/go-full/internal/pingpong/app/query"
	"github.com/cooperlutz/go-full/internal/pingpong/domain/repository"
	"github.com/cooperlutz/go-full/pkg/telemetree"
	"github.com/cooperlutz/go-full/pkg/types"
)

// ensure that we've conformed to the `IPingPongService` with a compile-time check.
var _ IPingPongService = (*PingPongService)(nil)

// IPingPongService is the interface that describes the pingpong service.
type IPingPongService interface {
	PingPong(ctx context.Context, cmd command.PingPongCommand) (command.PingPongCommandResult, error) // creates a new pingpong message

	// STEP 4.1. Implement Service Interface
	// here we define what we want the service interface to do, we provide the Query struct,
	// and return the resulting response containing the relevant entity attributes
	FindOneByID(ctx context.Context, q query.FindOneByID) (query.FindOneByIDResponse, error)        // returns one ping pong according to the id provided
	FindAllPings(ctx context.Context) (query.FindAllQueryResponse, error)                           // returns all ping messages
	FindAllPongs(ctx context.Context) (query.FindAllQueryResponse, error)                           // returns all pong messages
	FindAll(ctx context.Context) (query.FindAllQueryResponseRaw, error)                             // returns all ping and pong messages
	TotalNumberOfPingPongs(ctx context.Context) (types.QuantityMetric, error)                       // returns the total number of pingpong
	TotalNumberOfPings(ctx context.Context) (types.QuantityMetric, error)                           // returns the total number of pings
	TotalNumberOfPongs(ctx context.Context) (types.QuantityMetric, error)                           // returns the total number of pongs
	TotalNumberOfPingPongsPerDay(ctx context.Context) ([]types.MeasureCountbyDateTimeMetric, error) // returns the total number of pingpongs created per day
}

type PingPongService struct {
	Persist repository.IPingPongRepository
}

func NewPingPongService(repo repository.IPingPongRepository) *PingPongService {
	return &PingPongService{
		Persist: repo,
	}
}

func (s *PingPongService) PingPong(ctx context.Context, cmd command.PingPongCommand) (command.PingPongCommandResult, error) {
	ctx, span := telemetree.AddSpan(ctx, "service.pingpong")
	defer span.End()

	inputEntity, err := mapper.MapFromCommandPingPong(cmd)
	if err != nil {
		return command.PingPongCommandResult{}, err
	}

	if err := s.Persist.SavePingPong(ctx, inputEntity); err != nil {
		return command.PingPongCommandResult{}, err
	}

	outputResponseMessage := inputEntity.DetermineResponseMessage()

	// Create the result to return
	result := command.NewPingPongCommandResult(outputResponseMessage)

	return result, nil
}

// STEP 4.3. Implement Service Logic
// here we implement the service layer logic.
func (s *PingPongService) FindOneByID(ctx context.Context, q query.FindOneByID) (query.FindOneByIDResponse, error) {
	// update the context with a new span
	ctx, span := telemetree.AddSpan(ctx, "service.findOneById")
	defer span.End()

	// execute the relevant method at the repository persistence layer
	pp, err := s.Persist.FindOneByID(ctx, q.ID)
	if err != nil {
		return query.FindOneByIDResponse{}, err
	}

	// map the entity to the common result object
	result := mapper.MapToRawResult(pp)

	// create the response object that we need to return from our method
	response := query.FindOneByIDResponse{PingPongRawResult: result}

	// finally, we return the response object, and a `nil` for error
	return response, nil
}

func (s *PingPongService) FindAll(ctx context.Context) (query.FindAllQueryResponseRaw, error) {
	ctx, span := telemetree.AddSpan(ctx, "service.findall")
	defer span.End()

	allPings, err := s.Persist.FindAll(ctx)
	if err != nil {
		return query.FindAllQueryResponseRaw{}, err
	}

	response := mapper.MapListToQueryResponseRaw(allPings)

	return response, nil
}

func (s *PingPongService) FindAllPings(ctx context.Context) (query.FindAllQueryResponse, error) {
	ctx, span := telemetree.AddSpan(ctx, "service.findallpings")
	defer span.End()

	allPings, err := s.Persist.FindAllPings(ctx)
	if err != nil {
		return query.FindAllQueryResponse{}, err
	}

	response := mapper.MapListToQueryResponse(allPings)

	return response, nil
}

func (s *PingPongService) FindAllPongs(ctx context.Context) (query.FindAllQueryResponse, error) {
	ctx, span := telemetree.AddSpan(ctx, "service.findallpongs")
	defer span.End()

	allPongs, err := s.Persist.FindAllPongs(ctx)
	if err != nil {
		return query.FindAllQueryResponse{}, err
	}

	response := mapper.MapListToQueryResponse(allPongs)

	return response, nil
}

func (s *PingPongService) TotalNumberOfPingPongs(ctx context.Context) (types.QuantityMetric, error) {
	ctx, span := telemetree.AddSpan(ctx, "service.totalnumberofpingpongs")
	defer span.End()

	count, err := s.Persist.TotalNumberOfPingPongs(ctx)
	if err != nil {
		return types.QuantityMetric{Quantity: 0}, err
	}

	return count, nil
}

func (s *PingPongService) TotalNumberOfPings(ctx context.Context) (types.QuantityMetric, error) {
	ctx, span := telemetree.AddSpan(ctx, "service.totalnumberofpings")
	defer span.End()

	count, err := s.Persist.TotalNumberOfPings(ctx)
	if err != nil {
		return types.QuantityMetric{Quantity: 0}, err
	}

	return count, nil
}

func (s *PingPongService) TotalNumberOfPongs(ctx context.Context) (types.QuantityMetric, error) {
	ctx, span := telemetree.AddSpan(ctx, "service.totalnumberofpongs")
	defer span.End()

	count, err := s.Persist.TotalNumberOfPongs(ctx)
	if err != nil {
		return types.QuantityMetric{Quantity: 0}, err
	}

	return count, nil
}

func (s *PingPongService) TotalNumberOfPingPongsPerDay(ctx context.Context) ([]types.MeasureCountbyDateTimeMetric, error) {
	ctx, span := telemetree.AddSpan(ctx, "service.totalnumberofpingpongspersday")
	defer span.End()

	pingPongsPerDay, err := s.Persist.TotalNumberOfPingPongsCreatedPerDay(ctx)
	if err != nil {
		return nil, err
	}

	return pingPongsPerDay, nil
}
