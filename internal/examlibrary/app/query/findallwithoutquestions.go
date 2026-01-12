package query

// FindAllExamsWithoutQuestions represents the query to find all exams without their questions.
type FindAllExamsWithoutQuestions struct{}

// NewFindAllExamsWithoutQuestions creates a new FindAllExamsWithoutQuestions query.
type FindAllExamsWithoutQuestionsResponse struct {
	Exams []ExamWithoutQuestions
}

// ExamWithoutQuestions represents an exam without its questions.
type ExamWithoutQuestions struct {
	ExamID     string
	Name       string
	GradeLevel int
}
