package persist

import (
	"context"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/internal/examlibrary/domain/entity"
	"github.com/cooperlutz/go-full/internal/examlibrary/infra/persist/mapper"
	persist_postgres "github.com/cooperlutz/go-full/internal/examlibrary/infra/persist/postgres"
	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

// FindExamByID finds an exam by its ID.
func (r *examLibraryPersistPostgresRepository) FindExamByID(ctx context.Context, examID uuid.UUID) (entity.Exam, error) {
	ctx, span := telemetree.AddSpan(ctx, "persist.postgres.find_exam_by_id")
	defer span.End()

	dbExam, err := r.query.FindExamByID(ctx,
		persist_postgres.FindExamByIDParams{
			ExamID: pgxutil.UUIDToPgtypeUUID(examID),
		},
	)
	if err != nil {
		telemetree.RecordError(ctx, err, "failed to find exam by ID")

		return entity.Exam{}, err
	}

	examEntity := mapper.FromDBExamToDomain(dbExam)

	dbExamQuestions, err := r.query.FindAllExamQuestions(ctx,
		persist_postgres.FindAllExamQuestionsParams{
			ExamID: pgxutil.UUIDToPgtypeUUID(examEntity.GetIdUUID()),
		})
	if err != nil {
		telemetree.RecordError(ctx, err, "failed to find all exam questions")

		return entity.Exam{}, err
	}

	var questionEntities []entity.ExamQuestion

	for _, dbQuestion := range dbExamQuestions {
		questionEntity, err := mapper.FromDBExamQuestionToDomain(dbQuestion)
		if err != nil {
			telemetree.RecordError(ctx, err, "failed to map exam question from DB to domain")

			return entity.Exam{}, err
		}

		questionEntities = append(questionEntities, questionEntity)
	}

	for _, questionEntity := range questionEntities {
		examEntity.AddQuestion(questionEntity)
	}

	return examEntity, nil
}
