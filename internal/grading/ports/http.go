package ports

import (
	"context"

	"github.com/cooperlutz/go-full/internal/grading/app"
	"github.com/cooperlutz/go-full/internal/grading/app/command"
	"github.com/cooperlutz/go-full/internal/grading/app/event"
	"github.com/cooperlutz/go-full/internal/grading/app/query"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

// HttpServer represents the HTTP server for the Grading module.
type HttpServer struct {
	app app.Application
}

// NewHttpServer creates a new HttpServer instance with the provided Grading application.
func NewHttpServer(application app.Application) HttpServer {
	return HttpServer{
		app: application,
	}
}

// StrictHandler returns a strict HTTP handler for the Grading module.
func (h HttpServer) StrictHandler() ServerInterface {
	return NewStrictHandler(h, nil)
}

// (GET /v1/exams/{examId}).
func (h HttpServer) GetExam(ctx context.Context, request GetExamRequestObject) (GetExamResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "grading.adapters.inbound.http.get_exam")
	defer span.End()

	exam, err := h.app.Queries.FindExam.Handle(ctx, request.ExamId)
	if err != nil {
		return nil, err
	}

	return GetExam200JSONResponse(queryExamToHttpExam(exam)), nil
}

// (GET /v1/exams/{examId}/questions/{questionIndex}).
func (h HttpServer) GetExamQuestion(ctx context.Context, request GetExamQuestionRequestObject) (GetExamQuestionResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "grading.adapters.inbound.http.get_exam_question")
	defer span.End()

	question, err := h.app.Queries.FindExamQuestion.Handle(
		ctx,
		request.ExamId,
		request.QuestionIndex,
	)
	if err != nil {
		return nil, err
	}

	return GetExamQuestion200JSONResponse(
		queryQuestionToHttpQuestion(question),
	), nil
}

// (PATCH /v1/exams/{examId}/questions/{questionIndex}/grade).
func (h HttpServer) GradeExamQuestion(ctx context.Context, request GradeExamQuestionRequestObject) (GradeExamQuestionResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "grading.adapters.inbound.http.grade_exam_question")
	defer span.End()

	err := h.app.Commands.GradeQuestion.Handle(
		ctx,
		command.GradeQuestion{
			ExamId:        request.ExamId,
			QuestionIndex: request.QuestionIndex,
			Points:        request.Body.Points,
			Feedback:      request.Body.Feedback,
		},
	)
	if err != nil {
		return nil, err
	}

	exam, err := h.app.Queries.FindExam.Handle(ctx, request.ExamId)
	if err != nil {
		return nil, err
	}

	if exam.State == "completed" {
		err = h.app.Events.GradingCompleted.Handle(ctx, event.GradingCompleted{ExamId: exam.ExamId})
		if err != nil {
			return nil, err
		}
	}

	return GradeExamQuestion201JSONResponse{}, nil
}

// (GET /v1/exams/ungraded).
func (h HttpServer) GetFindIncompleteExams(ctx context.Context, request GetFindIncompleteExamsRequestObject) (GetFindIncompleteExamsResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "grading.adapters.inbound.http.get_find_incomplete_exams")
	defer span.End()

	exams, err := h.app.Queries.FindIncompleteExams.Handle(ctx)
	if err != nil {
		return nil, err
	}

	var httpExams []Exam
	for _, exam := range exams {
		httpExams = append(httpExams, queryExamToHttpExam(exam))
	}

	return GetFindIncompleteExams200JSONResponse(httpExams), nil
}

func queryExamToHttpExam(exam query.Exam) Exam {
	var questions []Question
	for _, q := range exam.Questions {
		questions = append(questions, queryQuestionToHttpQuestion(q))
	}

	return Exam{
		ExamId:              exam.ExamId,
		State:               exam.State,
		TotalPointsPossible: exam.TotalPointsPossible,
		TotalPointsEarned:   exam.TotalPointsReceived,
		Questions:           questions,
	}
}

func queryQuestionToHttpQuestion(question query.Question) Question {
	return Question{
		ExamId:         question.ExamId,
		Feedback:       question.Feedback,
		Graded:         question.Graded,
		PointsEarned:   question.PointsReceived,
		PointsPossible: &question.PointsPossible,
		ProvidedAnswer: &question.ProvidedAnswer,
		QuestionId:     question.QuestionId,
		QuestionIndex:  question.Index,
		QuestionType:   question.QuestionType,
	}
}
