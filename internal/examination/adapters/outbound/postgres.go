package outbound

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/cooperlutz/go-full/internal/examination/app/query"
	"github.com/cooperlutz/go-full/internal/examination/domain/examination"
	"github.com/cooperlutz/go-full/pkg/deebee"
	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

// PostgresAdapter implements the examination repository using Postgres as the data store.
type PostgresAdapter struct {
	Handler IQuerierExamination
}

// NewPostgresAdapter creates a new instance of PostgresAdapter.
func NewPostgresAdapter(db deebee.IDatabase) PostgresAdapter {
	return PostgresAdapter{
		Handler: NewQueriesWrapper(db),
	}
}

// FindAll retrieves all exams from the database and maps them to domain entities.
func (p PostgresAdapter) FindAll(ctx context.Context) ([]query.Exam, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.outbound.postgres.findall")
	defer span.End()

	exams, err := p.Handler.FindAllExams(ctx)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return examinationExamsToQuery(exams)
}

// FindQuestion retrieves a question by its exam ID and question index from the database.
func (p PostgresAdapter) FindQuestion(ctx context.Context, id uuid.UUID, questionIndex int32) (query.Question, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.outbound.postgres.findquestion")
	defer span.End()

	question, err := p.Handler.GetQuestionByExamAndIndex(
		ctx,
		GetQuestionByExamAndIndexParams{
			ExamID: pgxutil.UUIDToPgtypeUUID(id),
			Index:  questionIndex,
		},
	)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return query.Question{}, err
	}

	return question.toQueryQuestion(), nil
}

// FindExam retrieves an exam by its ID and maps it to a query.Exam.
func (p PostgresAdapter) FindExam(ctx context.Context, id uuid.UUID) (query.Exam, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.outbound.postgres.findexam")
	defer span.End()

	exam, err := p.GetExam(ctx, id)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return query.Exam{}, err
	}

	return mapEntityExamToQuery(exam), nil
}

// FindExam retrieves an exam by its ID from the database.
func (p PostgresAdapter) GetExam(ctx context.Context, id uuid.UUID) (*examination.Exam, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.outbound.postgres.getexam")
	defer span.End()

	exam, err := p.Handler.GetExam(
		ctx,
		GetExamParams{ExamID: pgxutil.UUIDToPgtypeUUID(id)},
	)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	questions, err := p.Handler.GetQuestionsByExam(
		ctx,
		GetQuestionsByExamParams{ExamID: pgxutil.UUIDToPgtypeUUID(id)},
	)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return exam.toDomain(questions...)
}

// AddExam adds a new exam to the database.
func (p PostgresAdapter) AddExam(ctx context.Context, exam *examination.Exam) error {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.outbound.postgres.addexam")
	defer span.End()

	dbExam := mapEntityExamToDB(exam)

	err := p.Handler.AddExam(ctx, AddExamParams(dbExam))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	err = p.addQuestions(ctx, exam)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

// addQuestions adds questions associated with an exam to the database.
func (p PostgresAdapter) addQuestions(ctx context.Context, exam *examination.Exam) error {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.outbound.postgres.addquestions")
	defer span.End()

	for _, question := range exam.GetQuestions() {
		dbQuestion := mapEntityQuestionToDB(question, exam.GetIdUUID())

		err := p.Handler.AddQuestion(ctx, AddQuestionParams(dbQuestion))
		if err != nil {
			telemetree.RecordError(ctx, err)

			return err
		}
	}

	return nil
}

// UpdateExam updates an existing exam in the database.
func (p PostgresAdapter) UpdateExam(
	ctx context.Context,
	examId uuid.UUID,
	updateFn func(e *examination.Exam) (*examination.Exam, error),
) error {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.outbound.postgres.updateexam")
	defer span.End()

	tx, err := p.Handler.Begin(ctx)
	if err != nil {
		return err
	}

	exam, err := p.GetExam(ctx, examId)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	defer func() {
		err = p.finishTransaction(ctx, err, tx)
	}()

	updatedExam, err := updateFn(exam)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	dbExam := mapEntityExamToDB(updatedExam)
	err = p.Handler.SaveExam(ctx, SaveExamParams(dbExam))

	for _, question := range updatedExam.GetQuestions() {
		dbQuestion := mapEntityQuestionToDB(question, updatedExam.GetIdUUID())

		err = p.Handler.SaveQuestion(ctx, SaveQuestionParams(dbQuestion))
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
