package inbound

import (
	"context"

	"github.com/cooperlutz/go-full/internal/examination/app"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(application app.Application) HttpServer {
	return HttpServer{
		app: application,
	}
}

func (h HttpServer) StrictHandler() ServerInterface {
	return NewStrictHandler(h, nil)
}

// (GET /api/v1/exams/available).
func (h HttpServer) GetAvailableExams(ctx context.Context, request GetAvailableExamsRequestObject) (GetAvailableExamsResponseObject, error) {
	exams, err := h.app.Queries.AvailableExams.Handle(ctx)
	if err != nil {
		return nil, err
	}

	var responseExams []Exam
	for _, e := range exams {
		responseExams = append(responseExams, Exam{
			ExamId:    &e.ExamId,
			StudentId: &e.StudentId,
		})
	}

	return GetAvailableExams200JSONResponse(responseExams), nil
}
