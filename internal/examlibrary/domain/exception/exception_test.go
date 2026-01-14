package exception_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/internal/examlibrary/domain/exception"
)

func TestErrQuestionNotFound_Error(t *testing.T) {
	err := exception.ErrQuestionNotFound{}
	assert.Equal(t, "question not found", err.Error())
}

func TestErrInvalidIndex_Error(t *testing.T) {
	err := exception.ErrInvalidIndex{}
	assert.Equal(t, "invalid index", err.Error())
}

func TestErrInvalidQuestionType_Error(t *testing.T) {
	err := exception.ErrInvalidQuestionType{}
	assert.Equal(t, "invalid question type", err.Error())
}
