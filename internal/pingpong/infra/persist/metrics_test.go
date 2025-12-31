package persist

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	persist_postgres "github.com/cooperlutz/go-full/internal/pingpong/infra/persist/postgres"
	"github.com/cooperlutz/go-full/pkg/types"
	mocks "github.com/cooperlutz/go-full/test/mocks"
)

func TestTotalNumberOfPingPongs_Success(t *testing.T) {
	// Arrange
	mQuerier := mocks.NewMockIQuerierPingPong(t)
	repo := &PingPongPersistPostgresRepository{
		query: mQuerier,
	}
	expectedCount := int64(5)
	mQuerier.On(
		"TotalNumberOfPingPongs",
		mock.Anything,
	).Return(expectedCount, nil)

	// Act
	returnedCount, err := repo.TotalNumberOfPingPongs(context.Background())

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedCount, returnedCount.Quantity)
}

func TestTotalNumberOfPingPongs_Failure(t *testing.T) {
	// Arrange
	mQuerier := mocks.NewMockIQuerierPingPong(t)
	repo := &PingPongPersistPostgresRepository{
		query: mQuerier,
	}
	mQuerier.On(
		"TotalNumberOfPingPongs",
		mock.Anything,
	).Return(int64(0), assert.AnError)

	// Act
	returnedCount, err := repo.TotalNumberOfPingPongs(context.Background())

	// Assert
	assert.Error(t, err)
	assert.Equal(t, int64(0), returnedCount.Quantity)
}

func TestTotalNumberOfPings_Success(t *testing.T) {
	// Arrange
	mQuerier := mocks.NewMockIQuerierPingPong(t)
	repo := &PingPongPersistPostgresRepository{
		query: mQuerier,
	}
	expectedCount := int64(5)
	mQuerier.On(
		"TotalNumberOfPings",
		mock.Anything,
	).Return(expectedCount, nil)

	// Act
	returnedCount, err := repo.TotalNumberOfPings(context.Background())

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedCount, returnedCount.Quantity)
}

func TestTotalNumberOfPings_Failure(t *testing.T) {
	// Arrange
	mQuerier := mocks.NewMockIQuerierPingPong(t)
	repo := &PingPongPersistPostgresRepository{
		query: mQuerier,
	}
	mQuerier.On(
		"TotalNumberOfPings",
		mock.Anything,
	).Return(int64(0), assert.AnError)

	// Act
	returnedCount, err := repo.TotalNumberOfPings(context.Background())

	// Assert
	assert.Error(t, err)
	assert.Equal(t, int64(0), returnedCount.Quantity)
}

func TestTotalNumberOfPongs_Success(t *testing.T) {
	// Arrange
	mQuerier := mocks.NewMockIQuerierPingPong(t)
	repo := &PingPongPersistPostgresRepository{
		query: mQuerier,
	}
	expectedCount := int64(5)
	mQuerier.On(
		"TotalNumberOfPongs",
		mock.Anything,
	).Return(expectedCount, nil)
	// Act
	returnedCount, err := repo.TotalNumberOfPongs(context.Background())
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedCount, returnedCount.Quantity)
}

func TestTotalNumberOfPongs_Failure(t *testing.T) {
	// Arrange
	mQuerier := mocks.NewMockIQuerierPingPong(t)
	repo := &PingPongPersistPostgresRepository{
		query: mQuerier,
	}
	mQuerier.On(
		"TotalNumberOfPongs",
		mock.Anything,
	).Return(int64(0), assert.AnError)

	// Act
	returnedCount, err := repo.TotalNumberOfPongs(context.Background())

	// Assert
	assert.Error(t, err)
	assert.Equal(t, int64(0), returnedCount.Quantity)
}

func TestAverageNumberOfPingPongsCreatedPerDay_Success(t *testing.T) {
	// Arrange
	mQuerier := mocks.NewMockIQuerierPingPong(t)
	repo := &PingPongPersistPostgresRepository{
		query: mQuerier,
	}
	expectedOutput := types.QuantityMetric{Quantity: 8}
	mockResponse := []persist_postgres.FrequencyDistributionByDayRow{
		{
			CreationDate: pgtype.Date{Time: time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC), Valid: true},
			PingOrPong:   pgtype.Text{String: "ping", Valid: true},
			Frequency:    6,
		},
		{
			CreationDate: pgtype.Date{Time: time.Date(2023, 10, 2, 0, 0, 0, 0, time.UTC), Valid: true},
			PingOrPong:   pgtype.Text{String: "pong", Valid: true},
			Frequency:    10,
		},
	}
	expectedCount := int64(16)
	mQuerier.On(
		"TotalNumberOfPingPongs",
		mock.Anything,
	).Return(expectedCount, nil)
	mQuerier.On(
		"FrequencyDistributionByDay",
		mock.Anything,
	).Return(mockResponse, nil)
	// Act
	returnedCount, err := repo.AverageNumberOfPingPongsCreatedPerDay(context.Background())
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, returnedCount)
}

func TestAverageNumberOfPingPongsCreatedPerDay_Failure(t *testing.T) {
	// Arrange
	mQuerier := mocks.NewMockIQuerierPingPong(t)
	repo := &PingPongPersistPostgresRepository{
		query: mQuerier,
	}
	mQuerier.On(
		"FrequencyDistributionByDay",
		mock.Anything,
	).Return([]persist_postgres.FrequencyDistributionByDayRow{}, assert.AnError)

	// Act
	returnedCount, err := repo.AverageNumberOfPingPongsCreatedPerDay(context.Background())

	// Assert
	assert.Error(t, err)
	assert.Equal(t, types.QuantityMetric{Quantity: 0}, returnedCount)
}

func TestAverageNumberOfPingPongsCreatedPerDay_Failure_TotalNumberOfPingPongs(t *testing.T) {
	// Arrange
	mQuerier := mocks.NewMockIQuerierPingPong(t)
	repo := &PingPongPersistPostgresRepository{
		query: mQuerier,
	}
	mQuerier.On(
		"FrequencyDistributionByDay",
		mock.Anything,
	).Return([]persist_postgres.FrequencyDistributionByDayRow{}, nil)
	mQuerier.On(
		"TotalNumberOfPingPongs",
		mock.Anything,
	).Return(int64(0), assert.AnError)
	// Act
	returnedCount, err := repo.AverageNumberOfPingPongsCreatedPerDay(context.Background())

	// Assert
	assert.Error(t, err)
	assert.Equal(t, types.QuantityMetric{Quantity: 0}, returnedCount)
}

func TestTotalNumberOfPingPongsCreatedPerDay_Success(t *testing.T) {
	// Arrange
	mQuerier := mocks.NewMockIQuerierPingPong(t)
	repo := &PingPongPersistPostgresRepository{
		query: mQuerier,
	}
	expectedOutput := []types.MeasureCountbyDateTimeMetric{
		{
			Count:    6,
			DateTime: time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			Count:    10,
			DateTime: time.Date(2023, 10, 2, 0, 0, 0, 0, time.UTC),
		},
	}
	mockResponse := []persist_postgres.CountPerDayRow{
		{
			CreationDate: pgtype.Date{Time: time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC), Valid: true},
			CountCreated: 6,
		},
		{
			CreationDate: pgtype.Date{Time: time.Date(2023, 10, 2, 0, 0, 0, 0, time.UTC), Valid: true},
			CountCreated: 10,
		},
	}
	mQuerier.On(
		"CountPerDay",
		mock.Anything,
	).Return(mockResponse, nil)

	// Act
	returnedCount, err := repo.TotalNumberOfPingPongsCreatedPerDay(context.Background())

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, returnedCount)
}

func TestTotalNumberOfPingPongsCreatedPerDay_Failure(t *testing.T) {
	// Arrange
	mQuerier := mocks.NewMockIQuerierPingPong(t)
	repo := &PingPongPersistPostgresRepository{
		query: mQuerier,
	}

	mQuerier.On(
		"CountPerDay",
		mock.Anything,
	).Return([]persist_postgres.CountPerDayRow{}, assert.AnError)

	// Act
	returnedCount, err := repo.TotalNumberOfPingPongsCreatedPerDay(context.Background())

	// Assert
	assert.Error(t, err)
	assert.Empty(t, returnedCount)
}
