package mapper

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/internal/pingpong/domain/entity"
	postgresql "github.com/cooperlutz/go-full/internal/pingpong/infra/persist/postgres"
)

var (
	sampleTime = time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
	sampleUUID = uuid.New()
)

func TestTranslateFromDB(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		db   postgresql.Pingpong
		want entity.PingPongEntity
	}{
		{
			name: "Valid PingPongEntity from DB",
			db: postgresql.Pingpong{
				PingOrPong: pgtype.Text{String: "Ping!", Valid: true},
			},
			want: entity.PingPongEntity{Message: "Ping!"},
		},
		{
			name: "Invalid PingPongEntity from DB",
			db: postgresql.Pingpong{
				PingOrPong: pgtype.Text{String: "", Valid: false},
			},
			want: entity.PingPongEntity{Message: ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TranslateFromDB(tt.db)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTranslatePingPongsFromDB_EmptyInput(t *testing.T) {
	result := TranslatePingPongsFromDB([]postgresql.Pingpong{})
	if result == nil {
		t.Fatal("Expected non-nil result")
	}
	if len(result.PingPongs) != 0 {
		t.Errorf("Expected empty PingPongs slice, got %d", len(result.PingPongs))
	}
}

func TestTranslatePingPongsFromDB_SingleItem(t *testing.T) {
	input := []postgresql.Pingpong{
		{PingOrPong: pgtype.Text{String: "Ping"}},
	}
	expected := "Ping"

	result := TranslatePingPongsFromDB(input)
	if result == nil {
		t.Fatal("Expected non-nil result")
	}
	if len(result.PingPongs) != 1 {
		t.Fatalf("Expected 1 PingPongEntity, got %d", len(result.PingPongs))
	}
	if result.PingPongs[0].Message != expected {
		t.Errorf("Expected message %q, got %q", expected, result.PingPongs[0].Message)
	}
}

func TestTranslatePingPongsFromDB_MultipleItems(t *testing.T) {
	input := []postgresql.Pingpong{
		{PingOrPong: pgtype.Text{String: "Ping", Valid: true}},
		{PingOrPong: pgtype.Text{String: "Pong", Valid: true}},
	}
	expected := []string{"Ping", "Pong"}

	result := TranslatePingPongsFromDB(input)
	if result == nil {
		t.Fatal("Expected non-nil result")
	}
	if len(result.PingPongs) != len(expected) {
		t.Fatalf("Expected %d PingPongEntities, got %d", len(expected), len(result.PingPongs))
	}
	for i, msg := range expected {
		if result.PingPongs[i].Message != msg {
			t.Errorf("At index %d, expected %q, got %q", i, msg, result.PingPongs[i].Message)
		}
	}
}

func TestTranslateFromDBRaw_DeletedAtZero(t *testing.T) {
	t.Parallel()

	id := uuid.New()
	now := time.Now()
	p := postgresql.Pingpong{
		PingpongID: pgtype.UUID{Bytes: id, Valid: true},
		PingOrPong: pgtype.Text{String: "ping", Valid: true},
		CreatedAt:  pgtype.Timestamptz{Time: now, Valid: true},
		UpdatedAt:  pgtype.Timestamptz{Time: now, Valid: true},
		DeletedAt:  pgtype.Timestamptz{Time: time.Time{}, Valid: false},
		Deleted:    false,
	}

	result := TranslateFromDBRaw(p)

	if result.PingPongMetadata == nil {
		t.Fatal("PingPongMetadata should not be nil")
	}
	if result.PingPongMetadata.PingPongID != id {
		t.Errorf("expected PingPongID %v, got %v", id, result.PingPongMetadata.PingPongID)
	}
	if !result.PingPongMetadata.CreatedAt.Equal(now) {
		t.Errorf("expected CreatedAt %v, got %v", now, result.PingPongMetadata.CreatedAt)
	}
	if result.PingPongMetadata.DeletedAt != nil {
		t.Errorf("expected DeletedAt nil, got %v", result.PingPongMetadata.DeletedAt)
	}
	if result.PingPongMetadata.Deleted != false {
		t.Errorf("expected Deleted false, got %v", result.PingPongMetadata.Deleted)
	}
	if result.Message != "ping" {
		t.Errorf("expected Message 'ping', got %v", result.Message)
	}
}

func TestTranslateFromDBRaw_DeletedAtSet(t *testing.T) {
	t.Parallel()

	id := uuid.New()
	now := time.Now()
	deletedAt := now.Add(1 * time.Hour)
	p := postgresql.Pingpong{
		PingpongID: pgtype.UUID{Bytes: id, Valid: true},
		PingOrPong: pgtype.Text{String: "pong", Valid: true},
		CreatedAt:  pgtype.Timestamptz{Time: now, Valid: true},
		UpdatedAt:  pgtype.Timestamptz{Time: now, Valid: true},
		DeletedAt:  pgtype.Timestamptz{Time: deletedAt, Valid: true},
		Deleted:    true,
	}

	result := TranslateFromDBRaw(p)

	if result.PingPongMetadata == nil {
		t.Fatal("PingPongMetadata should not be nil")
	}
	if result.PingPongMetadata.PingPongID != id {
		t.Errorf("expected PingPongID %v, got %v", id, result.PingPongMetadata.PingPongID)
	}
	if !result.PingPongMetadata.CreatedAt.Equal(now) {
		t.Errorf("expected CreatedAt %v, got %v", now, result.PingPongMetadata.CreatedAt)
	}
	if result.PingPongMetadata.DeletedAt == nil {
		t.Error("expected DeletedAt not nil")
	} else if !result.PingPongMetadata.DeletedAt.Equal(deletedAt) {
		t.Errorf("expected DeletedAt %v, got %v", deletedAt, result.PingPongMetadata.DeletedAt)
	}
	if result.PingPongMetadata.Deleted != true {
		t.Errorf("expected Deleted true, got %v", result.PingPongMetadata.Deleted)
	}
	if result.Message != "pong" {
		t.Errorf("expected Message 'pong', got %v", result.Message)
	}
}
