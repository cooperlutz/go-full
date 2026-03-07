package outbound

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueriesWrapper_WithTx(t *testing.T) {
	t.Parallel()
	// Arrange & Act
	qw := NewQueriesWrapper(nil)
	qwWithTx := qw.WithTx(nil)
	// Assert
	assert.NotNil(t, qwWithTx)
}
