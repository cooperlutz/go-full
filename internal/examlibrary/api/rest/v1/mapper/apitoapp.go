package mapper

import (
	"github.com/cooperlutz/go-full/internal/examlibrary/api/rest/v1/server"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/command"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/common"
)

func FromApiExamQuestionToAppExamQuestion(eq server.ExamQuestion) common.ExamQuestion {
	return common.NewExamQuestion(
		*eq.Index,
		*eq.QuestionText,
		string(*eq.QuestionType),
		*eq.PossiblePoints,
		eq.CorrectAnswer,
		eq.PossibleAnswers,
	)
}

func FromApiExamQuestionsToAppExamQuestions(eqs []server.ExamQuestion) []common.ExamQuestion {
	var questions []common.ExamQuestion
	for _, eq := range eqs {
		questions = append(questions, FromApiExamQuestionToAppExamQuestion(eq))
	}
	return questions
}

func FromAppExamQuestionToApiExamQuestion(eq common.ExamQuestion) server.ExamQuestion {
	questionType := server.QuestionType(eq.QuestionType)
	return server.ExamQuestion{
		Index:           &eq.Index,
		QuestionText:    &eq.QuestionText,
		QuestionType:    &questionType,
		PossiblePoints:  &eq.PossiblePoints,
		CorrectAnswer:   eq.CorrectAnswer,
		PossibleAnswers: eq.ResponseOptions,
	}
}

func FromAppExamQuestionsToApiExamQuestions(eqs []common.ExamQuestion) []server.ExamQuestion {
	var questions []server.ExamQuestion
	for _, eq := range eqs {
		questions = append(questions, FromAppExamQuestionToApiExamQuestion(eq))
	}
	return questions
}

func FromApiExamToAppAddExamToLibrary(exam server.Exam) (command.AddExamToLibrary, error) {
	if exam.Name == nil || exam.GradeLevel == nil || exam.Questions == nil {
		return command.AddExamToLibrary{}, &server.RequiredParamError{ParamName: "Name, GradeLevel, and Questions are required"}
	}
	questions := FromApiExamQuestionsToAppExamQuestions(*exam.Questions)
	return command.NewAddExamToLibrary(
		*exam.Name,
		*exam.GradeLevel,
		questions,
	), nil
}
