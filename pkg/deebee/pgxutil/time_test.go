package pgxutil_test

import (
	"database/sql"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
)

func TestPgtypeTimeNullToTimePtr(t *testing.T) {
	tests := []struct {
		name     string
		input    sql.NullTime
		expected *time.Time
	}{
		{
			name:     "valid time",
			input:    sql.NullTime{Time: time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC), Valid: true},
			expected: &time.Time{},
		},
		{
			name:     "null time",
			input:    sql.NullTime{Time: time.Time{}, Valid: false},
			expected: nil,
		},
		{
			name:     "valid time with different date",
			input:    sql.NullTime{Time: time.Date(2023, 12, 25, 23, 59, 59, 999999999, time.UTC), Valid: true},
			expected: &time.Time{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pgxutil.PgtypeTimeNullToTimePtr(tt.input)

			if tt.expected == nil {
				if result != nil {
					t.Errorf("expected nil, got %v", result)
				}
				return
			}

			if result == nil {
				t.Errorf("expected non-nil pointer, got nil")
				return
			}

			if !result.Equal(tt.input.Time) {
				t.Errorf("expected %v, got %v", tt.input.Time, *result)
			}
		})
	}
}

func TestTimeToPgtypeTimeNull(t *testing.T) {
	tests := []struct {
		name     string
		input    *time.Time
		expected sql.NullTime
	}{
		{
			name:     "valid time pointer",
			input:    &time.Time{},
			expected: sql.NullTime{Time: time.Time{}, Valid: true},
		},
		{
			name:     "nil pointer",
			input:    nil,
			expected: sql.NullTime{Time: time.Time{}, Valid: false},
		},
		{
			name:     "valid time with specific date",
			input:    new(time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)),
			expected: sql.NullTime{Time: time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC), Valid: true},
		},
		{
			name:     "valid time with different date",
			input:    new(time.Date(2023, 12, 25, 23, 59, 59, 999999999, time.UTC)),
			expected: sql.NullTime{Time: time.Date(2023, 12, 25, 23, 59, 59, 999999999, time.UTC), Valid: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pgxutil.TimeToPgtypeTimeNull(tt.input)

			if result.Valid != tt.expected.Valid {
				t.Errorf("expected Valid=%v, got Valid=%v", tt.expected.Valid, result.Valid)
				return
			}

			if result.Valid && !result.Time.Equal(tt.expected.Time) {
				t.Errorf("expected %v, got %v", tt.expected.Time, result.Time)
			}
		})
	}
}

func TestTimeToPgtypeDate(t *testing.T) {
	tests := []struct {
		name     string
		input    *time.Time
		expected pgtype.Date
	}{
		{
			name:     "valid time pointer",
			input:    new(time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)),
			expected: pgtype.Date{Time: time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC), Valid: true},
		},
		{
			name:     "nil pointer",
			input:    nil,
			expected: pgtype.Date{Time: time.Time{}, Valid: false},
		},
		{
			name:     "valid time with different date",
			input:    new(time.Date(2023, 12, 25, 23, 59, 59, 999999999, time.UTC)),
			expected: pgtype.Date{Time: time.Date(2023, 12, 25, 23, 59, 59, 999999999, time.UTC), Valid: true},
		},
		{
			name:     "valid time with time component ignored",
			input:    new(time.Date(2024, 6, 10, 15, 45, 30, 123456789, time.UTC)),
			expected: pgtype.Date{Time: time.Date(2024, 6, 10, 15, 45, 30, 123456789, time.UTC), Valid: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pgxutil.TimeToPgtypeDate(tt.input)

			if result.Valid != tt.expected.Valid {
				t.Errorf("expected Valid=%v, got Valid=%v", tt.expected.Valid, result.Valid)
				return
			}

			if result.Valid && !result.Time.Equal(tt.expected.Time) {
				t.Errorf("expected %v, got %v", tt.expected.Time, result.Time)
			}
		})
	}
}

func TestTimeToPgtypeTime(t *testing.T) {
	tests := []struct {
		name     string
		input    *time.Time
		expected pgtype.Time
	}{
		{
			name:     "nil pointer",
			input:    nil,
			expected: pgtype.Time{Microseconds: 0, Valid: false},
		},
		{
			name:     "valid time with morning time",
			input:    new(time.Date(2024, 1, 15, 10, 30, 45, 500000000, time.UTC)),
			expected: pgtype.Time{Microseconds: 10*3600*1000000 + 30*60*1000000 + 45*1000000 + 500000, Valid: true},
		},
		{
			name:     "valid time at midnight",
			input:    new(time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)),
			expected: pgtype.Time{Microseconds: 0, Valid: true},
		},
		{
			name:     "valid time at end of day",
			input:    new(time.Date(2024, 1, 15, 23, 59, 59, 999999000, time.UTC)),
			expected: pgtype.Time{Microseconds: 23*3600*1000000 + 59*60*1000000 + 59*1000000 + 999999, Valid: true},
		},
		{
			name:     "valid time with microseconds",
			input:    new(time.Date(2024, 6, 10, 15, 45, 30, 123456000, time.UTC)),
			expected: pgtype.Time{Microseconds: 15*3600*1000000 + 45*60*1000000 + 30*1000000 + 123456, Valid: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pgxutil.TimeToPgtypeTime(tt.input)

			if result.Valid != tt.expected.Valid {
				t.Errorf("expected Valid=%v, got Valid=%v", tt.expected.Valid, result.Valid)
				return
			}

			if result.Valid && result.Microseconds != tt.expected.Microseconds {
				t.Errorf("expected Microseconds=%v, got Microseconds=%v", tt.expected.Microseconds, result.Microseconds)
			}
		})
	}
}
