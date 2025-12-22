package exception_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/internal/pingpong/domain/exception"
)

func TestErrPingPongMsgValidation_Error(t *testing.T) {
	t.Parallel()
	// Arrange
	err := exception.ErrPingPongMsgValidation{}
	expectedMessage := "ya gotta send a ping or a pong"
	// Act
	returnedMessage := err.Error()
	// Assert
	assert.Equal(t, expectedMessage, returnedMessage)
}
