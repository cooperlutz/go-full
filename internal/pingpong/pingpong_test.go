package pingpong_test

import (
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/internal/pingpong"
)

func TestNewModule(t *testing.T) {
	// Arrange
	var dummyPool *pgxpool.Pool

	// Act
	module := pingpong.NewModule(dummyPool)

	// Assert
	assert.NotNil(t, module)
	assert.NotNil(t, module.RestApi)
	assert.NotNil(t, module.UseCase)
	assert.NotNil(t, module.PersistentRepo)
}
