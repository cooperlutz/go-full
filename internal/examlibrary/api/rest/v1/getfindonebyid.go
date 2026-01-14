package v1

import (
	"context"

	"go.opentelemetry.io/otel/trace"

	"github.com/cooperlutz/go-full/internal/examlibrary/api/rest/v1/mapper"
	"github.com/cooperlutz/go-full/internal/examlibrary/api/rest/v1/server"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/query"
)

// Get an Exam by ID
// (GET /exams/{examID})
func (c *ExamLibraryRestAPIControllerV1) GetFindOneByID(ctx context.Context, request server.GetFindOneByIDRequestObject) (server.GetFindOneByIDResponseObject, error) {
	spanCtx := trace.SpanContextFromContext(ctx)

	qry := query.FindOneExamByID{
		ExamID: request.ExamID.String(),
	}

	qryResponse, err := c.UseCase.FindOneExamByID(ctx, qry)
	if err != nil {
		return server.GetFindOneByID400Response{}, err
	}

	exam := mapper.FromAppFindOneExamByIDResponseToApiExam(qryResponse)

	response := server.GetFindOneByID200JSONResponse{
		Body:    exam,
		Headers: server.GetFindOneByID200ResponseHeaders{XRequestId: spanCtx.TraceID().String()},
	}

	return response, nil
}
