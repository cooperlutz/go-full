package mapper

import (
	"github.com/cooperlutz/go-full/internal/examlibrary/api/rest/v1/server"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/command"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/query"
)

func FromAppExamWithoutQuestionsToApiExamMetadata(e query.ExamWithoutQuestions) server.ExamMetadata {
	return server.ExamMetadata{
		Id:         &e.ExamID,
		Name:       &e.Name,
		GradeLevel: &e.GradeLevel,
		TimeLimit:  &e.TimeLimit,
	}
}

func FromAppFindOneExamByIDResponseToApiExam(e query.FindOneExamByIDResponse) server.Exam {
	var questions []server.ExamQuestion
	for _, q := range *e.Questions {
		questions = append(questions, FromAppExamQuestionToApiExamQuestion(q))
	}
	exam := server.Exam{
		Id:         &e.ExamID,
		Name:       &e.Name,
		GradeLevel: &e.GradeLevel,
		TimeLimit:  &e.TimeLimit,
		Questions:  &questions,
	}
	return exam
}

func FromAppExamsWithoutQuestionsToApiExamMetadataList(exams []query.ExamWithoutQuestions) []server.ExamMetadata {
	var examMetadataList []server.ExamMetadata
	for _, e := range exams {
		examMetadataList = append(examMetadataList, FromAppExamWithoutQuestionsToApiExamMetadata(e))
	}
	return examMetadataList
}

func FromAppAddExamToLibraryResultToApiExam(cmd command.AddExamToLibraryResult) server.Exam {
	var questions []server.ExamQuestion
	for _, eq := range cmd.Questions {
		questions = append(questions, FromAppExamQuestionToApiExamQuestion(eq))
	}
	exam := server.Exam{
		Id:         &cmd.ExamID,
		Name:       &cmd.Name,
		GradeLevel: &cmd.GradeLevel,
		TimeLimit:  &cmd.TimeLimit,
		Questions:  &questions,
	}
	return exam
}
