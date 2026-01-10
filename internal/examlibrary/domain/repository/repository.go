package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/internal/examlibrary/domain/entity"
)

type IExamLibraryRepository interface {
	// Commands
	SaveExam(ctx context.Context, exam entity.Exam) error
	SaveExamQuestion(ctx context.Context, question entity.ExamQuestion) error

	// Queries
	FindExamByID(ctx context.Context, examID uuid.UUID) (entity.Exam, error)
	FindAllExams(ctx context.Context) ([]entity.Exam, error)
	FindAllExamsWithQuestions(ctx context.Context) ([]entity.Exam, error)
}
