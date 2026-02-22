package mapper

import (
	"time"

	"github.com/cooperlutz/go-full/internal/examlibrary/app/command"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/common"
	"github.com/cooperlutz/go-full/internal/examlibrary/domain/entity"
	"github.com/cooperlutz/go-full/internal/examlibrary/domain/valueobject"
)

// FromAppQuestionTypeToDomainQuestionType maps an app QuestionType string to a domain QuestionType value object.
func FromAppQuestionTypeToDomainQuestionType(qt string) (valueobject.QuestionType, error) {
	return valueobject.QuestionTypeFromString(qt)
}

// FromAppExamQuestionToDomainExamQuestion maps an app ExamQuestion to a domain ExamQuestion.
func FromAppExamQuestionToDomainExamQuestion(eq common.ExamQuestion) (entity.ExamQuestion, error) {
	questionType, err := FromAppQuestionTypeToDomainQuestionType(eq.QuestionType)
	if err != nil {
		return entity.ExamQuestion{}, err
	}

	newQuestion := *entity.NewExamQuestion(
		eq.Index,
		eq.QuestionText,
		questionType,
		eq.PossiblePoints,
		eq.CorrectAnswer,
		eq.ResponseOptions,
	)

	return newQuestion, nil
}

// FromAppExamQuestionsToDomainExamQuestions maps a slice of app ExamQuestions to a slice of domain ExamQuestions.
func FromAppExamQuestionsToDomainExamQuestions(eqs []common.ExamQuestion) ([]entity.ExamQuestion, error) {
	var questions []entity.ExamQuestion

	for _, eq := range eqs {
		domainQuestion, err := FromAppExamQuestionToDomainExamQuestion(eq)
		if err != nil {
			return nil, err
		}

		questions = append(questions, domainQuestion)
	}

	return questions, nil
}

// FromAppAddExamToLibraryToDomainExam maps an app AddExamToLibrary command to a domain Exam.
func FromAppAddExamToLibraryToDomainExam(cmd command.AddExamToLibrary) (entity.Exam, error) {
	domainQuestions, err := FromAppExamQuestionsToDomainExamQuestions(cmd.Questions)
	if err != nil {
		return entity.Exam{}, err
	}

	newExam := entity.NewExam(
		cmd.Name,
		valueobject.GradeLevel(cmd.GradeLevel),
		time.Duration(cmd.TimeLimit)*time.Second,
		&domainQuestions,
	)

	return *newExam, nil
}
