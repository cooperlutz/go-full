package query

type Exam struct {
	ExamId            string
	LibraryExamId     string
	StudentId         string
	State             string
	AnsweredQuestions int32
	TotalQuestions    int32
	Questions         []Question
}

type Question struct {
	ExamId          string
	Answered        bool
	QuestionID      string
	QuestionIndex   int32
	QuestionText    string
	QuestionType    string
	ResponseOptions *[]string
	ProvidedAnswer  *string
}
