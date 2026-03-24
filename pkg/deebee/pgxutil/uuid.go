package pgxutil

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// UUIDToPgtypeUUID converts a uuid.UUID to a pgx UUID type.
func UUIDToPgtypeUUID(id uuid.UUID) pgtype.UUID {
	return pgtype.UUID{Bytes: id, Valid: true}
}

// UUIDPtrToPgtypeUUID converts a *uuid.UUID to a pgx UUID type, handling nil pointers appropriately.
func UUIDPtrToPgtypeUUID(id *uuid.UUID) pgtype.UUID {
	if id == nil {
		return pgtype.UUID{Bytes: uuid.Nil, Valid: false}
	}

	return pgtype.UUID{Bytes: *id, Valid: true}
}

// UUIDSliceToPgtypeUUIDSlice converts a slice of uuid.UUID values to a slice of pgx UUID types.
func UUIDSliceToPgtypeUUIDSlice(ids []uuid.UUID) []pgtype.UUID {
	pgUUIDs := make([]pgtype.UUID, len(ids))
	for i, id := range ids {
		pgUUIDs[i] = UUIDToPgtypeUUID(id)
	}

	return pgUUIDs
}

// UUIDSliceOfPtrsToPgtypeUUIDSlice converts a slice of *uuid.UUID pointers to a slice of pgx UUID types, handling nil pointers appropriately.
func UUIDSliceOfPtrsToPgtypeUUIDSlice(ids []*uuid.UUID) []pgtype.UUID {
	pgUUIDs := make([]pgtype.UUID, len(ids))
	for i, id := range ids {
		pgUUIDs[i] = UUIDPtrToPgtypeUUID(id)
	}

	return pgUUIDs
}

// PgtypeUUIDToUUID converts a pgx UUID type to a uuid.UUID value, returning uuid.Nil if the pgx UUID is null.
func PgtypeUUIDToUUID(id pgtype.UUID) uuid.UUID {
	if !id.Valid {
		return uuid.Nil
	}

	return id.Bytes
}

// PgtypeUUIDToUUIDPtr converts a pgx UUID type to a *uuid.UUID pointer, returning nil if the pgx UUID is null.
func PgtypeUUIDToUUIDPtr(id pgtype.UUID) *uuid.UUID {
	if !id.Valid {
		return nil
	}

	bytes := uuid.UUID(id.Bytes)

	return &bytes
}

// PgtypeUUIDSliceToUUIDSlice converts a slice of pgx UUID types to a slice of uuid.UUID values, returning uuid.Nil for any null pgx UUID values.
func PgtypeUUIDSliceToUUIDSlice(ids []pgtype.UUID) []uuid.UUID {
	uuids := make([]uuid.UUID, len(ids))
	for i, id := range ids {
		uuids[i] = PgtypeUUIDToUUID(id)
	}

	return uuids
}

// PgtypeUUIDSliceToUUIDSlicePtr converts a slice of pgx UUID types to a slice of *uuid.UUID pointers, returning nil for any null pgx UUID values.
func PgtypeUUIDSliceToUUIDSliceOfPtrs(ids []pgtype.UUID) []*uuid.UUID {
	uuids := make([]*uuid.UUID, len(ids))
	for i, id := range ids {
		uuids[i] = PgtypeUUIDToUUIDPtr(id)
	}

	return uuids
}
