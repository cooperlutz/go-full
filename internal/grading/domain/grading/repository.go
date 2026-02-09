package grading

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	AddExam(ctx context.Context, exam *Exam) error

	GetExam(ctx context.Context, id uuid.UUID) (*Exam, error)

	UpdateExam(
		ctx context.Context,
		exam *Exam,
		updateFn func(h *Exam) (*Exam, error),
	) error
}
