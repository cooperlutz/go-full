package persist

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pashagolub/pgxmock/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	persist_postgres "github.com/cooperlutz/go-full/internal/pingpong/infra/persist/postgres"
	"github.com/cooperlutz/go-full/test/fixtures"
	mocks "github.com/cooperlutz/go-full/test/mocks/pingpong"
)

var timeNow = time.Now()

var validPingPongReturn = persist_postgres.Pingpong{
	PingpongID: pgtype.UUID{Bytes: fixtures.InvalidPingPong.GetIdUUID(), Valid: true},
	PingOrPong: pgtype.Text{String: "ping", Valid: true},
	CreatedAt: pgtype.Timestamptz{
		Time: timeNow, Valid: true,
	},
	UpdatedAt: pgtype.Timestamptz{
		Time: timeNow, Valid: true,
	},
	DeletedAt: pgtype.Timestamptz{},
	Deleted:   false,
}

func TestSavePingPong_Success(t *testing.T) {
	// Arrange
	mockDB, err := pgxmock.NewPool()
	if err != nil {
		t.Fatal(err)
	}
	defer mockDB.Close()
	mockQuerier := mocks.NewMockIQuerierPingPong(t)
	ctx := context.Background()
	repo := &PingPongPersistPostgresRepository{
		db:    mockDB,
		query: mockQuerier,
	}
	mockDB.ExpectBegin()
	mockQuerier.On("CreatePingPong", mock.Anything, mock.Anything).Return(validPingPongReturn, nil)
	mockQuerier.On("WithTx", mock.Anything).Return(mockQuerier)
	mockDB.ExpectCommit()
	p := fixtures.ValidPing

	// Act
	err = repo.SavePingPong(ctx, p)

	// Assert
	assert.NoError(t, err)
}
