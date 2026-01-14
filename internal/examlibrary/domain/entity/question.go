package entity

import (
	"time"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/internal/examlibrary/domain/exception"
	"github.com/cooperlutz/go-full/internal/examlibrary/domain/valueobject"
	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

// ExamQuestion represents a question entity within an exam.
type ExamQuestion struct {
	*baseentitee.EntityMetadata
	index           int
	questionText    string
	questionType    valueobject.QuestionType
	possiblePoints  int
	correctAnswer   *string
	responseOptions *[]string
}

// GetIndex returns the index of the question in the exam.
func (eq ExamQuestion) GetIndex() int {
	return eq.index
}

// GetQuestionText returns the text of the question.
func (eq ExamQuestion) GetQuestionText() string {
	return eq.questionText
}

// GetQuestionType returns the type of the question.
func (eq ExamQuestion) GetQuestionType() valueobject.QuestionType {
	return eq.questionType
}

// GetPossiblePoints returns the possible points for the question.
func (eq ExamQuestion) GetPossiblePoints() int {
	return eq.possiblePoints
}

// GetCorrectAnswer returns the correct answer for the question.
func (eq ExamQuestion) GetCorrectAnswer() *string {
	return eq.correctAnswer
}

// GetResponseOptions returns the response options for the question.
func (eq ExamQuestion) GetResponseOptions() *[]string {
	return eq.responseOptions
}

// NewExamQuestion creates a new ExamQuestion entity.
func NewExamQuestion(
	index int,
	questionText string,
	questionType valueobject.QuestionType,
	possiblePoints int,
	correctAnswer *string,
	options *[]string,
) *ExamQuestion {
	return &ExamQuestion{
		EntityMetadata:  baseentitee.NewEntityMetadata(),
		index:           index,
		questionText:    questionText,
		questionType:    questionType,
		possiblePoints:  possiblePoints,
		correctAnswer:   correctAnswer,
		responseOptions: options,
	}
}

// MapToExamQuestion maps raw data to an ExamQuestion entity.
// This is typically used when reconstructing an ExamQuestion from persistent storage.
func MapToExamQuestion(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	deleted bool,
	deletedAt *time.Time,
	questionText string,
	questionType string,
	possiblePoints int,
	correctAnswer *string,
	responseOptions *[]string,
	index int,
) (ExamQuestion, error) {
	qtype, err := valueobject.QuestionTypeFromString(questionType)
	if err != nil {
		return ExamQuestion{}, exception.ErrInvalidQuestionType{}
	}

	return ExamQuestion{
		EntityMetadata: baseentitee.MapToEntityMetadataFromCommonTypes(
			id,
			createdAt,
			updatedAt,
			deleted,
			deletedAt,
		),
		questionText:    questionText,
		questionType:    qtype,
		possiblePoints:  possiblePoints,
		correctAnswer:   correctAnswer,
		responseOptions: responseOptions,
		index:           index,
	}, nil
}
