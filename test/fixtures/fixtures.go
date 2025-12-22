package fixtures

import (
	"time"

	"github.com/google/uuid"

	rest_api_V1 "github.com/cooperlutz/go-full/internal/pingpong/api/rest/v1/server"
	"github.com/cooperlutz/go-full/internal/pingpong/domain/entity"
	"github.com/cooperlutz/go-full/pkg/baseentitee"
	"github.com/cooperlutz/go-full/pkg/utilitee"
)

var (
	ValidMetadata = baseentitee.MapToEntityMetadataFromCommonTypes(
		uuid.UUID{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		time.Date(2023, time.January, 1, 12, 0, 0, 0, time.UTC),
		time.Date(2023, time.January, 1, 12, 0, 0, 0, time.UTC),
		false,
		nil,
	)
	ValidPing = entity.MapToEntity(
		"ping",
		ValidMetadata,
	)
	ValidPong = entity.MapToEntity(
		"pong",
		ValidMetadata,
	)
	ValidReturningPing = entity.MapToEntity(
		"Ping!",
		ValidMetadata,
	)
	ValidReturningPong = entity.MapToEntity(
		"Pong!",
		ValidMetadata,
	)
	InvalidPingPong = entity.MapToEntity(
		"ring",
		ValidMetadata,
	)
	RestApiV1PingPongRaw = []rest_api_V1.PingPongRaw{
		{
			Message:   utilitee.StrPtr("ping"),
			CreatedAt: utilitee.TimePtr(time.Now()),
			Deleted:   utilitee.BoolPtr(false),
			UpdatedAt: utilitee.TimePtr(time.Now()),
			DeletedAt: nil,
			Id:        utilitee.StrPtr(uuid.New().String()),
		},
		{
			Message:   utilitee.StrPtr("pong"),
			CreatedAt: utilitee.TimePtr(time.Now()),
			Deleted:   utilitee.BoolPtr(false),
			UpdatedAt: utilitee.TimePtr(time.Now()),
			DeletedAt: nil,
			Id:        utilitee.StrPtr(uuid.New().String()),
		},
	}
	RestApiV1PingPongsRaw = rest_api_V1.PingPongsRaw{
		Pingpongs: &RestApiV1PingPongRaw,
	}
)
