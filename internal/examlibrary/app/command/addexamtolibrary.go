package command

import "github.com/cooperlutz/go-full/internal/examlibrary/app/common"

// AddExamToLibrary represents the command to add a new exam to the library.
type AddExamToLibrary struct {
	Name       string
	GradeLevel int
	TimeLimit  int64
	Questions  []common.ExamQuestion
}

// NewAddExamToLibrary creates a new AddExamToLibrary command.
func NewAddExamToLibrary(name string, gradeLevel int, timeLimit int64, questions []common.ExamQuestion) AddExamToLibrary {
	return AddExamToLibrary{
		Name:       name,
		GradeLevel: gradeLevel,
		TimeLimit:  timeLimit,
		Questions:  questions,
	}
}

// AddExamToLibraryResult represents the result of adding a new exam to the library.
type AddExamToLibraryResult struct {
	ExamID     string
	Name       string
	GradeLevel int
	TimeLimit  int64
	Questions  []common.ExamQuestion
}

// NewAddExamToLibraryResult creates a new AddExamToLibraryResult.
func NewAddExamToLibraryResult(examID, name string, gradeLevel int, timeLimit int64, questions []common.ExamQuestion) AddExamToLibraryResult {
	return AddExamToLibraryResult{
		ExamID:     examID,
		Name:       name,
		GradeLevel: gradeLevel,
		TimeLimit:  timeLimit,
		Questions:  questions,
	}
}
