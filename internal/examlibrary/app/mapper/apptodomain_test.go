package mapper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/internal/examlibrary/app/command"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/mapper"
	"github.com/cooperlutz/go-full/internal/examlibrary/domain/entity"
	"github.com/cooperlutz/go-full/test/fixtures"
)

func TestFromAppAddExamToLibraryToDomainExam(t *testing.T) {
	// Arrange
	tests := []struct {
		name        string
		input       command.AddExamToLibrary
		expected    entity.Exam
		errExpected bool
	}{
		{
			name:        "maps AddExamToLibrary command to domain Exam",
			input:       fixtures.ValidAppCommandAddExamToLibrary,
			expected:    fixtures.ValidDomainExam,
			errExpected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			result, err := mapper.FromAppAddExamToLibraryToDomainExam(tt.input)
			// Assert
			if tt.errExpected {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expected.GetName(), result.GetName())
			assert.WithinDuration(t, time.Now(), result.GetCreatedAtTime(), time.Second)
			assert.Nil(t, result.GetDeletedAtTime())
			assert.False(t, result.IsDeleted())
			assert.Equal(t, tt.expected.GetGradeLevel(), result.GetGradeLevel())
			assert.Equal(t, len(tt.expected.GetQuestions()), len(result.GetQuestions()))
		})
	}
}
