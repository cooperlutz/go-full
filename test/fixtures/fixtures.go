package fixtures

import (
	"time"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/internal/pingpong/domain/entity"
	"github.com/cooperlutz/go-full/pkg/baseentitee"
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
)
