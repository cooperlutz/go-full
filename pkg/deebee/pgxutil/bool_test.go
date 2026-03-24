package pgxutil_test

import (
	"testing"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
)

func TestBoolToPgtypeBool(t *testing.T) {
	tests := []struct {
		name string
		b    *bool
		want pgtype.Bool
	}{
		{
			name: "nil pointer returns invalid Bool",
			b:    nil,
			want: pgtype.Bool{Bool: false, Valid: false},
		},
		{
			name: "true pointer returns valid Bool with true value",
			b:    func() *bool { v := true; return &v }(),
			want: pgtype.Bool{Bool: true, Valid: true},
		},
		{
			name: "false pointer returns valid Bool with false value",
			b:    func() *bool { v := false; return &v }(),
			want: pgtype.Bool{Bool: false, Valid: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pgxutil.BoolToPgtypeBool(tt.b)
			if got != tt.want {
				t.Errorf("BoolToPgtypeBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPgtypeBoolToBoolPtr(t *testing.T) {
	tests := []struct {
		name string
		b    pgtype.Bool
		want *bool
	}{
		{
			name: "invalid Bool returns nil pointer",
			b:    pgtype.Bool{Bool: false, Valid: false},
			want: nil,
		},
		{
			name: "valid Bool with true value returns true pointer",
			b:    pgtype.Bool{Bool: true, Valid: true},
			want: func() *bool { v := true; return &v }(),
		},
		{
			name: "valid Bool with false value returns false pointer",
			b:    pgtype.Bool{Bool: false, Valid: true},
			want: func() *bool { v := false; return &v }(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pgxutil.PgtypeBoolToBoolPtr(tt.b)
			if (got == nil && tt.want == nil) || (got != nil && tt.want != nil && *got == *tt.want) {
				return
			}
			t.Errorf("PgtypeBoolToBoolPtr() = %v, want %v", got, tt.want)
		})
	}
}
