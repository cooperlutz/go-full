package examination

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type Repository interface {
	AddExam(ctx context.Context, exam *Exam) error

	GetQuestion(ctx context.Context, id uuid.UUID) (*Question, error)

	GetExam(ctx context.Context, id uuid.UUID) (*Exam, error)

	UpdateExam(
		ctx context.Context,
		exam *Exam,
		updateFn func(h *Exam) (*Exam, error),
	) error
}

// MapToExam creates an Exam domain object from the given parameters.
// This should ONLY BE USED when reconstructing an Exam from a repository.
func MapToExam(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	deleted bool,
	deletedAt *time.Time,
	studentId uuid.UUID,
	startedAt *time.Time,
	completedAt *time.Time,
	completed bool,
	questions []*Question,
) *Exam {
	return &Exam{
		EntityMetadata: baseentitee.MapToEntityMetadataFromCommonTypes(
			id,
			createdAt,
			updatedAt,
			deleted,
			deletedAt,
		),
		studentId:   studentId,
		startedAt:   startedAt,
		completedAt: completedAt,
		completed:   completed,
		questions:   questions,
	}
}

// MapToQuestion maps raw data to a Question entity.
// This should ONLY BE USED when reconstructing a Question from a repository.
func MapToQuestion(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	deleted bool,
	deletedAt *time.Time,
	examId uuid.UUID,
	questionText string,
	questionType string,
	providedAnswer *string,
	responseOptions *[]string,
	index int32,
	answered bool,
) (*Question, error) {
	qtype, err := QuestionTypeFromString(questionType)
	if err != nil {
		return nil, err
	}

	return &Question{
		EntityMetadata: baseentitee.MapToEntityMetadataFromCommonTypes(
			id,
			createdAt,
			updatedAt,
			deleted,
			deletedAt,
		),
		examId:          examId,
		questionText:    questionText,
		questionType:    qtype,
		providedAnswer:  providedAnswer,
		responseOptions: responseOptions,
		index:           index,
		answered:        answered,
	}, nil
}
