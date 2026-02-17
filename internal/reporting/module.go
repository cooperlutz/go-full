package reporting

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cooperlutz/go-full/internal/reporting/adapters/inbound"
	"github.com/cooperlutz/go-full/internal/reporting/app"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/hteeteepee"
)

type ReportingModule struct {
	RestApi http.Handler
}

// NewModule - Initializes the Reporting module with its needed dependencies.
func NewModule(
	pgConn *pgxpool.Pool,
	pubSub *eeventdriven.BasePgsqlPubSubProcessor,
) (*ReportingModule, error) {
	application, err := app.NewApplication(
		pgConn,
		pubSub,
	)
	if err != nil {
		return nil, err
	}

	apiServer := inbound.NewHttpServer(
		application,
	)

	inbound.RegisterEventHandlers(application.Events, pubSub)

	module := &ReportingModule{
		RestApi: inbound.HandlerFromMux(
			apiServer.StrictHandler(),
			hteeteepee.NewRouter(
				"reporting",
			),
		),
	}

	return module, nil
}
