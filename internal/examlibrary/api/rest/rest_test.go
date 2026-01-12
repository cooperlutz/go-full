package rest_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/cooperlutz/go-full/internal/examlibrary/api/rest"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/command"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/common"
	"github.com/cooperlutz/go-full/pkg/utilitee"
	"github.com/cooperlutz/go-full/test/fixtures"
	mocks "github.com/cooperlutz/go-full/test/mocks"
)

// Test that /api/v1/ping returns 200 and "pong"
func TestNewExamLibraryAPIRouter(t *testing.T) {
	t.Parallel()

	// Arrange
	useCase := mocks.NewMockIExamLibraryUseCase(t)
	router := rest.NewExamLibraryAPIRouter(useCase)
	useCase.Mock.On(
		"AddExamToLibrary",
		mock.Anything,
		command.NewAddExamToLibrary(
			"Sample Exam",
			5,
			[]common.ExamQuestion{
				common.NewExamQuestion(
					1,
					"What animal is known to bark?",
					"multiple-choice",
					5,
					utilitee.StrPtr("dog"),
					&[]string{"dog", "cat", "bird", "fish"},
				),
			},
		),
	).Return(
		command.AddExamToLibraryResult{
			ExamID:     fixtures.ValidUUID.String(),
			Name:       "Sample Exam",
			GradeLevel: 5,
			Questions: []common.ExamQuestion{
				common.NewExamQuestion(
					1,
					"What animal is known to bark?",
					"multiple-choice",
					5,
					utilitee.StrPtr("dog"),
					&[]string{"dog", "cat", "bird", "fish"},
				),
			},
		},
		nil,
	)
	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/exams",
		bytes.NewBufferString(`
{
	"name": "Sample Exam",
	"gradeLevel": 5,
	"questions": [
		{
			"index": 1,
			"questionText": "What animal is known to bark?",
			"questionType": "multiple-choice",
			"possiblePoints": 5,
			"possibleAnswers": ["dog", "cat", "bird", "fish"],
			"correctAnswer": "dog"
		}
	]
}`),
	)
	w := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, string(data), "Sample Exam")
}
