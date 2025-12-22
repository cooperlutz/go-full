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

func TestStrPtr(t *testing.T) {
	t.Parallel()
	s := "hello"
	ptr := utilitee.StrPtr(s)
	assert.NotNil(t, ptr)
	assert.Equal(t, &s, ptr)
}

func TestBoolPtr(t *testing.T) {
	t.Parallel()
	b := true
	ptr := utilitee.BoolPtr(b)
	assert.NotNil(t, ptr)
	assert.Equal(t, &b, ptr)
}

func TestTimePtr(t *testing.T) {
	t.Parallel()
	timeVal := time.Now()
	ptr := utilitee.TimePtr(timeVal)
	assert.NotNil(t, ptr)
	assert.Equal(t, &timeVal, ptr)
}
