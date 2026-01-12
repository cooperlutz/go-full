package v1_test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	v1 "github.com/cooperlutz/go-full/internal/examlibrary/api/rest/v1"
	v1_server "github.com/cooperlutz/go-full/internal/examlibrary/api/rest/v1/server"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/query"
	"github.com/cooperlutz/go-full/pkg/utilitee"
	"github.com/cooperlutz/go-full/test/fixtures"
	"github.com/cooperlutz/go-full/test/mocks"
)

func TestExamLibraryRestAPIControllerV1_GetFindAllExams(t *testing.T) {
	t.Parallel()

	// Arrange
	ctx := t.Context()
	mock_svc := mocks.NewMockIExamLibraryUseCase(t)
	controller := v1.NewRestAPIController(mock_svc)

	tests := []struct {
		testCaseName         string
		param                v1_server.GetFindAllExamsRequestObject
		expectedResponse     v1_server.GetFindAllExamsResponseObject
		expectedResponseCode int
		expectedServiceError error
		mockFunction         *mock.Call
	}{
		{
			"GET all exams successfully",
			v1_server.GetFindAllExamsRequestObject{},
			v1_server.GetFindAllExams200JSONResponse{
				Body: v1_server.Exams{
					{
						Id:         utilitee.StrPtr(fixtures.ValidUUID.String()),
						Name:       utilitee.StrPtr("Sample Exam"),
						GradeLevel: utilitee.IntPtr(5),
					},
				},
				Headers: v1_server.GetFindAllExams200ResponseHeaders{
					XRequestId: "00000000000000000000000000000000",
				},
			},
			200,
			nil,
			mock_svc.On(
				"FindAllExamsWithoutQuestions",
				ctx,
				query.FindAllExamsWithoutQuestions{},
			).Return(
				query.FindAllExamsWithoutQuestionsResponse{
					Exams: []query.ExamWithoutQuestions{
						{
							ExamID:     fixtures.ValidUUID.String(),
							Name:       "Sample Exam",
							GradeLevel: 5,
						},
					},
				},
				nil,
			),
		},
	}

	// Ensure that the mock expectations are met
	defer mock_svc.AssertExpectations(t)

	for _, tt := range tests {
		t.Run(tt.testCaseName, func(t *testing.T) {
			rr := httptest.NewRecorder()

			// Act
			response, err := controller.GetFindAllExams(
				ctx,
				tt.param,
			)
			// Assert
			assert.Equal(t, tt.expectedServiceError, err)
			assert.Equal(t, tt.expectedResponse, response)
			// Assertions of the written response
			response.VisitGetFindAllExamsResponse(rr)
			assert.Equal(t, tt.expectedResponseCode, rr.Code)
		})
	}
}

func TestExamLibraryRestAPIControllerV1_GetFindAllExams_Failure(t *testing.T) {
	t.Parallel()

	// Arrange
	ctx := t.Context()
	mock_svc := mocks.NewMockIExamLibraryUseCase(t)
	controller := v1.NewRestAPIController(mock_svc)

	tests := []struct {
		testCaseName         string
		param                v1_server.GetFindAllExamsRequestObject
		expectedResponse     v1_server.GetFindAllExamsResponseObject
		expectedResponseCode int
		expectedServiceError error
		mockFunction         *mock.Call
	}{
		{
			"GET all exams, usecase returns error",
			v1_server.GetFindAllExamsRequestObject{},
			v1_server.GetFindAllExams400Response{},
			400,
			assert.AnError,
			mock_svc.On(
				"FindAllExamsWithoutQuestions",
				ctx,
				query.FindAllExamsWithoutQuestions{},
			).Return(
				query.FindAllExamsWithoutQuestionsResponse{},
				assert.AnError,
			),
		},
	}

	// Ensure that the mock expectations are met
	defer mock_svc.AssertExpectations(t)

	for _, tt := range tests {
		t.Run(tt.testCaseName, func(t *testing.T) {
			rr := httptest.NewRecorder()
			// Act
			response, err := controller.GetFindAllExams(
				ctx,
				tt.param,
			)
			// Assert
			assert.Equal(t, tt.expectedServiceError, err)
			assert.Equal(t, tt.expectedResponse, response)
			// Assertions of the written response
			response.VisitGetFindAllExamsResponse(rr)
			assert.Equal(t, tt.expectedResponseCode, rr.Code)
		})
	}
}
