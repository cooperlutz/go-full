package v1

import (
	"context"

	"go.opentelemetry.io/otel/trace"

	"github.com/cooperlutz/go-full/internal/examlibrary/api/rest/v1/mapper"
	"github.com/cooperlutz/go-full/internal/examlibrary/api/rest/v1/server"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

// Add a new Exam to the Library
// (POST /exams)
func (c *ExamLibraryRestAPIControllerV1) PostAddExamToLibrary(ctx context.Context, request server.PostAddExamToLibraryRequestObject) (server.PostAddExamToLibraryResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "examlibrary.api.rest.v1.post_add_exam_to_library")
	defer span.End()
	spanCtx := trace.SpanContextFromContext(ctx)

	cmd, err := mapper.FromApiExamToAppAddExamToLibrary(*request.Body)
	if err != nil {
		return server.PostAddExamToLibrary400Response{}, err
	}

	cmdResult, err := c.UseCase.AddExamToLibrary(ctx, cmd)
	if err != nil {
		return server.PostAddExamToLibrary400Response{}, err
	}

	exam := mapper.FromAppAddExamToLibraryResultToApiExam(cmdResult)

	response := server.PostAddExamToLibrary200JSONResponse{
		Body:    exam,
		Headers: server.PostAddExamToLibrary200ResponseHeaders{XRequestId: spanCtx.TraceID().String()},
	}
	return response, nil
}
