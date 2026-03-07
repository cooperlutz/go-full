package grading

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type Repository interface {
	AddExam(ctx context.Context, exam *Exam) error

	GetExam(ctx context.Context, id uuid.UUID) (*Exam, error)

	UpdateExam(
		ctx context.Context,
		examId uuid.UUID,
		updateFn func(e *Exam) (*Exam, error),
	) error
}

func MapToExam(
	id uuid.UUID,
	createdAt, updatedAt time.Time,
	deleted bool,
	deletedAt *time.Time,
	studentId uuid.UUID,
	libraryExamId uuid.UUID,
	examinationExamId uuid.UUID,
	questions []*Question,
	totalPointsPossible int32,
	totalPointsReceived *int32,
	state string,
) (*Exam, error) {
	examState, err := ExamStateFromString(state)
	if err != nil {
		return nil, err
	}

	return &Exam{
		EntityMetadata: baseentitee.MapToEntityMetadataFromCommonTypes(
			id,
			createdAt,
			updatedAt,
			deleted,
			deletedAt,
		),
		studentId:           studentId,
		libraryExamId:       libraryExamId,
		examinationExamId:   examinationExamId,
		questions:           questions,
		state:               examState,
		totalPossiblePoints: totalPointsPossible,
		totalPointsReceived: totalPointsReceived,
	}, nil
}

func MapToQuestion(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	deleted bool,
	deletedAt *time.Time,
	examId uuid.UUID,
	index int32,
	questionType string,
	graded bool,
	feedback *string,
	providedAnswer string,
	correctAnswer *string,
	correctlyAnswered *bool,
	pointsPossible int32,
	pointsReceived *int32,
) (*Question, error) {
	qType, err := QuestionTypeFromString(questionType)
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
		examId:            examId,
		questionType:      qType,
		index:             index,
		graded:            graded,
		feedback:          feedback,
		providedAnswer:    providedAnswer,
		correctAnswer:     correctAnswer,
		correctlyAnswered: correctlyAnswered,
		pointsPossible:    pointsPossible,
		pointsReceived:    pointsReceived,
	}, nil
}
