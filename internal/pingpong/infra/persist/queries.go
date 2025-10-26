package persist

import (
	"context"

	"github.com/cooperlutz/go-full/internal/pingpong/domain/entity"
	"github.com/cooperlutz/go-full/internal/pingpong/infra/persist/mapper"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

// FindAll - Query all pingpongs from DB and return a list of PingPongEntity.
func (r *PingPongPersistPostgresRepository) FindAll(ctx context.Context) (*entity.ListOfPingPongs, error) {
	// telemetree: Add a tracing span for the SavePingPong operation
	ctx, span := telemetree.AddSpan(ctx, "persist.postgres.findall")
	defer span.End()

	pingpongs, err := r.query.FindAll(ctx)
	if err != nil {
		return &entity.ListOfPingPongs{}, err
	}

	result := mapper.MapFromDBPingPongs(pingpongs)

	return result, nil
}

// FindAllPings - Query all pings from the database and return them as a list of PingPongEntity.
func (r *PingPongPersistPostgresRepository) FindAllPings(ctx context.Context) (*entity.ListOfPingPongs, error) {
	// telemetree: Add a tracing span for the SavePingPong operation
	ctx, span := telemetree.AddSpan(ctx, "persist.postgres.findallpings")
	defer span.End()

	pingpongs, err := r.query.FindAllPing(ctx)
	if err != nil {
		return &entity.ListOfPingPongs{}, err
	}

	result := mapper.MapFromDBPingPongs(pingpongs)

	return result, nil
}

// FindAllPongs - Query all pongs from the database and return them as a list of PingPongEntity.
func (r *PingPongPersistPostgresRepository) FindAllPongs(ctx context.Context) (*entity.ListOfPingPongs, error) {
	// telemetree: Add a tracing span for the SavePingPong operation
	ctx, span := telemetree.AddSpan(ctx, "persist.postgres.findallpongs")
	defer span.End()

	pingpongs, err := r.query.FindAllPong(ctx)
	if err != nil {
		return &entity.ListOfPingPongs{}, err
	}

	result := mapper.MapFromDBPingPongs(pingpongs)

	return result, nil
}
