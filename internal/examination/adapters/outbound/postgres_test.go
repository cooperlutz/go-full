package outbound_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/cooperlutz/go-full/internal/examination/adapters/outbound"
	"github.com/cooperlutz/go-full/test/mocks"
)

func TestNewPostgresAdapter(t *testing.T) {
	testDB := mocks.NewMockDBTX(t)
	repo := outbound.NewPostgresAdapter(testDB)
	assert.NotNil(t, repo)
}

func TestPostgres_FindAll(t *testing.T) {
	ctx := context.Background()
	testDB := mocks.NewMockIQuerierExamination(t)
	repo := outbound.PostgresAdapter{
		Handler: testDB,
	}
	testDB.On("FindAllExams", mock.Anything, mock.Anything).Return(
		outbound.FixtureExams, nil,
	)
	exams, err := repo.FindAll(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, exams)
}

func TestPostgres_FindAll_Error(t *testing.T) {
	ctx := context.Background()
	testDB := mocks.NewMockIQuerierExamination(t)
	repo := outbound.PostgresAdapter{
		Handler: testDB,
	}
	testDB.On("FindAllExams", mock.Anything, mock.Anything).Return(
		nil, assert.AnError,
	)
	exams, err := repo.FindAll(ctx)
	assert.Error(t, err)
	assert.Nil(t, exams)
}
