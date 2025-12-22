package baseentitee_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

func TestDeletedAtFromTime(t *testing.T) {
	nilTime := (*time.Time)(nil)
	deletedAtNil := baseentitee.DeletedAtFromTime(nilTime)
	assert.Nil(t, deletedAtNil)
}
