package mapper_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/internal/examlibrary/app/command"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/common"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/mapper"
	"github.com/cooperlutz/go-full/internal/examlibrary/domain/entity"
	"github.com/cooperlutz/go-full/test/fixtures"
)

func TestFromDomainExamToAppAddExamToLibraryResult(t *testing.T) {
	// Arrange
	tests := []struct {
		name        string
		input       entity.Exam
		expected    command.AddExamToLibraryResult
		errExpected bool
	}{
		{
			name:        "maps AddExamToLibrary command to domain Exam",
			input:       fixtures.ValidDomainExam,
			expected:    fixtures.ValidAppCommandAddExamToLibraryResult,
			errExpected: false,
		},
		{
			name:        "maps AddExamToLibrary command to domain Exam with no questions",
			input:       fixtures.ValidDomainExamWithNoQuestions,
			expected:    fixtures.ValidAppCommandAddExamToLibraryResultWithNoQuestions,
			errExpected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			result := mapper.FromDomainExamToAppAddExamToLibraryResult(tt.input)
			// Assert
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFromDomainExamQuestionsToAppExamQuestions(t *testing.T) {
	// Arrange
	tests := []struct {
		name     string
		input    []entity.ExamQuestion
		expected []common.ExamQuestion
	}{
		{
			name:     "maps domain ExamQuestions to app ExamQuestions",
			input:    fixtures.ValidDomainExamQuestions,
			expected: fixtures.ValidAppExamQuestions,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			result := mapper.FromDomainExamQuestionsToAppExamQuestions(tt.input)
			// Assert
			assert.Equal(t, tt.expected, result)
		})
	}
}
