package pingpong

import (
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
)

func TestNewModule(t *testing.T) {
	var dummyPool *pgxpool.Pool

	handler := NewModule(dummyPool)

	assert.NotNil(t, handler)
}
