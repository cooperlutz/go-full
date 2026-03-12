package patientmanagement

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cooperlutz/go-full/internal/patientmanagement/adapters/inbound"
	"github.com/cooperlutz/go-full/internal/patientmanagement/app"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/hteeteepee"
)

type PatientManagementModule struct {
	RestApi http.Handler
}

// NewModule - Initializes the PatientManagement module with its needed dependencies.
func NewModule(
	pgconn *pgxpool.Pool,
	pubSub *eeventdriven.BasePgsqlPubSubProcessor,
) (*PatientManagementModule, error) {
	application, err := app.NewApplication(
		pgconn,
		pubSub,
	)
	if err != nil {
		return nil, err
	}

	apiServer := inbound.NewHttpAdapter(application)

	inbound.NewSqlSubscriberAdapter(application.Events, pubSub).RegisterEventHandlers()

	module := &PatientManagementModule{
		RestApi: inbound.HandlerFromMux(
			apiServer.StrictHandler(),
			hteeteepee.NewRouter(
				"patientmanagement",
			),
		),
	}

	return module, nil
}
