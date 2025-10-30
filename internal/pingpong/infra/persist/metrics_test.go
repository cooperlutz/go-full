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
	mocks "github.com/cooperlutz/go-full/test/mocks/pingpong"
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

	returnedCount, err := repo.TotalNumberOfPingPongs(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, expectedCount, returnedCount.Quantity)
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

	returnedCount, err := repo.TotalNumberOfPings(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, expectedCount, returnedCount.Quantity)
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

	returnedCount, err := repo.TotalNumberOfPongs(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, expectedCount, returnedCount.Quantity)
}

func TestAverageNumberOfPingPongsCreatedPerDay_Success(t *testing.T) {
	// Arrange
	mQuerier := mocks.NewMockIQuerierPingPong(t)

	repo := &PingPongPersistPostgresRepository{
		query: mQuerier,
	}
	expectedOutput := types.QuantityMetric{Quantity: 8}

	resp := []persist_postgres.FrequencyDistributionByDayRow{
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
	).Return(resp, nil)
	// Act
	returnedCount, err := repo.AverageNumberOfPingPongsCreatedPerDay(context.Background())
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, returnedCount)
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

	resp := []persist_postgres.CountPerDayRow{
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
	).Return(resp, nil)
	// Act
	returnedCount, err := repo.TotalNumberOfPingPongsCreatedPerDay(context.Background())
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, returnedCount)
}
