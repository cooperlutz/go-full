/*
Package examination contains the module associated with the Examination domain.
the examination module is responsible for managing the lifecycle of exams throughout the test-taking process

The Examination module adopts a hexagonal architecture pattern
The Examination module prioritizes and focuses on e2e tests, rather than unit tests
*/
package examination

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cooperlutz/go-full/internal/examination/adapters/inbound"
	"github.com/cooperlutz/go-full/internal/examination/app"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/usecase"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/hteeteepee"
	"github.com/cooperlutz/go-full/pkg/workerbee"
)

type ExaminationModule struct {
	RestApi http.Handler
}

// NewModule - Initializes the Examination module with its needed dependencies.
func NewModule(
	pgconn *pgxpool.Pool,
	pubSub *eeventdriven.BasePgsqlPubSubProcessor,
	examLibraryUseCase usecase.IExamLibraryUseCase,
	backgroundWorker *workerbee.Worker,
) (*ExaminationModule, error) {
	application, err := app.NewApplication(
		pgconn,
		pubSub,
		examLibraryUseCase,
	)
	if err != nil {
		return nil, err
	}

	apiServer := inbound.NewHttpServer(application)

	inbound.NewTaskWorkerAdapter(backgroundWorker, application)

	inbound.NewSqlSubscriberAdapter(pubSub).RegisterEventHandlers()

	module := &ExaminationModule{
		RestApi: inbound.HandlerFromMux(
			apiServer.StrictHandler(),
			hteeteepee.NewRouter(
				"examination",
			),
		),
	}

	return module, nil
}
