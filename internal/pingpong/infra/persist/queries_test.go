package persist

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/cooperlutz/go-full/internal/pingpong/domain/entity"
	persist_postgres "github.com/cooperlutz/go-full/internal/pingpong/infra/persist/postgres"
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
			PingpongID: pgtype.UUID{Bytes: validPingPongID, Valid: true},
			PingOrPong: pgtype.Text{String: "ping", Valid: true},
			CreatedAt:  pgtype.Timestamptz{Time: timeNow, Valid: true},
			UpdatedAt:  pgtype.Timestamptz{Time: timeNow, Valid: true},
			DeletedAt:  pgtype.Timestamptz{},
			Deleted:    false,
		},
		{
			PingpongID: pgtype.UUID{Bytes: validPingPongIDTwo, Valid: true},
			PingOrPong: pgtype.Text{String: "pong", Valid: true},
			CreatedAt:  pgtype.Timestamptz{Time: timeNow, Valid: true},
			UpdatedAt:  pgtype.Timestamptz{Time: timeNow, Valid: true},
			DeletedAt:  pgtype.Timestamptz{},
			Deleted:    false,
		},
	}
	expectedOutput := entity.ListOfPingPongs{
		PingPongs: []entity.PingPongEntity{
			{
				Message: "ping",
				PingPongMetadata: &entity.PingPongMetadata{
					PingPongID: validPingPongID,
					CreatedAt:  timeNow,
					UpdatedAt:  timeNow,
					DeletedAt:  nil,
					Deleted:    false,
				},
			},
			{
				Message: "pong",
				PingPongMetadata: &entity.PingPongMetadata{
					PingPongID: validPingPongIDTwo,
					CreatedAt:  timeNow,
					UpdatedAt:  timeNow,
					DeletedAt:  nil,
					Deleted:    false,
				},
			},
		},
	}

	mQuerier.On(
		"FindAll",
		mock.Anything,
	).Return(mockResponse, nil)

	returnedCount, err := repo.FindAll(context.Background())
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
			PingpongID: pgtype.UUID{Bytes: validPingPongID, Valid: true},
			PingOrPong: pgtype.Text{String: "ping", Valid: true},
			CreatedAt:  pgtype.Timestamptz{Time: timeNow, Valid: true},
			UpdatedAt:  pgtype.Timestamptz{Time: timeNow, Valid: true},
			DeletedAt:  pgtype.Timestamptz{},
			Deleted:    false,
		},
		{
			PingpongID: pgtype.UUID{Bytes: validPingPongIDTwo, Valid: true},
			PingOrPong: pgtype.Text{String: "ping", Valid: true},
			CreatedAt:  pgtype.Timestamptz{Time: timeNow, Valid: true},
			UpdatedAt:  pgtype.Timestamptz{Time: timeNow, Valid: true},
			DeletedAt:  pgtype.Timestamptz{},
			Deleted:    false,
		},
	}
	expectedOutput := entity.ListOfPingPongs{
		PingPongs: []entity.PingPongEntity{
			{
				Message: "ping",
				PingPongMetadata: &entity.PingPongMetadata{
					PingPongID: validPingPongID,
					CreatedAt:  timeNow,
					UpdatedAt:  timeNow,
					DeletedAt:  nil,
					Deleted:    false,
				},
			},
			{
				Message: "ping",
				PingPongMetadata: &entity.PingPongMetadata{
					PingPongID: validPingPongIDTwo,
					CreatedAt:  timeNow,
					UpdatedAt:  timeNow,
					DeletedAt:  nil,
					Deleted:    false,
				},
			},
		},
	}

	mQuerier.On(
		"FindAllPing",
		mock.Anything,
	).Return(mockResponse, nil)

	returnedCount, err := repo.FindAllPings(context.Background())
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
			PingpongID: pgtype.UUID{Bytes: validPingPongID, Valid: true},
			PingOrPong: pgtype.Text{String: "pong", Valid: true},
			CreatedAt:  pgtype.Timestamptz{Time: timeNow, Valid: true},
			UpdatedAt:  pgtype.Timestamptz{Time: timeNow, Valid: true},
			DeletedAt:  pgtype.Timestamptz{},
			Deleted:    false,
		},
		{
			PingpongID: pgtype.UUID{Bytes: validPingPongIDTwo, Valid: true},
			PingOrPong: pgtype.Text{String: "pong", Valid: true},
			CreatedAt:  pgtype.Timestamptz{Time: timeNow, Valid: true},
			UpdatedAt:  pgtype.Timestamptz{Time: timeNow, Valid: true},
			DeletedAt:  pgtype.Timestamptz{},
			Deleted:    false,
		},
	}
	expectedOutput := entity.ListOfPingPongs{
		PingPongs: []entity.PingPongEntity{
			{
				Message: "pong",
				PingPongMetadata: &entity.PingPongMetadata{
					PingPongID: validPingPongID,
					CreatedAt:  timeNow,
					UpdatedAt:  timeNow,
					DeletedAt:  nil,
					Deleted:    false,
				},
			},
			{
				Message: "pong",
				PingPongMetadata: &entity.PingPongMetadata{
					PingPongID: validPingPongIDTwo,
					CreatedAt:  timeNow,
					UpdatedAt:  timeNow,
					DeletedAt:  nil,
					Deleted:    false,
				},
			},
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
