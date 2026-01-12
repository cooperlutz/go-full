package query

import "github.com/cooperlutz/go-full/internal/examlibrary/app/common"

type FindAllExamsWithQuestions struct{}

type FindAllExamsWithQuestionsResponse struct {
	Exams []ExamWithQuestions
}

type ExamWithQuestions struct {
	ExamID     string
	Name       string
	GradeLevel int
	Questions  []common.ExamQuestion
}
