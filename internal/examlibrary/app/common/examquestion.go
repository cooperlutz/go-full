package common

// ExamQuestion represents a question in an exam.
type ExamQuestion struct {
	Index           int
	QuestionText    string
	QuestionType    string
	PossiblePoints  int
	CorrectAnswer   *string
	ResponseOptions *[]string
}

// NewExamQuestion creates a new ExamQuestion.
func NewExamQuestion(
	index int,
	questionText string,
	questionType string,
	possiblePoints int,
	correctAnswer *string,
	options *[]string,
) ExamQuestion {
	return ExamQuestion{
		Index:           index,
		QuestionText:    questionText,
		QuestionType:    questionType,
		PossiblePoints:  possiblePoints,
		CorrectAnswer:   correctAnswer,
		ResponseOptions: options,
	}
}
