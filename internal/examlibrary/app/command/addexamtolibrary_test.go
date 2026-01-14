package command_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/internal/examlibrary/app/command"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/common"
	"github.com/cooperlutz/go-full/pkg/utilitee"
)

func TestNewAddExamToLibrary(t *testing.T) {
	// Arrange
	name := "Sample Exam"
	gradeLevel := 10
	questions := []common.ExamQuestion{
		{
			Index:           1,
			QuestionText:    "What is 2 + 2?",
			QuestionType:    "multiple-choice",
			PossiblePoints:  5,
			ResponseOptions: &[]string{"3", "4", "5"},
			CorrectAnswer:   utilitee.StrPtr("4"),
		},
	}

	// Act
	cmd := command.NewAddExamToLibrary(name, gradeLevel, questions)

	// Assert
	assert.Equal(t, name, cmd.Name)
	assert.Equal(t, gradeLevel, cmd.GradeLevel)
	assert.Equal(t, questions, cmd.Questions)
	assert.Equal(t, 1, cmd.Questions[0].Index)
	assert.Equal(t, "What is 2 + 2?", cmd.Questions[0].QuestionText)
	assert.Equal(t, "multiple-choice", cmd.Questions[0].QuestionType)
	assert.Equal(t, 5, cmd.Questions[0].PossiblePoints)
	assert.Equal(t, &[]string{"3", "4", "5"}, cmd.Questions[0].ResponseOptions)
	assert.Equal(t, utilitee.StrPtr("4"), cmd.Questions[0].CorrectAnswer)
}

func TestNewAddExamToLibraryResult(t *testing.T) {
	// Arrange
	examID := "exam-123"
	name := "Sample Exam"
	gradeLevel := 10
	questions := []common.ExamQuestion{
		{
			Index:           1,
			QuestionText:    "What is 2 + 2?",
			QuestionType:    "multiple-choice",
			PossiblePoints:  5,
			ResponseOptions: &[]string{"3", "4", "5"},
			CorrectAnswer:   utilitee.StrPtr("4"),
		},
	}

	// Act
	result := command.NewAddExamToLibraryResult(examID, name, gradeLevel, questions)

	// Assert
	assert.Equal(t, examID, result.ExamID)
	assert.Equal(t, name, result.Name)
	assert.Equal(t, gradeLevel, result.GradeLevel)
	assert.Equal(t, questions, result.Questions)
	assert.Equal(t, 1, result.Questions[0].Index)
	assert.Equal(t, "What is 2 + 2?", result.Questions[0].QuestionText)
	assert.Equal(t, "multiple-choice", result.Questions[0].QuestionType)
	assert.Equal(t, 5, result.Questions[0].PossiblePoints)
	assert.Equal(t, &[]string{"3", "4", "5"}, result.Questions[0].ResponseOptions)
	assert.Equal(t, utilitee.StrPtr("4"), result.Questions[0].CorrectAnswer)
}
