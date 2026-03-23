package pgxutil_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
)

func TestSliceOfPtrsToPgtype(t *testing.T) {
	t.Parallel()

	// Test with nil slice
	var nilSlice []*int
	result := pgxutil.SliceOfPtrsToPgtype(nilSlice)
	assert.NotNil(t, result)
	assert.Empty(t, result)

	// Test with empty slice
	emptySlice := []*int{}
	result = pgxutil.SliceOfPtrsToPgtype(emptySlice)
	assert.NotNil(t, result)
	assert.Empty(t, result)

	// Test with valid slice of pointers
	val1 := 10
	val2 := 20
	val3 := 30
	sliceOfPtrs := []*int{&val1, &val2, &val3}
	result = pgxutil.SliceOfPtrsToPgtype(sliceOfPtrs)
	assert.Len(t, result, 3)
	assert.Equal(t, 10, result[0])
	assert.Equal(t, 20, result[1])
	assert.Equal(t, 30, result[2])

	// Test with nil pointers in slice
	result = pgxutil.SliceOfPtrsToPgtype([]*int{&val1, nil, &val3})
	assert.Len(t, result, 3)
	assert.Equal(t, 10, result[0])
	assert.Equal(t, 0, result[1])
	assert.Equal(t, 30, result[2])

	// Test with string pointers
	str1 := "hello"
	str2 := "world"
	stringResult := pgxutil.SliceOfPtrsToPgtype([]*string{&str1, &str2})
	assert.Len(t, stringResult, 2)
	assert.Equal(t, "hello", stringResult[0])
	assert.Equal(t, "world", stringResult[1])
}
