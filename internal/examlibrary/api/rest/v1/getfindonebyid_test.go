package v1_test

import (
	"net/http/httptest"
	"testing"

	v1 "github.com/cooperlutz/go-full/internal/examlibrary/api/rest/v1"
	v1_server "github.com/cooperlutz/go-full/internal/examlibrary/api/rest/v1/server"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/common"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/query"
	"github.com/cooperlutz/go-full/pkg/utilitee"
	"github.com/cooperlutz/go-full/test/fixtures"
	"github.com/cooperlutz/go-full/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestExamLibraryRestAPIControllerV1_GetFindOneByID(t *testing.T) {
	t.Parallel()

	// Arrange
	ctx := t.Context()
	mock_svc := mocks.NewMockIExamLibraryUseCase(t)
	controller := v1.NewRestAPIController(mock_svc)

	mcptr := v1_server.MultipleChoice

	tests := []struct {
		testCaseName         string
		param                v1_server.GetFindOneByIDRequestObject
		expectedResponse     v1_server.GetFindOneByIDResponseObject
		expectedResponseCode int
		expectedServiceError error
		mockFunction         *mock.Call
	}{
		{
			"GET one exam by ID successfully",
			v1_server.GetFindOneByIDRequestObject{},
			v1_server.GetFindOneByID200JSONResponse{
				Body: v1_server.Exam{
					Id:         utilitee.StrPtr(fixtures.ValidUUID.String()),
					Name:       utilitee.StrPtr("Sample Exam"),
					GradeLevel: utilitee.IntPtr(5),
					Questions: &[]v1_server.ExamQuestion{
						{
							PossiblePoints:  utilitee.IntPtr(5),
							QuestionType:    &mcptr,
							Index:           utilitee.IntPtr(1),
							QuestionText:    utilitee.StrPtr("What is 2 + 2?"),
							PossibleAnswers: &[]string{"3", "4", "5", "6"},
							CorrectAnswer:   utilitee.StrPtr("4"),
						},
					},
				},
				Headers: v1_server.GetFindOneByID200ResponseHeaders{
					XRequestId: "00000000000000000000000000000000",
				},
			},
			200,
			nil,
			mock_svc.On(
				"FindOneExamByID",
				ctx,
				query.FindOneExamByID{
					ExamID: fixtures.ValidUUID.String(),
				},
			).Return(
				query.FindOneExamByIDResponse{
					ExamID:     fixtures.ValidUUID.String(),
					Name:       "Sample Exam",
					GradeLevel: 5,
					Questions: &[]common.ExamQuestion{
						{
							Index:           1,
							PossiblePoints:  5,
							QuestionType:    "multiple-choice",
							QuestionText:    "What is 2 + 2?",
							ResponseOptions: &[]string{"3", "4", "5", "6"},
							CorrectAnswer:   utilitee.StrPtr("4"),
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
			response, err := controller.GetFindOneByID(
				ctx,
				tt.param,
			)
			// Assert
			assert.Equal(t, tt.expectedServiceError, err)
			assert.Equal(t, tt.expectedResponse, response)
			// Assertions of the written response
			response.VisitGetFindOneByIDResponse(rr)
			assert.Equal(t, tt.expectedResponseCode, rr.Code)
		})
	}
}

func TestExamLibraryRestAPIControllerV1_GetFindOneByID_Failure(t *testing.T) {
	t.Parallel()

	// Arrange
	ctx := t.Context()
	mock_svc := mocks.NewMockIExamLibraryUseCase(t)
	controller := v1.NewRestAPIController(mock_svc)

	tests := []struct {
		testCaseName         string
		param                v1_server.GetFindOneByIDRequestObject
		expectedResponse     v1_server.GetFindOneByIDResponseObject
		expectedResponseCode int
		expectedServiceError error
		mockFunction         *mock.Call
	}{
		{
			"GET one exam by ID failure",
			v1_server.GetFindOneByIDRequestObject{},
			v1_server.GetFindOneByID400Response{},
			400,
			assert.AnError,
			mock_svc.On(
				"FindOneExamByID",
				ctx,
				query.FindOneExamByID{
					ExamID: fixtures.ValidUUID.String(),
				},
			).Return(
				query.FindOneExamByIDResponse{},
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
			response, err := controller.GetFindOneByID(
				ctx,
				tt.param,
			)
			// Assert
			assert.Equal(t, tt.expectedServiceError, err)
			assert.Equal(t, tt.expectedResponse, response)
			// Assertions of the written response
			response.VisitGetFindOneByIDResponse(rr)
			assert.Equal(t, tt.expectedResponseCode, rr.Code)
		})
	}
}
