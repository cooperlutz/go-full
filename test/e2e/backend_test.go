package e2e_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	pingpong_api_client_v1 "github.com/cooperlutz/go-full/api/rest/pingpong/v1/client"
)

/*
Tests in this file do not adhere to encapsulate full end to end functionality, nor do they necessarily follow Behavior Driven Development,
but instead encapsulate tests associated with backend rest API as part of the e2e test suite
*/

func TestPostPingPongWithPing(t *testing.T) {
	// Arrange
	ctx := context.Background()
	val := "ping"
	req := pingpong_api_client_v1.PingPongJSONRequestBody{Message: &val}
	currentPings, err := pingpongApiClient.GetPingsWithResponse(ctx)
	numPingsBefore := len(*currentPings.JSON200.Pingpongs)

	// Act
	response, err := pingpongApiClient.PingPongWithResponse(ctx, req)

	// Assert
	assert.NoError(t, err, "Error calling PingPong: %v", err)
	assert.Equal(t, 200, response.StatusCode(), "Expected response message to be 'pong'")
	assert.Equal(t, "Pong!", *response.JSON200.Message, "Expected response message to be 'pong'")
	numPingsAfterAction, err := pingpongApiClient.GetPingsWithResponse(ctx)
	assert.Equal(t, numPingsBefore+1, len(*numPingsAfterAction.JSON200.Pingpongs), "Expected number of pings to increase by 1")
}
