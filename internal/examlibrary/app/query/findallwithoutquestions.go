package query

type FindAllExamsWithoutQuestions struct{}

type FindAllExamsWithoutQuestionsResponse struct {
	Exams []ExamWithoutQuestions
}

type ExamWithoutQuestions struct {
	ExamID     string
	Name       string
	GradeLevel int
}
