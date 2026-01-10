package entity

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/internal/examlibrary/domain/valueobject"
)

func TestNewExamQuestion(t *testing.T) {
	questionText := "What is 2 + 2?"
	questionType := valueobject.QuestionMultipleChoice
	possiblePoints := 5
	correctAnswer := "4"
	options := []string{"3", "4", "5", "6"}

	examQuestion := NewExamQuestion(
		questionText,
		questionType,
		possiblePoints,
		&correctAnswer,
		&options,
	)

	assert.Equal(t, questionText, examQuestion.questionText)
	assert.Equal(t, questionType, examQuestion.questionType)
	assert.Equal(t, possiblePoints, examQuestion.possiblePoints)
	assert.Equal(t, &correctAnswer, examQuestion.correctAnswer)
	assert.Equal(t, &options, examQuestion.responseOptions)
}

func TestMapToExamQuestion(t *testing.T) {
	id := uuid.New()
	createdAt := time.Now()
	updatedAt := time.Now()
	deleted := false
	var deletedAt *time.Time = nil
	questionText := "What is 2 + 2?"
	questionType := valueobject.QuestionMultipleChoice
	possiblePoints := 5
	correctAnswer := "4"
	options := []string{"3", "4", "5", "6"}

	examQuestion := MapToExamQuestion(
		id,
		createdAt,
		updatedAt,
		deleted,
		deletedAt,
		questionText,
		questionType,
		possiblePoints,
		&correctAnswer,
		&options,
		0,
	)

	assert.Equal(t, id, examQuestion.GetIdUUID())
	assert.Equal(t, questionText, examQuestion.questionText)
	assert.Equal(t, questionType, examQuestion.questionType)
	assert.Equal(t, possiblePoints, examQuestion.possiblePoints)
	assert.Equal(t, &correctAnswer, examQuestion.correctAnswer)
	assert.Equal(t, &options, examQuestion.responseOptions)
}
