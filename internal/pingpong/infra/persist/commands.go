package persist

import (
	"context"

	"github.com/cooperlutz/go-full/internal/pingpong/domain/entity"
	"github.com/cooperlutz/go-full/internal/pingpong/infra/persist/mapper"
	"github.com/cooperlutz/go-full/pkg/telemetree"
	"go.opentelemetry.io/otel/attribute"
)

// SavePingPong - Save a PingPong entity to the database.
func (r *PingPongPersistPostgresRepository) SavePingPong(ctx context.Context, p entity.PingPongEntity) error {
	// telemetree: Add a tracing span for the SavePingPong operation
	ctx, span := telemetree.AddSpan(ctx, "persist.postgres.savepingpong",
		[]attribute.KeyValue{
			attribute.String("pingpong.message", p.GetMessage()),
		}...,
	)
	defer span.End()

	queryParams := mapper.MapToDB(p)

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(ctx); rbErr != nil {
				telemetree.RecordError(ctx, rbErr, "failed to rollback transaction")
			}

			return
		}

		if cmErr := tx.Commit(ctx); cmErr != nil {
			telemetree.RecordError(ctx, cmErr, "failed to commit transaction")
			err = cmErr
		}
	}()

	q := r.query.WithTx(tx)

	_, err = q.CreatePingPong(ctx, queryParams)
	if err != nil {
		telemetree.RecordError(ctx, err, "failed to create pingpong record")

		return err
	}

	return nil
}
