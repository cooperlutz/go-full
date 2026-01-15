package outbound

import (
	"github.com/cooperlutz/go-full/internal/examination/domain/examination"
)

// ToDomain maps the ExaminationExam to the domain entity.
func (e ExaminationExam) ToDomain() examination.Exam {
	return examination.Exam{
		Name:       e.Name,
		GradeLevel: e.GradeLevel.Int32,
	}
}

// ExaminationExamsToDomain maps a slice of ExaminationExam to a slice of domain Exam entities.
func ExaminationExamsToDomain(exams []ExaminationExam) []examination.Exam {
	domainExams := make([]examination.Exam, len(exams))
	for i, exam := range exams {
		domainExams[i] = exam.ToDomain()
	}

	return domainExams
}
