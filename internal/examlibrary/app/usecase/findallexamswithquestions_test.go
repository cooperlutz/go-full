package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/cooperlutz/go-full/internal/examlibrary/app/query"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/usecase"
	"github.com/cooperlutz/go-full/internal/examlibrary/domain/entity"
	"github.com/cooperlutz/go-full/test/fixtures"
	"github.com/cooperlutz/go-full/test/mocks"
)

func TestFindAllExamsWithQuestions(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := mocks.NewMockIExamLibraryRepository(t)
	mockPubSub := mocks.NewMockIPubSubEventProcessor(t)
	mockRepo.On("FindAllExams",
		mock.Anything,
	).Return(
		[]entity.Exam{
			fixtures.ValidDomainExamWithNoQuestions,
		},
		nil,
	)
	defer mockRepo.AssertExpectations(t)

	useCase := usecase.NewExamLibraryUseCase(
		mockRepo,
		mockPubSub,
	)

	// Act
	result, err := useCase.FindAllExamsWithQuestions(
		ctx,
		query.FindAllExamsWithQuestions{},
	)
	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result.Exams, 1)
}

func TestFindAllExamsWithQuestions_FailurePersist(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := mocks.NewMockIExamLibraryRepository(t)
	mockPubSub := mocks.NewMockIPubSubEventProcessor(t)
	mockRepo.On("FindAllExams",
		mock.Anything,
	).Return(
		[]entity.Exam{},
		assert.AnError,
	)
	defer mockRepo.AssertExpectations(t)

	useCase := usecase.NewExamLibraryUseCase(
		mockRepo,
		mockPubSub,
	)

	// Act
	result, err := useCase.FindAllExamsWithQuestions(
		ctx,
		query.FindAllExamsWithQuestions{},
	)
	// Assert
	assert.Error(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result.Exams, 0)
}
