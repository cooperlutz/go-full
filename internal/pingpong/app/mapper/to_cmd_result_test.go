package mapper_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/internal/pingpong/app/command"
	"github.com/cooperlutz/go-full/internal/pingpong/app/common"
	"github.com/cooperlutz/go-full/internal/pingpong/app/mapper"
	"github.com/cooperlutz/go-full/internal/pingpong/domain/entity"
)

func TestMapToCommandResult(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		entity entity.PingPongEntity
		want   *command.PingPongCommandResult
	}{
		{
			name:   "maps entity to command result",
			entity: entity.PingPongEntity{Message: "pong"},
			want: &command.PingPongCommandResult{
				PingPongResult: &common.PingPongResult{Message: "pong"},
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
