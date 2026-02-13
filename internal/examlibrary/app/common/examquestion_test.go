package common_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/internal/examlibrary/app/common"
)

func TestNewExamQuestion(t *testing.T) {
	// Arrange
	index := 1
	questionText := "What is the capital of France?"
	questionType := "multiple-choice"
	possiblePoints := 5
	correctAnswer := new("Paris")
	options := &[]string{"Berlin", "Madrid", "Paris", "Rome"}

	// Act
	question := common.NewExamQuestion(index, questionText, questionType, possiblePoints, correctAnswer, options)

	// Assert
	assert.Equal(t, index, question.Index)
	assert.Equal(t, questionText, question.QuestionText)
	assert.Equal(t, questionType, question.QuestionType)
	assert.Equal(t, possiblePoints, question.PossiblePoints)
	assert.Equal(t, correctAnswer, question.CorrectAnswer)
	assert.Equal(t, options, question.ResponseOptions)
}
