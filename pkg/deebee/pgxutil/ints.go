package pgxutil

import (
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/cooperlutz/go-full/pkg/utilitee"
)

// IntToPgtypeInt4 converts an int pointer to a pgx Int4 type.
func IntToPgtypeInt4(i *int) pgtype.Int4 {
	// Check for overflow/underflow
	if i == nil || utilitee.SafeIntToInt32(i) == 0 {
		return pgtype.Int4{Int32: 0, Valid: false}
	}

	return pgtype.Int4{Int32: utilitee.SafeIntToInt32(i), Valid: true}
}

// Int32ToPgtypeInt4 converts an int32 pointer to a pgx Int4 type.
func Int32ToPgtypeInt4(i *int32) pgtype.Int4 {
	if i == nil {
		return pgtype.Int4{Int32: 0, Valid: false}
	}

	return pgtype.Int4{Int32: *i, Valid: true}
}

// PgtypeInt4ToInt32Ptr converts a pgx Int4 type to an int32 pointer, returning nil if the pgx Int4 is null.
func PgtypeInt4ToInt32Ptr(i pgtype.Int4) *int32 {
	if !i.Valid {
		return nil
	}

	return &i.Int32
}

// Int16ToPgtypeInt2 converts an int16 pointer to a pgx Int2 type.
func Int16ToPgtypeInt2(i *int16) pgtype.Int2 {
	if i == nil {
		return pgtype.Int2{Int16: 0, Valid: false}
	}

	return pgtype.Int2{Int16: *i, Valid: true}
}

// PgtypeInt2ToInt16Ptr converts a pgx Int2 type to an int16 pointer, returning nil if the pgx Int2 is null.
func PgtypeInt2ToInt16Ptr(i pgtype.Int2) *int16 {
	if !i.Valid {
		return nil
	}

	return &i.Int16
}

// Int64ToPgtypeInt8 converts an int64 pointer to a pgx Int8 type.
func PgtypeInt8ToInt64Ptr(i pgtype.Int8) *int64 {
	if !i.Valid {
		return nil
	}

	return &i.Int64
}

// Int64ToPgtypeInt8 converts an int64 pointer to a pgx Int8 type.
func Int64ToPgtypeInt8(i *int64) pgtype.Int8 {
	if i == nil {
		return pgtype.Int8{Int64: 0, Valid: false}
	}

	return pgtype.Int8{Int64: *i, Valid: true}
}
