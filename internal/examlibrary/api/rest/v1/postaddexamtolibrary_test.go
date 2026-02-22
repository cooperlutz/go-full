package v1_test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	v1 "github.com/cooperlutz/go-full/internal/examlibrary/api/rest/v1"
	v1_server "github.com/cooperlutz/go-full/internal/examlibrary/api/rest/v1/server"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/command"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/common"
	"github.com/cooperlutz/go-full/test/fixtures"
	"github.com/cooperlutz/go-full/test/mocks"
)

func TestExamLibraryRestAPIControllerV1_PostAddExamToLibrary(t *testing.T) {
	t.Parallel()

	// Arrange
	ctx := t.Context()
	// Mock the service
	mock_svc := mocks.NewMockIExamLibraryUseCase(t)
	controller := v1.NewRestAPIController(mock_svc)

	mcptr := v1_server.MultipleChoice

	tests := []struct {
		testCaseName         string
		param                v1_server.PostAddExamToLibraryRequestObject
		expectedResponse     v1_server.PostAddExamToLibraryResponseObject
		expectedResponseCode int
		expectedServiceError error
		mockFunction         *mock.Call
	}{
		{
			"POST a new exam to the library successfully",
			v1_server.PostAddExamToLibraryRequestObject{
				Body: &v1_server.Exam{
					Name:       new("Sample Exam"),
					GradeLevel: new(5),
					TimeLimit:  new(int64(3600)),
					Questions: &[]v1_server.ExamQuestion{
						{
							Index:          new(1),
							QuestionText:   new("What animal is known to bark?"),
							QuestionType:   &mcptr,
							PossiblePoints: new(5),
							PossibleAnswers: &[]string{
								"dog",
								"cat",
								"bird",
								"fish",
							},
							CorrectAnswer: new("dog"),
						},
					},
				},
			},
			v1_server.PostAddExamToLibrary200JSONResponse{
				Body: v1_server.Exam{
					Id:         new(fixtures.ValidUUID.String()),
					Name:       new("Sample Exam"),
					GradeLevel: new(5),
					TimeLimit:  new(int64(3600)),
					Questions: &[]v1_server.ExamQuestion{
						{
							Index:          new(1),
							QuestionText:   new("What animal is known to bark?"),
							QuestionType:   &mcptr,
							PossiblePoints: new(5),
							PossibleAnswers: &[]string{
								"dog",
								"cat",
								"bird",
								"fish",
							},
							CorrectAnswer: new("dog"),
						},
					},
				},
				Headers: v1_server.PostAddExamToLibrary200ResponseHeaders{
					XRequestId: "00000000000000000000000000000000",
				},
			},
			200,
			nil,
			mock_svc.On(
				"AddExamToLibrary",
				mock.Anything,
				command.NewAddExamToLibrary(
					"Sample Exam",
					5,
					3600,
					[]common.ExamQuestion{
						common.NewExamQuestion(
							1,
							"What animal is known to bark?",
							"multiple-choice",
							5,
							new("dog"),
							&[]string{"dog", "cat", "bird", "fish"},
						),
					},
				),
			).Return(
				command.AddExamToLibraryResult{
					ExamID:     fixtures.ValidUUID.String(),
					Name:       "Sample Exam",
					GradeLevel: 5,
					TimeLimit:  3600,
					Questions: []common.ExamQuestion{
						common.NewExamQuestion(
							1,
							"What animal is known to bark?",
							"multiple-choice",
							5,
							new("dog"),
							&[]string{"dog", "cat", "bird", "fish"},
						),
					},
				},
				nil,
			),
		},
		{
			"POST a new exam to the library with missing name results in error",
			v1_server.PostAddExamToLibraryRequestObject{
				Body: &v1_server.Exam{
					GradeLevel: new(5),
					TimeLimit:  new(int64(3600)),
					Questions:  &[]v1_server.ExamQuestion{},
				},
			},
			v1_server.PostAddExamToLibrary400Response{},
			400,
			&v1_server.RequiredParamError{
				ParamName: "Name, GradeLevel, and Questions are required",
			},
			nil,
		},
		{
			"POST a new exam to the library, use case returns invalid question type error",
			v1_server.PostAddExamToLibraryRequestObject{
				Body: &v1_server.Exam{
					Name:       new("Sample Exam"),
					GradeLevel: new(5),
					TimeLimit:  new(int64(3600)),
					Questions: &[]v1_server.ExamQuestion{
						{
							Index:          new(1),
							QuestionText:   new("What animal is known to meow?"),
							QuestionType:   &mcptr,
							PossiblePoints: new(5),
							PossibleAnswers: &[]string{
								"dog",
								"cat",
								"bird",
								"fish",
							},
							CorrectAnswer: new("cat"),
						},
					},
				},
			},
			v1_server.PostAddExamToLibrary400Response{},
			400,
			assert.AnError,
			mock_svc.On(
				"AddExamToLibrary",
				mock.Anything,
				command.NewAddExamToLibrary(
					"Sample Exam",
					5,
					3600,
					[]common.ExamQuestion{
						common.NewExamQuestion(
							1,
							"What animal is known to meow?",
							"multiple-choice",
							5,
							new("cat"),
							&[]string{"dog", "cat", "bird", "fish"},
						),
					},
				),
			).Return(
				command.AddExamToLibraryResult{},
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
			response, err := controller.PostAddExamToLibrary(
				ctx,
				tt.param,
			)
			// Assert
			assert.Equal(t, tt.expectedServiceError, err)
			assert.Equal(t, tt.expectedResponse, response)
			// Assertions of the written response
			response.VisitPostAddExamToLibraryResponse(rr)
			assert.Equal(t, tt.expectedResponseCode, rr.Code)
		})
	}
}
