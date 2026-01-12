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
