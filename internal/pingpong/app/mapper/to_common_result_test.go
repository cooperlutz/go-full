package mapper_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/internal/pingpong/app/common"
	"github.com/cooperlutz/go-full/internal/pingpong/app/mapper"
	"github.com/cooperlutz/go-full/internal/pingpong/domain/entity"
)

func TestMapToCommonResult(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		entity entity.PingPongEntity
		want   common.PingPongResult
	}{
		{
			name:   "maps entity to common result",
			entity: entity.PingPongEntity{Message: "ping"},
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
	id := uuid.New()
	createdAt := time.Now().Add(-time.Hour)
	updatedAt := time.Now()
	e := entity.PingPongEntity{
		Message: "pong",
		PingPongMetadata: &entity.PingPongMetadata{
			PingPongID: id,
			CreatedAt:  createdAt,
			UpdatedAt:  updatedAt,
			Deleted:    true,
			DeletedAt:  nil,
		},
	}

	// Act
	result := mapper.MapToRawResult(e)

	// Assert
	assert.NotNil(t, result)
	assert.Equal(t, id.String(), result.ID)
	assert.Equal(t, "pong", result.Message)
	assert.WithinDuration(t, createdAt, result.CreatedAt, time.Second)
	assert.WithinDuration(t, updatedAt, result.UpdatedAt, time.Second)
	assert.Equal(t, true, result.Deleted)
	assert.Nil(t, result.DeletedAt)
}
