package rest

import (
	"net/http"

	v1 "github.com/cooperlutz/go-full/internal/pingpong/api/rest/v1"
	v1_server "github.com/cooperlutz/go-full/internal/pingpong/api/rest/v1/server"
	"github.com/cooperlutz/go-full/internal/pingpong/app/usecase"
	"github.com/cooperlutz/go-full/pkg/hteeteepee"
)

// NewPingPongRouter sets up the routing for the PingPong API.
// The router will be mounted according to the service name defined in app/run.go
// according to the pattern `/{service-name}/api/v1`
//
// Currently, this is "/ping-pong", so the full path to the v1 API will be `/ping-pong/api/v1`
// Any handlers mounted will then be relative to that path, e.g. `/ping-pong/api/v1/ping`
func NewPingPongAPIRouter(uc usecase.IPingPongUseCase) http.Handler {
	/*
		Setup API Versions
	*/

	// Create a new Chi based router instance to be used by the PingPong API
	pingPongRouter := hteeteepee.NewRouter("pingpong.api")

	// Create the v1 controller
	controller := v1.NewRestAPIController(uc)

	// Create the v1 handler from the controller
	// We're using the "strict" handler which enforces request validation
	// and reduces boilerplate code in the controller methods.
	// So we create a handler from our custom controller
	v1Handler := v1_server.NewStrictHandler(controller, nil)

	// Additional versions of the Ping Pong API could be added here in the future
	// e.g., v2Handler := pingpong_api_v2.Handler(controller)

	// Finally, mount the versioned handlers onto the root router
	pingPongRouter.Mount("/api/v1", v1_server.Handler(v1Handler))

	// rootRouter.Mount("/api/v2", v2Handler) // Example for future versions
	return pingPongRouter
}
