package mapper

import (
	"time"

	"github.com/cooperlutz/go-full/internal/pingpong/domain/entity"
	postgresql "github.com/cooperlutz/go-full/internal/pingpong/infra/persist/postgres"
)

func TranslatePingPongsFromDB(pingponglist []postgresql.Pingpong) *entity.ListOfPingPongs {
	if len(pingponglist) == 0 {
		return &entity.ListOfPingPongs{PingPongs: []entity.PingPongEntity{}}
	}

	outputList := make([]entity.PingPongEntity, 0, len(pingponglist))

	for _, ppp := range pingponglist {
		translatedPingPong := entity.PingPongEntity{
			Message: ppp.PingOrPong.String,
		}

		outputList = append(outputList, translatedPingPong)
	}

	return &entity.ListOfPingPongs{PingPongs: outputList}
}

func TranslateFromDB(p postgresql.Pingpong) entity.PingPongEntity {
	return entity.PingPongEntity{
		Message: p.PingOrPong.String,
	}
}

func TranslateFromDBRaw(p postgresql.Pingpong) entity.PingPongEntity {
	var deletedAtTime *time.Time
	if p.DeletedAt.Time.IsZero() {
		deletedAtTime = nil
	} else {
		deletedAtTime = &p.DeletedAt.Time
	}

	return entity.PingPongEntity{
		PingPongMetadata: &entity.PingPongMetadata{
			PingPongID: p.PingpongID.Bytes,
			CreatedAt:  p.CreatedAt.Time,
			UpdatedAt:  p.UpdatedAt.Time,
			DeletedAt:  deletedAtTime,
			Deleted:    p.Deleted,
		},
		Message: p.PingOrPong.String,
	}
}
