package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/cooperlutz/go-full/internal/examlibrary/app/query"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/usecase"
	"github.com/cooperlutz/go-full/test/fixtures"
	"github.com/cooperlutz/go-full/test/mocks"
)

func TestFindOneExamByID(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := mocks.NewMockIExamLibraryRepository(t)
	mockPubSub := mocks.NewMockIPubSubEventProcessor(t)
	mockRepo.On("FindExamByID",
		mock.Anything,
		fixtures.ValidUUID,
	).Return(
		fixtures.ValidDomainExam,
		nil,
	)
	defer mockRepo.AssertExpectations(t)

	useCase := usecase.NewExamLibraryUseCase(
		mockRepo,
		mockPubSub,
	)

	// Act
	result, err := useCase.FindOneExamByID(
		ctx,
		query.FindOneExamByID{
			ExamID: fixtures.ValidUUID.String(),
		},
	)
	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, fixtures.ValidAppFindOneExamByIDResponse.GradeLevel, result.GradeLevel)
	assert.Equal(t, fixtures.ValidAppFindOneExamByIDResponse.Name, result.Name)
	assert.IsType(t, fixtures.ValidAppFindOneExamByIDResponse.ExamID, result.ExamID)
}

func TestFindOneExamByID_FailurePersist(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := mocks.NewMockIExamLibraryRepository(t)
	mockPubSub := mocks.NewMockIPubSubEventProcessor(t)
	mockRepo.On("FindExamByID",
		mock.Anything,
		fixtures.ValidUUID,
	).Return(
		fixtures.ValidDomainExam,
		assert.AnError,
	)
	defer mockRepo.AssertExpectations(t)

	useCase := usecase.NewExamLibraryUseCase(
		mockRepo,
		mockPubSub,
	)

	// Act
	result, err := useCase.FindOneExamByID(
		ctx,
		query.FindOneExamByID{
			ExamID: fixtures.ValidUUID.String(),
		},
	)
	// Assert
	assert.Error(t, err)
	assert.NotNil(t, result)
}
