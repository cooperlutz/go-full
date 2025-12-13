package persist

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/cooperlutz/go-full/internal/pingpong/domain/entity"
	persist_postgres "github.com/cooperlutz/go-full/internal/pingpong/infra/persist/postgres"
	"github.com/cooperlutz/go-full/test/fixtures"
	mocks "github.com/cooperlutz/go-full/test/mocks/pingpong"
)

func TestFindAll_Success(t *testing.T) {
	// Arrange
	mQuerier := mocks.NewMockIQuerierPingPong(t)
	repo := &PingPongPersistPostgresRepository{
		query: mQuerier,
	}
	mockResponse := []persist_postgres.Pingpong{
		{
			PingpongID: pgtype.UUID{Bytes: fixtures.ValidPing.GetIdUUID(), Valid: true},
			PingOrPong: pgtype.Text{String: "ping", Valid: true},
			CreatedAt:  pgtype.Timestamptz{Time: fixtures.ValidPing.GetCreatedAtTime(), Valid: true},
			UpdatedAt:  pgtype.Timestamptz{Time: fixtures.ValidPing.GetUpdatedAtTime(), Valid: true},
			DeletedAt:  pgtype.Timestamptz{},
			Deleted:    false,
		},
		{
			PingpongID: pgtype.UUID{Bytes: fixtures.ValidPong.GetIdUUID(), Valid: true},
			PingOrPong: pgtype.Text{String: "pong", Valid: true},
			CreatedAt:  pgtype.Timestamptz{Time: fixtures.ValidPong.GetCreatedAtTime(), Valid: true},
			UpdatedAt:  pgtype.Timestamptz{Time: fixtures.ValidPong.GetUpdatedAtTime(), Valid: true},
			DeletedAt:  pgtype.Timestamptz{},
			Deleted:    false,
		},
	}
	expectedOutput := entity.ListOfPingPongs{
		PingPongs: []entity.PingPongEntity{
			fixtures.ValidPing,
			fixtures.ValidPong,
		},
	}
	mQuerier.On(
		"FindAll",
		mock.Anything,
	).Return(mockResponse, nil)

	// Act
	returnedCount, err := repo.FindAll(context.Background())

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, returnedCount)
}

func TestFindAllPings_Success(t *testing.T) {
	// Arrange
	mQuerier := mocks.NewMockIQuerierPingPong(t)
	repo := &PingPongPersistPostgresRepository{
		query: mQuerier,
	}
	mockResponse := []persist_postgres.Pingpong{
		{
			PingpongID: pgtype.UUID{Bytes: fixtures.ValidPing.GetIdUUID(), Valid: true},
			PingOrPong: pgtype.Text{String: "ping", Valid: true},
			CreatedAt:  pgtype.Timestamptz{Time: fixtures.ValidPing.GetCreatedAtTime(), Valid: true},
			UpdatedAt:  pgtype.Timestamptz{Time: fixtures.ValidPing.GetUpdatedAtTime(), Valid: true},
			DeletedAt:  pgtype.Timestamptz{},
			Deleted:    false,
		},
		{
			PingpongID: pgtype.UUID{Bytes: fixtures.ValidPing.GetIdUUID(), Valid: true},
			PingOrPong: pgtype.Text{String: "ping", Valid: true},
			CreatedAt:  pgtype.Timestamptz{Time: fixtures.ValidPing.GetCreatedAtTime(), Valid: true},
			UpdatedAt:  pgtype.Timestamptz{Time: fixtures.ValidPing.GetUpdatedAtTime(), Valid: true},
			DeletedAt:  pgtype.Timestamptz{},
			Deleted:    false,
		},
	}
	expectedOutput := entity.ListOfPingPongs{
		PingPongs: []entity.PingPongEntity{
			fixtures.ValidPing,
			fixtures.ValidPing,
		},
	}
	mQuerier.On(
		"FindAllPing",
		mock.Anything,
	).Return(mockResponse, nil)

	// Act
	returnedCount, err := repo.FindAllPings(context.Background())

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, returnedCount)
}

func TestFindAllPongs_Success(t *testing.T) {
	// Arrange
	mQuerier := mocks.NewMockIQuerierPingPong(t)
	repo := &PingPongPersistPostgresRepository{
		query: mQuerier,
	}
	mockResponse := []persist_postgres.Pingpong{
		{
			PingpongID: pgtype.UUID{Bytes: fixtures.ValidPong.GetIdUUID(), Valid: true},
			PingOrPong: pgtype.Text{String: "pong", Valid: true},
			CreatedAt:  pgtype.Timestamptz{Time: fixtures.ValidPong.GetCreatedAtTime(), Valid: true},
			UpdatedAt:  pgtype.Timestamptz{Time: fixtures.ValidPong.GetUpdatedAtTime(), Valid: true},
			DeletedAt:  pgtype.Timestamptz{},
			Deleted:    false,
		},
		{
			PingpongID: pgtype.UUID{Bytes: fixtures.ValidPong.GetIdUUID(), Valid: true},
			PingOrPong: pgtype.Text{String: "pong", Valid: true},
			CreatedAt:  pgtype.Timestamptz{Time: fixtures.ValidPong.GetCreatedAtTime(), Valid: true},
			UpdatedAt:  pgtype.Timestamptz{Time: fixtures.ValidPong.GetUpdatedAtTime(), Valid: true},
			DeletedAt:  pgtype.Timestamptz{},
			Deleted:    false,
		},
	}
	expectedOutput := entity.ListOfPingPongs{
		PingPongs: []entity.PingPongEntity{
			fixtures.ValidPong,
			fixtures.ValidPong,
		},
	}
	mQuerier.On(
		"FindAllPong",
		mock.Anything,
	).Return(mockResponse, nil)

	returnedCount, err := repo.FindAllPongs(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, returnedCount)
}

/*
STEP 2.3. Implement Repository Logic Tests

here we implement the relevant tests for the repository layer
*/
func TestFindOneByID_Success(t *testing.T) {
	// Arrange
	mQuerier := mocks.NewMockIQuerierPingPong(t)

	repo := &PingPongPersistPostgresRepository{
		query: mQuerier,
	}
	mockResponse := persist_postgres.Pingpong{
		PingpongID: pgtype.UUID{Bytes: fixtures.ValidPing.GetIdUUID(), Valid: true},
		PingOrPong: pgtype.Text{String: "ping", Valid: true},
		CreatedAt:  pgtype.Timestamptz{Time: fixtures.ValidPing.GetCreatedAtTime(), Valid: true},
		UpdatedAt:  pgtype.Timestamptz{Time: fixtures.ValidPing.GetUpdatedAtTime(), Valid: true},
		DeletedAt:  pgtype.Timestamptz{},
		Deleted:    false,
	}
	expectedOutput := fixtures.ValidPing
	mQuerier.On(
		"FindOneByID",
		mock.Anything,
		mock.Anything,
	).Return(mockResponse, nil)

	// Act
	returnedCount, err := repo.FindOneByID(context.Background(), fixtures.ValidPing.GetIdUUID())

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, returnedCount)
}

func TestFindOneByID_Failure(t *testing.T) {
	// Arrange
	mQuerier := mocks.NewMockIQuerierPingPong(t)
	ctx := context.Background()
	repo := &PingPongPersistPostgresRepository{
		query: mQuerier,
	}
	mockResponse := persist_postgres.Pingpong{}
	tempError := errors.New("this error happened")
	expectedOutput := entity.PingPongEntity{}
	mQuerier.On(
		"FindOneByID",
		mock.Anything,
		persist_postgres.FindOneByIDParams{PingpongID: pgtype.UUID{Bytes: fixtures.ValidPing.GetIdUUID(), Valid: true}},
	).Return(mockResponse, tempError)

	// Act
	returnedEntity, err := repo.FindOneByID(ctx, fixtures.ValidPing.GetIdUUID())

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedOutput, returnedEntity)
}
