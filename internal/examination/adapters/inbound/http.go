package inbound

import (
	"context"

	"github.com/cooperlutz/go-full/internal/examination/app"
	"github.com/cooperlutz/go-full/internal/examination/app/command"
	"github.com/cooperlutz/go-full/internal/examination/app/event"
	"github.com/cooperlutz/go-full/internal/examination/app/query"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

// HttpServer represents the HTTP server for the Examination module.
type HttpServer struct {
	app app.Application
}

// NewHttpServer creates a new HttpServer instance with the provided Examination application.
func NewHttpServer(application app.Application) HttpServer {
	return HttpServer{
		app: application,
	}
}

// StrictHandler returns a strict HTTP handler for the Examination module.
func (h HttpServer) StrictHandler() ServerInterface {
	return NewStrictHandler(h, nil)
}

// (GET /v1/exams).
func (h HttpServer) GetAvailableExams(ctx context.Context, request GetAvailableExamsRequestObject) (GetAvailableExamsResponseObject, error) {
	exams, err := h.app.Queries.AvailableExams.Handle(ctx)
	if err != nil {
		return nil, err
	}

	var responseExams []Exam
	for _, e := range exams {
		responseExams = append(responseExams, queryExamToHttpExam(e))
	}

	return GetAvailableExams200JSONResponse(responseExams), nil
}

// (POST /v1/exams)
func (h HttpServer) StartNewExam(ctx context.Context, request StartNewExamRequestObject) (StartNewExamResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.inbound.http.startnewexam")
	defer span.End()

	exam, err := h.app.Commands.StartExam.Handle(ctx, command.StartExam{
		StudentId:     request.Body.StudentId,
		ExamLibraryID: request.Body.LibraryExamId,
	})
	if err != nil {
		return nil, err
	}

	err = h.app.Events.ExamStarted.Handle(ctx, event.ExamStarted{
		ExamID:    exam.ExamId,
		StudentID: exam.StudentId,
	})
	if err != nil {
		return nil, err
	}

	return StartNewExam201JSONResponse{
		ExamId:    exam.ExamId,
		StudentId: exam.StudentId,
	}, nil
}

// (POST /v1/exams/{examId}/questions/{questionIndex})
func (h HttpServer) AnswerQuestion(ctx context.Context, request AnswerQuestionRequestObject) (AnswerQuestionResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.inbound.http.answerquestion")
	defer span.End()

	err := h.app.Commands.AnswerQuestion.Handle(ctx, command.AnswerQuestion{
		ExamID:        request.ExamId,
		QuestionIndex: request.QuestionIndex,
		Answer:        request.Body.ProvidedAnswer,
	})
	if err != nil {
		return nil, err
	}

	return AnswerQuestion200JSONResponse{}, nil
}

// (GET /v1/exams/{examId}/questions/{questionIndex})
func (h HttpServer) GetExamQuestion(ctx context.Context, request GetExamQuestionRequestObject) (GetExamQuestionResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.inbound.http.getexamquestion")
	defer span.End()

	question, err := h.app.Queries.FindQuestion.Handle(ctx, query.FindQuestion{
		ExamID:        request.ExamId,
		QuestionIndex: request.QuestionIndex,
	})
	if err != nil {
		return nil, err
	}

	questionResponse := queryQuestionToHttpQuestion(question)

	return GetExamQuestion200JSONResponse(questionResponse), nil
}

// (GET /v1/exams/{examId}/progress)
func (h HttpServer) GetExamProgress(ctx context.Context, request GetExamProgressRequestObject) (GetExamProgressResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.inbound.http.getexamprogress")
	defer span.End()

	exam, err := h.app.Queries.FindExam.Handle(ctx, query.FindExam{
		ExamID: request.ExamId,
	})
	if err != nil {
		return nil, err
	}

	return GetExamProgress200JSONResponse{
		AnsweredQuestions: exam.AnsweredQuestions,
		TotalQuestions:    exam.TotalQuestions,
	}, nil
}

// (GET /v1/exams/{examId})
func (h HttpServer) GetExam(ctx context.Context, request GetExamRequestObject) (GetExamResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.inbound.http.getexam")
	defer span.End()

	exam, err := h.app.Queries.FindExam.Handle(ctx, query.FindExam{
		ExamID: request.ExamId,
	})
	if err != nil {
		return nil, err
	}

	examResponse := queryExamToHttpExam(exam)

	return GetExam200JSONResponse(examResponse), nil
}

// (POST /v1/exams/{examId}/submit)
func (h HttpServer) SubmitExam(ctx context.Context, request SubmitExamRequestObject) (SubmitExamResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.inbound.http.submitexam")
	defer span.End()

	err := h.app.Commands.SubmitExam.Handle(ctx, command.SubmitExam{
		ExamID: request.ExamId,
	})
	if err != nil {
		return nil, err
	}

	err = h.app.Events.ExamSubmitted.Handle(ctx, event.ExamSubmitted{
		ExamID: request.ExamId,
	})
	if err != nil {
		return nil, err
	}

	return SubmitExam200Response{}, nil
}

func queryExamToHttpExam(e query.Exam) Exam {
	var questions []Question
	for _, q := range e.Questions {
		questions = append(questions, queryQuestionToHttpQuestion(q))
	}

	return Exam{
		ExamId:            e.ExamId,
		LibraryExamId:     &e.LibraryExamId,
		StudentId:         e.StudentId,
		Completed:         e.Completed,
		AnsweredQuestions: &e.AnsweredQuestions,
		TotalQuestions:    &e.TotalQuestions,
		Questions:         &questions,
	}
}

func queryQuestionToHttpQuestion(q query.Question) Question {
	return Question{
		QuestionId:      q.QuestionID,
		QuestionIndex:   q.QuestionIndex,
		ExamId:          q.ExamId,
		Answered:        q.Answered,
		QuestionText:    q.QuestionText,
		QuestionType:    q.QuestionType,
		ResponseOptions: q.ResponseOptions,
		ProvidedAnswer:  q.ProvidedAnswer,
	}
}
