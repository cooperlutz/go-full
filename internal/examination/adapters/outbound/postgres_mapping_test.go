package outbound

import (
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/internal/examination/domain/examination"
)

var (
	FixtureExamSixthGrade = ExaminationExam{
		Name:       "Exam One",
		GradeLevel: pgtype.Int4{Int32: 6, Valid: true},
	}
	FixtureExamSeventhGrade = ExaminationExam{
		Name:       "Exam Two",
		GradeLevel: pgtype.Int4{Int32: 7, Valid: true},
	}
	FixtureExams = []ExaminationExam{
		FixtureExamSixthGrade,
		FixtureExamSeventhGrade,
	}
)

func TestExaminationExamToDomain(t *testing.T) {
	domainExam := FixtureExamSixthGrade.ToDomain()

	expectedExam := examination.Exam{
		Name:       "Exam One",
		GradeLevel: 6,
	}

	assert.Equal(t, expectedExam, domainExam)
}

func TestExaminationExamsToDomain(t *testing.T) {
	domainExams := ExaminationExamsToDomain(FixtureExams)

	expectedExams := []examination.Exam{
		{
			Name:       "Exam One",
			GradeLevel: 6,
		},
		{
			Name:       "Exam Two",
			GradeLevel: 7,
		},
	}

	assert.Equal(t, expectedExams, domainExams)
}
