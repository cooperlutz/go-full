package grading

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cooperlutz/go-full/internal/examlibrary/app/usecase"
	"github.com/cooperlutz/go-full/internal/grading/adapters/inbound"
	"github.com/cooperlutz/go-full/internal/grading/app"
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
		examLibraryUseCase,
	)
	if err != nil {
		return nil, err
	}

	apiServer := inbound.NewHttpServer(
		application,
	)

	inbound.RegisterEventHandlers(application.Events, pubSub)

	module := &GradingModule{
		RestApi: inbound.HandlerFromMux(
			apiServer.StrictHandler(),
			hteeteepee.NewRouter(
				"grading",
			),
		),
	}

	return module, nil
}
