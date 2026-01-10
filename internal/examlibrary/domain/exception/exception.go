package exception

// ErrQuestionNotFound is returned when a question is not found.
type ErrQuestionNotFound struct{}

// Error implements the error interface for ErrQuestionNotFound.
func (e ErrQuestionNotFound) Error() string {
	return "question not found"
}

// ErrInvalidIndex is returned when an invalid index is used.
type ErrInvalidIndex struct{}

// Error implements the error interface for ErrInvalidIndex.
func (e ErrInvalidIndex) Error() string {
	return "invalid index"
}

// ErrInvalidQuestionType is returned when an invalid question type is used.
type ErrInvalidQuestionType struct{}

// Error implements the error interface for ErrInvalidQuestionType.
func (e ErrInvalidQuestionType) Error() string {
	return "invalid question type"
}
