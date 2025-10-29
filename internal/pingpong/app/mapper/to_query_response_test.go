package mapper_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/internal/pingpong/app/common"
	"github.com/cooperlutz/go-full/internal/pingpong/app/mapper"
	"github.com/cooperlutz/go-full/internal/pingpong/app/query"
	"github.com/cooperlutz/go-full/internal/pingpong/domain/entity"
)

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
					{Message: "ping"},
					{Message: "pong"},
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
					{Message: "ping"},
					{Message: "pong"},
				},
			},
			want: query.FindAllQueryResponseRaw{
				Entities: []entity.PingPongEntity{
					{Message: "ping"},
					{Message: "pong"},
				},
			},
		},
		{
			name: "maps empty list to raw query response",
			list: entity.ListOfPingPongs{
				PingPongs: []entity.PingPongEntity{},
			},
			want: query.FindAllQueryResponseRaw{
				Entities: []entity.PingPongEntity{},
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
