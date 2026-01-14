package persist

import (
	"context"

	"github.com/cooperlutz/go-full/internal/examlibrary/domain/entity"
	"github.com/cooperlutz/go-full/internal/examlibrary/infra/persist/mapper"
	persist_postgres "github.com/cooperlutz/go-full/internal/examlibrary/infra/persist/postgres"
	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

// FindAllExams finds all exams in the database.
func (r *examLibraryPersistPostgresRepository) FindAllExams(ctx context.Context) ([]entity.Exam, error) {
	ctx, span := telemetree.AddSpan(ctx, "persist.postgres.findAllExams")
	defer span.End()

	dbExams, err := r.query.FindAllExams(ctx)
	if err != nil {
		telemetree.RecordError(ctx, err, "failed to find all exams")

		return nil, err
	}

	var examEntities []entity.Exam

	for _, dbExam := range dbExams {
		examEntity := mapper.FromDBExamToDomain(dbExam)

		dbExamQuestions, err := r.query.FindAllExamQuestions(ctx,
			persist_postgres.FindAllExamQuestionsParams{
				ExamID: pgxutil.UUIDToPgtypeUUID(examEntity.GetIdUUID()),
			})
		if err != nil {
			telemetree.RecordError(ctx, err, "failed to find all exams")

			return nil, err
		}

		var questionEntities []entity.ExamQuestion

		for _, dbQuestion := range dbExamQuestions {
			questionEntity, err := mapper.FromDBExamQuestionToDomain(dbQuestion)
			if err != nil {
				telemetree.RecordError(ctx, err, "failed to map exam question from DB to domain")

				return nil, err
			}

			questionEntities = append(questionEntities, questionEntity)
		}

		for _, questionEntity := range questionEntities {
			examEntity.AddQuestion(questionEntity)
		}

		examEntities = append(examEntities, examEntity)
	}

	return examEntities, nil
}
