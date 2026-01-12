package common

import "github.com/cooperlutz/go-full/internal/examlibrary/domain/entity"

type ExamQuestion struct {
	Index           int
	QuestionText    string
	QuestionType    string
	PossiblePoints  int
	CorrectAnswer   *string
	ResponseOptions *[]string
}

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
