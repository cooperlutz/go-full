package repository

import (
	"context"

	"github.com/cooperlutz/go-full/internal/examlibrary/domain/entity"
	"github.com/google/uuid"
)

type IExamLibraryRepository interface {
	// Commands
	SaveExam(ctx context.Context, exam entity.Exam) error

	// Queries
	FindExamByID(ctx context.Context, examID uuid.UUID) (entity.Exam, error)
	FindAllExams(ctx context.Context) ([]entity.Exam, error)
}
