package pgxutil_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
)

func TestUUIDToPgtypeUUID(t *testing.T) {
	id := uuid.New()
	pgUUID := pgxutil.UUIDToPgtypeUUID(id)

	assert.True(t, pgUUID.Valid)
	assert.Equal(t, [16]byte(id[:]), pgUUID.Bytes)
}

func TestUUIDPtrToPgtypeUUID(t *testing.T) {
	t.Run("with valid UUID pointer", func(t *testing.T) {
		id := uuid.New()
		pgUUID := pgxutil.UUIDPtrToPgtypeUUID(&id)

		assert.True(t, pgUUID.Valid)
		assert.Equal(t, [16]byte(id[:]), pgUUID.Bytes)
	})

	t.Run("with nil pointer", func(t *testing.T) {
		pgUUID := pgxutil.UUIDPtrToPgtypeUUID(nil)

		assert.False(t, pgUUID.Valid)
		assert.Equal(t, [16]byte(uuid.Nil[:]), pgUUID.Bytes)
	})
}

func TestUUIDSliceToPgtypeUUIDSlice(t *testing.T) {
	t.Run("with empty slice", func(t *testing.T) {
		ids := []uuid.UUID{}
		pgUUIDs := pgxutil.UUIDSliceToPgtypeUUIDSlice(ids)

		assert.Empty(t, pgUUIDs)
	})

	t.Run("with single UUID", func(t *testing.T) {
		id := uuid.New()
		ids := []uuid.UUID{id}
		pgUUIDs := pgxutil.UUIDSliceToPgtypeUUIDSlice(ids)

		assert.Len(t, pgUUIDs, 1)
		assert.True(t, pgUUIDs[0].Valid)
		assert.Equal(t, [16]byte(id[:]), pgUUIDs[0].Bytes)
	})

	t.Run("with multiple UUIDs", func(t *testing.T) {
		id1 := uuid.New()
		id2 := uuid.New()
		id3 := uuid.New()
		ids := []uuid.UUID{id1, id2, id3}
		pgUUIDs := pgxutil.UUIDSliceToPgtypeUUIDSlice(ids)

		assert.Len(t, pgUUIDs, 3)
		for i, id := range ids {
			assert.True(t, pgUUIDs[i].Valid)
			assert.Equal(t, [16]byte(id[:]), pgUUIDs[i].Bytes)
		}
	})
}

func TestUUIDSliceOfPtrsToPgtypeUUIDSlice(t *testing.T) {
	t.Run("with empty slice", func(t *testing.T) {
		ids := []*uuid.UUID{}
		pgUUIDs := pgxutil.UUIDSliceOfPtrsToPgtypeUUIDSlice(ids)

		assert.Empty(t, pgUUIDs)
	})

	t.Run("with single valid UUID pointer", func(t *testing.T) {
		id := uuid.New()
		ids := []*uuid.UUID{&id}
		pgUUIDs := pgxutil.UUIDSliceOfPtrsToPgtypeUUIDSlice(ids)

		assert.Len(t, pgUUIDs, 1)
		assert.True(t, pgUUIDs[0].Valid)
		assert.Equal(t, [16]byte(id[:]), pgUUIDs[0].Bytes)
	})

	t.Run("with single nil pointer", func(t *testing.T) {
		ids := []*uuid.UUID{nil}
		pgUUIDs := pgxutil.UUIDSliceOfPtrsToPgtypeUUIDSlice(ids)

		assert.Len(t, pgUUIDs, 1)
		assert.False(t, pgUUIDs[0].Valid)
		assert.Equal(t, [16]byte(uuid.Nil[:]), pgUUIDs[0].Bytes)
	})

	t.Run("with multiple mixed pointers", func(t *testing.T) {
		id1 := uuid.New()
		id2 := uuid.New()
		ids := []*uuid.UUID{&id1, nil, &id2}
		pgUUIDs := pgxutil.UUIDSliceOfPtrsToPgtypeUUIDSlice(ids)

		assert.Len(t, pgUUIDs, 3)
		assert.True(t, pgUUIDs[0].Valid)
		assert.Equal(t, [16]byte(id1[:]), pgUUIDs[0].Bytes)
		assert.False(t, pgUUIDs[1].Valid)
		assert.Equal(t, [16]byte(uuid.Nil[:]), pgUUIDs[1].Bytes)
		assert.True(t, pgUUIDs[2].Valid)
		assert.Equal(t, [16]byte(id2[:]), pgUUIDs[2].Bytes)
	})
}

func TestPgtypeUUIDToUUID(t *testing.T) {
	t.Run("with valid UUID", func(t *testing.T) {
		id := uuid.New()
		pgUUID := pgtype.UUID{Bytes: id, Valid: true}
		result := pgxutil.PgtypeUUIDToUUID(pgUUID)

		assert.Equal(t, id, result)
	})

	t.Run("with invalid UUID", func(t *testing.T) {
		pgUUID := pgtype.UUID{Bytes: uuid.Nil, Valid: false}
		result := pgxutil.PgtypeUUIDToUUID(pgUUID)

		assert.Equal(t, uuid.Nil, result)
	})
}

func TestPgtypeUUIDToUUIDPtr(t *testing.T) {
	t.Run("with valid UUID", func(t *testing.T) {
		id := uuid.New()
		pgUUID := pgtype.UUID{Bytes: id, Valid: true}
		result := pgxutil.PgtypeUUIDToUUIDPtr(pgUUID)

		assert.NotNil(t, result)
		assert.Equal(t, id, *result)
	})

	t.Run("with invalid UUID", func(t *testing.T) {
		pgUUID := pgtype.UUID{Bytes: uuid.Nil, Valid: false}
		result := pgxutil.PgtypeUUIDToUUIDPtr(pgUUID)

		assert.Nil(t, result)
	})
}

func TestPgtypeUUIDSliceToUUIDSlice(t *testing.T) {
	t.Run("with empty slice", func(t *testing.T) {
		ids := []pgtype.UUID{}
		uuids := pgxutil.PgtypeUUIDSliceToUUIDSlice(ids)

		assert.Empty(t, uuids)
	})

	t.Run("with single valid UUID", func(t *testing.T) {
		id := uuid.New()
		ids := []pgtype.UUID{{Bytes: id, Valid: true}}
		uuids := pgxutil.PgtypeUUIDSliceToUUIDSlice(ids)

		assert.Len(t, uuids, 1)
		assert.Equal(t, id, uuids[0])
	})

	t.Run("with single invalid UUID", func(t *testing.T) {
		ids := []pgtype.UUID{{Bytes: uuid.Nil, Valid: false}}
		uuids := pgxutil.PgtypeUUIDSliceToUUIDSlice(ids)

		assert.Len(t, uuids, 1)
		assert.Equal(t, uuid.Nil, uuids[0])
	})

	t.Run("with multiple valid UUIDs", func(t *testing.T) {
		id1 := uuid.New()
		id2 := uuid.New()
		id3 := uuid.New()
		ids := []pgtype.UUID{
			{Bytes: id1, Valid: true},
			{Bytes: id2, Valid: true},
			{Bytes: id3, Valid: true},
		}
		uuids := pgxutil.PgtypeUUIDSliceToUUIDSlice(ids)

		assert.Len(t, uuids, 3)
		assert.Equal(t, id1, uuids[0])
		assert.Equal(t, id2, uuids[1])
		assert.Equal(t, id3, uuids[2])
	})

	t.Run("with multiple mixed valid and invalid UUIDs", func(t *testing.T) {
		id1 := uuid.New()
		id2 := uuid.New()
		ids := []pgtype.UUID{
			{Bytes: id1, Valid: true},
			{Bytes: uuid.Nil, Valid: false},
			{Bytes: id2, Valid: true},
		}
		uuids := pgxutil.PgtypeUUIDSliceToUUIDSlice(ids)

		assert.Len(t, uuids, 3)
		assert.Equal(t, id1, uuids[0])
		assert.Equal(t, uuid.Nil, uuids[1])
		assert.Equal(t, id2, uuids[2])
	})
}

func TestPgtypeUUIDSliceToUUIDSliceOfPtrs(t *testing.T) {
	t.Run("with empty slice", func(t *testing.T) {
		ids := []pgtype.UUID{}
		uuids := pgxutil.PgtypeUUIDSliceToUUIDSliceOfPtrs(ids)

		assert.Empty(t, uuids)
	})

	t.Run("with single valid UUID", func(t *testing.T) {
		id := uuid.New()
		ids := []pgtype.UUID{{Bytes: id, Valid: true}}
		uuids := pgxutil.PgtypeUUIDSliceToUUIDSliceOfPtrs(ids)

		assert.Len(t, uuids, 1)
		assert.NotNil(t, uuids[0])
		assert.Equal(t, id, *uuids[0])
	})

	t.Run("with single invalid UUID", func(t *testing.T) {
		ids := []pgtype.UUID{{Bytes: uuid.Nil, Valid: false}}
		uuids := pgxutil.PgtypeUUIDSliceToUUIDSliceOfPtrs(ids)

		assert.Len(t, uuids, 1)
		assert.Nil(t, uuids[0])
	})

	t.Run("with multiple valid UUIDs", func(t *testing.T) {
		id1 := uuid.New()
		id2 := uuid.New()
		id3 := uuid.New()
		ids := []pgtype.UUID{
			{Bytes: id1, Valid: true},
			{Bytes: id2, Valid: true},
			{Bytes: id3, Valid: true},
		}
		uuids := pgxutil.PgtypeUUIDSliceToUUIDSliceOfPtrs(ids)

		assert.Len(t, uuids, 3)
		assert.NotNil(t, uuids[0])
		assert.Equal(t, id1, *uuids[0])
		assert.NotNil(t, uuids[1])
		assert.Equal(t, id2, *uuids[1])
		assert.NotNil(t, uuids[2])
		assert.Equal(t, id3, *uuids[2])
	})

	t.Run("with multiple mixed valid and invalid UUIDs", func(t *testing.T) {
		id1 := uuid.New()
		id2 := uuid.New()
		ids := []pgtype.UUID{
			{Bytes: id1, Valid: true},
			{Bytes: uuid.Nil, Valid: false},
			{Bytes: id2, Valid: true},
		}
		uuids := pgxutil.PgtypeUUIDSliceToUUIDSliceOfPtrs(ids)

		assert.Len(t, uuids, 3)
		assert.NotNil(t, uuids[0])
		assert.Equal(t, id1, *uuids[0])
		assert.Nil(t, uuids[1])
		assert.NotNil(t, uuids[2])
		assert.Equal(t, id2, *uuids[2])
	})
}
