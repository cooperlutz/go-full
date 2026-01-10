package persist_postgres_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	persist_postgres "github.com/cooperlutz/go-full/internal/examlibrary/infra/persist/postgres"
)

func TestQueriesWrapper_WithTx(t *testing.T) {
	t.Parallel()
	// Arrange & Act
	var qw persist_postgres.IQuerierExamLibrary = persist_postgres.NewQueriesWrapper(nil)
	qwWithTx := qw.WithTx(nil)
	// Assert
	assert.NotNil(t, qwWithTx)
}
