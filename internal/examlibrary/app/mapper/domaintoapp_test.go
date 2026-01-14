package mapper_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/internal/examlibrary/app/command"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/common"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/mapper"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/query"
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

func TestFromDomainExamToAppFindOneExamByIDResponse(t *testing.T) {
	// Arrange
	tests := []struct {
		name     string
		input    entity.Exam
		expected query.FindOneExamByIDResponse
	}{
		{
			name:     "maps domain Exam to app FindOneExamByIDResponse",
			input:    fixtures.ValidDomainExam,
			expected: fixtures.ValidAppFindOneExamByIDResponse,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			result := mapper.FromDomainExamToAppFindOneExamByIDResponse(tt.input)
			// Assert
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFromDomainExamsToAppFindAllExamsWithoutQuestionsResponse(t *testing.T) {
	// Arrange
	tests := []struct {
		name     string
		input    []entity.Exam
		expected query.FindAllExamsWithoutQuestionsResponse
	}{
		{
			name:  "maps domain Exams to app FindAllExamsWithoutQuestionsResponse",
			input: []entity.Exam{fixtures.ValidDomainExam},
			expected: query.FindAllExamsWithoutQuestionsResponse{
				Exams: []query.ExamWithoutQuestions{
					fixtures.ValidAppExamWithoutQuestions,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			result := mapper.FromDomainExamsToAppFindAllExamsWithoutQuestionsResponse(tt.input)
			// Assert
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFromDomainExamsToAppFindAllExamsWithQuestionsResponse(t *testing.T) {
	// Arrange
	tests := []struct {
		name     string
		input    []entity.Exam
		expected query.FindAllExamsWithQuestionsResponse
	}{
		{
			name:  "maps domain Exams to app FindAllExamsWithQuestionsResponse",
			input: []entity.Exam{fixtures.ValidDomainExam},
			expected: query.FindAllExamsWithQuestionsResponse{
				Exams: []query.ExamWithQuestions{
					fixtures.ValidAppExamWithQuestions,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			result := mapper.FromDomainExamsToAppFindAllExamsWithQuestionsResponse(tt.input)
			// Assert
			assert.Equal(t, tt.expected, result)
		})
	}
}
