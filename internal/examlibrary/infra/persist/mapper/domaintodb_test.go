package mapper_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/internal/examlibrary/domain/entity"
	"github.com/cooperlutz/go-full/internal/examlibrary/infra/persist/mapper"
	persist_postgres "github.com/cooperlutz/go-full/internal/examlibrary/infra/persist/postgres"
	"github.com/cooperlutz/go-full/test/fixtures"
)

func TestFromDomainExamToDB(t *testing.T) {
	// Arrange
	tests := []struct {
		name     string
		input    entity.Exam
		expected persist_postgres.SaveExamParams
	}{
		{
			name:     "maps AddExamToLibrary command to domain Exam",
			input:    fixtures.ValidDomainExam,
			expected: fixtures.ValidDBExam,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			result := mapper.FromDomainExamToDB(tt.input)
			// Assert
			// assert.NotNil(t, questions)
			assert.Equal(t, tt.expected.ExamID, result.ExamID)
			assert.Equal(t, tt.expected.Name, result.Name)
			assert.Equal(t, tt.expected.GradeLevel, result.GradeLevel)
			assert.Equal(t, tt.expected.CreatedAt, result.CreatedAt)
			assert.Equal(t, tt.expected.DeletedAt, result.DeletedAt)
		})
	}
}

func TestFromDomainExamQuestionToDB(t *testing.T) {
	// Arrange
	tests := []struct {
		name     string
		examId   string
		input    entity.ExamQuestion
		expected persist_postgres.SaveExamQuestionParams
	}{
		{
			name:     "maps ExamQuestion entity to DB params",
			examId:   fixtures.ValidDomainExam.GetIdUUID().String(),
			input:    fixtures.ValidDomainExam.GetQuestions()[0],
			expected: fixtures.ValidDBExamQuestion,
		},
		// {
		// 	name:     "maps another ExamQuestion entity to DB params",
		// 	examId:   fixtures.ValidDomainExam.GetIdUUID().String(),
		// 	input:    fixtures.ValidDomainExam.GetQuestions()[1],
		// 	expected: fixtures.ValidDBExamQuestion,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			result := mapper.FromDomainExamQuestionToDB(
				fixtures.ValidDomainExam.GetIdUUID(),
				tt.input,
			)
			// Assert
			assert.Equal(t, tt.expected.ExamQuestionID, result.ExamQuestionID)
			assert.Equal(t, tt.expected.ExamID, result.ExamID)
			assert.Equal(t, tt.expected.QuestionText, result.QuestionText)
			assert.Equal(t, tt.expected.PossiblePoints, result.PossiblePoints)
			assert.Equal(t, tt.expected.AnswerText, result.AnswerText)
			assert.Equal(t, tt.expected.CreatedAt, result.CreatedAt)
			assert.Equal(t, tt.expected.DeletedAt, result.DeletedAt)
		})
	}
}
