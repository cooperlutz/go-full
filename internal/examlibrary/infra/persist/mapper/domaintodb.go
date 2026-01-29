package mapper

import (
	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/internal/examlibrary/domain/entity"
	"github.com/cooperlutz/go-full/internal/examlibrary/domain/valueobject"
	persist_postgres "github.com/cooperlutz/go-full/internal/examlibrary/infra/persist/postgres"
	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
	"github.com/cooperlutz/go-full/pkg/utilitee"
)

// FromDomainExamToDB maps a domain exam entity to a database exam.
func FromDomainExamToDB(exam entity.Exam) persist_postgres.SaveExamParams {
	createdAt := exam.GetCreatedAtTime()
	updatedAt := exam.GetUpdatedAtTime()
	gradeLevel := exam.GetGradeLevel().Int()

	examParams := persist_postgres.SaveExamParams{
		ExamID:     pgxutil.UUIDToPgtypeUUID(exam.GetIdUUID()),
		CreatedAt:  pgxutil.TimeToTimestampz(&createdAt),
		UpdatedAt:  pgxutil.TimeToTimestampz(&updatedAt),
		DeletedAt:  pgxutil.TimeToTimestampz(exam.GetDeletedAtTime()),
		Deleted:    exam.IsDeleted(),
		Name:       exam.GetName(),
		GradeLevel: pgxutil.IntToPgtypeInt4(&gradeLevel),
	}

	return examParams
}

// FromDomainExamQuestionToDB maps a domain exam question entity to a database exam question.
func FromDomainExamQuestionToDB(examId uuid.UUID, question entity.ExamQuestion) persist_postgres.SaveExamQuestionParams {
	createdAt := question.GetCreatedAtTime()
	updatedAt := question.GetUpdatedAtTime()
	questionIndex := question.GetIndex()
	questionIndexInt32 := utilitee.SafeIntToInt32(&questionIndex)
	possiblePoints := question.GetPossiblePoints()
	possiblePointsInt32 := utilitee.SafeIntToInt32(&possiblePoints)

	var responseOptions []string
	if question.GetQuestionType() == valueobject.QuestionMultipleChoice {
		responseOptions = *question.GetResponseOptions()
	}

	return persist_postgres.SaveExamQuestionParams{
		ExamQuestionID:  pgxutil.UUIDToPgtypeUUID(question.GetIdUUID()),
		CreatedAt:       pgxutil.TimeToTimestampz(&createdAt),
		UpdatedAt:       pgxutil.TimeToTimestampz(&updatedAt),
		DeletedAt:       pgxutil.TimeToTimestampz(question.GetDeletedAtTime()),
		Deleted:         question.IsDeleted(),
		ExamID:          pgxutil.UUIDToPgtypeUUID(examId),
		Index:           questionIndexInt32,
		QuestionText:    question.GetQuestionText(),
		AnswerText:      pgxutil.StrToPgtypeText(question.GetCorrectAnswer()),
		PossiblePoints:  possiblePointsInt32,
		QuestionType:    question.GetQuestionType().String(),
		ResponseOptions: responseOptions,
	}
}
