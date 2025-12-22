package usecase_test

import (
	"context"
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
	"github.com/cooperlutz/go-full/internal/pingpong/app/usecase"
	"github.com/cooperlutz/go-full/internal/pingpong/domain/entity"
	"github.com/cooperlutz/go-full/pkg/types"
	"github.com/cooperlutz/go-full/test/fixtures"
	mocks "github.com/cooperlutz/go-full/test/mocks/pingpong"
)

var (
	testTracerProvider *trace.TracerProvider
	testExporter       *tracetest.InMemoryExporter
	validPingPongID    = uuid.New()
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
	uc := usecase.NewPingPongUseCase(
		// Use the mock repository
		mocks.NewMockIPingPongRepository(t),
	)
	assert.NotNil(t, uc)
}

// Test successful PingPong call
func TestPingPongUseCase_PingPong_Success(t *testing.T) {
	// Arrange
	mockRepo := mocks.NewMockIPingPongRepository(t)
	mockRepo.On("SavePingPong", mock.Anything, mock.AnythingOfType("entity.PingPongEntity")).Return(nil)
	defer mockRepo.AssertExpectations(t)
	useCase := usecase.NewPingPongUseCase(mockRepo)
	cmd := command.PingPongCommand{Message: "ping"}
	// Act
	result, err := useCase.PingPong(context.Background(), cmd)
	// Assert
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

	useCase := usecase.NewPingPongUseCase(mockRepo)
	_, err := useCase.PingPong(context.Background(), cmd)

	assert.Error(t, err)
}

// Test error from repository
func TestPingPongUseCase_PingPong_RepoError(t *testing.T) {
	mockRepo := mocks.NewMockIPingPongRepository(t)
	mockRepo.On("SavePingPong", mock.Anything, mock.AnythingOfType("entity.PingPongEntity")).Return(assert.AnError)
	defer mockRepo.AssertExpectations(t)

	cmd := command.PingPongCommand{Message: "ping"}
	useCase := usecase.NewPingPongUseCase(mockRepo)

	_, err := useCase.PingPong(context.Background(), cmd)

	assert.Error(t, err)
}

func TestPingPongUseCase_PingPong_OtelSpan(t *testing.T) {
	// Clear all spans from the exporter before the test starts
	t.Parallel()
	testExporter.Reset()

	mockRepo := mocks.NewMockIPingPongRepository(t)
	mockRepo.On("SavePingPong", mock.Anything, mock.AnythingOfType("entity.PingPongEntity")).Return(nil)
	defer mockRepo.AssertExpectations(t)

	useCase := usecase.NewPingPongUseCase(mockRepo)
	cmd := command.PingPongCommand{Message: "ping"}

	// Act
	ctx := context.Background()
	_, err := useCase.PingPong(ctx, cmd)

	// Assert
	assert.NoError(t, err)

	// Check that a span was created
	spans := testExporter.GetSpans()
	assert.Len(t, spans, 1)
	assert.Equal(t, "pingpong.usecase.pingpong", spans[0].Name)
}

// STEP 4.2. Implement service. Logic Tests
// here we define our tests for the service. layer logic
func TestPingPongUseCase_FindOneByID_Success(t *testing.T) {
	// Arrange
	mockRepo := mocks.NewMockIPingPongRepository(t)
	mockRepo.On("FindOneByID", mock.Anything, validPingPongID).Return(fixtures.ValidPong, nil)
	defer mockRepo.AssertExpectations(t)

	// Act
	useCase := usecase.NewPingPongUseCase(mockRepo)
	resp, err := useCase.FindOneByID(context.Background(), query.FindOneByID{ID: validPingPongID})

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestPingPongUseCase_FindOneByID_RepoError(t *testing.T) {
	// Arrange
	mockRepo := mocks.NewMockIPingPongRepository(t)
	ctx := context.Background()

	expectedResponse := query.FindOneByIDResponse{}
	mockRepo.On("FindOneByID", mock.Anything, validPingPongID).Return(entity.PingPongEntity{}, assert.AnError)
	defer mockRepo.AssertExpectations(t)

	// Act
	useCase := usecase.NewPingPongUseCase(mockRepo)
	resp, err := useCase.FindOneByID(ctx, query.FindOneByID{ID: validPingPongID})

	// Assert
	assert.Equal(t, expectedResponse, resp)
	assert.Error(t, err)
}

func TestPingPongUseCase_FindAll_Success(t *testing.T) {
	// Arrange
	mockRepo := mocks.NewMockIPingPongRepository(t)
	mockRepo.On("FindAll", mock.Anything).Return(entity.ListOfPingPongs{
		PingPongs: []entity.PingPongEntity{
			fixtures.ValidPong,
			fixtures.ValidPing,
		},
	}, nil)
	defer mockRepo.AssertExpectations(t)
	useCase := usecase.NewPingPongUseCase(mockRepo)
	// Act
	resp, err := useCase.FindAll(context.Background())
	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestPingPongUseCase_FindAll_RepoError(t *testing.T) {
	// Arrange
	mockRepo := mocks.NewMockIPingPongRepository(t)
	mockRepo.On("FindAll", mock.Anything).Return(entity.ListOfPingPongs{}, assert.AnError)
	defer mockRepo.AssertExpectations(t)
	useCase := usecase.NewPingPongUseCase(mockRepo)
	// Act
	_, err := useCase.FindAll(context.Background())
	// Assert
	assert.Error(t, err)
}

func TestPingPongUseCase_FindAllPings_Success(t *testing.T) {
	// Arrange
	mockRepo := mocks.NewMockIPingPongRepository(t)
	mockRepo.On("FindAllPings", mock.Anything).Return(entity.ListOfPingPongs{
		PingPongs: []entity.PingPongEntity{
			fixtures.ValidPing,
		},
	}, nil)
	defer mockRepo.AssertExpectations(t)
	useCase := usecase.NewPingPongUseCase(mockRepo)
	// Act
	resp, err := useCase.FindAllPings(context.Background())
	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestPingPongUseCase_FindAllPings_RepoError(t *testing.T) {
	// Arrange
	mockRepo := mocks.NewMockIPingPongRepository(t)
	mockRepo.On("FindAllPings", mock.Anything).Return(entity.ListOfPingPongs{}, assert.AnError)
	defer mockRepo.AssertExpectations(t)
	useCase := usecase.NewPingPongUseCase(mockRepo)
	// Act
	_, err := useCase.FindAllPings(context.Background())
	// Assert
	assert.Error(t, err)
}

func TestPingPongUseCase_FindAllPongs_Success(t *testing.T) {
	// Arrange
	mockRepo := mocks.NewMockIPingPongRepository(t)
	mockRepo.On("FindAllPongs", mock.Anything).Return(entity.ListOfPingPongs{
		PingPongs: []entity.PingPongEntity{
			fixtures.ValidPong,
		},
	}, nil)
	defer mockRepo.AssertExpectations(t)
	useCase := usecase.NewPingPongUseCase(mockRepo)
	// Act
	resp, err := useCase.FindAllPongs(context.Background())
	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestPingPongUseCase_FindAllPongs_RepoError(t *testing.T) {
	// Arrange
	mockRepo := mocks.NewMockIPingPongRepository(t)
	mockRepo.On("FindAllPongs", mock.Anything).Return(entity.ListOfPingPongs{}, assert.AnError)
	defer mockRepo.AssertExpectations(t)
	useCase := usecase.NewPingPongUseCase(mockRepo)
	// Act
	_, err := useCase.FindAllPongs(context.Background())
	// Assert
	assert.Error(t, err)
}

func TestPingPongUseCase_TotalNumberOfPingPongs_Success(t *testing.T) {
	// Arrange
	mockRepo := mocks.NewMockIPingPongRepository(t)
	mockRepo.On("TotalNumberOfPingPongs", mock.Anything).Return(types.QuantityMetric{Quantity: 10}, nil)
	defer mockRepo.AssertExpectations(t)
	useCase := usecase.NewPingPongUseCase(mockRepo)

	// Act
	count, err := useCase.TotalNumberOfPingPongs(context.Background())

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, types.QuantityMetric{Quantity: 10}, count)
}

func TestPingPongUseCase_TotalNumberOfPingPongs_RepoError(t *testing.T) {
	// Arrange
	mockRepo := mocks.NewMockIPingPongRepository(t)
	mockRepo.On("TotalNumberOfPingPongs", mock.Anything).Return(types.QuantityMetric{Quantity: 0}, assert.AnError)
	defer mockRepo.AssertExpectations(t)
	useCase := usecase.NewPingPongUseCase(mockRepo)
	// Act
	count, err := useCase.TotalNumberOfPingPongs(context.Background())
	// Assert
	assert.Equal(t, types.QuantityMetric{Quantity: 0}, count)
	assert.Error(t, err)
}

func TestPingPongUseCase_TotalNumberOfPings_Success(t *testing.T) {
	// Arrange
	mockRepo := mocks.NewMockIPingPongRepository(t)
	mockRepo.On("TotalNumberOfPings", mock.Anything).Return(types.QuantityMetric{Quantity: 10}, nil)
	defer mockRepo.AssertExpectations(t)
	useCase := usecase.NewPingPongUseCase(mockRepo)
	// Act
	count, err := useCase.TotalNumberOfPings(context.Background())
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, types.QuantityMetric{Quantity: 10}, count)
}

func TestPingPongUseCase_TotalNumberOfPings_RepoError(t *testing.T) {
	// Arrange
	mockRepo := mocks.NewMockIPingPongRepository(t)
	mockRepo.On("TotalNumberOfPings", mock.Anything).Return(types.QuantityMetric{Quantity: 0}, assert.AnError)
	defer mockRepo.AssertExpectations(t)
	useCase := usecase.NewPingPongUseCase(mockRepo)
	// Act
	count, err := useCase.TotalNumberOfPings(context.Background())
	// Assert
	assert.Equal(t, types.QuantityMetric{Quantity: 0}, count)
	assert.Error(t, err)
}

func TestPingPongUseCase_TotalNumberOfPongs_Success(t *testing.T) {
	// Arrange
	mockRepo := mocks.NewMockIPingPongRepository(t)
	mockRepo.On("TotalNumberOfPongs", mock.Anything).Return(types.QuantityMetric{Quantity: 15}, nil)
	defer mockRepo.AssertExpectations(t)
	useCase := usecase.NewPingPongUseCase(mockRepo)
	// Act
	count, err := useCase.TotalNumberOfPongs(context.Background())
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, types.QuantityMetric{Quantity: 15}, count)
}

func TestPingPongUseCase_TotalNumberOfPongs_RepoError(t *testing.T) {
	// Arrange
	mockRepo := mocks.NewMockIPingPongRepository(t)
	mockRepo.On("TotalNumberOfPongs", mock.Anything).Return(types.QuantityMetric{Quantity: 0}, assert.AnError)
	defer mockRepo.AssertExpectations(t)

	useCase := usecase.NewPingPongUseCase(mockRepo)
	// Act
	count, err := useCase.TotalNumberOfPongs(context.Background())
	// Assert
	assert.Equal(t, types.QuantityMetric{Quantity: 0}, count)
	assert.Error(t, err)
}

func TestPingPongUseCase_TotalNumberOfPingPongsPerDay_Success(t *testing.T) {
	// Arrange
	mockRepo := mocks.NewMockIPingPongRepository(t)
	date1, _ := time.Parse("2006-01-02", "2024-06-01")
	date2, _ := time.Parse("2006-01-02", "2024-06-02")
	expected := []types.MeasureCountbyDateTimeMetric{
		{DateTime: date1, Count: 5},
		{DateTime: date2, Count: 7},
	}
	mockRepo.On("TotalNumberOfPingPongsCreatedPerDay", mock.Anything).Return(expected, nil)
	defer mockRepo.AssertExpectations(t)
	useCase := usecase.NewPingPongUseCase(mockRepo)
	// Act
	result, err := useCase.TotalNumberOfPingPongsPerDay(context.Background())
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestPingPongUseCase_TotalNumberOfPingPongsPerDay_RepoError(t *testing.T) {
	// Arrange
	mockRepo := mocks.NewMockIPingPongRepository(t)
	mockRepo.On("TotalNumberOfPingPongsCreatedPerDay", mock.Anything).Return(nil, assert.AnError)
	defer mockRepo.AssertExpectations(t)
	useCase := usecase.NewPingPongUseCase(mockRepo)
	// Act
	result, err := useCase.TotalNumberOfPingPongsPerDay(context.Background())
	// Assert
	assert.Nil(t, result)
	assert.Error(t, err)
}
