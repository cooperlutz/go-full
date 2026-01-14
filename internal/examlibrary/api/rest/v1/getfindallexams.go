package v1

import (
	"context"

	"go.opentelemetry.io/otel/trace"

	"github.com/cooperlutz/go-full/internal/examlibrary/api/rest/v1/mapper"
	"github.com/cooperlutz/go-full/internal/examlibrary/api/rest/v1/server"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/query"
)

// Get all Exams
// (GET /exams)
func (c *ExamLibraryRestAPIControllerV1) GetFindAllExams(ctx context.Context, request server.GetFindAllExamsRequestObject) (server.GetFindAllExamsResponseObject, error) {
	spanCtx := trace.SpanContextFromContext(ctx)

	qry := query.FindAllExamsWithoutQuestions{}

	qryResponse, err := c.UseCase.FindAllExamsWithoutQuestions(ctx, qry)
	if err != nil {
		return server.GetFindAllExams400Response{}, err
	}

	exams := mapper.FromAppExamsWithoutQuestionsToApiExamMetadataList(qryResponse.Exams)

	response := server.GetFindAllExams200JSONResponse{
		Body:    exams,
		Headers: server.GetFindAllExams200ResponseHeaders{XRequestId: spanCtx.TraceID().String()},
	}
	return response, nil
}
