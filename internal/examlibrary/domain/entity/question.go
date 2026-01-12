package entity

import (
	"time"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/internal/examlibrary/domain/valueobject"
	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type ExamQuestion struct {
	*baseentitee.EntityMetadata
	index           int
	questionText    string
	questionType    valueobject.QuestionType
	possiblePoints  int
	correctAnswer   *string
	responseOptions *[]string
}

func (eq ExamQuestion) GetIndex() int {
	return eq.index
}

func (eq ExamQuestion) GetQuestionText() string {
	return eq.questionText
}

func (eq ExamQuestion) GetQuestionType() valueobject.QuestionType {
	return eq.questionType
}

func (eq ExamQuestion) GetPossiblePoints() int {
	return eq.possiblePoints
}

func (eq ExamQuestion) GetCorrectAnswer() *string {
	return eq.correctAnswer
}

func (eq ExamQuestion) GetResponseOptions() *[]string {
	return eq.responseOptions
}

func NewExamQuestion(
	questionText string,
	questionType valueobject.QuestionType,
	possiblePoints int,
	correctAnswer *string,
	options *[]string,
) *ExamQuestion {
	return &ExamQuestion{
		EntityMetadata:  baseentitee.NewEntityMetadata(),
		questionText:    questionText,
		questionType:    questionType,
		possiblePoints:  possiblePoints,
		correctAnswer:   correctAnswer,
		responseOptions: options,
	}
}

func MapToExamQuestion(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	deleted bool,
	deletedAt *time.Time,
	questionText string,
	questionType valueobject.QuestionType,
	possiblePoints int,
	correctAnswer *string,
	responseOptions *[]string,
	index int,
) ExamQuestion {
	return ExamQuestion{
		EntityMetadata: baseentitee.MapToEntityMetadataFromCommonTypes(
			id,
			createdAt,
			updatedAt,
			deleted,
			deletedAt,
		),
		questionText:    questionText,
		questionType:    questionType,
		possiblePoints:  possiblePoints,
		correctAnswer:   correctAnswer,
		responseOptions: responseOptions,
		index:           index,
	}
}
