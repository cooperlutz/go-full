package rest_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/cooperlutz/go-full/internal/pingpong/api/rest"
	"github.com/cooperlutz/go-full/internal/pingpong/app/command"
	"github.com/cooperlutz/go-full/internal/pingpong/app/common"
	mocks "github.com/cooperlutz/go-full/test/mocks"
)

// Test that /v1/ping-pongs returns 200 and "pong"
func TestNewPingPongAPIRouter(t *testing.T) {
	t.Parallel()

	// Arrange
	useCase := mocks.NewMockIPingPongUseCase(t)
	router := rest.NewPingPongAPIRouter(useCase)
	useCase.Mock.On("PingPong", mock.Anything, mock.Anything).Return(command.PingPongCommandResult{
		PingPongResult: &common.PingPongResult{Message: "pong"},
	}, nil)
	req := httptest.NewRequest(
		http.MethodPost,
		"/v1/ping-pongs",
		bytes.NewBufferString(`{"message":"ping"}`),
	)
	w := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, string(data), "pong")
}
