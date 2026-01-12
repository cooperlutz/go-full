package mapper

import (
	"github.com/cooperlutz/go-full/internal/examlibrary/domain/entity"
	"github.com/cooperlutz/go-full/internal/examlibrary/domain/valueobject"
	persist_postgres "github.com/cooperlutz/go-full/internal/examlibrary/infra/persist/postgres"
	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
)

func FromDBExamToDomain(dbExam persist_postgres.ExamLibraryExam) entity.Exam {
	examEntity := entity.MapToExamEntity(
		dbExam.ExamID.Bytes,
		dbExam.CreatedAt.Time,
		dbExam.UpdatedAt.Time,
		dbExam.Deleted,
		&dbExam.DeletedAt.Time,
		dbExam.Name,
		valueobject.GradeLevel(int(dbExam.GradeLevel.Int32)),
		nil, // Questions to be mapped later
	)

	return examEntity
}

func FromDBExamQuestionToDomain(dbQuestion persist_postgres.ExamLibraryExamQuestion) (entity.ExamQuestion, error) {
	questionEntity, err := entity.MapToExamQuestion(
		dbQuestion.ExamQuestionID.Bytes,
		dbQuestion.CreatedAt.Time,
		dbQuestion.UpdatedAt.Time,
		dbQuestion.Deleted,
		pgxutil.TimestampzToTimePtr(dbQuestion.DeletedAt),
		dbQuestion.QuestionText,
		dbQuestion.QuestionType,
		int(dbQuestion.PossiblePoints),
		&dbQuestion.AnswerText.String,
		&dbQuestion.ResponseOptions,
		int(dbQuestion.Index),
	)

	return questionEntity, err
}
