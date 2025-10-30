package repository

import (
	"context"

	"github.com/cooperlutz/go-full/internal/pingpong/domain/entity"
	"github.com/cooperlutz/go-full/pkg/types"
)

type IPingPongRepository interface {
	// Commands
	SavePingPong(ctx context.Context, p *entity.PingPongEntity) error

	// Queries
	FindAll(ctx context.Context) (entity.ListOfPingPongs, error)
	FindAllPings(ctx context.Context) (entity.ListOfPingPongs, error)
	FindAllPongs(ctx context.Context) (entity.ListOfPingPongs, error)

	// Metrics
	TotalNumberOfPingPongs(ctx context.Context) (types.QuantityMetric, error)
	TotalNumberOfPings(ctx context.Context) (types.QuantityMetric, error)
	TotalNumberOfPongs(ctx context.Context) (types.QuantityMetric, error)
	TotalNumberOfPingPongsCreatedPerDay(ctx context.Context) ([]types.MeasureCountbyDateTimeMetric, error)
	AverageNumberOfPingPongsCreatedPerDay(ctx context.Context) (types.QuantityMetric, error)
}
