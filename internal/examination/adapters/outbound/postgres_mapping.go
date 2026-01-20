package outbound

import (
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
