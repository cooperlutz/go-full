// Package utilitee provides utility functions that can be used across the project
package utilitee

import (
	"math"
	"time"
)

// RightNow is a simple utility function to ensure consistency of returning a "now" timestamp being localized to UTC.
func RightNow() time.Time {
	return time.Now().UTC()
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
