package pgxutil_test

import (
	"testing"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
)

func TestFloat32ToPgtypeFloat4(t *testing.T) {
	tests := []struct {
		name     string
		input    *float32
		expected pgtype.Float4
	}{
		{
			name:     "nil pointer",
			input:    nil,
			expected: pgtype.Float4{Float32: 0, Valid: false},
		},
		{
			name:     "zero value",
			input:    new(float32(0)),
			expected: pgtype.Float4{Float32: 0, Valid: true},
		},
		{
			name:     "positive value",
			input:    new(float32(3.14)),
			expected: pgtype.Float4{Float32: 3.14, Valid: true},
		},
		{
			name:     "negative value",
			input:    new(float32(-2.71)),
			expected: pgtype.Float4{Float32: -2.71, Valid: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pgxutil.Float32ToPgtypeFloat4(tt.input)
			if result.Valid != tt.expected.Valid {
				t.Errorf("Valid: got %v, want %v", result.Valid, tt.expected.Valid)
			}
			if result.Float32 != tt.expected.Float32 {
				t.Errorf("Float32: got %v, want %v", result.Float32, tt.expected.Float32)
			}
		})
	}
}

func TestFloat64ToPgtypeFloat8(t *testing.T) {
	tests := []struct {
		name     string
		input    *float64
		expected pgtype.Float8
	}{
		{
			name:     "nil pointer",
			input:    nil,
			expected: pgtype.Float8{Float64: 0, Valid: false},
		},
		{
			name:     "zero value",
			input:    new(float64(0)),
			expected: pgtype.Float8{Float64: 0, Valid: true},
		},
		{
			name:     "positive value",
			input:    new(float64(3.14159)),
			expected: pgtype.Float8{Float64: 3.14159, Valid: true},
		},
		{
			name:     "negative value",
			input:    new(float64(-2.71828)),
			expected: pgtype.Float8{Float64: -2.71828, Valid: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pgxutil.Float64ToPgtypeFloat8(tt.input)
			if result.Valid != tt.expected.Valid {
				t.Errorf("Valid: got %v, want %v", result.Valid, tt.expected.Valid)
			}
			if result.Float64 != tt.expected.Float64 {
				t.Errorf("Float64: got %v, want %v", result.Float64, tt.expected.Float64)
			}
		})
	}
}

func TestPgtypeFloat4ToFloat32Ptr(t *testing.T) {
	tests := []struct {
		name     string
		input    pgtype.Float4
		expected *float32
	}{
		{
			name:     "invalid float4",
			input:    pgtype.Float4{Float32: 0, Valid: false},
			expected: nil,
		},
		{
			name:     "zero value",
			input:    pgtype.Float4{Float32: 0, Valid: true},
			expected: new(float32(0)),
		},
		{
			name:     "positive value",
			input:    pgtype.Float4{Float32: 3.14, Valid: true},
			expected: new(float32(3.14)),
		},
		{
			name:     "negative value",
			input:    pgtype.Float4{Float32: -2.71, Valid: true},
			expected: new(float32(-2.71)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pgxutil.PgtypeFloat4ToFloat32Ptr(tt.input)
			if result == nil && tt.expected == nil {
				return
			}
			if result == nil || tt.expected == nil {
				t.Errorf("got %v, want %v", result, tt.expected)
				return
			}
			if *result != *tt.expected {
				t.Errorf("got %v, want %v", *result, *tt.expected)
			}
		})
	}
}

func TestPgtypeFloat8ToFloat64Ptr(t *testing.T) {
	tests := []struct {
		name     string
		input    pgtype.Float8
		expected *float64
	}{
		{
			name:     "invalid float8",
			input:    pgtype.Float8{Float64: 0, Valid: false},
			expected: nil,
		},
		{
			name:     "zero value",
			input:    pgtype.Float8{Float64: 0, Valid: true},
			expected: new(float64(0)),
		},
		{
			name:     "positive value",
			input:    pgtype.Float8{Float64: 3.14159, Valid: true},
			expected: new(float64(3.14159)),
		},
		{
			name:     "negative value",
			input:    pgtype.Float8{Float64: -2.71828, Valid: true},
			expected: new(float64(-2.71828)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pgxutil.PgtypeFloat8ToFloat64Ptr(tt.input)
			if result == nil && tt.expected == nil {
				return
			}
			if result == nil || tt.expected == nil {
				t.Errorf("got %v, want %v", result, tt.expected)
				return
			}
			if *result != *tt.expected {
				t.Errorf("got %v, want %v", *result, *tt.expected)
			}
		})
	}
}
