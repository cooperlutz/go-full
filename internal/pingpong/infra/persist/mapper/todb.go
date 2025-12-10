package mapper

import (
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/cooperlutz/go-full/internal/pingpong/domain/entity"
	postgresql "github.com/cooperlutz/go-full/internal/pingpong/infra/persist/postgres"
)

func MapToDB(pp entity.PingPongEntity) postgresql.CreatePingPongParams {
	return postgresql.CreatePingPongParams{
		PingpongID: pgtype.UUID{Bytes: pp.GetIdUUID(), Valid: pp.Valid()},
		PingOrPong: pgtype.Text{String: pp.GetMessage(), Valid: pp.Valid()},
		CreatedAt:  pgtype.Timestamptz{Time: pp.GetCreatedAtTime(), InfinityModifier: pgtype.Finite, Valid: pp.Valid()},
		DeletedAt:  pgtype.Timestamptz{},
		Deleted:    false,
	}
}
