package mapper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/internal/pingpong/api/rest/v1/mapper"
	"github.com/cooperlutz/go-full/internal/pingpong/api/rest/v1/server"
	"github.com/cooperlutz/go-full/internal/pingpong/app/common"
	"github.com/cooperlutz/go-full/internal/pingpong/app/query"
	"github.com/cooperlutz/go-full/pkg/types"
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
	res := query.FindAllQueryResponse{
		PingPongs: []common.PingPongResult{
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

func TestMapMeasureCountByDateTimeToTrend(t *testing.T) {
	dt1 := time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC)
	dt2 := time.Date(2021, 6, 7, 8, 9, 10, 0, time.UTC)

	input := []types.MeasureCountbyDateTime{
		{DateTime: dt1, Count: 7},
		{DateTime: dt2, Count: 13},
	}

	trend := mapper.MapMeasureCountByDateTimeToTrend(input)

	assert.NotNil(t, trend.DimensionKeys)
	assert.NotNil(t, trend.DimensionValues)
	assert.Len(t, *trend.DimensionKeys, 2)
	assert.Len(t, *trend.DimensionValues, 2)

	expectedKey1 := dt1.String()
	expectedKey2 := dt2.String()

	assert.Equal(t, expectedKey1, string((*trend.DimensionKeys)[0]))
	assert.Equal(t, expectedKey2, string((*trend.DimensionKeys)[1]))

	assert.Equal(t, server.TrendValue(7), (*trend.DimensionValues)[0])
	assert.Equal(t, server.TrendValue(13), (*trend.DimensionValues)[1])
}
