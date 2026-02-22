package e2e_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	examLibrary_api_client_v1 "github.com/cooperlutz/go-full/api/rest/examlibrary/v1/client"
)

func TestRestAPIPostExam(t *testing.T) {
	// Arrange
	ctx := context.Background()
	val := 8
	examName := "Midterm Exam"
	mc := examLibrary_api_client_v1.QuestionType("multiple-choice")

	req := examLibrary_api_client_v1.Exam{
		GradeLevel: &val,
		Name:       &examName,
		TimeLimit:  new(int64(3600)),
		Questions: &[]examLibrary_api_client_v1.ExamQuestion{
			{
				Index:          new(1),
				QuestionText:   new("What animal is known to bark?"),
				QuestionType:   &mc,
				PossiblePoints: new(5),
				CorrectAnswer:  new("dog"),
				PossibleAnswers: &[]string{
					"dog", "cat", "bird", "fish",
				},
			},
		},
	}

	// Act
	response, err := examLibraryApiClient.PostAddExamToLibraryWithResponse(ctx, req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 200, response.StatusCode())
}
