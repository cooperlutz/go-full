package pingpong_test

import (
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/internal/pingpong"
)

func TestNewModule(t *testing.T) {
	var dummyPool *pgxpool.Pool

	module := pingpong.NewModule(dummyPool)

	assert.NotNil(t, module)
	assert.NotNil(t, module.RestApi)
	assert.NotNil(t, module.UseCase)
	assert.NotNil(t, module.PersistentRepo)
}
