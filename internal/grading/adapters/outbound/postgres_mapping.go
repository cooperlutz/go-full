package outbound

import (
	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/internal/grading/app/query"
	"github.com/cooperlutz/go-full/internal/grading/domain/grading"
	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
)

// toDomain maps the GradingExam to the domain entity.
func (e GradingExam) toDomain(questions ...GradingQuestion) (*grading.Exam, error) {
	var domainQuestions []*grading.Question

	for i := range questions {
		q, err := questions[i].toDomain()
		if err != nil {
			return nil, err
		}

		domainQuestions = append(domainQuestions, q)
	}

	return grading.MapToExam(
		e.ExamID.Bytes,
		e.CreatedAt.Time,
		e.UpdatedAt.Time,
		e.Deleted,
		pgxutil.TimestampzToTimePtr(e.DeletedAt),
		e.StudentID.Bytes,
		e.LibraryExamID.Bytes,
		e.ExaminationExamID.Bytes,
		domainQuestions,
	), nil
}

func (e GradingExam) toQuery(questions ...GradingQuestion) query.Exam {
	var qs []query.Question
	for _, q := range questions {
		qs = append(qs, q.toQuery())
	}

	return query.Exam{
		ExamId:              e.ExamID.String(),
		GradingCompleted:    e.GradingCompleted,
		TotalPointsPossible: e.TotalPointsPossible.Int32,
		TotalPointsReceived: &e.TotalPointsReceived.Int32,
		Questions:           qs,
	}
}

// toDomain maps the GradingQuestion to the domain entity.
func (q GradingQuestion) toDomain() (*grading.Question, error) {
	return grading.MapToQuestion(
		q.QuestionID.Bytes,
		q.CreatedAt.Time,
		q.UpdatedAt.Time,
		q.Deleted,
		pgxutil.TimestampzToTimePtr(q.DeletedAt),
		q.ExamID.Bytes,
		q.Index,
		q.QuestionType,
		q.Graded,
		&q.Feedback.String,
		q.ProvidedAnswer,
		&q.CorrectAnswer.String,
		&q.CorrectlyAnswered.Bool,
		q.PointsPossible,
		&q.PointsReceived.Int32,
	)
}

func (q GradingQuestion) toQuery() query.Question {
	return query.Question{
		QuestionId:     q.QuestionID.String(),
		ExamId:         q.ExamID.String(),
		Index:          q.Index,
		QuestionType:   q.QuestionType,
		Graded:         q.Graded,
		Feedback:       &q.Feedback.String,
		ProvidedAnswer: q.ProvidedAnswer,
		CorrectAnswer:  &q.CorrectAnswer.String,
		PointsPossible: q.PointsPossible,
		PointsReceived: &q.PointsReceived.Int32,
	}
}

// mapEntityExamToDB maps a domain Exam entity to the GradingExam database model.
func mapEntityExamToDB(exam *grading.Exam) GradingExam {
	createdAt := exam.GetCreatedAtTime()
	updatedAt := exam.GetUpdatedAtTime()
	pointsPossible := exam.GetTotalPointsPossible()

	return GradingExam{
		ExamID:              pgxutil.UUIDToPgtypeUUID(exam.GetIdUUID()),
		CreatedAt:           pgxutil.TimeToTimestampz(&createdAt),
		UpdatedAt:           pgxutil.TimeToTimestampz(&updatedAt),
		Deleted:             exam.IsDeleted(),
		DeletedAt:           pgxutil.TimeToTimestampz(exam.GetDeletedAtTime()),
		StudentID:           pgxutil.UUIDToPgtypeUUID(exam.GetStudentId()),
		LibraryExamID:       pgxutil.UUIDToPgtypeUUID(exam.GetExamLibraryExamId()),
		ExaminationExamID:   pgxutil.UUIDToPgtypeUUID(exam.GetExaminationExamId()),
		GradingCompleted:    exam.IsCompleted(),
		TotalPointsPossible: pgxutil.Int32ToPgtypeInt4(&pointsPossible),
		TotalPointsReceived: pgxutil.Int32ToPgtypeInt4(exam.GetTotalPointsReceived()),
	}
}

// mapEntityQuestionToDB maps a domain Question entity to the GradingQuestion database model.
func mapEntityQuestionToDB(question *grading.Question, examID uuid.UUID) GradingQuestion {
	createdAt := question.GetCreatedAtTime()
	updatedAt := question.GetUpdatedAtTime()

	return GradingQuestion{
		QuestionID:        pgxutil.UUIDToPgtypeUUID(question.GetIdUUID()),
		CreatedAt:         pgxutil.TimeToTimestampz(&createdAt),
		UpdatedAt:         pgxutil.TimeToTimestampz(&updatedAt),
		Deleted:           question.IsDeleted(),
		DeletedAt:         pgxutil.TimeToTimestampz(question.GetDeletedAtTime()),
		ExamID:            pgxutil.UUIDToPgtypeUUID(examID),
		Index:             question.GetIndex(),
		QuestionType:      question.GetQuestionType().String(),
		Graded:            question.IsGraded(),
		Feedback:          pgxutil.StrToPgtypeText(question.GetFeedback()),
		ProvidedAnswer:    question.GetProvidedAnswer(),
		CorrectAnswer:     pgxutil.StrToPgtypeText(question.GetCorrectAnswer()),
		CorrectlyAnswered: pgxutil.BoolToPgtypeBool(question.IsCorrectlyAnswered()),
		PointsPossible:    question.GetPointsPossible(),
		PointsReceived:    pgxutil.Int32ToPgtypeInt4(question.GetPointsReceived()),
	}
}
