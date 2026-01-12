package pubsub_test

import (
	"testing"

	"github.com/pashagolub/pgxmock/v4"
	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/internal/examlibrary/infra/pubsub"
	"github.com/cooperlutz/go-full/test/mocks"
)

func TestNew_Success(t *testing.T) {
	pgxmock, err := pgxmock.NewPool()
	assert.NoError(t, err)
	defer pgxmock.Close()

	mockRepo := mocks.NewMockIExamLibraryRepository(t)

	ps, err := pubsub.New(pgxmock, mockRepo)

	assert.NoError(t, err)
	assert.NotNil(t, ps)
}
