package mapper

import (
	"github.com/cooperlutz/go-full/internal/examlibrary/app/command"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/common"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/query"
	"github.com/cooperlutz/go-full/internal/examlibrary/domain/entity"
)

// FromDomainExamQuestionToAppExamQuestion maps a domain ExamQuestion to an app ExamQuestion.
func FromDomainExamQuestionToAppExamQuestion(eq entity.ExamQuestion) common.ExamQuestion {
	return common.NewExamQuestion(
		eq.GetIndex(),
		eq.GetQuestionText(),
		eq.GetQuestionType().String(),
		eq.GetPossiblePoints(),
		eq.GetCorrectAnswer(),
		eq.GetResponseOptions(),
	)
}

// FromDomainExamQuestionsToAppExamQuestions maps a slice of domain ExamQuestions to a slice of app ExamQuestions.
func FromDomainExamQuestionsToAppExamQuestions(eqs []entity.ExamQuestion) []common.ExamQuestion {
	var questions []common.ExamQuestion
	for _, eq := range eqs {
		questions = append(questions, FromDomainExamQuestionToAppExamQuestion(eq))
	}

	return questions
}

// FromDomainExamToAppAddExamToLibraryResult maps a domain Exam to an app AddExamToLibraryResult.
func FromDomainExamToAppAddExamToLibraryResult(exam entity.Exam) command.AddExamToLibraryResult {
	if len(exam.GetQuestions()) <= 0 {
		return command.NewAddExamToLibraryResult(
			exam.GetIdString(),
			exam.GetName(),
			exam.GetGradeLevel().Int(),
			[]common.ExamQuestion{},
		)
	}

	var questions []common.ExamQuestion
	for _, eq := range exam.GetQuestions() {
		questions = append(questions, FromDomainExamQuestionToAppExamQuestion(eq))
	}

	return command.NewAddExamToLibraryResult(
		exam.GetIdString(),
		exam.GetName(),
		exam.GetGradeLevel().Int(),
		questions,
	)
}

// FromDomainExamToAppFindOneExamByIDResponse maps a domain Exam to an app FindOneExamByIDResponse.
func FromDomainExamToAppFindOneExamByIDResponse(exam entity.Exam) query.FindOneExamByIDResponse {
	var questions []common.ExamQuestion

	if len(exam.GetQuestions()) > 0 {
		for _, eq := range exam.GetQuestions() {
			questions = append(questions, FromDomainExamQuestionToAppExamQuestion(eq))
		}
	}

	return query.FindOneExamByIDResponse{
		ExamID:     exam.GetIdString(),
		Name:       exam.GetName(),
		GradeLevel: exam.GetGradeLevel().Int(),
		Questions:  &questions,
	}
}

// FromDomainExamsToAppFindAllExamsWithoutQuestionsResponse maps a slice of domain Exams to an app FindAllExamsWithoutQuestionsResponse.
func FromDomainExamsToAppFindAllExamsWithoutQuestionsResponse(exams []entity.Exam) query.FindAllExamsWithoutQuestionsResponse {
	var results []query.ExamWithoutQuestions
	for _, exam := range exams {
		results = append(results, query.ExamWithoutQuestions{
			ExamID:     exam.GetIdString(),
			Name:       exam.GetName(),
			GradeLevel: exam.GetGradeLevel().Int(),
		})
	}

	return query.FindAllExamsWithoutQuestionsResponse{
		Exams: results,
	}
}

// FromDomainExamsToAppFindAllExamsWithQuestionsResponse maps a slice of domain Exams to an app FindAllExamsWithQuestionsResponse.
func FromDomainExamsToAppFindAllExamsWithQuestionsResponse(exams []entity.Exam) query.FindAllExamsWithQuestionsResponse {
	var results []query.ExamWithQuestions

	for _, exam := range exams {
		var questions []common.ExamQuestion

		if len(exam.GetQuestions()) > 0 {
			for _, eq := range exam.GetQuestions() {
				questions = append(questions, FromDomainExamQuestionToAppExamQuestion(eq))
			}
		}

		results = append(results, query.ExamWithQuestions{
			ExamID:     exam.GetIdString(),
			Name:       exam.GetName(),
			GradeLevel: exam.GetGradeLevel().Int(),
			Questions:  questions,
		})
	}

	return query.FindAllExamsWithQuestionsResponse{
		Exams: results,
	}
}
