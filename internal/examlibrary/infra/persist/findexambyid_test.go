package persist

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	persist_postgres "github.com/cooperlutz/go-full/internal/examlibrary/infra/persist/postgres"
	"github.com/cooperlutz/go-full/test/fixtures"
	"github.com/cooperlutz/go-full/test/mocks"
)

func TestFindOneByID_Success(t *testing.T) {
	// Arrange
	mQuerier := mocks.NewMockIQuerierExamLibrary(t)
	repo := &examLibraryPersistPostgresRepository{
		query: mQuerier,
	}
	mockResponseExam := fixtures.ValidDBExamLibraryExam
	mockResponseExamQuestions := []persist_postgres.ExamLibraryExamQuestion{
		fixtures.ValidDBExamQuestionMultipleChoice,
	}
	mQuerier.On(
		"FindExamByID",
		mock.Anything,
		mock.Anything,
	).Return(
		mockResponseExam,
		nil,
	)
	mQuerier.On(
		"FindAllExamQuestions",
		mock.Anything,
		mock.Anything,
	).Return(
		mockResponseExamQuestions,
		nil,
	)

	// Act
	exam, err := repo.FindExamByID(context.Background(), fixtures.ValidUUID)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, exam)
}

func TestFindOneByID_Failure_FindExamByIDReturnsError(t *testing.T) {
	// Arrange
	mQuerier := mocks.NewMockIQuerierExamLibrary(t)
	repo := &examLibraryPersistPostgresRepository{
		query: mQuerier,
	}
	mockResponseExam := fixtures.ValidDBExamLibraryExam
	mQuerier.On(
		"FindExamByID",
		mock.Anything,
		mock.Anything,
	).Return(
		mockResponseExam,
		assert.AnError,
	)

	// Act
	_, err := repo.FindExamByID(context.Background(), fixtures.ValidUUID)

	// Assert
	assert.Error(t, err)
}

func TestFindOneByID_Failure_FindAllExamQuestionsReturnsError(t *testing.T) {
	// Arrange
	mQuerier := mocks.NewMockIQuerierExamLibrary(t)
	repo := &examLibraryPersistPostgresRepository{
		query: mQuerier,
	}
	mockResponseExam := fixtures.ValidDBExamLibraryExam
	mQuerier.On(
		"FindExamByID",
		mock.Anything,
		mock.Anything,
	).Return(
		mockResponseExam,
		nil,
	)
	mQuerier.On(
		"FindAllExamQuestions",
		mock.Anything,
		mock.Anything,
	).Return(
		[]persist_postgres.ExamLibraryExamQuestion{},
		assert.AnError,
	)

	// Act
	_, err := repo.FindExamByID(context.Background(), fixtures.ValidUUID)

	// Assert
	assert.Error(t, err)
	// assert.Nil(t, exam)
}
