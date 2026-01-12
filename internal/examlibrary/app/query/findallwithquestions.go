package query

type FindAllExamsWithQuestions struct{}

type FindAllExamsWithQuestionsResponse struct {
	Exams []ExamWithQuestions
}

type ExamWithQuestions struct {
	ExamID    string
	Questions []Question
}

type Question struct {
	QuestionID string
	Content    string
}
