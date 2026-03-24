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

// SafeIntToInt64 safely converts an *int to int64, returning 0 if the pointer is nil or if the value is out of int64 bounds.
func SliceOfPointersToPointerSlice[T any](slice []*T) *[]T {
	result := SliceOfPointersToSlice(slice)

	return &result
}

// PointerSliceToSliceOfPointers converts a pointer to a slice of values to a slice of pointers. If the input pointer is nil, it returns nil.
func SliceOfPointersToSlice[T any](slice []*T) []T {
	result := make([]T, len(slice))
	for i, v := range slice {
		if v != nil {
			result[i] = *v
		}
	}

	return result
}

// SliceOfValuesToSliceOfPointers converts a slice of values to a slice of pointers. If the input slice is nil, it returns nil.
func SliceOfValuesToSliceOfPointers[T any](slice []T) []*T {
	if slice == nil {
		return nil
	}

	result := make([]*T, len(slice))
	for i, v := range slice {
		result[i] = &v
	}

	return result
}
