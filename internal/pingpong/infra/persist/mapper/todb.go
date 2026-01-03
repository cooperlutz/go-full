package mapper

import (
	"github.com/cooperlutz/go-full/internal/pingpong/domain/entity"
	postgresql "github.com/cooperlutz/go-full/internal/pingpong/infra/persist/postgres"
	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
)

func MapToDB(pp entity.PingPongEntity) postgresql.CreatePingPongParams {
	createdAt := pp.GetCreatedAtTime()
	msg := pp.GetMessage()

	return postgresql.CreatePingPongParams{
		PingpongID: pgxutil.UUIDToPgtypeUUID(pp.GetIdUUID()),
		PingOrPong: pgxutil.StrToPgtypeText(&msg),
		CreatedAt:  pgxutil.TimeToTimestampz(&createdAt),
		DeletedAt:  pgxutil.TimeToTimestampz(pp.GetDeletedAtTime()),
		Deleted:    pp.IsDeleted(),
	}
}
