package persist

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pashagolub/pgxmock/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/cooperlutz/go-full/internal/pingpong/domain/entity"
	persist_postgres "github.com/cooperlutz/go-full/internal/pingpong/infra/persist/postgres"
	mocks "github.com/cooperlutz/go-full/test/mocks/pingpong"
)

var (
	timeNow            = time.Now()
	validPingPongID    = uuid.New()
	validPingPongIDTwo = uuid.New()
)

var validPingPongReturn = persist_postgres.Pingpong{
	PingpongID: pgtype.UUID{Bytes: validPingPongID},
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
	mockDB, err := pgxmock.NewPool()
	if err != nil {
		t.Fatal(err)
	}
	defer mockDB.Close()

	mockQuerier := mocks.NewMockIQuerierPingPong(t)
	// mQuerier := mocks.NewMockQuerier(t)
	ctx := context.Background()
	timeNow := time.Now()
	validPingPongID := uuid.New()

	repo := &PingPongPersistPostgresRepository{
		db:    mockDB,
		query: mockQuerier,
	}

	mockDB.ExpectBegin()

	// // mockTx := &MockTx{}
	// // mockTx.On("Commit", ctx).Return(nil)
	// // mockTx.On("Rollback", ctx).Return(nil)

	// // // mockDB := &MockDB{}
	// mockDB.On("Begin", ctx).Return(mockTx, nil)
	// mockQuerier.On("WithTx", mock.Anything).Return(mQuerier)
	mockQuerier.On("CreatePingPong", mock.Anything, mock.Anything).Return(validPingPongReturn, nil)
	mockQuerier.On("WithTx", mock.Anything).Return(mockQuerier)
	mockDB.ExpectCommit()

	p := &entity.PingPongEntity{
		Message: "ping",
		PingPongMetadata: &entity.PingPongMetadata{
			CreatedAt:  timeNow,
			UpdatedAt:  timeNow,
			Deleted:    false,
			DeletedAt:  nil,
			PingPongID: validPingPongID,
		},
	}

	err = repo.SavePingPong(ctx, p)
	assert.NoError(t, err)
}
