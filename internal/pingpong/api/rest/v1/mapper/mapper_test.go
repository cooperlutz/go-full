package mapper_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/internal/pingpong/api/rest/v1/mapper"
	"github.com/cooperlutz/go-full/internal/pingpong/api/rest/v1/server"
	"github.com/cooperlutz/go-full/internal/pingpong/app/common"
	"github.com/cooperlutz/go-full/internal/pingpong/app/query"
	"github.com/cooperlutz/go-full/pkg/types"
	"github.com/cooperlutz/go-full/test/fixtures"
)

func TestMapFindAllToResponseRaw(t *testing.T) {
	// Arrange
	res := query.FindAllQueryResponseRaw{
		Entities: []common.PingPongRawResult{
			{
				ID:        *fixtures.RestApiV1PingPongRaw[0].Id,
				Message:   *fixtures.RestApiV1PingPongRaw[0].Message,
				CreatedAt: *fixtures.RestApiV1PingPongRaw[0].CreatedAt,
				UpdatedAt: *fixtures.RestApiV1PingPongRaw[0].UpdatedAt,
				DeletedAt: fixtures.RestApiV1PingPongRaw[0].DeletedAt,
				Deleted:   *fixtures.RestApiV1PingPongRaw[0].Deleted,
			},
			{
				ID:        *fixtures.RestApiV1PingPongRaw[1].Id,
				Message:   *fixtures.RestApiV1PingPongRaw[1].Message,
				CreatedAt: *fixtures.RestApiV1PingPongRaw[1].CreatedAt,
				UpdatedAt: *fixtures.RestApiV1PingPongRaw[1].UpdatedAt,
				DeletedAt: fixtures.RestApiV1PingPongRaw[1].DeletedAt,
				Deleted:   *fixtures.RestApiV1PingPongRaw[1].Deleted,
			},
		},
	}
	// Act
	httpRes := mapper.MapFindAllToResponseRaw(res)
	// Assert
	assert.NotNil(t, httpRes.Pingpongs)
	assert.Equal(t, fixtures.RestApiV1PingPongsRaw, httpRes)
	assert.Len(t, *httpRes.Pingpongs, 2)
	assert.Equal(t, *fixtures.RestApiV1PingPongRaw[0].Message, *(*httpRes.Pingpongs)[0].Message)
	assert.Equal(t, *fixtures.RestApiV1PingPongRaw[1].Message, *(*httpRes.Pingpongs)[1].Message)
}

func TestMapPingPongToCommand(t *testing.T) {
	// Arrange
	msg := "hello"
	req := server.PingPongRequestObject{
		JSONBody: &server.PingPongJSONRequestBody{
			Message: &msg,
		},
	}
	// Act
	cmd := mapper.MapPingPongToCommand(req)
	// Assert
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
	// Arrange
	dt1 := time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC)
	dt2 := time.Date(2021, 6, 7, 8, 9, 10, 0, time.UTC)
	input := []types.MeasureCountbyDateTimeMetric{
		{DateTime: dt1, Count: 7},
		{DateTime: dt2, Count: 13},
	}
	expectedKey1 := dt1.String()
	expectedKey2 := dt2.String()

	// Act
	trend := mapper.MapMeasureCountByDateTimeToTrend(input)

	// Assert
	assert.NotNil(t, trend.DimensionKeys)
	assert.NotNil(t, trend.DimensionValues)
	assert.Len(t, *trend.DimensionKeys, 2)
	assert.Len(t, *trend.DimensionValues, 2)
	assert.Equal(t, expectedKey1, string((*trend.DimensionKeys)[0]))
	assert.Equal(t, expectedKey2, string((*trend.DimensionKeys)[1]))
	assert.Equal(t, server.TrendValue(7), (*trend.DimensionValues)[0])
	assert.Equal(t, server.TrendValue(13), (*trend.DimensionValues)[1])
}

// STEP 3.2. Implement API Handlers & Mappers Tests
// here, we write the test that implements our logic for mapping objects from Service Layer to API Layer
func TestMapToQueryFindOneByID(t *testing.T) {
	// Arrange
	randomUUID, err := uuid.NewUUID()
	expectedOutput := query.FindOneByID{
		ID: randomUUID,
	}
	input := server.GetFindOneByIDRequestObject{
		PingPongID: randomUUID,
	}

	// Act
	result := mapper.MapToQueryFindOneByID(input)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, result)
}
