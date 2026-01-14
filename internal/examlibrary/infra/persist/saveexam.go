package persist

import (
	"context"

	"github.com/cooperlutz/go-full/internal/examlibrary/domain/entity"
	"github.com/cooperlutz/go-full/internal/examlibrary/infra/persist/mapper"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

// SaveExam saves an exam to the database.
func (r *examLibraryPersistPostgresRepository) SaveExam(ctx context.Context, exam entity.Exam) error {
	ctx, span := telemetree.AddSpan(ctx, "persist.postgres.saveExam")
	defer span.End()

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(ctx); rbErr != nil {
				telemetree.RecordError(ctx, rbErr, "failed to rollback transaction")
			}

			return
		}

		if cmErr := tx.Commit(ctx); cmErr != nil {
			telemetree.RecordError(ctx, cmErr, "failed to commit transaction")
			err = cmErr
		}
	}()

	q := r.query.WithTx(tx)

	queryParams := mapper.FromDomainExamToDB(exam)

	err = q.SaveExam(ctx, queryParams)
	if err != nil {
		telemetree.RecordError(ctx, err, "failed to save exam record")

		return err
	}

	examQuestions := exam.GetQuestions()

	if len(examQuestions) > 0 {
		for _, question := range examQuestions {
			questionParams := mapper.FromDomainExamQuestionToDB(
				exam.GetIdUUID(),
				question,
			)

			err = q.SaveExamQuestion(ctx, questionParams)
			if err != nil {
				telemetree.RecordError(ctx, err, "failed to save exam question record")

				return err
			}
		}
	}

	return nil
}
