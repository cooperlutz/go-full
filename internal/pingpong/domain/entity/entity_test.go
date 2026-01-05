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

func TestNewPingPong(t *testing.T) {
	t.Parallel()
	// Arrange
	tests := []struct {
		name              string
		input             string
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
			assert.IsType(t, uuid.New(), result.GetIdUUID())
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedMessage, result.GetMessage())
			assert.Equal(t, tt.expectedDeleted, false)
			assert.WithinDuration(t, tt.expectedCreatedAt, result.GetCreatedAtTime(), time.Second)
			assert.WithinDuration(t, tt.expectedUpdatedAt, result.GetUpdatedAtTime(), time.Second)
			assert.Equal(t, tt.expectedDeletedAt, result.GetDeletedAtTime())
		})
	}
}

func TestPingPongEntity_Validate(t *testing.T) {
	t.Parallel()

	validPing, _ := entity.New("ping")
	validPong, _ := entity.New("pong")

	tests := []struct {
		name     string
		entity   entity.PingPongEntity
		expected error
	}{
		{
			"check if pingpong entity ping is valid",
			validPing,
			nil,
		},
		{
			"check if pingpong entity pong is valid",
			validPong,
			nil,
		},
		{
			"invalid ping pong",
			fixtures.InvalidPingPong,
			exception.ErrPingPongMsgValidation{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			err := tt.entity.Validate()
			// Assert
			assert.Equal(t, tt.expected, err)
		})
	}
}

func TestPingPongEntity_DetermineResponse(t *testing.T) {
	t.Parallel()
	validPing, _ := entity.New("ping")
	validPong, _ := entity.New("pong")

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
			fixtures.InvalidPingPong,
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			result := tt.entity.DetermineResponseMessage()
			// Assert
			assert.Equal(t, tt.expectedReturnMessage, result)
		})
	}
}

func TestPingPongEntity_GetMessage(t *testing.T) {
	t.Parallel()
	validPing, _ := entity.New("ping")
	validPong, _ := entity.New("pong")

	tests := []struct {
		name     string
		entity   entity.PingPongEntity
		expected string
	}{
		{
			"get valid ping message",
			validPing,
			"ping",
		},
		{
			"get valid pong message",
			validPong,
			"pong",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			result := tt.entity.GetMessage()
			// Assert
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestPingPongEntity_SetMessage(t *testing.T) {
	t.Parallel()
	validPing, _ := entity.New("ping")
	validPong, _ := entity.New("pong")

	tests := []struct {
		name        string
		entity      entity.PingPongEntity
		newMessage  string
		expectedMsg string
	}{
		{
			"set valid ping message",
			validPing,
			"pong",
			"pong",
		},
		{
			"set valid pong message",
			validPong,
			"ping",
			"ping",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			tt.entity.SetMessage(tt.newMessage)
			result := tt.entity.GetMessage()
			// Assert
			assert.Equal(t, tt.expectedMsg, result)
		})
	}
}

func TestPingPongEntity_MultipleMutations(t *testing.T) {
	t.Parallel()
	validPing, _ := entity.New("ping")
	validPong, _ := entity.New("pong")

	tests := []struct {
		name                   string
		entity                 entity.PingPongEntity
		expectedInitialMessage string
		newMessage             string
		expectedNewMessage     string
	}{
		{
			"run multiple mutations on valid ping",
			validPing,
			"ping",
			"pong",
			"pong",
		},
		{
			"run multiple mutations on valid pong",
			validPong,
			"pong",
			"ping",
			"ping",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Assert
			assert.Equal(t, tt.expectedInitialMessage, tt.entity.GetMessage())

			// Act
			tt.entity.SetMessage(tt.newMessage)

			// Act
			updatedMessage := tt.entity.GetMessage()

			// Assert
			assert.Equal(t, tt.expectedNewMessage, updatedMessage)

			// Act
			tt.entity.MarkDeleted()

			// Assert
			assert.True(t, tt.entity.IsDeleted())
			assert.NotNil(t, tt.entity.GetDeletedAt())
			assert.WithinDuration(t, time.Now(), tt.entity.GetUpdatedAtTime(), time.Second)
		})
	}
}
