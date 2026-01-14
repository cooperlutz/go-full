package v1

import (
	"github.com/cooperlutz/go-full/internal/examlibrary/api/rest/v1/server"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/usecase"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ server.StrictServerInterface = (*ExamLibraryRestAPIControllerV1)(nil)

// ExamLibraryRestAPIControllerV1 is the controller for the ExamLibrary API
type ExamLibraryRestAPIControllerV1 struct {
	UseCase usecase.IExamLibraryUseCase
}

// NewRestAPIController creates a new ExamLibraryRestAPIControllerV1
func NewRestAPIController(uc usecase.IExamLibraryUseCase) *ExamLibraryRestAPIControllerV1 {
	return &ExamLibraryRestAPIControllerV1{
		UseCase: uc,
	}
}
