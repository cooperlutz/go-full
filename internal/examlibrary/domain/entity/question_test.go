package entity

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/internal/examlibrary/domain/valueobject"
	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

var ValidExamQuestion = ExamQuestion{
	EntityMetadata:  baseentitee.NewEntityMetadata(),
	index:           1,
	questionText:    "What is the capital of France?",
	questionType:    valueobject.QuestionMultipleChoice,
	possiblePoints:  5,
	correctAnswer:   nil,
	responseOptions: &[]string{"Berlin", "Madrid", "Paris", "Rome"},
}

func TestNewExamQuestion(t *testing.T) {
	ind := 1
	questionText := "What is 2 + 2?"
	questionType := valueobject.QuestionMultipleChoice
	possiblePoints := 5
	correctAnswer := "4"
	options := []string{"3", "4", "5", "6"}

	examQuestion := NewExamQuestion(
		ind,
		questionText,
		questionType,
		possiblePoints,
		&correctAnswer,
		&options,
	)

	assert.Equal(t, ind, examQuestion.index)
	assert.Equal(t, questionText, examQuestion.questionText)
	assert.Equal(t, questionType, examQuestion.questionType)
	assert.Equal(t, possiblePoints, examQuestion.possiblePoints)
	assert.Equal(t, &correctAnswer, examQuestion.correctAnswer)
	assert.Equal(t, &options, examQuestion.responseOptions)
}

func TestMapToExamQuestion(t *testing.T) {
	id := uuid.New()
	ind := 1
	createdAt := time.Now()
	updatedAt := time.Now()
	deleted := false
	var deletedAt *time.Time = nil
	questionText := "What is 2 + 2?"
	questionType := "multiple-choice"
	possiblePoints := 5
	correctAnswer := "4"
	options := []string{"3", "4", "5", "6"}

	examQuestion, err := MapToExamQuestion(
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
		ind,
	)
	assert.NoError(t, err)
	assert.Equal(t, id, examQuestion.GetIdUUID())
	assert.Equal(t, ind, examQuestion.GetIndex())
	assert.Equal(t, questionText, examQuestion.questionText)
	assert.Equal(t, valueobject.QuestionMultipleChoice, examQuestion.questionType)
	assert.Equal(t, possiblePoints, examQuestion.possiblePoints)
	assert.Equal(t, &correctAnswer, examQuestion.correctAnswer)
	assert.Equal(t, &options, examQuestion.responseOptions)
}

func TestExamQuestion_GetPossiblePoints(t *testing.T) {
	assert.Equal(t, 5, ValidExamQuestion.GetPossiblePoints())
}

func TestExamQuestion_GetQuestionType(t *testing.T) {
	assert.Equal(t, valueobject.QuestionMultipleChoice, ValidExamQuestion.GetQuestionType())
}

func TestExamQuestion_GetQuestionText(t *testing.T) {
	assert.Equal(t, "What is the capital of France?", ValidExamQuestion.GetQuestionText())
}

func TestExamQuestion_GetIndex(t *testing.T) {
	assert.Equal(t, 1, ValidExamQuestion.GetIndex())
}

func TestExamQuestion_GetCorrectAnswer(t *testing.T) {
	assert.Nil(t, ValidExamQuestion.GetCorrectAnswer())
}

func TestExamQuestion_GetResponseOptions(t *testing.T) {
	expectedOptions := &[]string{"Berlin", "Madrid", "Paris", "Rome"}
	assert.Equal(t, expectedOptions, ValidExamQuestion.GetResponseOptions())
}
