package pgxutil

import "github.com/jackc/pgx/v5/pgtype"

// SliceOfPtrsToPgtype converts a slice of pointers to a slice of values, treating nil pointers as null values.
func Float32ToPgtypeFloat4(f *float32) pgtype.Float4 {
	if f == nil {
		return pgtype.Float4{Float32: 0, Valid: false}
	}

	return pgtype.Float4{Float32: *f, Valid: true}
}

// Float64ToPgtypeFloat8 converts a float64 pointer to a pgx Float8 type.
func Float64ToPgtypeFloat8(f *float64) pgtype.Float8 {
	if f == nil {
		return pgtype.Float8{Float64: 0, Valid: false}
	}

	return pgtype.Float8{Float64: *f, Valid: true}
}

// PgtypeFloat4ToFloat32Ptr converts a pgx Float4 type to a float32 pointer, returning nil if the pgx Float4 is null.
func PgtypeFloat4ToFloat32Ptr(f pgtype.Float4) *float32 {
	if !f.Valid {
		return nil
	}

	return &f.Float32
}

// PgtypeFloat8ToFloat64Ptr converts a pgx Float8 type to a float64 pointer, returning nil if the pgx Float8 is null.
func PgtypeFloat8ToFloat64Ptr(f pgtype.Float8) *float64 {
	if !f.Valid {
		return nil
	}

	return &f.Float64
}
