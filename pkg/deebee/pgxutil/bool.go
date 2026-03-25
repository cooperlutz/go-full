package pgxutil

import "github.com/jackc/pgx/v5/pgtype"

// BoolToPgtypeBool converts a bool pointer to a pgx Bool type.
func BoolToPgtypeBool(b *bool) pgtype.Bool {
	if b == nil {
		return pgtype.Bool{Bool: false, Valid: false}
	}

	return pgtype.Bool{Bool: *b, Valid: true}
}

// PgtypeBoolToBoolPtr converts a pgx Bool type to a bool pointer, returning nil if the pgx Bool is null.
func PgtypeBoolToBoolPtr(b pgtype.Bool) *bool {
	if !b.Valid {
		return nil
	}

	return &b.Bool
}
