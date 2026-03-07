package query

type Exam struct {
	ExamId              string
	State               string
	TotalPointsPossible int32
	TotalPointsReceived *int32
	Questions           []Question
}

type Question struct {
	QuestionId     string
	ExamId         string
	Index          int32
	QuestionType   string
	Graded         bool
	Feedback       *string
	ProvidedAnswer string
	CorrectAnswer  *string
	PointsPossible int32
	PointsReceived *int32
}
