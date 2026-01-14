package mapper_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/internal/examlibrary/domain/entity"
	"github.com/cooperlutz/go-full/internal/examlibrary/infra/persist/mapper"
	persist_postgres "github.com/cooperlutz/go-full/internal/examlibrary/infra/persist/postgres"
	"github.com/cooperlutz/go-full/test/fixtures"
)

func TestFromDBExamToDomain(t *testing.T) {
	// Arrange
	tests := []struct {
		name     string
		input    persist_postgres.ExamLibraryExam
		expected entity.Exam
	}{
		{
			name:     "maps AddExamToLibrary command to domain Exam",
			input:    fixtures.ValidDBExamLibraryExam,
			expected: fixtures.ValidDomainExam,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			result := mapper.FromDBExamToDomain(tt.input)
			// Assert
			assert.Equal(t, tt.expected.GetIdUUID(), result.GetIdUUID())
			assert.Equal(t, tt.expected.GetName(), result.GetName())
			assert.Equal(t, tt.expected.GetGradeLevel(), result.GetGradeLevel())
			assert.Equal(t, tt.expected.GetCreatedAtTime(), result.GetCreatedAtTime())
			assert.Equal(t, tt.expected.GetUpdatedAtTime(), result.GetUpdatedAtTime())
		})
	}
}

func TestFromDBExamQuestionToDomain(t *testing.T) {
	// Arrange
	tests := []struct {
		name     string
		input    persist_postgres.ExamLibraryExamQuestion
		expected entity.ExamQuestion
	}{
		{
			name:     "maps DB ExamQuestion to domain ExamQuestion",
			input:    fixtures.ValidDBExamQuestionMultipleChoice,
			expected: fixtures.ValidDomainExamQuestionMultipleChoice,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			result, err := mapper.FromDBExamQuestionToDomain(tt.input)
			assert.NoError(t, err)
			// Assert
			assert.Equal(t, tt.expected.GetIndex(), result.GetIndex())
			assert.Equal(t, tt.expected.GetQuestionText(), result.GetQuestionText())
			assert.Equal(t, tt.expected.GetQuestionType(), result.GetQuestionType())
			assert.Equal(t, tt.expected.GetPossiblePoints(), result.GetPossiblePoints())
			assert.Equal(t, tt.expected.GetCorrectAnswer(), result.GetCorrectAnswer())
			assert.Equal(t, tt.expected.GetResponseOptions(), result.GetResponseOptions())
		})
	}
}
