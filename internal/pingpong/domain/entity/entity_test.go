package entity_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/internal/pingpong/domain/entity"
	"github.com/cooperlutz/go-full/internal/pingpong/domain/exception"
	"github.com/cooperlutz/go-full/test/fixtures"
)

var randomUUID = uuid.New()

func TestPingPongEntity_Validate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		entity   entity.PingPongEntity
		expected error
	}{
		{
			"valid ping",
			fixtures.ValidPing,
			nil,
		},
		{
			"valid pong",
			fixtures.ValidPong,
			nil,
		},
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
			fixtures.ValidPing,
			"Pong!",
		},
		{
			"pong returns Ping!",
			fixtures.ValidPong,
			"Ping!",
		},
		{
			"a message that is not a ping or a pong returns an empty string",
			fixtures.InvalidPingPong,
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
		{
			"valid ping",
			fixtures.ValidPing,
			true,
		},
		{
			"valid pong",
			fixtures.ValidPong,
			true,
		},
		{
			"invalid ping pong", fixtures.InvalidPingPong, false,
		},
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
		{
			"valid ping",
			fixtures.ValidPing,
			"ping",
		},
		{
			"valid pong",
			fixtures.ValidPong,
			"pong",
		},
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
			fixtures.ValidPing,
			"ping",
		},
		{
			"valid pong",
			fixtures.ValidPong,
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
