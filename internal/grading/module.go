/*
Package grading contains the module associated with the Grading domain.
The grading module is responsible for managing the lifecycle of exams throughout the test-taking process.

The Grading module adopts a ports and adapters architecture pattern.
*/
package grading

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cooperlutz/go-full/internal/examlibrary/app/usecase"
	"github.com/cooperlutz/go-full/internal/grading/app"
	"github.com/cooperlutz/go-full/internal/grading/ports"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/hteeteepee"
)

type GradingModule struct {
	RestApi http.Handler
}

// NewModule - Initializes the Grading module with its needed dependencies.
func NewModule(
	pgconn *pgxpool.Pool,
	pubSub *eeventdriven.BasePgsqlPubSubProcessor,
	examLibraryUseCase usecase.IExamLibraryUseCase,
) (*GradingModule, error) {
	application, err := app.NewApplication(
		pgconn,
		pubSub,
		examLibraryUseCase,
	)
	if err != nil {
		return nil, err
	}

	apiServer := ports.NewHttpServer(
		application,
	)

	ports.RegisterEventHandlers(application.Events, pubSub)

	module := &GradingModule{
		RestApi: ports.HandlerFromMux(
			apiServer.StrictHandler(),
			hteeteepee.NewRouter(
				"grading",
			),
		),
	}

	return module, nil
}
