package mapper_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/internal/pingpong/api/rest/v1/mapper"
	"github.com/cooperlutz/go-full/internal/pingpong/api/rest/v1/server"
	"github.com/cooperlutz/go-full/internal/pingpong/app/common"
	"github.com/cooperlutz/go-full/internal/pingpong/app/query"
)

func TestMapPingPongToCommand(t *testing.T) {
	msg := "hello"
	req := server.PingPongRequestObject{
		JSONBody: &server.PingPongJSONRequestBody{
			Message: &msg,
		},
	}
	cmd := mapper.MapPingPongToCommand(req)
	assert.NotNil(t, cmd)
	assert.Equal(t, msg, cmd.Message)
}

func TestMapFindAllToResponse(t *testing.T) {
	res := &query.FindAllQueryResponse{
		PingPongs: []*common.PingPongResult{
			{
				Message: "ping",
			},
			{
				Message: "pong",
			},
		},
	}
	httpRes := mapper.MapFindAllToResponse(res)
	assert.NotNil(t, httpRes.Pingpongs)
	assert.Len(t, *httpRes.Pingpongs, 2)
	assert.Equal(t, "ping", *(*httpRes.Pingpongs)[0].Message)
	assert.Equal(t, "pong", *(*httpRes.Pingpongs)[1].Message)
}
