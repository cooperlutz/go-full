package utilitee

import (
	"math"
	"time"
)

// RightNow is a simple utility function to ensure consistency of returning a "now" timestamp being localized to UTC.
func RightNow() time.Time {
	return time.Now().UTC()
}

// StrPtr returns a pointer to the given string value.
//
//go:fix inline
func StrPtr(s string) *string {
	return new(s)
}

// BoolPtr returns a pointer to the given bool value.
//
//go:fix inline
func BoolPtr(b bool) *bool {
	return new(b)
}

// TimePtr returns a pointer to the given time.Time value.
//
//go:fix inline
func TimePtr(t time.Time) *time.Time {
	return new(t)
}

// IntPtr returns a pointer to the given int value.
//
//go:fix inline
func IntPtr(i int) *int {
	return new(i)
}

// SafeIntToInt32 safely converts an *int to int32, returning 0 if the pointer is nil or if the value is out of int32 bounds.
func SafeIntToInt32(i *int) int32 {
	if i == nil {
		return 0
	}
	// Check for overflow/underflow
	if *i > math.MaxInt32 || *i < math.MinInt32 {
		return 0
	}

	return int32(*i)
}
