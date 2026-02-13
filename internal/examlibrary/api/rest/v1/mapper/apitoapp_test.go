package mapper_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/internal/examlibrary/api/rest/v1/mapper"
	"github.com/cooperlutz/go-full/internal/examlibrary/api/rest/v1/server"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/command"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/common"
	"github.com/cooperlutz/go-full/test/fixtures"
)

func TestFromApiExamQuestionToAppExamQuestion(t *testing.T) {
	// Arrange
	tests := []struct {
		name     string
		input    server.ExamQuestion
		expected common.ExamQuestion
	}{
		{
			name:     "converts API exam question to app exam question",
			input:    fixtures.ValidApiExamQuestionMultipleChoice,
			expected: fixtures.ValidAppExamQuestionMultipleChoice,
		},
		{
			name:     "converts short answer question",
			input:    fixtures.ValidApiExamQuestionShortAnswer,
			expected: fixtures.ValidAppExamQuestionShortAnswer,
		},
		{
			name:     "converts essay question",
			input:    fixtures.ValidApiExamQuestionEssay,
			expected: fixtures.ValidAppExamQuestionEssay,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			result := mapper.FromApiExamQuestionToAppExamQuestion(tt.input)
			// Assert
			assert.Equal(t, tt.expected.Index, result.Index)
			assert.Equal(t, tt.expected.QuestionText, result.QuestionText)
			assert.Equal(t, tt.expected.QuestionType, result.QuestionType)
			assert.Equal(t, tt.expected.PossiblePoints, result.PossiblePoints)
		})
	}
}

func TestFromApiExamQuestionsToAppExamQuestions(t *testing.T) {
	// Arrange
	tests := []struct {
		name     string
		input    []server.ExamQuestion
		expected []common.ExamQuestion
	}{
		{
			name:     "converts list of API exam questions to app exam questions",
			input:    fixtures.ValidApiExamQuestions,
			expected: fixtures.ValidAppExamQuestions,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			result := mapper.FromApiExamQuestionsToAppExamQuestions(tt.input)
			// Assert
			assert.Equal(t, len(tt.expected), len(result))
			for i := range tt.expected {
				assert.Equal(t, tt.expected[i].Index, result[i].Index)
				assert.Equal(t, tt.expected[i].QuestionText, result[i].QuestionText)
				assert.Equal(t, tt.expected[i].QuestionType, result[i].QuestionType)
				assert.Equal(t, tt.expected[i].PossiblePoints, result[i].PossiblePoints)
			}
		})
	}
}

func TestFromAppExamQuestionToApiExamQuestion(t *testing.T) {
	// Arrange
	tests := []struct {
		name     string
		input    common.ExamQuestion
		expected server.ExamQuestion
	}{
		{
			name:     "converts app exam question to API exam question",
			input:    fixtures.ValidAppExamQuestionMultipleChoice,
			expected: fixtures.ValidApiExamQuestionMultipleChoice,
		},
		{
			name:     "converts short answer question",
			input:    fixtures.ValidAppExamQuestionShortAnswer,
			expected: fixtures.ValidApiExamQuestionShortAnswer,
		},
		{
			name:     "converts essay question",
			input:    fixtures.ValidAppExamQuestionEssay,
			expected: fixtures.ValidApiExamQuestionEssay,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			result := mapper.FromAppExamQuestionToApiExamQuestion(tt.input)
			// Assert
			assert.Equal(t, *tt.expected.Index, *result.Index)
			assert.Equal(t, *tt.expected.QuestionText, *result.QuestionText)
			assert.Equal(t, *tt.expected.QuestionType, *result.QuestionType)
			assert.Equal(t, *tt.expected.PossiblePoints, *result.PossiblePoints)
		})
	}
}

func TestFromAppExamQuestionsToApiExamQuestions(t *testing.T) {
	// Arrange
	tests := []struct {
		name     string
		input    []common.ExamQuestion
		expected []server.ExamQuestion
	}{
		{
			name:     "converts list of app exam questions to API exam questions",
			input:    fixtures.ValidAppExamQuestions,
			expected: fixtures.ValidApiExamQuestions,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			result := mapper.FromAppExamQuestionsToApiExamQuestions(tt.input)
			// Assert
			assert.Equal(t, len(tt.expected), len(result))
			for i := range tt.expected {
				assert.Equal(t, *tt.expected[i].Index, *result[i].Index)
				assert.Equal(t, *tt.expected[i].QuestionText, *result[i].QuestionText)
				assert.Equal(t, *tt.expected[i].QuestionType, *result[i].QuestionType)
				assert.Equal(t, *tt.expected[i].PossiblePoints, *result[i].PossiblePoints)
			}
		})
	}
}

func TestFromApiExamToAppAddExamToLibrary(t *testing.T) {
	// Arrange
	tests := []struct {
		name        string
		input       server.Exam
		expected    command.AddExamToLibrary
		expectError bool
	}{
		{
			name: "converts API exam to app AddExamToLibrary command",
			input: server.Exam{
				Name:       new("Sample Exam"),
				GradeLevel: new(5),
				Questions:  &fixtures.ValidApiExamQuestions,
			},
			expected: command.NewAddExamToLibrary(
				"Sample Exam",
				5,
				fixtures.ValidAppExamQuestions,
			),
			expectError: false,
		},
		{
			name: "returns error when required fields are missing",
			input: server.Exam{
				Name:      nil,
				Questions: &fixtures.ValidApiExamQuestions,
			},
			expected:    command.AddExamToLibrary{},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			result, err := mapper.FromApiExamToAppAddExamToLibrary(tt.input)
			// Assert
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected.Name, result.Name)
				assert.Equal(t, tt.expected.GradeLevel, result.GradeLevel)
				assert.Equal(t, len(tt.expected.Questions), len(result.Questions))
				for i := range tt.expected.Questions {
					assert.Equal(t, tt.expected.Questions[i].Index, result.Questions[i].Index)
					assert.Equal(t, tt.expected.Questions[i].QuestionText, result.Questions[i].QuestionText)
					assert.Equal(t, tt.expected.Questions[i].QuestionType, result.Questions[i].QuestionType)
					assert.Equal(t, tt.expected.Questions[i].PossiblePoints, result.Questions[i].PossiblePoints)
				}
			}
		})
	}
}
