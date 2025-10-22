package mapper

import (
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/internal/pingpong/domain/entity"
	postgresql "github.com/cooperlutz/go-full/internal/pingpong/infra/persist/postgres"
)

func TestMapToDB(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		pp   *entity.PingPongEntity
		want postgresql.CreatePingPongParams
	}{
		{
			name: "Valid PingPongEntity to DB Create Params",
			pp: &entity.PingPongEntity{
				Message: "ping",
				PingPongMetadata: &entity.PingPongMetadata{
					PingPongID: sampleUUID,
					CreatedAt:  sampleTime,
					UpdatedAt:  sampleTime,
					DeletedAt:  nil,
					Deleted:    false,
				},
			},
			want: postgresql.CreatePingPongParams{
				PingpongID: pgtype.UUID{Bytes: sampleUUID, Valid: true},
				PingOrPong: pgtype.Text{String: "ping", Valid: true},
				CreatedAt:  pgtype.Timestamptz{Time: sampleTime, InfinityModifier: pgtype.Finite, Valid: true},
				DeletedAt:  pgtype.Timestamptz{},
				Deleted:    false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapToDB(tt.pp)
			assert.IsType(t, pgtype.UUID{}, got.PingpongID)
			assert.Equal(t, tt.want.PingOrPong, got.PingOrPong)
			assert.Equal(t, tt.want.Deleted, got.Deleted)
			assert.Equal(t, tt.want.DeletedAt, got.DeletedAt)
			assert.WithinDuration(t, tt.want.CreatedAt.Time, got.CreatedAt.Time, time.Minute)
			assert.True(t, got.CreatedAt.Valid)
		})
	}
}
