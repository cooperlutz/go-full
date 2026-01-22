package outbound

import (
	"context"

	"github.com/cooperlutz/go-full/internal/examination/domain/examination"
	"github.com/cooperlutz/go-full/pkg/telemetree"
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
func (p PostgresAdapter) FindAll(ctx context.Context) ([]examination.Exam, error) {
	exams, err := p.Handler.FindAllExams(ctx)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return ExaminationExamsToDomain(exams), nil
}

// AddExam adds a new exam to the database.
func (p PostgresAdapter) AddExam(ctx context.Context, exam *examination.Exam) error {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.outbound.postgres.addexam")
	defer span.End()

	dbExam := ExaminationExamToDB(exam)

	err := p.Handler.AddExam(ctx, AddExamParams(dbExam))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}
