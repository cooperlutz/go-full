package common

import "github.com/cooperlutz/go-full/internal/examlibrary/domain/entity"

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

// MapToDomainExamQuestion maps a common ExamQuestion to a domain ExamQuestion.
func MapToCommonExamQuestion(eq entity.ExamQuestion) ExamQuestion {
	return ExamQuestion{
		Index:           eq.GetIndex(),
		QuestionText:    eq.GetQuestionText(),
		QuestionType:    eq.GetQuestionType().String(),
		PossiblePoints:  eq.GetPossiblePoints(),
		CorrectAnswer:   eq.GetCorrectAnswer(),
		ResponseOptions: eq.GetResponseOptions(),
	}
}
