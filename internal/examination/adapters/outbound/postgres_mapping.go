package outbound

import (
	"context"

	"github.com/cooperlutz/go-full/internal/examination/domain/examination"
)

func (q *Queries) FindAll(ctx context.Context) ([]examination.Exam, error) {
	exams, err := q.FindAllExams(ctx)
	if err != nil {
		return nil, err
	}

	return ExaminationExamsToDomain(exams), nil
}

func (e ExaminationExam) ToDomain() examination.Exam {
	return examination.Exam{
		Name:       e.Name,
		GradeLevel: e.GradeLevel.Int32,
	}
}

func ExaminationExamsToDomain(exams []ExaminationExam) []examination.Exam {
	domainExams := make([]examination.Exam, len(exams))
	for i, exam := range exams {
		domainExams[i] = exam.ToDomain()
	}

	return domainExams
}
