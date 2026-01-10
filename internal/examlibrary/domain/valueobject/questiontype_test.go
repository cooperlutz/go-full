package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/internal/examlibrary/domain/exception"
)

func TestQuestionTypeString(t *testing.T) {
	tests := []struct {
		name     string
		qt       QuestionType
		expected string
	}{
		{
			name:     "multiple choice",
			qt:       QuestionMultipleChoice,
			expected: "multiple_choice",
		},
		{
			name:     "essay",
			qt:       QuestionEssay,
			expected: "essay",
		},
		{
			name:     "short answer",
			qt:       QuestionShortAnswer,
			expected: "short_answer",
		},
		{
			name:     "unknown",
			qt:       QuestionUnknown,
			expected: "unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.qt.String()
			if got != tt.expected {
				t.Errorf("String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestQuestionTypeFromString(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    QuestionType
		expectedErr error
	}{
		{
			name:        "valid multiple choice",
			input:       "multiple_choice",
			expected:    QuestionMultipleChoice,
			expectedErr: nil,
		},
		{
			name:        "valid essay",
			input:       "essay",
			expected:    QuestionEssay,
			expectedErr: nil,
		},
		{
			name:        "valid short answer",
			input:       "short_answer",
			expected:    QuestionShortAnswer,
			expectedErr: nil,
		},
		{
			name:        "valid unknown",
			input:       "unknown",
			expected:    QuestionUnknown,
			expectedErr: nil,
		},
		{
			name:        "invalid string",
			input:       "invalid_type",
			expected:    QuestionUnknown,
			expectedErr: exception.ErrInvalidQuestionType{},
		},
		{
			name:        "empty string",
			input:       "",
			expected:    QuestionUnknown,
			expectedErr: exception.ErrInvalidQuestionType{},
		},
		{
			name:        "case sensitive",
			input:       "ESSAY",
			expected:    QuestionUnknown,
			expectedErr: exception.ErrInvalidQuestionType{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QuestionTypeFromString(tt.input)
			assert.NotNil(t, got)
			assert.Equal(t, tt.expected, got)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
