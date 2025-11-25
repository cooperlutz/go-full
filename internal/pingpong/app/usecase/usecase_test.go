package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"

	"github.com/cooperlutz/go-full/internal/pingpong/app/command"
	"github.com/cooperlutz/go-full/internal/pingpong/app/common"
	"github.com/cooperlutz/go-full/internal/pingpong/app/query"
	service "github.com/cooperlutz/go-full/internal/pingpong/app/usecase"
	"github.com/cooperlutz/go-full/internal/pingpong/domain/entity"
	"github.com/cooperlutz/go-full/pkg/types"
	mocks "github.com/cooperlutz/go-full/test/mocks/pingpong"
)

var (
	testTracerProvider *trace.TracerProvider
	testExporter       *tracetest.InMemoryExporter
	timeNow            = time.Now()
	validPingPongID    = uuid.New()
	validPingPongIDTwo = uuid.New()
)

func TestMain(m *testing.M) {
	// Create an in-memory span exporter
	testExporter = tracetest.NewInMemoryExporter()
	testTracerProvider = trace.NewTracerProvider(
		trace.WithSyncer(testExporter),
	)
	defer testTracerProvider.Shutdown(context.Background())

	// Register the test tracer provider globally
	otel.SetTracerProvider(testTracerProvider)

	// Run the tests
	m.Run()
}

func TestNewPingPongUseCaseWithMockRepo(t *testing.T) {
	service := service.NewPingPongUseCase(
		// Use the mock repository
		mocks.NewMockIPingPongRepository(t),
	)
	assert.NotNil(t, service)
}

// Test successful PingPong call
func TestPingPongUseCase_PingPong_Success(t *testing.T) {
	mockRepo := mocks.NewMockIPingPongRepository(t)
	mockRepo.On("SavePingPong", mock.Anything, mock.AnythingOfType("entity.PingPongEntity")).Return(nil)
	defer mockRepo.AssertExpectations(t)

	svc := service.NewPingPongUseCase(mockRepo)
	cmd := command.PingPongCommand{Message: "ping"}

	result, err := svc.PingPong(context.Background(), cmd)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, command.PingPongCommandResult{
		PingPongResult: &common.PingPongResult{Message: "Pong!"},
	}, result)
}

// Test error from mapper
func TestPingPongUseCase_PingPong_MapperError(t *testing.T) {
	mockRepo := mocks.NewMockIPingPongRepository(t)
	defer mockRepo.AssertExpectations(t)

	// Simulate invalid command (assuming mapper returns error for invalid message)
	cmd := command.PingPongCommand{Message: "invalid"}

	svc := service.NewPingPongUseCase(mockRepo)
	_, err := svc.PingPong(context.Background(), cmd)

	assert.Error(t, err)
}

// Test error from repository
func TestPingPongUseCase_PingPong_RepoError(t *testing.T) {
	mockRepo := mocks.NewMockIPingPongRepository(t)
	mockRepo.On("SavePingPong", mock.Anything, mock.AnythingOfType("entity.PingPongEntity")).Return(errors.New("repo error"))
	defer mockRepo.AssertExpectations(t)

	cmd := command.PingPongCommand{Message: "ping"}
	svc := service.NewPingPongUseCase(mockRepo)

	_, err := svc.PingPong(context.Background(), cmd)

	assert.EqualError(t, err, "repo error")
}

func TestPingPongUseCase_PingPong_OtelSpan(t *testing.T) {
	// Clear all spans from the exporter before the test starts
	t.Parallel()
	testExporter.Reset()

	mockRepo := mocks.NewMockIPingPongRepository(t)
	mockRepo.On("SavePingPong", mock.Anything, mock.AnythingOfType("entity.PingPongEntity")).Return(nil)
	defer mockRepo.AssertExpectations(t)

	svc := service.NewPingPongUseCase(mockRepo)
	cmd := command.PingPongCommand{Message: "ping"}

	// Act
	ctx := context.Background()
	_, err := svc.PingPong(ctx, cmd)

	// Assert
	assert.NoError(t, err)

	// Check that a span was created
	spans := testExporter.GetSpans()
	assert.Len(t, spans, 1)
	assert.Equal(t, "service.pingpong", spans[0].Name)
}

// STEP 4.2. Implement Service Logic Tests
// here we define our tests for the service layer logic
func TestPingPongUseCase_FindOneByID_Success(t *testing.T) {
	// Arrange
	mockRepo := mocks.NewMockIPingPongRepository(t)
	mockRepo.On("FindOneByID", mock.Anything, validPingPongID).Return(entity.PingPongEntity{
		Message: "pong",
		PingPongMetadata: &entity.PingPongMetadata{
			PingPongID: validPingPongID,
			CreatedAt:  timeNow,
			UpdatedAt:  timeNow,
			DeletedAt:  nil,
			Deleted:    false,
		},
	}, nil)
	defer mockRepo.AssertExpectations(t)

	// Act
	svc := service.NewPingPongUseCase(mockRepo)
	resp, err := svc.FindOneByID(context.Background(), query.FindOneByID{ID: validPingPongID})

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestPingPongUseCase_FindOneByID_RepoError(t *testing.T) {
	// Arrange
	mockRepo := mocks.NewMockIPingPongRepository(t)
	ctx := context.Background()
	tempError := errors.New("findone error")
	expectedResponse := query.FindOneByIDResponse{}
	mockRepo.On("FindOneByID", mock.Anything, validPingPongID).Return(entity.PingPongEntity{}, tempError)
	defer mockRepo.AssertExpectations(t)

	// Act
	svc := service.NewPingPongUseCase(mockRepo)
	resp, err := svc.FindOneByID(ctx, query.FindOneByID{ID: validPingPongID})

	// Assert
	assert.Equal(t, expectedResponse, resp)
	assert.Error(t, err)
	assert.EqualError(t, err, "findone error")
}

func TestPingPongUseCase_FindAll_Success(t *testing.T) {
	mockRepo := mocks.NewMockIPingPongRepository(t)
	mockRepo.On("FindAll", mock.Anything).Return(entity.ListOfPingPongs{
		PingPongs: []entity.PingPongEntity{
			{Message: "pong"},
			{Message: "ping"},
		},
	}, nil)
	defer mockRepo.AssertExpectations(t)

	svc := service.NewPingPongUseCase(mockRepo)
	resp, err := svc.FindAll(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestPingPongUseCase_FindAll_RepoError(t *testing.T) {
	mockRepo := mocks.NewMockIPingPongRepository(t)
	mockRepo.On("FindAll", mock.Anything).Return(entity.ListOfPingPongs{}, errors.New("findall error"))
	defer mockRepo.AssertExpectations(t)

	svc := service.NewPingPongUseCase(mockRepo)
	_, err := svc.FindAll(context.Background())
	// assert.Nil(t, resp)
	assert.EqualError(t, err, "findall error")
}

func TestPingPongUseCase_FindAllPings_Success(t *testing.T) {
	mockRepo := mocks.NewMockIPingPongRepository(t)
	mockRepo.On("FindAllPings", mock.Anything).Return(entity.ListOfPingPongs{
		PingPongs: []entity.PingPongEntity{
			{Message: "ping"},
		},
	}, nil)
	defer mockRepo.AssertExpectations(t)

	svc := service.NewPingPongUseCase(mockRepo)
	resp, err := svc.FindAllPings(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	// assert.Len(t, resp.Results, 1)entity
}

func TestPingPongUseCase_FindAllPings_RepoError(t *testing.T) {
	mockRepo := mocks.NewMockIPingPongRepository(t)
	mockRepo.On("FindAllPings", mock.Anything).Return(entity.ListOfPingPongs{}, errors.New("findallpings error"))
	defer mockRepo.AssertExpectations(t)

	svc := service.NewPingPongUseCase(mockRepo)
	_, err := svc.FindAllPings(context.Background())

	assert.EqualError(t, err, "findallpings error")
}

func TestPingPongUseCase_FindAllPongs_Success(t *testing.T) {
	mockRepo := mocks.NewMockIPingPongRepository(t)
	mockRepo.On("FindAllPongs", mock.Anything).Return(entity.ListOfPingPongs{
		PingPongs: []entity.PingPongEntity{
			{Message: "pong"},
		},
	}, nil)
	defer mockRepo.AssertExpectations(t)

	svc := service.NewPingPongUseCase(mockRepo)
	resp, err := svc.FindAllPongs(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	// assert.Len(t, resp.Results, 1)
}

func TestPingPongUseCase_FindAllPongs_RepoError(t *testing.T) {
	mockRepo := mocks.NewMockIPingPongRepository(t)
	mockRepo.On("FindAllPongs", mock.Anything).Return(entity.ListOfPingPongs{}, errors.New("findallpongs error"))
	defer mockRepo.AssertExpectations(t)

	svc := service.NewPingPongUseCase(mockRepo)
	_, err := svc.FindAllPongs(context.Background())

	assert.EqualError(t, err, "findallpongs error")
}

func TestPingPongUseCase_TotalNumberOfPingPongs_Success(t *testing.T) {
	mockRepo := mocks.NewMockIPingPongRepository(t)
	mockRepo.On("TotalNumberOfPingPongs", mock.Anything).Return(types.QuantityMetric{Quantity: 10}, nil)
	defer mockRepo.AssertExpectations(t)

	svc := service.NewPingPongUseCase(mockRepo)
	count, err := svc.TotalNumberOfPingPongs(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, types.QuantityMetric{Quantity: 10}, count)
}

func TestPingPongUseCase_TotalNumberOfPingPongs_RepoError(t *testing.T) {
	mockRepo := mocks.NewMockIPingPongRepository(t)
	mockRepo.On("TotalNumberOfPingPongs", mock.Anything).Return(types.QuantityMetric{Quantity: 0}, errors.New("count error"))
	defer mockRepo.AssertExpectations(t)

	svc := service.NewPingPongUseCase(mockRepo)
	count, err := svc.TotalNumberOfPingPongs(context.Background())
	assert.Equal(t, types.QuantityMetric{Quantity: 0}, count)
	assert.EqualError(t, err, "count error")
}

func TestPingPongUseCase_TotalNumberOfPings_Success(t *testing.T) {
	mockRepo := mocks.NewMockIPingPongRepository(t)
	mockRepo.On("TotalNumberOfPings", mock.Anything).Return(types.QuantityMetric{Quantity: 10}, nil)
	defer mockRepo.AssertExpectations(t)

	svc := service.NewPingPongUseCase(mockRepo)
	count, err := svc.TotalNumberOfPings(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, types.QuantityMetric{Quantity: 10}, count)
}

func TestPingPongUseCase_TotalNumberOfPings_RepoError(t *testing.T) {
	mockRepo := mocks.NewMockIPingPongRepository(t)
	mockRepo.On("TotalNumberOfPings", mock.Anything).Return(types.QuantityMetric{Quantity: 0}, errors.New("pings error"))
	defer mockRepo.AssertExpectations(t)

	svc := service.NewPingPongUseCase(mockRepo)
	count, err := svc.TotalNumberOfPings(context.Background())
	assert.Equal(t, types.QuantityMetric{Quantity: 0}, count)
	assert.EqualError(t, err, "pings error")
}

func TestPingPongUseCase_TotalNumberOfPongs_Success(t *testing.T) {
	mockRepo := mocks.NewMockIPingPongRepository(t)
	mockRepo.On("TotalNumberOfPongs", mock.Anything).Return(types.QuantityMetric{Quantity: 15}, nil)
	defer mockRepo.AssertExpectations(t)

	svc := service.NewPingPongUseCase(mockRepo)
	count, err := svc.TotalNumberOfPongs(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, types.QuantityMetric{Quantity: 15}, count)
}

func TestPingPongUseCase_TotalNumberOfPongs_RepoError(t *testing.T) {
	mockRepo := mocks.NewMockIPingPongRepository(t)
	mockRepo.On("TotalNumberOfPongs", mock.Anything).Return(types.QuantityMetric{Quantity: 0}, errors.New("pongs error"))
	defer mockRepo.AssertExpectations(t)

	svc := service.NewPingPongUseCase(mockRepo)
	count, err := svc.TotalNumberOfPongs(context.Background())
	assert.Equal(t, types.QuantityMetric{Quantity: 0}, count)
	assert.EqualError(t, err, "pongs error")
}

func TestPingPongUseCase_TotalNumberOfPingPongsPerDay_Success(t *testing.T) {
	mockRepo := mocks.NewMockIPingPongRepository(t)
	date1, _ := time.Parse("2006-01-02", "2024-06-01")
	date2, _ := time.Parse("2006-01-02", "2024-06-02")
	expected := []types.MeasureCountbyDateTimeMetric{
		{DateTime: date1, Count: 5},
		{DateTime: date2, Count: 7},
	}
	mockRepo.On("TotalNumberOfPingPongsCreatedPerDay", mock.Anything).Return(expected, nil)
	defer mockRepo.AssertExpectations(t)

	svc := service.NewPingPongUseCase(mockRepo)
	result, err := svc.TotalNumberOfPingPongsPerDay(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestPingPongUseCase_TotalNumberOfPingPongsPerDay_RepoError(t *testing.T) {
	mockRepo := mocks.NewMockIPingPongRepository(t)
	mockRepo.On("TotalNumberOfPingPongsCreatedPerDay", mock.Anything).Return(nil, errors.New("perday error"))
	defer mockRepo.AssertExpectations(t)

	svc := service.NewPingPongUseCase(mockRepo)
	result, err := svc.TotalNumberOfPingPongsPerDay(context.Background())
	assert.Nil(t, result)
	assert.EqualError(t, err, "perday error")
}
