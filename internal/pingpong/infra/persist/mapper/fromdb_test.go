package mapper

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/internal/pingpong/domain/entity"
	postgresql "github.com/cooperlutz/go-full/internal/pingpong/infra/persist/postgres"
	"github.com/cooperlutz/go-full/test/fixtures"
)

var (
	sampleTime = time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)
	sampleUUID = uuid.New()
)

func TestMapFromDB(t *testing.T) {
	t.Parallel()

	// Arrange
	tests := []struct {
		name string
		db   postgresql.Pingpong
		want entity.PingPongEntity
	}{
		{
			name: "Valid PingPongEntity from DB",
			db: postgresql.Pingpong{
				PingpongID: pgtype.UUID{Bytes: fixtures.ValidPing.GetIdUUID(), Valid: true},
				PingOrPong: pgtype.Text{String: "ping", Valid: true},
				CreatedAt:  pgtype.Timestamptz{Time: fixtures.ValidPing.GetCreatedAtTime(), Valid: true},
				UpdatedAt:  pgtype.Timestamptz{Time: fixtures.ValidPing.GetUpdatedAtTime(), Valid: true},
				DeletedAt:  pgtype.Timestamptz{Time: time.Time{}, Valid: false},
				Deleted:    false,
			},
			want: fixtures.ValidPing,
		},
		{
			name: "Invalid PingPongEntity from DB",
			db: postgresql.Pingpong{
				PingpongID: pgtype.UUID{Bytes: fixtures.InvalidPingPong.GetIdUUID(), Valid: true},
				PingOrPong: pgtype.Text{String: "ring", Valid: true},
				CreatedAt:  pgtype.Timestamptz{Time: fixtures.InvalidPingPong.GetCreatedAtTime(), Valid: true},
				UpdatedAt:  pgtype.Timestamptz{Time: fixtures.InvalidPingPong.GetUpdatedAtTime(), Valid: true},
				DeletedAt:  pgtype.Timestamptz{Time: time.Time{}, Valid: false},
				Deleted:    false,
			},
			want: fixtures.InvalidPingPong,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			got := MapFromDB(tt.db)
			// Assert
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMapFromDBPingPongs_EmptyInput(t *testing.T) {
	// Arrange
	// Act
	result := MapFromDBPingPongs([]postgresql.Pingpong{})
	// Assert
	assert.NotNil(t, result.PingPongs)
	assert.Equal(t, 0, len(result.PingPongs))
}

func TestMapFromDBPingPongs_SingleItem(t *testing.T) {
	// Arrange
	input := []postgresql.Pingpong{
		{PingOrPong: pgtype.Text{String: "Ping"}},
	}
	expected := "Ping"

	// Act
	result := MapFromDBPingPongs(input)

	// Assert
	assert.NotNil(t, result.PingPongs)
	assert.Equal(t, 1, len(result.PingPongs))
	assert.Equal(t, expected, result.PingPongs[0].GetMessage())
}

func TestMapFromDBPingPongs_MultipleItems(t *testing.T) {
	// Arrange
	input := []postgresql.Pingpong{
		{PingOrPong: pgtype.Text{String: "Ping", Valid: true}},
		{PingOrPong: pgtype.Text{String: "Pong", Valid: true}},
	}
	expected := []string{"Ping", "Pong"}

	// Act
	result := MapFromDBPingPongs(input)

	// Assert
	assert.NotNil(t, result.PingPongs)
	assert.Equal(t, len(expected), len(result.PingPongs))
	assert.Equal(t, expected[0], result.PingPongs[0].GetMessage())
	assert.Equal(t, expected[1], result.PingPongs[1].GetMessage())
}

func TestMapFromDBPingPongRaw_DeletedAtSet(t *testing.T) {
	t.Parallel()

	// Arrange
	id := uuid.New()
	now := time.Now()
	deletedAt := now.Add(1 * time.Hour)
	p := postgresql.Pingpong{
		PingpongID: pgtype.UUID{Bytes: id, Valid: true},
		PingOrPong: pgtype.Text{String: "pong", Valid: true},
		CreatedAt:  pgtype.Timestamptz{Time: now, Valid: true},
		UpdatedAt:  pgtype.Timestamptz{Time: deletedAt, Valid: true},
		DeletedAt:  pgtype.Timestamptz{Time: deletedAt, Valid: true},
		Deleted:    true,
	}

	// Act
	result := MapFromDB(p)

	// Assert
	assert.NotNil(t, result)
	assert.Equal(t, "pong", result.GetMessage())
	assert.Equal(t, id, result.GetIdUUID())
	assert.Equal(t, now, result.GetCreatedAtTime())
	assert.Equal(t, deletedAt, result.GetUpdatedAtTime())
	assert.Equal(t, &deletedAt, result.GetDeletedAtTime())
	assert.True(t, result.IsDeleted())
}
