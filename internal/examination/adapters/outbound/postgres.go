package outbound

import (
	"context"

	"github.com/cooperlutz/go-full/internal/examination/domain/examination"
)

// PostgresAdapter implements the examination repository using Postgres as the data store.
type PostgresAdapter struct {
	Handler IQuerierExamination
}

// NewPostgresAdapter creates a new instance of PostgresAdapter.
func NewPostgresAdapter(db DBTX) PostgresAdapter {
	return PostgresAdapter{
		Handler: NewQueriesWrapper(db),
	}
}

// FindAll retrieves all exams from the database and maps them to domain entities.
func (q PostgresAdapter) FindAll(ctx context.Context) ([]examination.Exam, error) {
	exams, err := q.Handler.FindAllExams(ctx)
	if err != nil {
		return nil, err
	}

	return ExaminationExamsToDomain(exams), nil
}
