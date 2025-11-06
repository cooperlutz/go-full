package repository

import (
	"context"

	"github.com/google/uuid"

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

	/* STEP 1.2. Implement Domain Repository Interface
	here we define what the Repository should allow us to do, we want to input a uuid, and return the relevant entity associated with that uuid
	NOTE: no additional logic is required to be added to an entity, as the added functinoality is keying off of existing entity fields
	*/
	FindOneByID(ctx context.Context, id uuid.UUID) (entity.PingPongEntity, error)

	// Metrics
	TotalNumberOfPingPongs(ctx context.Context) (types.QuantityMetric, error)
	TotalNumberOfPings(ctx context.Context) (types.QuantityMetric, error)
	TotalNumberOfPongs(ctx context.Context) (types.QuantityMetric, error)
	TotalNumberOfPingPongsCreatedPerDay(ctx context.Context) ([]types.MeasureCountbyDateTimeMetric, error)
	AverageNumberOfPingPongsCreatedPerDay(ctx context.Context) (types.QuantityMetric, error)
}
