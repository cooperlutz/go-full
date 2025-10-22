package mapper

import (
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/cooperlutz/go-full/internal/pingpong/domain/entity"
	postgresql "github.com/cooperlutz/go-full/internal/pingpong/infra/persist/postgres"
)

func MapToDB(pp *entity.PingPongEntity) postgresql.CreatePingPongParams {
	return postgresql.CreatePingPongParams{
		PingpongID: pgtype.UUID{Bytes: pp.PingPongID, Valid: true},
		PingOrPong: pgtype.Text{String: pp.Message, Valid: pp.Valid()},
		CreatedAt:  pgtype.Timestamptz{Time: pp.CreatedAt, InfinityModifier: pgtype.Finite, Valid: true},
		DeletedAt:  pgtype.Timestamptz{},
		Deleted:    false,
	}
}
