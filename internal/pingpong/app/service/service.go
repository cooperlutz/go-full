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
	PingPong(ctx context.Context, cmd *command.PingPongCommand) (*command.PingPongCommandResult, error) // creates a new pingpong message
	FindAllPings(ctx context.Context) (*query.FindAllQueryResponse, error)                              // returns all ping messages
	FindAllPongs(ctx context.Context) (*query.FindAllQueryResponse, error)                              // returns all pong messages
	FindAll(ctx context.Context) (*query.FindAllQueryResponseRaw, error)                                // returns all ping and pong messages
	TotalNumberOfPingPongs(ctx context.Context) (int64, error)                                          // returns the total number of pingpong
	TotalNumberOfPings(ctx context.Context) (int64, error)                                              // returns the total number of pings
	TotalNumberOfPongs(ctx context.Context) (int64, error)                                              // returns the total number of pongs
	TotalNumberOfPingPongsPerDay(ctx context.Context) ([]types.MeasureCountbyDateTime, error)           // returns the total number of pingpongs created per day
}

type PingPongService struct {
	Persist repository.IPingPongRepository
}

func NewPingPongService(repo repository.IPingPongRepository) *PingPongService {
	return &PingPongService{
		Persist: repo,
	}
}

func (s *PingPongService) PingPong(ctx context.Context, cmd *command.PingPongCommand) (*command.PingPongCommandResult, error) {
	ctx, span := telemetree.AddSpan(ctx, "service.pingpong")
	defer span.End()

	inputEntity, err := mapper.MapFromPingPongCommand(cmd)
	if err != nil {
		return nil, err
	}

	if err := s.Persist.SavePingPong(ctx, inputEntity); err != nil {
		return nil, err
	}

	outputResponseMessage := inputEntity.DetermineResponseMessage()

	// Create the result to return
	result := command.NewPingPongCommandResult(outputResponseMessage)

	return result, nil
}

func (s *PingPongService) FindAll(ctx context.Context) (*query.FindAllQueryResponseRaw, error) {
	ctx, span := telemetree.AddSpan(ctx, "service.findall")
	defer span.End()

	allPings, err := s.Persist.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	response := mapper.MapListToQueryResponseRaw(allPings)

	return response, nil
}

func (s *PingPongService) FindAllPings(ctx context.Context) (*query.FindAllQueryResponse, error) {
	ctx, span := telemetree.AddSpan(ctx, "service.findallpings")
	defer span.End()

	allPings, err := s.Persist.FindAllPings(ctx)
	if err != nil {
		return nil, err
	}

	response := mapper.MapListToQueryResponse(allPings)

	return response, nil
}

func (s *PingPongService) FindAllPongs(ctx context.Context) (*query.FindAllQueryResponse, error) {
	ctx, span := telemetree.AddSpan(ctx, "service.findallpongs")
	defer span.End()

	allPongs, err := s.Persist.FindAllPongs(ctx)
	if err != nil {
		return nil, err
	}

	response := mapper.MapListToQueryResponse(allPongs)

	return response, nil
}

func (s *PingPongService) TotalNumberOfPingPongs(ctx context.Context) (int64, error) {
	ctx, span := telemetree.AddSpan(ctx, "service.totalnumberofpingpongs")
	defer span.End()

	count, err := s.Persist.TotalNumberOfPingPongs(ctx)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *PingPongService) TotalNumberOfPings(ctx context.Context) (int64, error) {
	ctx, span := telemetree.AddSpan(ctx, "service.totalnumberofpings")
	defer span.End()

	count, err := s.Persist.TotalNumberOfPings(ctx)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *PingPongService) TotalNumberOfPongs(ctx context.Context) (int64, error) {
	ctx, span := telemetree.AddSpan(ctx, "service.totalnumberofpongs")
	defer span.End()

	count, err := s.Persist.TotalNumberOfPongs(ctx)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *PingPongService) TotalNumberOfPingPongsPerDay(ctx context.Context) ([]types.MeasureCountbyDateTime, error) {
	ctx, span := telemetree.AddSpan(ctx, "service.totalnumberofpingpongspersday")
	defer span.End()

	pingPongsPerDay, err := s.Persist.TotalNumberOfPingPongsCreatedPerDay(ctx)
	if err != nil {
		return nil, err
	}

	return pingPongsPerDay, nil
}
