package mapper

import (
	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/internal/examlibrary/domain/entity"
	persist_postgres "github.com/cooperlutz/go-full/internal/examlibrary/infra/persist/postgres"
	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
	"github.com/cooperlutz/go-full/pkg/utilitee"
)

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

func FromDomainExamQuestionToDB(examId uuid.UUID, question entity.ExamQuestion) persist_postgres.SaveExamQuestionParams {
	createdAt := question.GetCreatedAtTime()
	updatedAt := question.GetUpdatedAtTime()
	questionIndex := question.GetIndex()
	questionIndexInt32 := utilitee.SafeIntToInt32(&questionIndex)
	possiblePoints := question.GetPossiblePoints()
	possiblePointsInt32 := utilitee.SafeIntToInt32(&possiblePoints)

	return persist_postgres.SaveExamQuestionParams{
		ExamQuestionID: pgxutil.UUIDToPgtypeUUID(question.GetIdUUID()),
		CreatedAt:      pgxutil.TimeToTimestampz(&createdAt),
		UpdatedAt:      pgxutil.TimeToTimestampz(&updatedAt),
		DeletedAt:      pgxutil.TimeToTimestampz(question.GetDeletedAtTime()),
		Deleted:        question.IsDeleted(),
		ExamID:         pgxutil.UUIDToPgtypeUUID(examId),
		Index:          questionIndexInt32,
		QuestionText:   question.GetQuestionText(),
		AnswerText:     pgxutil.StrToPgtypeText(question.GetCorrectAnswer()),
		PossiblePoints: possiblePointsInt32,
		QuestionType:   question.GetQuestionType().String(),
	}
}

// func MapFromDBQuestion(
// 	question persist_postgres.ExamQuestion,
// 	options []persist_postgres.ExamQuestionOption,
// ) entity.ExamQuestion {
// 	var optionEntities []entity.ExamQuestionOption

// 	for _, o := range options {
// 		optionEntity := MapFromDBQuestionOption(o)
// 		optionEntities = append(optionEntities, optionEntity)
// 	}

// 	qType, ok := vo.QuestionTypeFromString(question.QuestionType)
// 	if !ok {
// 		qType = vo.QuestionUnknown
// 	}

// 	questionEntity := entity.ExamQuestionFromRaw(
// 		question.ExamQuestionID.Bytes,
// 		question.CreatedAt.Time,
// 		question.UpdatedAt.Time,
// 		question.Deleted,
// 		pgxutil.TimestampzToTimePtr(question.DeletedAt),
// 		question.QuestionText,
// 		qType,
// 		question.PointValue,
// 		optionEntities,
// 	)

// 	return questionEntity
// }

// func MapFromDB(exam persist_postgres.Exam, questions *[]entity.ExamQuestion) entity.Exam {
// 	e := entity.MapToExamEntity(
// 		exam.ExamID.Bytes,
// 		exam.CreatedAt.Time,
// 		exam.UpdatedAt.Time,
// 		exam.Deleted,
// 		pgxutil.TimestampzToTimePtr(exam.DeletedAt),
// 		exam.Name,
// 		questions,
// 	)

// 	return e
// }
