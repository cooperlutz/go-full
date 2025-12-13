package mapper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/internal/pingpong/app/command"
	"github.com/cooperlutz/go-full/internal/pingpong/app/common"
	"github.com/cooperlutz/go-full/internal/pingpong/app/mapper"
	"github.com/cooperlutz/go-full/internal/pingpong/app/query"
	"github.com/cooperlutz/go-full/internal/pingpong/domain/entity"
	"github.com/cooperlutz/go-full/internal/pingpong/domain/exception"
	"github.com/cooperlutz/go-full/test/fixtures"
)

func TestMapFromPingPongCommand(t *testing.T) {
	t.Parallel()

	// Arrange
	tests := []struct {
		name                string
		command             command.PingPongCommand
		wantEntityMsg       string
		wantEntityUpdatedAt time.Time
		wantEntityDeletedAt *time.Time
		wantEntityDeleted   bool
		wantErr             error
	}{
		{
			name: "maps command to entity",
			command: command.PingPongCommand{
				Message: "ping",
			},
			wantEntityMsg:       "ping",
			wantEntityUpdatedAt: time.Now(),
			wantEntityDeletedAt: nil,
			wantEntityDeleted:   false,
			wantErr:             nil,
		},
		{
			name: "an invalid command message returns a validation error",
			command: command.PingPongCommand{
				Message: "NONSENSE_MESSAGE",
			},
			wantEntityMsg:       "NONSENSE_MESSAGE",
			wantEntityUpdatedAt: time.Now(),
			wantEntityDeletedAt: nil,
			wantEntityDeleted:   false,
			wantErr:             exception.ErrPingPongMsgValidation{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			got, err := mapper.MapFromCommandPingPong(tt.command)
			// Assert
			// if we receive an error, we should have wanted one, and the error we received should match the one we wanted
			if err != nil {
				assert.NotNil(t, tt.wantErr)
				assert.Equal(t, tt.wantErr.Error(), err.Error())
				assert.IsType(t, tt.wantErr, err)
				return
			}

			// the returning message should match the input message
			assert.Equal(t, tt.wantEntityMsg, got.GetMessage())
			// the returning entity updatedAt timestamp should be close to now
			assert.WithinDuration(t, tt.wantEntityUpdatedAt, got.GetUpdatedAtTime(), time.Second)
			// the returning entity createdAt timestamp should be close to now
			assert.WithinDuration(t, time.Now(), got.GetCreatedAtTime(), time.Second)
			// the returning entity deletedAt timestamp should be nil
			assert.Nil(t, got.GetDeletedAt())
			// the returning entity deleted flag should be false
			assert.Equal(t, tt.wantEntityDeleted, got.IsDeleted())
		})
	}
}

func TestMapToCommandResult(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		entity entity.PingPongEntity
		want   command.PingPongCommandResult
	}{
		{
			name:   "maps entity to command result",
			entity: fixtures.ValidPing,
			want: command.PingPongCommandResult{
				PingPongResult: &common.PingPongResult{Message: "ping"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mapper.MapToCommandResult(tt.entity)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMapListToQueryResponse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		list entity.ListOfPingPongs
		want query.FindAllQueryResponse
	}{
		{
			name: "maps list of entities to list of common results",
			list: entity.ListOfPingPongs{
				PingPongs: []entity.PingPongEntity{
					fixtures.ValidPing,
					fixtures.ValidPong,
				},
			},
			want: query.FindAllQueryResponse{
				PingPongs: []common.PingPongResult{
					{Message: "ping"},
					{Message: "pong"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mapper.MapListToQueryResponse(tt.list)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMapListToQueryResponseRaw(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		list entity.ListOfPingPongs
		want query.FindAllQueryResponseRaw
	}{
		{
			name: "maps list of entities to raw query response",
			list: entity.ListOfPingPongs{
				PingPongs: []entity.PingPongEntity{
					fixtures.ValidPing,
					fixtures.ValidPong,
				},
			},
			want: query.FindAllQueryResponseRaw{
				Entities: []common.PingPongRawResult{
					{
						ID:        fixtures.ValidPing.GetIdString(),
						Message:   "ping",
						CreatedAt: fixtures.ValidPing.GetCreatedAtTime(),
						UpdatedAt: fixtures.ValidPing.GetUpdatedAtTime(),
						Deleted:   false,
						DeletedAt: nil,
					},
					{
						ID:        fixtures.ValidPong.GetIdString(),
						Message:   "pong",
						CreatedAt: fixtures.ValidPong.GetCreatedAtTime(),
						UpdatedAt: fixtures.ValidPong.GetUpdatedAtTime(),
						Deleted:   false,
						DeletedAt: nil,
					},
				},
			},
		},
		{
			name: "maps empty list to raw query response",
			list: entity.ListOfPingPongs{
				PingPongs: []entity.PingPongEntity{},
			},
			want: query.FindAllQueryResponseRaw{
				Entities: []common.PingPongRawResult{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mapper.MapListToQueryResponseRaw(tt.list)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMapToCommonResult(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		entity entity.PingPongEntity
		want   common.PingPongResult
	}{
		{
			name:   "maps entity to common result",
			entity: fixtures.ValidPing,
			want: common.PingPongResult{
				Message: "ping",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mapper.MapToResult(tt.entity)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMapToRawResult(t *testing.T) {
	t.Parallel()

	// Arrange
	e := fixtures.ValidPong

	// Act
	result := mapper.MapToRawResult(e)

	// Assert
	assert.NotNil(t, result)
	assert.Equal(t, e.GetIdString(), result.ID)
	assert.Equal(t, "pong", result.Message)
	assert.Equal(t, e.GetCreatedAtTime(), result.CreatedAt)
	assert.Equal(t, e.GetUpdatedAtTime(), result.UpdatedAt)
	assert.Equal(t, false, result.Deleted)
	assert.Nil(t, result.DeletedAt)
}
