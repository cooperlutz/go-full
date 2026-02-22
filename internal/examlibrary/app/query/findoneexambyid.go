package query

import "github.com/cooperlutz/go-full/internal/examlibrary/app/common"

// FindOneExamByID represents the query to find one exam by its ID.
type FindOneExamByID struct {
	ExamID string
}

// FindOneExamByIDResponse represents the response for the FindOneExamByID query.
type FindOneExamByIDResponse struct {
	ExamID     string
	Name       string
	GradeLevel int
	TimeLimit  int64
	Questions  *[]common.ExamQuestion
}
