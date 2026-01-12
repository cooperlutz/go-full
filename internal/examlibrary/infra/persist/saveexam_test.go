package persist

import (
	"context"
	"testing"

	"github.com/pashagolub/pgxmock/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/cooperlutz/go-full/test/fixtures"
	"github.com/cooperlutz/go-full/test/mocks"
)

func TestSaveExam_Success(t *testing.T) {
	// Arrange
	mockDB, err := pgxmock.NewPool()
	if err != nil {
		t.Fatal(err)
	}
	defer mockDB.Close()
	mockQuerier := mocks.NewMockIQuerierExamLibrary(t)
	ctx := context.Background()
	repo := &examLibraryPersistPostgresRepository{
		db:    mockDB,
		query: mockQuerier,
	}
	mockDB.ExpectBegin()
	mockQuerier.On("SaveExam",
		mock.Anything,
		mock.Anything,
	).Return(
		nil,
	)
	mockQuerier.On("SaveExamQuestion",
		mock.Anything,
		mock.Anything,
	).Return(
		nil,
	)
	mockQuerier.On(
		"WithTx",
		mock.Anything,
	).Return(
		mockQuerier,
	)
	mockDB.ExpectCommit()
	p := fixtures.ValidDomainExam

	// Act
	err = repo.SaveExam(ctx, p)

	// Assert
	assert.NoError(t, err)
}
