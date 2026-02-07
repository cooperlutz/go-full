package rest

import (
	"net/http"

	v1 "github.com/cooperlutz/go-full/internal/examlibrary/api/rest/v1"
	v1_server "github.com/cooperlutz/go-full/internal/examlibrary/api/rest/v1/server"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/usecase"
	"github.com/cooperlutz/go-full/pkg/hteeteepee"
)

// NewExamLibraryRouter sets up the routing for the ExamLibrary API.
// The router will be mounted according to the service name defined in app/run.go
// according to the pattern `/{service-name}/api/v1`
func NewExamLibraryAPIRouter(uc usecase.IExamLibraryUseCase) http.Handler {
	// Create a new Chi based router instance to be used by the ExamLibrary API
	examLibraryRouter := hteeteepee.NewRouter("examlibrary.api")

	// Create the v1 controller
	controller := v1.NewRestAPIController(uc)

	// Create the v1 handler from the controller
	// We're using the "strict" handler which enforces request validation
	// and reduces boilerplate code in the controller methods.
	// So we create a handler from our custom controller
	v1Handler := v1_server.NewStrictHandler(controller, nil)

	// Additional versions of the ExamLibrary API could be added here in the future
	// e.g., v2Handler := examlibrary_api_v2.Handler(controller)

	// Finally, mount the versioned handlers onto the root router
	examLibraryRouter.Mount("/v1", v1_server.Handler(v1Handler))

	// rootRouter.Mount("/api/v2", v2Handler) // Example for future versions
	return examLibraryRouter
}
