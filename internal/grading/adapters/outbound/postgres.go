package outbound

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/cooperlutz/go-full/internal/grading/app/query"
	"github.com/cooperlutz/go-full/internal/grading/domain/grading"
	"github.com/cooperlutz/go-full/pkg/deebee"
	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
	"github.com/cooperlutz/go-full/pkg/telemetree"
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

func (p PostgresAdapter) AddExam(ctx context.Context, exam *grading.Exam) error {
	ctx, span := telemetree.AddSpan(ctx, "grading.adapters.outbound.postgres.addexam")
	defer span.End()

	examToDb := mapEntityExamToDB(exam)

	err := p.postgres.AddExam(ctx, AddExamParams(examToDb))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	for _, question := range exam.GetQuestions() {
		dbQuestion := mapEntityQuestionToDB(question, exam.GetIdUUID())

		err = p.postgres.AddQuestion(ctx, AddQuestionParams(dbQuestion))
		if err != nil {
			telemetree.RecordError(ctx, err)

			return err
		}
	}

	return nil
}

func (p PostgresAdapter) GetExam(ctx context.Context, examId uuid.UUID) (*grading.Exam, error) {
	ctx, span := telemetree.AddSpan(ctx, "grading.adapters.outbound.postgres.getexam")
	defer span.End()

	examFromDb, err := p.postgres.GetExam(ctx, GetExamParams{
		ExamID: pgxutil.UUIDToPgtypeUUID(examId),
	})
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	questionsFromDb, err := p.postgres.GetQuestionsForExam(ctx, GetQuestionsForExamParams{
		ExamID: pgxutil.UUIDToPgtypeUUID(examId),
	})
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	exam, err := examFromDb.toDomain(questionsFromDb...)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return exam, nil
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

	exam, err := p.GetExam(ctx, examId)
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

func (p PostgresAdapter) FindExam(ctx context.Context, examId string) (query.Exam, error) {
	ctx, span := telemetree.AddSpan(ctx, "grading.adapters.outbound.postgres.findexam")
	defer span.End()

	examFromDb, err := p.postgres.GetExam(ctx, GetExamParams{
		ExamID: pgxutil.UUIDToPgtypeUUID(uuid.MustParse(examId)),
	})
	if err != nil {
		telemetree.RecordError(ctx, err)

		return query.Exam{}, err
	}

	questionsFromDb, err := p.postgres.GetQuestionsForExam(ctx, GetQuestionsForExamParams{
		ExamID: pgxutil.UUIDToPgtypeUUID(uuid.MustParse(examId)),
	})
	if err != nil {
		telemetree.RecordError(ctx, err)

		return query.Exam{}, err
	}

	exam := examFromDb.toQuery(questionsFromDb...)

	return exam, nil
}

func (p PostgresAdapter) FindIncompleteExams(ctx context.Context) ([]query.Exam, error) {
	ctx, span := telemetree.AddSpan(ctx, "grading.adapters.outbound.postgres.findincompleteexams")
	defer span.End()

	examsFromDb, err := p.postgres.FindAllIncompleteExams(ctx)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	var exams []query.Exam

	for _, e := range examsFromDb {
		questionsFromDb, err := p.postgres.GetQuestionsForExam(ctx, GetQuestionsForExamParams{
			ExamID: e.ExamID,
		})
		if err != nil {
			telemetree.RecordError(ctx, err)

			return nil, err
		}

		exams = append(exams, e.toQuery(questionsFromDb...))
	}

	return exams, nil
}

func (p PostgresAdapter) FindExamQuestion(ctx context.Context, examId string, questionIndex int32) (query.Question, error) {
	ctx, span := telemetree.AddSpan(ctx, "grading.adapters.outbound.postgres.findexamquestion")
	defer span.End()

	questionFromDb, err := p.postgres.FindQuestionByExamIdAndQuestionIndex(ctx, FindQuestionByExamIdAndQuestionIndexParams{
		ExamID: pgxutil.UUIDToPgtypeUUID(uuid.MustParse(examId)),
		Index:  questionIndex,
	})
	if err != nil {
		telemetree.RecordError(ctx, err)

		return query.Question{}, err
	}

	return questionFromDb.toQuery(), nil
}
