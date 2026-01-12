package command

import "github.com/cooperlutz/go-full/internal/examlibrary/app/common"

type AddExamToLibrary struct {
	Name       string
	GradeLevel int
	Questions  []common.ExamQuestion
}

func NewAddExamToLibrary(name string, gradeLevel int, questions []common.ExamQuestion) AddExamToLibrary {
	return AddExamToLibrary{
		Name:       name,
		GradeLevel: gradeLevel,
		Questions:  questions,
	}
}

type AddExamToLibraryResult struct {
	ExamID     string
	Name       string
	GradeLevel int
	Questions  []common.ExamQuestion
}

func NewAddExamToLibraryResult(examID, name string, gradeLevel int, questions []common.ExamQuestion) AddExamToLibraryResult {
	return AddExamToLibraryResult{
		ExamID:     examID,
		Name:       name,
		GradeLevel: gradeLevel,
		Questions:  questions,
	}
}
