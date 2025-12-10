package entity_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/internal/pingpong/domain/entity"
	"github.com/cooperlutz/go-full/internal/pingpong/domain/exception"
	"github.com/cooperlutz/go-full/pkg/base"
)

var (
	varCreatedAt      = base.CreatedAtFromTime(time.Now())
	varUpdatedAt      = base.UpdatedAtFromTime(time.Now())
	varEntityMetadata = base.MapToEntityMetadata(
		base.EntityIdFromUUID(uuid.UUID{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}),
		varCreatedAt,
		varUpdatedAt,
		false,
		nil,
	)
	validPing = entity.MapToEntity(
		"ping",
		varEntityMetadata,
	)
	validPong = entity.MapToEntity(
		"pong",
		varEntityMetadata,
	)
	invalidPingPong = entity.MapToEntity(
		"ring",
		varEntityMetadata,
	)
	randomUUID = uuid.New()
)

func TestPingPongEntity_Validate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		entity   entity.PingPongEntity
		expected error
	}{
		{"valid ping", validPing, nil},
		{"valid pong", validPong, nil},
		// {"invalid ping pong", invalidPingPong, exception.ErrPingPongMsgValidation{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.entity.Validate()
			assert.Equal(t, tt.expected, err)
		})
	}
}

func TestPingPongEntity_DetermineResponse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                  string
		entity                entity.PingPongEntity
		expectedReturnMessage string
	}{
		{
			"ping returns Pong!",
			validPing,
			"Pong!",
		},
		{
			"pong returns Ping!",
			validPong,
			"Ping!",
		},
		{
			"a message that is not a ping or a pong returns an empty string",
			invalidPingPong,
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.entity.DetermineResponseMessage()
			assert.Equal(t, tt.expectedReturnMessage, result)
		})
	}
}

func TestPingPongEntity_Valid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		entity   entity.PingPongEntity
		expected bool
	}{
		{"valid ping", validPing, true},
		{"valid pong", validPong, true},
		// {"invalid ping pong", invalidPingPong, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.entity.Valid()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestPingPongEntity_GetMessage(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		entity   entity.PingPongEntity
		expected string
	}{
		{"valid ping", validPing, "ping"},
		{"valid pong", validPong, "pong"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.entity.GetMessage()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNewPingPong(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name              string
		input             string
		expectedIDType    uuid.UUID
		expectedMessage   string
		expectedCreatedAt time.Time
		expectedDeletedAt *time.Time
		expectedUpdatedAt time.Time
		expectedDeleted   bool
		expectedError     error
	}{
		{
			"create a new ping",
			"ping",
			randomUUID, // just checking type
			"ping",
			time.Now(),
			nil,
			time.Now(),
			false,
			nil,
		},
		{
			"try to create a new ping, receive an error",
			"purple",
			randomUUID, // just checking type
			"",
			time.Now(),
			nil,
			time.Now(),
			false,
			exception.ErrPingPongMsgValidation{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			result, err := entity.New(tt.input)

			// Assert
			if tt.expectedError != nil {
				assert.ErrorIs(
					t,
					err,
					tt.expectedError,
				)
				return
			}
			assert.IsType(t, tt.expectedIDType, result.GetIdUUID())
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedMessage, result.GetMessage())
			assert.WithinDuration(t, tt.expectedCreatedAt, time.Now(), time.Second)
		})
	}
}

func TestPingPongEntity_MultipleMutations(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		entity   entity.PingPongEntity
		expected string
	}{
		{
			"valid ping",
			validPing,
			"ping",
		},
		{
			"valid pong",
			validPong,
			"pong",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.entity.GetMessage()
			assert.Equal(t, tt.expected, result)
			tt.entity.MarkDeleted()
			assert.True(t, tt.entity.IsDeleted())
			assert.NotNil(t, tt.entity.GetDeletedAt())
			assert.WithinDuration(t, time.Now(), tt.entity.GetUpdatedAtTime(), time.Second)
		})
	}
}
