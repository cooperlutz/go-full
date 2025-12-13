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
	mocks "github.com/cooperlutz/go-full/test/mocks/pingpong"
)

func TestNewPingPongAPIRouter(t *testing.T) {
	t.Parallel()

	service := mocks.NewMockIPingPongUseCase(t)
	router := rest.NewPingPongAPIRouter(service)
	service.Mock.On("PingPong", mock.Anything, mock.Anything).Return(command.PingPongCommandResult{
		PingPongResult: &common.PingPongResult{Message: "pong"},
	}, nil)
	// Test that /api/v1/ping returns 200 and "pong"
	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/ping-pongs",
		bytes.NewBufferString(`{"message":"ping"}`),
	)
	w := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	assert.Equal(t, http.StatusOK, w.Code)
	// assert.Contains(t, w.Body.String(), "pong")
	assert.Contains(t, string(data), "pong")
}
