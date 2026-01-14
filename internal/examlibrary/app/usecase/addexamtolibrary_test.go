package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/cooperlutz/go-full/internal/examlibrary/app/usecase"
	"github.com/cooperlutz/go-full/test/fixtures"
	"github.com/cooperlutz/go-full/test/mocks"
)

func TestAddExamToLibrary(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := mocks.NewMockIExamLibraryRepository(t)
	mockPubSub := mocks.NewMockIPubSubEventProcessor(t)
	mockRepo.On("SaveExam", mock.Anything,
		mock.AnythingOfType("entity.Exam"),
	).Return(nil)
	mockPubSub.On(
		"EmitEvent",
		mock.Anything,
		mock.AnythingOfType("event.ExamAddedToLibrary"),
	).Return(nil)
	defer mockRepo.AssertExpectations(t)
	defer mockPubSub.AssertExpectations(t)

	useCase := usecase.NewExamLibraryUseCase(
		mockRepo,
		mockPubSub,
	)
	cmd := fixtures.ValidAppCommandAddExamToLibrary
	// Act
	result, err := useCase.AddExamToLibrary(
		ctx,
		cmd,
	)
	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestAddExamToLibrary_FailureEmitEvent(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := mocks.NewMockIExamLibraryRepository(t)
	mockPubSub := mocks.NewMockIPubSubEventProcessor(t)
	mockRepo.On("SaveExam", mock.Anything,
		mock.AnythingOfType("entity.Exam"),
	).Return(nil)
	mockPubSub.On(
		"EmitEvent",
		mock.Anything,
		mock.AnythingOfType("event.ExamAddedToLibrary"),
	).Return(
		assert.AnError,
	)
	defer mockRepo.AssertExpectations(t)
	defer mockPubSub.AssertExpectations(t)

	useCase := usecase.NewExamLibraryUseCase(
		mockRepo,
		mockPubSub,
	)
	cmd := fixtures.ValidAppCommandAddExamToLibrary
	// Act
	result, err := useCase.AddExamToLibrary(
		ctx,
		cmd,
	)
	// Assert
	assert.Error(t, err)
	assert.NotNil(t, result)
}

func TestAddExamToLibrary_FailurePersist(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockRepo := mocks.NewMockIExamLibraryRepository(t)
	mockPubSub := mocks.NewMockIPubSubEventProcessor(t)
	mockRepo.On("SaveExam", mock.Anything,
		mock.AnythingOfType("entity.Exam"),
	).Return(
		assert.AnError,
	)

	defer mockRepo.AssertExpectations(t)

	useCase := usecase.NewExamLibraryUseCase(
		mockRepo,
		mockPubSub,
	)
	cmd := fixtures.ValidAppCommandAddExamToLibrary
	// Act
	result, err := useCase.AddExamToLibrary(
		ctx,
		cmd,
	)
	// Assert
	assert.Error(t, err)
	assert.NotNil(t, result)
}
