package v1_test

import (
	"errors"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	v1 "github.com/cooperlutz/go-full/internal/pingpong/api/rest/v1"
	v1_server "github.com/cooperlutz/go-full/internal/pingpong/api/rest/v1/server"
	"github.com/cooperlutz/go-full/internal/pingpong/app/command"
	"github.com/cooperlutz/go-full/internal/pingpong/app/common"
	"github.com/cooperlutz/go-full/internal/pingpong/app/query"
	"github.com/cooperlutz/go-full/internal/pingpong/domain/entity"
	"github.com/cooperlutz/go-full/internal/pingpong/domain/exception"
	"github.com/cooperlutz/go-full/pkg/types"
	"github.com/cooperlutz/go-full/pkg/utilitee"
	"github.com/cooperlutz/go-full/test/fixtures"
	mocks "github.com/cooperlutz/go-full/test/mocks/pingpong"
)

var (
	testTime        = time.Date(2023, time.January, 1, 12, 0, 0, 0, time.UTC)
	testTimeString  = "2023-01-01 12:00:00 +0000 UTC"
	validPingPongID = uuid.UUID([16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
)

func TestPingPongRestAPIController_PingPong(t *testing.T) {
	t.Parallel()
	// Arrange
	var pingResponseMessage *string = new(string)
	*pingResponseMessage = "Ping!"
	var pongResponseMessage *string = new(string)
	*pongResponseMessage = "Pong!"

	// Unit tests for the PingPongRestAPIController
	tests := []struct {
		testCaseName         string
		param                v1_server.PingPongRequestObject
		expectedResponse     v1_server.PingPongResponseObject
		expectedResponseCode int
		expectedServiceError error
	}{
		{
			"GET a ping, receive a Pong!",
			v1_server.PingPongRequestObject{
				JSONBody: &v1_server.PingPongJSONRequestBody{
					Message: utilitee.StrPtr("ping"),
				},
			},
			v1_server.PingPong200JSONResponse{
				Body: v1_server.PingPong{
					Message: pongResponseMessage,
				},
				Headers: v1_server.PingPong200ResponseHeaders{
					XRequestId: "00000000000000000000000000000000",
				},
			},
			200,
			nil,
		},
		{
			"GET a pong, receive a Ping!",
			v1_server.PingPongRequestObject{
				JSONBody: &v1_server.PingPongJSONRequestBody{
					Message: utilitee.StrPtr("pong"),
				},
			},
			v1_server.PingPong200JSONResponse{
				Body: v1_server.PingPong{
					Message: pingResponseMessage,
				},
				Headers: v1_server.PingPong200ResponseHeaders{
					XRequestId: "00000000000000000000000000000000",
				},
			},
			200,
			nil,
		},
		{
			"GET a ring, receive an error",
			v1_server.PingPongRequestObject{
				JSONBody: &v1_server.PingPongJSONRequestBody{
					Message: utilitee.StrPtr("ring"),
				},
			},
			v1_server.PingPong400Response{},
			400,
			exception.ErrPingPongMsgValidation{},
		},
	}
	ctx := t.Context()

	// Mock the service
	mock_svc := mocks.NewMockIPingPongService(t)
	pingCmdResult := command.PingPongCommandResult{
		PingPongResult: &common.PingPongResult{Message: "Pong!"},
	}
	// Define the expected behavior of the mock service
	mock_svc.On(
		"PingPong",
		ctx,
		command.PingPongCommand{Message: "ping"},
	).Return(
		pingCmdResult,
		nil,
	)
	mock_svc.On(
		"PingPong",
		ctx,
		command.PingPongCommand{Message: "pong"},
	).Return(
		command.PingPongCommandResult{
			PingPongResult: &common.PingPongResult{Message: "Ping!"},
		},
		nil,
	)
	mock_svc.On(
		"PingPong",
		ctx,
		command.PingPongCommand{Message: "ring"},
	).Return(
		command.PingPongCommandResult{},
		exception.ErrPingPongMsgValidation{},
	)
	// Ensure that the mock expectations are met
	defer mock_svc.AssertExpectations(t)

	controller := v1.NewRestAPIController(mock_svc)
	/*
		Act & Assert
	*/
	for _, tt := range tests {
		t.Run(tt.testCaseName, func(t *testing.T) {
			ctx := t.Context()
			rr := httptest.NewRecorder()
			// Act
			response, err := controller.PingPong(ctx, tt.param)
			// Assert
			assert.Equal(t, tt.expectedServiceError, err)
			assert.Equal(t, tt.expectedResponse, response)
			// Assertions of the written response
			response.VisitPingPongResponse(rr)
			assert.Equal(t, tt.expectedResponseCode, rr.Code)
		})
	}
}

// GetFindAllPingPongs
func TestPingPongRestAPIController_GetFindAllPingPongs(t *testing.T) {
	t.Parallel()
	/*
		Arrange
	*/

	// Unit tests for the PingPongRestAPIController
	tests := []struct {
		testCaseName         string
		param                v1_server.GetFindAllPingPongsRequestObject
		expectedResponse     v1_server.GetFindAllPingPongsResponseObject
		expectedResponseCode int
		expectedServiceError error
	}{
		{
			testCaseName: "GET all pingpongs, receive Ping! and Pong!",
			param:        v1_server.GetFindAllPingPongsRequestObject{},
			expectedResponse: v1_server.GetFindAllPingPongs200JSONResponse{
				Body: v1_server.PingPongsRaw{
					Pingpongs: &[]v1_server.PingPongRaw{
						{
							Message:   utilitee.StrPtr("Ping!"),
							Deleted:   utilitee.BoolPtr(false),
							DeletedAt: nil,
							CreatedAt: &testTime,
							UpdatedAt: &testTime,
							Id:        utilitee.StrPtr("00000000-0000-0000-0000-000000000001"),
						},
						{
							Message:   utilitee.StrPtr("Pong!"),
							Deleted:   utilitee.BoolPtr(false),
							DeletedAt: nil,
							CreatedAt: &testTime,
							UpdatedAt: &testTime,
							Id:        utilitee.StrPtr("00000000-0000-0000-0000-000000000001"),
						},
					},
				},
				Headers: v1_server.GetFindAllPingPongs200ResponseHeaders{
					XRequestId: "00000000000000000000000000000000",
				},
			},
			expectedResponseCode: 200,
			expectedServiceError: nil,
		},
	}

	// tempUUID, _ := uuid.Parse("00000000-0000-0000-0000-000000000000")
	ctx := t.Context()
	// Mock the service
	mock_svc := mocks.NewMockIPingPongService(t)
	// Define the expected behavior of the mock service
	mock_svc.On(
		"FindAll",
		ctx,
	).Return(
		query.FindAllQueryResponseRaw{
			Entities: []entity.PingPongEntity{
				fixtures.ValidReturningPing,
				fixtures.ValidReturningPong,
			},
		},
		nil,
	)
	// Ensure that the mock expectations are met
	defer mock_svc.AssertExpectations(t)

	controller := v1.NewRestAPIController(mock_svc)
	/*
		Act & Assert
	*/
	for _, tt := range tests {
		t.Run(tt.testCaseName, func(t *testing.T) {
			ctx := t.Context()
			rr := httptest.NewRecorder()
			// Act
			response, err := controller.GetFindAllPingPongs(ctx, tt.param)
			// Assert
			assert.Equal(t, tt.expectedServiceError, err)
			assert.Equal(t, tt.expectedResponse, response)
			// Assertions of the written response
			response.VisitGetFindAllPingPongsResponse(rr)
			assert.Equal(t, tt.expectedResponseCode, rr.Code)
		})
	}
}

// STEP 3.2. Implement API Handlers & Mappers Tests
// Develop tests for the API handlers and mappers
// GetFindOneByID
func TestPingPongRestAPIController_GetFindOneByID_Success(t *testing.T) {
	t.Parallel()

	// Arrange

	// Unit tests for the PingPongRestAPIController
	tests := []struct {
		testCaseName         string
		param                v1_server.GetFindOneByIDRequestObject
		expectedResponse     v1_server.GetFindOneByIDResponseObject
		expectedResponseCode int
		expectedServiceError error
	}{
		{
			testCaseName: "GET all pingpongs, receive Ping! and Pong!",
			param: v1_server.GetFindOneByIDRequestObject{
				PingPongID: validPingPongID,
			},
			expectedResponse: v1_server.GetFindOneByID200JSONResponse{
				Body: v1_server.PingPongRaw{
					Message:   utilitee.StrPtr("Ping!"),
					Deleted:   utilitee.BoolPtr(false),
					DeletedAt: nil,
					CreatedAt: &testTime,
					UpdatedAt: &testTime,
					Id:        utilitee.StrPtr("00000000-0000-0000-0000-000000000000"),
				},
				Headers: v1_server.GetFindOneByID200ResponseHeaders{
					XRequestId: "0000000000000000",
				},
			},
			expectedResponseCode: 200,
			expectedServiceError: nil,
		},
	}

	tempUUID, _ := uuid.Parse("00000000-0000-0000-0000-000000000000")
	ctx := t.Context()

	// Mock the service
	mock_svc := mocks.NewMockIPingPongService(t)
	// Define the expected behavior of the mock service
	mock_svc.On(
		"FindOneByID",
		ctx,
		query.FindOneByID{ID: validPingPongID},
	).Return(
		query.FindOneByIDResponse{
			PingPongRawResult: common.PingPongRawResult{
				Message:   "Ping!",
				CreatedAt: testTime,
				UpdatedAt: testTime,
				ID:        tempUUID.String(),
				DeletedAt: nil,
				Deleted:   false,
			},
		},
		nil,
	)
	// Ensure that the mock expectations are met
	defer mock_svc.AssertExpectations(t)

	controller := v1.NewRestAPIController(mock_svc)

	for _, tt := range tests {
		t.Run(tt.testCaseName, func(t *testing.T) {
			ctx := t.Context()
			rr := httptest.NewRecorder()
			// Act
			response, err := controller.GetFindOneByID(ctx, tt.param)
			// Assert
			assert.Equal(t, tt.expectedServiceError, err)
			assert.Equal(t, tt.expectedResponse, response)
			// Assertions of the written response
			response.VisitGetFindOneByIDResponse(rr)
			assert.Equal(t, tt.expectedResponseCode, rr.Code)
		})
	}
}

func TestPingPongRestAPIController_GetFindOneByID_Failure(t *testing.T) {
	t.Parallel()

	// Arrange
	tempError := errors.New("there's an error")
	rr := httptest.NewRecorder()
	ctx := t.Context()
	// Unit tests for the PingPongRestAPIController
	tests := []struct {
		testCaseName         string
		param                v1_server.GetFindOneByIDRequestObject
		expectedResponse     v1_server.GetFindOneByIDResponseObject
		expectedResponseCode int
		expectedServiceError error
	}{
		{
			testCaseName: "GET all pingpongs, receive Ping! and Pong!",
			param: v1_server.GetFindOneByIDRequestObject{
				PingPongID: validPingPongID,
			},
			expectedResponse:     v1_server.GetFindOneByID400Response{},
			expectedResponseCode: 400,
			expectedServiceError: tempError,
		},
	}

	// Mock the service
	mock_svc := mocks.NewMockIPingPongService(t)
	// Define the expected behavior of the mock service
	mock_svc.On(
		"FindOneByID",
		ctx,
		query.FindOneByID{ID: validPingPongID},
	).Return(
		query.FindOneByIDResponse{},
		tempError,
	)
	// ensure that the mock expectations are met
	defer mock_svc.AssertExpectations(t)
	// setup the controller
	controller := v1.NewRestAPIController(mock_svc)

	for _, tt := range tests {
		t.Run(tt.testCaseName, func(t *testing.T) {
			// Act
			response, err := controller.GetFindOneByID(ctx, tt.param)

			// Assert
			assert.Equal(t, tt.expectedServiceError, err)
			assert.Equal(t, tt.expectedResponse, response)
			response.VisitGetFindOneByIDResponse(rr)
			assert.Equal(t, tt.expectedResponseCode, rr.Code)
		})
	}
}

// GetFindAllPingPongs
func TestPingPongRestAPIController_GetFindAllPingPongs_Failure(t *testing.T) {
	t.Parallel()
	/*
		Arrange
	*/
	errTemp := errors.New("temp error")

	// Unit tests for the PingPongRestAPIController
	tests := []struct {
		testCaseName         string
		param                v1_server.GetFindAllPingPongsRequestObject
		expectedResponse     v1_server.GetFindAllPingPongsResponseObject
		expectedResponseCode int
		expectedServiceError error
	}{
		{
			testCaseName:         "GET all pingpongs, receive Ping! and Pong!",
			param:                v1_server.GetFindAllPingPongsRequestObject{},
			expectedResponse:     v1_server.GetFindAllPingPongsResponseObject(nil),
			expectedResponseCode: 400,
			expectedServiceError: errTemp,
		},
	}

	ctx := t.Context()
	// Mock the service
	mock_svc := mocks.NewMockIPingPongService(t)
	// Define the expected behavior of the mock service
	mock_svc.On(
		"FindAll",
		ctx,
	).Return(
		query.FindAllQueryResponseRaw{}, errTemp,
	)
	// Ensure that the mock expectations are met
	defer mock_svc.AssertExpectations(t)

	controller := v1.NewRestAPIController(mock_svc)
	/*
		Act & Assert
	*/
	for _, tt := range tests {
		t.Run(tt.testCaseName, func(t *testing.T) {
			ctx := t.Context()
			rr := httptest.NewRecorder()
			// Act
			response, err := controller.GetFindAllPingPongs(ctx, tt.param)
			// Assert
			assert.Error(t, err)
			// assert.Equal(t, tt.expectedResponse, response)
			// Assertions of the written response
			response.VisitGetFindAllPingPongsResponse(rr)
			assert.Equal(t, tt.expectedResponseCode, rr.Code)
		})
	}
}

// GetPings
func TestPingPongRestAPIController_GetPings(t *testing.T) {
	t.Parallel()
	/*
		Arrange
	*/

	// Unit tests for the PingPongRestAPIController
	tests := []struct {
		testCaseName         string
		param                v1_server.GetPingsRequestObject
		expectedResponse     v1_server.GetPingsResponseObject
		expectedResponseCode int
		expectedServiceError error
	}{
		{
			testCaseName: "GET all pingpongs, receive Ping! and Pong!",
			param:        v1_server.GetPingsRequestObject{},
			expectedResponse: v1_server.GetPings200JSONResponse{
				Body: v1_server.PingPongs{
					Pingpongs: &[]v1_server.PingPong{
						{
							Message: utilitee.StrPtr("Ping!"),
						},
						{
							Message: utilitee.StrPtr("Ping!"),
						},
					},
				},
				Headers: v1_server.GetPings200ResponseHeaders{
					XRequestId: "00000000000000000000000000000000",
				},
			},
			expectedResponseCode: 200,
			expectedServiceError: nil,
		},
	}

	ctx := t.Context()
	// Mock the service
	mock_svc := mocks.NewMockIPingPongService(t)
	// Define the expected behavior of the mock service
	mock_svc.On(
		"FindAllPings",
		ctx,
	).Return(
		query.FindAllQueryResponse{
			PingPongs: []common.PingPongResult{
				{
					Message: "Ping!",
				},
				{
					Message: "Ping!",
				},
			},
		},
		nil,
	)
	// Ensure that the mock expectations are met
	defer mock_svc.AssertExpectations(t)

	controller := v1.NewRestAPIController(mock_svc)
	/*
		Act & Assert
	*/
	for _, tt := range tests {
		t.Run(tt.testCaseName, func(t *testing.T) {
			ctx := t.Context()
			rr := httptest.NewRecorder()
			// Act
			response, err := controller.GetPings(ctx, tt.param)
			// Assert
			assert.Equal(t, tt.expectedServiceError, err)
			assert.Equal(t, tt.expectedResponse, response)
			// Assertions of the written response
			response.VisitGetPingsResponse(rr)
			assert.Equal(t, tt.expectedResponseCode, rr.Code)
		})
	}
}

// GetPings
func TestPingPongRestAPIController_GetPings_Failure(t *testing.T) {
	t.Parallel()
	/*
		Arrange
	*/
	errTemp := errors.New("temp error")

	// Unit tests for the PingPongRestAPIController
	tests := []struct {
		testCaseName         string
		param                v1_server.GetPingsRequestObject
		expectedResponse     v1_server.GetPingsResponseObject
		expectedResponseCode int
		expectedServiceError error
	}{
		{
			testCaseName:         "GET all pingpongs, receive Ping! and Pong!",
			param:                v1_server.GetPingsRequestObject{},
			expectedResponse:     nil,
			expectedResponseCode: 400,
			expectedServiceError: errTemp,
		},
	}

	ctx := t.Context()
	// Mock the service
	mock_svc := mocks.NewMockIPingPongService(t)
	// Define the expected behavior of the mock service
	mock_svc.On(
		"FindAllPings",
		ctx,
	).Return(
		query.FindAllQueryResponse{}, errTemp,
	)
	// Ensure that the mock expectations are met
	defer mock_svc.AssertExpectations(t)

	controller := v1.NewRestAPIController(mock_svc)
	/*
		Act & Assert
	*/
	for _, tt := range tests {
		t.Run(tt.testCaseName, func(t *testing.T) {
			ctx := t.Context()
			rr := httptest.NewRecorder()
			// Act
			response, err := controller.GetPings(ctx, tt.param)
			// Assert
			assert.Error(t, err)
			// assert.Equal(t, tt.expectedResponse, response)
			// Assertions of the written response
			response.VisitGetPingsResponse(rr)
			assert.Equal(t, tt.expectedResponseCode, rr.Code)
		})
	}
}

// GetPongs
func TestPingPongRestAPIController_GetPongs(t *testing.T) {
	t.Parallel()
	/*
		Arrange
	*/

	// Unit tests for the PingPongRestAPIController
	tests := []struct {
		testCaseName         string
		param                v1_server.GetPongsRequestObject
		expectedResponse     v1_server.GetPongsResponseObject
		expectedResponseCode int
		expectedServiceError error
	}{
		{
			testCaseName: "GET all pongs, receive Pong! and Pong!",
			param:        v1_server.GetPongsRequestObject{},
			expectedResponse: v1_server.GetPongs200JSONResponse{
				Body: v1_server.PingPongs{
					Pingpongs: &[]v1_server.PingPong{
						{
							Message: utilitee.StrPtr("Pong!"),
						},
						{
							Message: utilitee.StrPtr("Pong!"),
						},
					},
				},
				Headers: v1_server.GetPongs200ResponseHeaders{
					XRequestId: "00000000000000000000000000000000",
				},
			},
			expectedResponseCode: 200,
			expectedServiceError: nil,
		},
	}

	ctx := t.Context()
	// Mock the service
	mock_svc := mocks.NewMockIPingPongService(t)
	// Define the expected behavior of the mock service
	mock_svc.On(
		"FindAllPongs",
		ctx,
	).Return(
		query.FindAllQueryResponse{
			PingPongs: []common.PingPongResult{
				{
					Message: "Pong!",
				},
				{
					Message: "Pong!",
				},
			},
		},
		nil,
	)
	// Ensure that the mock expectations are met
	defer mock_svc.AssertExpectations(t)

	controller := v1.NewRestAPIController(mock_svc)
	/*
		Act & Assert
	*/
	for _, tt := range tests {
		t.Run(tt.testCaseName, func(t *testing.T) {
			ctx := t.Context()
			rr := httptest.NewRecorder()
			// Act
			response, err := controller.GetPongs(ctx, tt.param)
			// Assert
			assert.Equal(t, tt.expectedServiceError, err)
			assert.Equal(t, tt.expectedResponse, response)
			// Assertions of the written response
			response.VisitGetPongsResponse(rr)
			assert.Equal(t, tt.expectedResponseCode, rr.Code)
		})
	}
}

// GetPongs
func TestPingPongRestAPIController_GetPongs_Failure(t *testing.T) {
	t.Parallel()
	/*
		Arrange
	*/
	errTemp := errors.New("temp error")

	// Unit tests for the PingPongRestAPIController
	tests := []struct {
		testCaseName         string
		param                v1_server.GetPongsRequestObject
		expectedResponse     v1_server.GetPongsResponseObject
		expectedResponseCode int
		expectedServiceError error
	}{
		{
			testCaseName:         "GET all pingpongs, receive Ping! and Pong!",
			param:                v1_server.GetPongsRequestObject{},
			expectedResponse:     v1_server.GetPongsResponseObject(nil),
			expectedResponseCode: 400,
			expectedServiceError: errTemp,
		},
	}

	ctx := t.Context()
	// Mock the service
	mock_svc := mocks.NewMockIPingPongService(t)
	// Define the expected behavior of the mock service
	mock_svc.On(
		"FindAllPongs",
		ctx,
	).Return(
		query.FindAllQueryResponse{}, errTemp,
	)
	// Ensure that the mock expectations are met
	defer mock_svc.AssertExpectations(t)

	controller := v1.NewRestAPIController(mock_svc)
	/*
		Act & Assert
	*/
	for _, tt := range tests {
		t.Run(tt.testCaseName, func(t *testing.T) {
			ctx := t.Context()
			rr := httptest.NewRecorder()
			// Act
			response, err := controller.GetPongs(ctx, tt.param)
			// Assert
			assert.Error(t, err)
			// assert.Equal(t, tt.expectedResponse, response)
			// Assertions of the written response
			response.VisitGetPongsResponse(rr)
			assert.Equal(t, tt.expectedResponseCode, rr.Code)
		})
	}
}

// GetDailyDistribution
func TestPingPongRestAPIController_GetDailyDistribution(t *testing.T) {
	t.Parallel()
	/*
		Arrange
	*/

	// Unit tests for the PingPongRestAPIController
	tests := []struct {
		testCaseName         string
		param                v1_server.GetDailyDistributionRequestObject
		expectedResponse     v1_server.GetDailyDistributionResponseObject
		expectedResponseCode int
		expectedServiceError error
	}{
		{
			testCaseName: "GET all pingpongs, receive Ping! and Pong!",
			param:        v1_server.GetDailyDistributionRequestObject{},
			expectedResponse: v1_server.GetDailyDistribution200JSONResponse{
				Body: v1_server.Trend{
					DimensionKeys: &[]v1_server.TrendKey{
						testTimeString,
					},
					DimensionValues: &[]v1_server.TrendValue{
						1,
					},
				},
				Headers: v1_server.GetDailyDistribution200ResponseHeaders{
					XRequestId: "00000000000000000000000000000000",
				},
			},
			expectedResponseCode: 200,
			expectedServiceError: nil,
		},
	}

	ctx := t.Context()
	// Mock the service
	mock_svc := mocks.NewMockIPingPongService(t)
	// Define the expected behavior of the mock service
	mock_svc.On(
		"TotalNumberOfPingPongsPerDay",
		ctx,
	).Return(
		[]types.MeasureCountbyDateTimeMetric{
			{
				DateTime: testTime,
				Count:    1,
			},
		},
		nil,
	)
	// Ensure that the mock expectations are met
	defer mock_svc.AssertExpectations(t)

	controller := v1.NewRestAPIController(mock_svc)
	/*
		Act & Assert
	*/
	for _, tt := range tests {
		t.Run(tt.testCaseName, func(t *testing.T) {
			ctx := t.Context()
			rr := httptest.NewRecorder()
			// Act
			response, err := controller.GetDailyDistribution(ctx, tt.param)
			// Assert
			assert.Equal(t, tt.expectedServiceError, err)
			assert.Equal(t, tt.expectedResponse, response)
			// Assertions of the written response
			response.VisitGetDailyDistributionResponse(rr)
			assert.Equal(t, tt.expectedResponseCode, rr.Code)
		})
	}
}

// GetTotalPingPongs
func TestPingPongRestAPIController_GetTotalPingPongs(t *testing.T) {
	t.Parallel()
	/*
		Arrange
	*/

	// Unit tests for the PingPongRestAPIController
	tests := []struct {
		testCaseName         string
		param                v1_server.GetTotalPingPongsRequestObject
		expectedResponse     v1_server.GetTotalPingPongsResponseObject
		expectedResponseCode int
		expectedServiceError error
	}{
		{
			testCaseName: "GET all pingpongs, receive Ping! and Pong!",
			param:        v1_server.GetTotalPingPongsRequestObject{},
			expectedResponse: v1_server.GetTotalPingPongs200JSONResponse{
				Body: int(12345),
				Headers: v1_server.GetTotalPingPongs200ResponseHeaders{
					XRequestId: "00000000000000000000000000000000",
				},
			},
			expectedResponseCode: 200,
			expectedServiceError: nil,
		},
	}

	ctx := t.Context()
	// Mock the service
	mock_svc := mocks.NewMockIPingPongService(t)
	// Define the expected behavior of the mock service
	mock_svc.On(
		"TotalNumberOfPingPongs",
		ctx,
	).Return(
		types.QuantityMetric{Quantity: 12345},
		nil,
	)
	// Ensure that the mock expectations are met
	defer mock_svc.AssertExpectations(t)

	controller := v1.NewRestAPIController(mock_svc)
	/*
		Act & Assert
	*/
	for _, tt := range tests {
		t.Run(tt.testCaseName, func(t *testing.T) {
			ctx := t.Context()
			rr := httptest.NewRecorder()
			// Act
			response, err := controller.GetTotalPingPongs(ctx, tt.param)
			// Assert
			assert.Equal(t, tt.expectedServiceError, err)
			assert.Equal(t, tt.expectedResponse, response)
			// Assertions of the written response
			response.VisitGetTotalPingPongsResponse(rr)
			assert.Equal(t, tt.expectedResponseCode, rr.Code)
		})
	}
}

// GetTotalPingPongs
func TestPingPongRestAPIController_GetTotalPingPongs_Failure(t *testing.T) {
	/*
		Arrange
	*/
	errTemp := errors.New("temp error")

	// Unit tests for the PingPongRestAPIController
	tests := []struct {
		testCaseName         string
		param                v1_server.GetTotalPingPongsRequestObject
		expectedResponse     v1_server.GetTotalPingPongsResponseObject
		expectedResponseCode int
		expectedServiceError error
	}{
		{
			testCaseName:         "GET all pingpongs, receive Ping! and Pong!",
			param:                v1_server.GetTotalPingPongsRequestObject{},
			expectedResponse:     v1_server.GetTotalPingPongsResponseObject(nil),
			expectedResponseCode: 400,
			expectedServiceError: errTemp,
		},
	}

	ctx := t.Context()
	// Mock the service
	mock_svc := mocks.NewMockIPingPongService(t)
	// Define the expected behavior of the mock service
	mock_svc.On(
		"TotalNumberOfPingPongs",
		ctx,
	).Return(
		types.QuantityMetric{Quantity: 0}, errTemp,
	)
	// Ensure that the mock expectations are met
	defer mock_svc.AssertExpectations(t)

	controller := v1.NewRestAPIController(mock_svc)
	/*
		Act & Assert
	*/
	for _, tt := range tests {
		t.Run(tt.testCaseName, func(t *testing.T) {
			ctx := t.Context()
			rr := httptest.NewRecorder()
			// Act
			response, err := controller.GetTotalPingPongs(ctx, tt.param)
			// Assert
			assert.Error(t, err)
			// assert.Equal(t, tt.expectedResponse, response)
			// Assertions of the written response
			response.VisitGetTotalPingPongsResponse(rr)
			assert.Equal(t, tt.expectedResponseCode, rr.Code)
		})
	}
}

// GetTotalPings
func TestPingPongRestAPIController_GetTotalPings(t *testing.T) {
	t.Parallel()
	/*
		Arrange
	*/

	// Unit tests for the PingPongRestAPIController
	tests := []struct {
		testCaseName         string
		param                v1_server.GetTotalPingsRequestObject
		expectedResponse     v1_server.GetTotalPingsResponseObject
		expectedResponseCode int
		expectedServiceError error
	}{
		{
			testCaseName: "GET all pingpongs, receive Ping! and Pong!",
			param:        v1_server.GetTotalPingsRequestObject{},
			expectedResponse: v1_server.GetTotalPings200JSONResponse{
				Body: int(12345),
				Headers: v1_server.GetTotalPings200ResponseHeaders{
					XRequestId: "00000000000000000000000000000000",
				},
			},
			expectedResponseCode: 200,
			expectedServiceError: nil,
		},
	}

	ctx := t.Context()
	// Mock the service
	mock_svc := mocks.NewMockIPingPongService(t)
	// Define the expected behavior of the mock service
	mock_svc.On(
		"TotalNumberOfPings",
		ctx,
	).Return(
		types.QuantityMetric{Quantity: 12345},
		nil,
	)
	// Ensure that the mock expectations are met
	defer mock_svc.AssertExpectations(t)

	controller := v1.NewRestAPIController(mock_svc)
	/*
		Act & Assert
	*/
	for _, tt := range tests {
		t.Run(tt.testCaseName, func(t *testing.T) {
			ctx := t.Context()
			rr := httptest.NewRecorder()
			// Act
			response, err := controller.GetTotalPings(ctx, tt.param)
			// Assert
			assert.Equal(t, tt.expectedServiceError, err)
			assert.Equal(t, tt.expectedResponse, response)
			// Assertions of the written response
			response.VisitGetTotalPingsResponse(rr)
			assert.Equal(t, tt.expectedResponseCode, rr.Code)
		})
	}
}

// GetTotalPingPongs
func TestPingPongRestAPIController_GetTotalPings_Failure(t *testing.T) {
	t.Parallel()
	/*
		Arrange
	*/
	errTemp := errors.New("temp error")

	// Unit tests for the PingPongRestAPIController
	tests := []struct {
		testCaseName         string
		param                v1_server.GetTotalPingsRequestObject
		expectedResponse     v1_server.GetTotalPingsResponseObject
		expectedResponseCode int
		expectedServiceError error
	}{
		{
			testCaseName:         "GET all pings, receive Ping! and Pong!",
			param:                v1_server.GetTotalPingsRequestObject{},
			expectedResponse:     nil,
			expectedResponseCode: 400,
			expectedServiceError: errTemp,
		},
	}

	ctx := t.Context()
	// Mock the service
	mock_svc := mocks.NewMockIPingPongService(t)
	// Define the expected behavior of the mock service
	mock_svc.On(
		"TotalNumberOfPings",
		ctx,
	).Return(
		types.QuantityMetric{Quantity: 0}, errTemp,
	)
	// Ensure that the mock expectations are met
	defer mock_svc.AssertExpectations(t)

	controller := v1.NewRestAPIController(mock_svc)
	/*
		Act & Assert
	*/
	for _, tt := range tests {
		t.Run(tt.testCaseName, func(t *testing.T) {
			ctx := t.Context()
			rr := httptest.NewRecorder()
			// Act
			response, err := controller.GetTotalPings(ctx, tt.param)
			// Assert
			assert.Error(t, err)
			// assert.Equal(t, tt.expectedResponse, response)
			// Assertions of the written response
			response.VisitGetTotalPingsResponse(rr)
			assert.Equal(t, tt.expectedResponseCode, rr.Code)
		})
	}
}

// GetTotalPongs
func TestPingPongRestAPIController_GetTotalPongs(t *testing.T) {
	t.Parallel()
	/*
		Arrange
	*/

	// Unit tests for the PingPongRestAPIController
	tests := []struct {
		testCaseName         string
		param                v1_server.GetTotalPongsRequestObject
		expectedResponse     v1_server.GetTotalPongsResponseObject
		expectedResponseCode int
		expectedServiceError error
	}{
		{
			testCaseName: "GET all pingpongs, receive Ping! and Pong!",
			param:        v1_server.GetTotalPongsRequestObject{},
			expectedResponse: v1_server.GetTotalPongs200JSONResponse{
				Body: int(12345),
				Headers: v1_server.GetTotalPongs200ResponseHeaders{
					XRequestId: "00000000000000000000000000000000",
				},
			},
			expectedResponseCode: 200,
			expectedServiceError: nil,
		},
	}

	ctx := t.Context()
	// Mock the service
	mock_svc := mocks.NewMockIPingPongService(t)
	// Define the expected behavior of the mock service
	mock_svc.On(
		"TotalNumberOfPongs",
		ctx,
	).Return(
		types.QuantityMetric{Quantity: 12345},
		nil,
	)
	// Ensure that the mock expectations are met
	defer mock_svc.AssertExpectations(t)

	controller := v1.NewRestAPIController(mock_svc)
	/*
		Act & Assert
	*/
	for _, tt := range tests {
		t.Run(tt.testCaseName, func(t *testing.T) {
			ctx := t.Context()
			rr := httptest.NewRecorder()
			// Act
			response, err := controller.GetTotalPongs(ctx, tt.param)
			// Assert
			assert.Equal(t, tt.expectedServiceError, err)
			assert.Equal(t, tt.expectedResponse, response)
			// Assertions of the written response
			response.VisitGetTotalPongsResponse(rr)
			assert.Equal(t, tt.expectedResponseCode, rr.Code)
		})
	}
}

// GetTotalPongs
func TestPingPongRestAPIController_GetTotalPongs_Failure(t *testing.T) {
	t.Parallel()
	/*
		Arrange
	*/
	errTemp := errors.New("temp error")

	// Unit tests for the PingPongRestAPIController
	tests := []struct {
		testCaseName         string
		param                v1_server.GetTotalPongsRequestObject
		expectedResponse     v1_server.GetTotalPongsResponseObject
		expectedResponseCode int
		expectedServiceError error
	}{
		{
			testCaseName:         "GET all pongs, receive Pong!",
			param:                v1_server.GetTotalPongsRequestObject{},
			expectedResponse:     nil,
			expectedResponseCode: 400,
			expectedServiceError: errTemp,
		},
	}

	ctx := t.Context()
	// Mock the service
	mock_svc := mocks.NewMockIPingPongService(t)
	// Define the expected behavior of the mock service
	mock_svc.On(
		"TotalNumberOfPongs",
		ctx,
	).Return(
		types.QuantityMetric{Quantity: 0}, errTemp,
	)
	// Ensure that the mock expectations are met
	defer mock_svc.AssertExpectations(t)

	controller := v1.NewRestAPIController(mock_svc)
	/*
		Act & Assert
	*/
	for _, tt := range tests {
		t.Run(tt.testCaseName, func(t *testing.T) {
			ctx := t.Context()
			rr := httptest.NewRecorder()
			// Act
			response, err := controller.GetTotalPongs(ctx, tt.param)
			// Assert
			assert.Error(t, err)
			// assert.Equal(t, tt.expectedResponse, response)
			// Assertions of the written response
			response.VisitGetTotalPongsResponse(rr)
			assert.Equal(t, tt.expectedResponseCode, rr.Code)
		})
	}
}
