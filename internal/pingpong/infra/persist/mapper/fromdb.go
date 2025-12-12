package mapper

import (
	"time"

	"github.com/cooperlutz/go-full/internal/pingpong/domain/entity"
	postgresql "github.com/cooperlutz/go-full/internal/pingpong/infra/persist/postgres"
	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

func MapFromDBPingPongs(pingponglist []postgresql.Pingpong) entity.ListOfPingPongs {
	if len(pingponglist) == 0 {
		return entity.ListOfPingPongs{PingPongs: []entity.PingPongEntity{}}
	}

	outputList := make([]entity.PingPongEntity, 0, len(pingponglist))

	for _, ppp := range pingponglist {
		translatedPingPong := MapFromDB(ppp)

		outputList = append(outputList, translatedPingPong)
	}

	return entity.ListOfPingPongs{PingPongs: outputList}
}

func MapFromDB(p postgresql.Pingpong) entity.PingPongEntity {
	var deletedAtTime *time.Time
	if p.DeletedAt.Time.IsZero() {
		deletedAtTime = nil
	} else {
		deletedAtTime = &p.DeletedAt.Time
	}

	entityMetadata := baseentitee.MapToEntityMetadataFromCommonTypes(
		p.PingpongID.Bytes,
		p.CreatedAt.Time,
		p.UpdatedAt.Time,
		p.Deleted,
		deletedAtTime,
	)

	return entity.MapToEntity(
		p.PingOrPong.String,
		entityMetadata,
	)
}
