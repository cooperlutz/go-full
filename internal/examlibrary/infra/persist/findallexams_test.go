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

func TestFindAll_Success(t *testing.T) {
	// Arrange
	mQuerier := mocks.NewMockIQuerierExamLibrary(t)
	repo := &examLibraryPersistPostgresRepository{
		query: mQuerier,
	}
	mockResponseExams := []persist_postgres.ExamLibraryExam{
		fixtures.ValidDBExamLibraryExam,
	}
	mockResponseExamQuestions := []persist_postgres.ExamLibraryExamQuestion{
		fixtures.ValidDBExamQuestionMultipleChoice,
	}
	mQuerier.On(
		"FindAllExams",
		mock.Anything,
	).Return(
		mockResponseExams,
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
	exam, err := repo.FindAllExams(context.Background())

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, exam)
}
