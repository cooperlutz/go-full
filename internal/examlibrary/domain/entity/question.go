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
