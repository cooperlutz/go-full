package persist

import (
	"context"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/internal/examlibrary/domain/entity"
)

func (r *examLibraryPersistPostgresRepository) FindExamByID(ctx context.Context, examID uuid.UUID) (entity.Exam, error) {
	return entity.Exam{}, nil
}
