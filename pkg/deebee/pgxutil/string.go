package pgxutil

import "github.com/jackc/pgx/v5/pgtype"

// StringToPgtypeText converts a string pointer to a pgx Text type, treating nil pointers as null values.
func StringToPgtypeText(s *string) pgtype.Text {
	if s == nil {
		return pgtype.Text{String: "", Valid: false}
	}

	return pgtype.Text{String: *s, Valid: true}
}

// PgtypeTextToStringPtr converts a pgx Text type to a string pointer, returning nil if the pgx Text is null.
func PgtypeTextToStringPtr(t pgtype.Text) *string {
	if !t.Valid {
		return nil
	}

	return &t.String
}

// StrToPgtypeText converts a string pointer to a pgx Text type.
func StrToPgtypeText(s *string) pgtype.Text {
	if s == nil {
		return pgtype.Text{String: "", Valid: false}
	}

	return pgtype.Text{String: *s, Valid: true}
}
