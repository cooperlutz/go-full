package mapper_test

import (
	"testing"

	"github.com/cooperlutz/go-full/internal/examlibrary/api/rest/v1/mapper"
	"github.com/cooperlutz/go-full/internal/examlibrary/api/rest/v1/server"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/command"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/query"
	"github.com/cooperlutz/go-full/test/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestFromAppExamWithoutQuestionsToApiExamMetadata(t *testing.T) {
	// Arrange
	tests := []struct {
		name     string
		input    query.ExamWithoutQuestions
		expected server.ExamMetadata
	}{
		{
			"Convert multiple-choice question from API to App model",
			fixtures.ValidAppExamWithoutQuestions,
			fixtures.ValidApiExamMetadata,
		},
		{
			"Convert short-answer question from API to App model",
			fixtures.ValidAppExamWithoutQuestions,
			fixtures.ValidApiExamMetadata,
		},
		{
			"Convert essay question from API to App model",
			fixtures.ValidAppExamWithoutQuestions,
			fixtures.ValidApiExamMetadata,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			result := mapper.FromAppExamWithoutQuestionsToApiExamMetadata(tt.input)
			// Assert
			assert.Equal(t, *tt.expected.Id, *result.Id)
			assert.Equal(t, *tt.expected.Name, *result.Name)
			assert.Equal(t, *tt.expected.GradeLevel, *result.GradeLevel)
		})
	}
}

func TestFromAppFindOneExamByIDResponseToApiExam(t *testing.T) {
	// Arrange
	tests := []struct {
		name     string
		input    query.FindOneExamByIDResponse
		expected server.Exam
	}{
		{
			"Convert FindOneExamByIDResponse from App to API model",
			fixtures.ValidAppFindOneExamByIDResponse,
			fixtures.ValidApiExam,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			result := mapper.FromAppFindOneExamByIDResponseToApiExam(tt.input)
			// Assert
			assert.Equal(t, *tt.expected.Id, *result.Id)
			assert.Equal(t, *tt.expected.Name, *result.Name)
			assert.Equal(t, *tt.expected.GradeLevel, *result.GradeLevel)
			assert.Equal(t, len(*tt.expected.Questions), len(*result.Questions))
			for i := range *tt.expected.Questions {
				expectedQ := (*tt.expected.Questions)[i]
				resultQ := (*result.Questions)[i]
				assert.Equal(t, *expectedQ.Index, *resultQ.Index)
				assert.Equal(t, *expectedQ.QuestionText, *resultQ.QuestionText)
				assert.Equal(t, *expectedQ.QuestionType, *resultQ.QuestionType)
				assert.Equal(t, *expectedQ.PossiblePoints, *resultQ.PossiblePoints)
				assert.Equal(t, expectedQ.CorrectAnswer, resultQ.CorrectAnswer)
				assert.Equal(t, expectedQ.PossibleAnswers, resultQ.PossibleAnswers)
			}
		})
	}
}

func TestFromAppExamsWithoutQuestionsToApiExamMetadataList(t *testing.T) {
	// Arrange
	tests := []struct {
		name     string
		input    []query.ExamWithoutQuestions
		expected []server.ExamMetadata
	}{
		{
			"Convert list of ExamWithoutQuestions from App to API model",
			fixtures.ValidAppExamsWithoutQuestions,
			fixtures.ValidApiExamsMetadataList,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			result := mapper.FromAppExamsWithoutQuestionsToApiExamMetadataList(tt.input)
			// Assert
			assert.Equal(t, len(tt.expected), len(result))
			for i := range tt.expected {
				assert.Equal(t, *tt.expected[i].Id, *result[i].Id)
				assert.Equal(t, *tt.expected[i].Name, *result[i].Name)
				assert.Equal(t, *tt.expected[i].GradeLevel, *result[i].GradeLevel)
			}
		})
	}
}

func TestFromAppAddExamToLibraryResultToApiExam(t *testing.T) {
	// Arrange
	tests := []struct {
		name     string
		input    command.AddExamToLibraryResult
		expected server.Exam
	}{
		{
			"Convert AddExamToLibraryResult from App to API model",
			fixtures.ValidAppCommandAddExamToLibraryResult,
			fixtures.ValidApiExam,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			result := mapper.FromAppAddExamToLibraryResultToApiExam(tt.input)
			// Assert
			assert.Equal(t, *tt.expected.Id, *result.Id)
			assert.Equal(t, *tt.expected.Name, *result.Name)
			assert.Equal(t, *tt.expected.GradeLevel, *result.GradeLevel)
			assert.Equal(t, len(*tt.expected.Questions), len(*result.Questions))
			for i := range *tt.expected.Questions {
				expectedQ := (*tt.expected.Questions)[i]
				resultQ := (*result.Questions)[i]
				assert.Equal(t, *expectedQ.Index, *resultQ.Index)
				assert.Equal(t, *expectedQ.QuestionText, *resultQ.QuestionText)
				assert.Equal(t, *expectedQ.QuestionType, *resultQ.QuestionType)
				assert.Equal(t, *expectedQ.PossiblePoints, *resultQ.PossiblePoints)
				assert.Equal(t, expectedQ.CorrectAnswer, resultQ.CorrectAnswer)
				assert.Equal(t, expectedQ.PossibleAnswers, resultQ.PossibleAnswers)
			}
		})
	}
}
