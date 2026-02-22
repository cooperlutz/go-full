package query

import "github.com/cooperlutz/go-full/internal/examlibrary/app/common"

// FindAllExamsWithQuestions represents the query to find all exams with their questions.
type FindAllExamsWithQuestions struct{}

// FindAllExamsWithQuestionsResponse represents the response for the FindAllExamsWithQuestions query.
type FindAllExamsWithQuestionsResponse struct {
	Exams []ExamWithQuestions
}

// ExamWithQuestions represents an exam with its questions.
type ExamWithQuestions struct {
	ExamID     string
	Name       string
	GradeLevel int
	TimeLimit  int64
	Questions  []common.ExamQuestion
}
