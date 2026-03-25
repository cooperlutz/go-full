package utilitee_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/pkg/utilitee"
)

func TestRightNow(t *testing.T) {
	t.Parallel()
	now := utilitee.RightNow()
	otherNow := time.Now()
	assert.WithinDuration(t, now, otherNow, 10*time.Millisecond)
}

func TestSafeIntToInt32(t *testing.T) {
	t.Parallel()

	// Test with nil pointer
	var nilIntPtr *int
	result := utilitee.SafeIntToInt32(nilIntPtr)
	assert.Equal(t, int32(0), result)

	// Test with valid int pointer within int32 range
	validInt := 12345
	validIntPtr := &validInt
	result = utilitee.SafeIntToInt32(validIntPtr)
	assert.Equal(t, int32(12345), result)

	// Test with int pointer exceeding int32 max value
	largeInt := int(^uint32(0)>>1) + 1 // One more than max int32
	largeIntPtr := &largeInt
	result = utilitee.SafeIntToInt32(largeIntPtr)
	assert.Equal(t, int32(0), result)

	// Test with int pointer below int32 min value
	smallInt := -int(^uint32(0)>>1) - 2 // One less than min int32
	smallIntPtr := &smallInt
	result = utilitee.SafeIntToInt32(smallIntPtr)
	assert.Equal(t, int32(0), result)
}

func TestSliceOfPointersToPointerSlice(t *testing.T) {
	t.Parallel()

	// Test with empty slice
	emptySlice := []*int{}
	result := utilitee.SliceOfPointersToPointerSlice(emptySlice)
	assert.NotNil(t, result)
	assert.Empty(t, result)

	// Test with valid slice of pointers
	int1 := 1
	int2 := 2
	int3 := 3
	sliceOfPointers := []*int{&int1, &int2, &int3}
	result = utilitee.SliceOfPointersToPointerSlice(sliceOfPointers)
	assert.NotNil(t, result)
	res := *result
	assert.Len(t, res, 3)
	assert.Equal(t, int(1), res[0])
	assert.Equal(t, int(2), res[1])
	assert.Equal(t, int(3), res[2])
}

func TestSliceOfValuesToSliceOfPointers(t *testing.T) {
	t.Parallel()

	var nilSlice []int
	result := utilitee.SliceOfValuesToSliceOfPointers(nilSlice)
	assert.Nil(t, result)

	// Test with empty slice
	emptySlice := []int{}
	result = utilitee.SliceOfValuesToSliceOfPointers(emptySlice)
	assert.NotNil(t, result)
	assert.Empty(t, result)

	// Test with valid slice of values
	values := []int{1, 2, 3}
	result = utilitee.SliceOfValuesToSliceOfPointers(values)
	assert.NotNil(t, result)
	assert.Len(t, result, 3)
	assert.Equal(t, int(1), *result[0])
	assert.Equal(t, int(2), *result[1])
	assert.Equal(t, int(3), *result[2])
}
