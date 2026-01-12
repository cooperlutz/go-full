package query

import "github.com/cooperlutz/go-full/internal/examlibrary/app/common"

type FindOneExamByID struct {
	ExamID string
}

type FindOneExamByIDResponse struct {
	ExamID     string
	Name       string
	GradeLevel int
	Questions  *[]common.ExamQuestion
}
