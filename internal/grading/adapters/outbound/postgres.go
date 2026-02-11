package outbound

import (
	"context"

	"github.com/cooperlutz/go-full/internal/grading/domain/grading"
	"github.com/cooperlutz/go-full/pkg/deebee"
	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
	"github.com/cooperlutz/go-full/pkg/telemetree"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// PostgresAdapter implements the grading repository using Postgres as the data store.
type PostgresAdapter struct {
	postgres IQuerierGrading
}

// NewPostgresAdapter creates a new instance of PostgresAdapter.
func NewPostgresAdapter(db deebee.IDatabase) PostgresAdapter {
	return PostgresAdapter{
		postgres: NewQueriesWrapper(db),
	}
}

// UpdateExam updates an existing exam in the database.
func (p PostgresAdapter) UpdateExam(
	ctx context.Context,
	examId uuid.UUID,
	updateFn func(e *grading.Exam) (*grading.Exam, error),
) error {
	ctx, span := telemetree.AddSpan(ctx, "grading.adapters.outbound.postgres.updateexam")
	defer span.End()

	tx, err := p.postgres.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		err = p.finishTransaction(ctx, err, tx)
	}()

	examFromDb, err := p.postgres.GetExam(ctx, GetExamParams{
		ExamID: pgxutil.UUIDToPgtypeUUID(examId),
	})
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	exam, err := examFromDb.toDomain()
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	updatedExam, err := updateFn(exam)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	examToDb := mapEntityExamToDB(updatedExam)

	err = p.postgres.UpdateExam(ctx, UpdateExamParams(examToDb))

	for _, question := range updatedExam.GetQuestions() {
		dbQuestion := mapEntityQuestionToDB(question, updatedExam.GetIdUUID())

		err = p.postgres.UpdateQuestion(ctx, UpdateQuestionParams(dbQuestion))
		if err != nil {
			telemetree.RecordError(ctx, err)

			return err
		}
	}

	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

// finishTransaction commits or rolls back the transaction based on the error state.
func (p PostgresAdapter) finishTransaction(ctx context.Context, err error, tx pgx.Tx) error {
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			telemetree.RecordError(ctx, rollbackErr, "failed to rollback tx")

			return rollbackErr
		}

		return err
	} else {
		if commitErr := tx.Commit(ctx); commitErr != nil {
			telemetree.RecordError(ctx, commitErr, "failed to commit tx")

			return commitErr
		}

		return nil
	}
}
