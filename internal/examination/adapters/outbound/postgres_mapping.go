package outbound

import (
	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/internal/examination/domain/examination"
	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
)

// ToDomain maps the ExaminationExam to the domain entity.
func (e ExaminationExam) ToDomain() examination.Exam {
	return examination.MapToExam(
		e.ExamID.Bytes,
		*pgxutil.TimestampzToTimePtr(e.CreatedAt),
		*pgxutil.TimestampzToTimePtr(e.UpdatedAt),
		e.Deleted,
		pgxutil.TimestampzToTimePtr(e.DeletedAt),
		e.StudentID.Bytes,
		pgxutil.TimestampzToTimePtr(e.StartedAt),
		pgxutil.TimestampzToTimePtr(e.CompletedAt),
		e.Completed,
		nil,
	)
}

// ExaminationExamsToDomain maps a slice of ExaminationExam to a slice of domain Exam entities.
func ExaminationExamsToDomain(exams []ExaminationExam) []examination.Exam {
	domainExams := make([]examination.Exam, len(exams))
	for i, exam := range exams {
		domainExams[i] = exam.ToDomain()
	}

	return domainExams
}

// ExaminationExamToDB maps a domain Exam entity to the ExaminationExam database model.
func ExaminationExamToDB(exam *examination.Exam) ExaminationExam {
	createdAt := exam.GetCreatedAtTime()
	updatedAt := exam.GetUpdatedAtTime()

	return ExaminationExam{
		ExamID:      pgxutil.UUIDToPgtypeUUID(exam.GetIdUUID()),
		CreatedAt:   pgxutil.TimeToTimestampz(&createdAt),
		UpdatedAt:   pgxutil.TimeToTimestampz(&updatedAt),
		Deleted:     exam.IsDeleted(),
		DeletedAt:   pgxutil.TimeToTimestampz(exam.GetDeletedAtTime()),
		StudentID:   pgxutil.UUIDToPgtypeUUID(exam.GetStudentIdUUID()),
		StartedAt:   pgxutil.TimeToTimestampz(exam.GetStartedAtTime()),
		CompletedAt: pgxutil.TimeToTimestampz(exam.GetCompletedAtTime()),
		Completed:   exam.IsCompleted(),
	}
}

func ExaminationQuestionToDB(question *examination.Question, examID uuid.UUID) ExaminationQuestion {
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
