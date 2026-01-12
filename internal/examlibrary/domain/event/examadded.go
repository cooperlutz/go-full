package event

// ExamAddedToLibrary is a domain event that signifies an exam has been added to the library.
type ExamAddedToLibrary struct {
	ExamID string
}
