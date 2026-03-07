package outbound

import (
	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/internal/examination/app/query"
	"github.com/cooperlutz/go-full/internal/examination/domain/examination"
	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
)

// toDomain maps the ExaminationExam to the domain entity.
func (e ExaminationExam) toDomain(questions ...ExaminationQuestion) (*examination.Exam, error) {
	var domainQuestions []*examination.Question

	for i := range questions {
		q, err := questions[i].toDomain()
		if err != nil {
			return nil, err
		}

		domainQuestions = append(domainQuestions, q)
	}

	return examination.MapToExam(
		e.ExamID.Bytes,
		e.CreatedAt.Time,
		e.UpdatedAt.Time,
		e.Deleted,
		pgxutil.TimestampzToTimePtr(e.DeletedAt),
		e.StudentID.Bytes,
		e.LibraryExamID.Bytes,
		pgxutil.TimestampzToTimePtr(e.StartedAt),
		pgxutil.TimestampzToTimePtr(e.CompletedAt),
		e.TimeLimit,
		pgxutil.TimestampzToTimePtr(e.TimeOfTimeLimit),
		e.State,
		domainQuestions,
	)
}

// toQueryExam maps the ExaminationExam to the query.Exam.
func (e ExaminationExam) toQueryExam() (query.Exam, error) {
	exam, err := e.toDomain()
	if err != nil {
		return query.Exam{}, err
	}

	return mapEntityExamToQuery(exam), nil
}

// toQueryQuestion maps the ExaminationQuestion to the query.Question.
func (q ExaminationQuestion) toQueryQuestion() query.Question {
	return query.Question{
		ExamId:          q.ExamID.String(),
		Answered:        q.Answered,
		QuestionID:      q.QuestionID.String(),
		QuestionIndex:   q.Index,
		QuestionText:    q.QuestionText,
		QuestionType:    q.QuestionType,
		ResponseOptions: &q.ResponseOptions,
		ProvidedAnswer:  &q.ProvidedAnswer.String,
	}
}

// toDomain maps the ExaminationQuestion to the domain entity.
func (q ExaminationQuestion) toDomain() (*examination.Question, error) {
	return examination.MapToQuestion(
		q.QuestionID.Bytes,
		q.CreatedAt.Time,
		q.UpdatedAt.Time,
		q.Deleted,
		pgxutil.TimestampzToTimePtr(q.DeletedAt),
		q.ExamID.Bytes,
		q.QuestionText,
		q.QuestionType,
		&q.ProvidedAnswer.String,
		&q.ResponseOptions,
		q.Index,
		q.Answered,
	)
}

// examinationExamsToQuery maps a slice of ExaminationExam to a slice of query.Exam entities.
func examinationExamsToQuery(exams []ExaminationExam) ([]query.Exam, error) {
	var domainExams []query.Exam

	for _, exam := range exams {
		queryExam, err := exam.toQueryExam()
		if err != nil {
			return nil, err
		}

		domainExams = append(domainExams, queryExam)
	}

	return domainExams, nil
}

// mapEntityExamToDB maps a domain Exam entity to the ExaminationExam database model.
func mapEntityExamToDB(exam *examination.Exam) ExaminationExam {
	createdAt := exam.GetCreatedAtTime()
	updatedAt := exam.GetUpdatedAtTime()

	return ExaminationExam{
		ExamID:          pgxutil.UUIDToPgtypeUUID(exam.GetIdUUID()),
		CreatedAt:       pgxutil.TimeToTimestampz(&createdAt),
		UpdatedAt:       pgxutil.TimeToTimestampz(&updatedAt),
		Deleted:         exam.IsDeleted(),
		DeletedAt:       pgxutil.TimeToTimestampz(exam.GetDeletedAtTime()),
		StudentID:       pgxutil.UUIDToPgtypeUUID(exam.GetStudentIdUUID()),
		LibraryExamID:   pgxutil.UUIDToPgtypeUUID(exam.GetLibraryExamIdUUID()),
		StartedAt:       pgxutil.TimeToTimestampz(exam.GetStartedAtTime()),
		CompletedAt:     pgxutil.TimeToTimestampz(exam.GetCompletedAtTime()),
		TimeLimit:       exam.GetTimeLimitSeconds(),
		TimeOfTimeLimit: pgxutil.TimeToTimestampz(exam.GetTimeOfTimeLimit()),
		State:           exam.GetState().String(),
	}
}

// mapEntityQuestionToDB maps a domain Question entity to the ExaminationQuestion database model.
func mapEntityQuestionToDB(question *examination.Question, examID uuid.UUID) ExaminationQuestion {
	createdAt := question.GetCreatedAtTime()
	updatedAt := question.GetUpdatedAtTime()

	return ExaminationQuestion{
		QuestionID:      pgxutil.UUIDToPgtypeUUID(question.GetIdUUID()),
		CreatedAt:       pgxutil.TimeToTimestampz(&createdAt),
		UpdatedAt:       pgxutil.TimeToTimestampz(&updatedAt),
		Deleted:         question.IsDeleted(),
		DeletedAt:       pgxutil.TimeToTimestampz(question.GetDeletedAtTime()),
		ExamID:          pgxutil.UUIDToPgtypeUUID(examID),
		Index:           question.GetIndex(),
		Answered:        question.IsAnswered(),
		QuestionText:    question.GetQuestionText(),
		QuestionType:    question.GetQuestionType().String(),
		ProvidedAnswer:  pgxutil.StrToPgtypeText(question.GetProvidedAnswer()),
		ResponseOptions: *question.GetResponseOptions(),
	}
}

// mapEntityExamToQuery maps a domain Exam entity to a query.Exam.
func mapEntityExamToQuery(exam *examination.Exam) query.Exam {
	var questions []query.Question
	for _, q := range exam.GetQuestions() {
		questions = append(questions, mapEntityQuestionToQuery(q, exam.GetIdUUID()))
	}

	return query.Exam{
		ExamId:            exam.GetIdString(),
		StudentId:         exam.GetStudentIdString(),
		LibraryExamId:     exam.GetLibraryExamIdUUID().String(),
		State:             exam.GetState().String(),
		AnsweredQuestions: exam.AnsweredQuestionsCount(),
		TotalQuestions:    exam.NumberOfQuestions(),
		StartedAt:         exam.GetStartedAtTime(),
		CompletedAt:       exam.GetCompletedAtTime(),
		TimeLimitSeconds:  exam.GetTimeLimitSeconds(),
		TimeOfTimeLimit:   exam.GetTimeOfTimeLimit(),
		Questions:         questions,
	}
}

// mapEntityQuestionToQuery maps a domain Question entity to a query.Question.
func mapEntityQuestionToQuery(question *examination.Question, examID uuid.UUID) query.Question {
	return query.Question{
		ExamId:          examID.String(),
		Answered:        question.IsAnswered(),
		QuestionID:      question.GetIdString(),
		QuestionIndex:   question.GetIndex(),
		QuestionText:    question.GetQuestionText(),
		QuestionType:    question.GetQuestionType().String(),
		ResponseOptions: question.GetResponseOptions(),
		ProvidedAnswer:  question.GetProvidedAnswer(),
	}
}
