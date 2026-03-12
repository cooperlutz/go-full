package veterinarystaff

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cooperlutz/go-full/internal/veterinarystaff/adapters/inbound"
	"github.com/cooperlutz/go-full/internal/veterinarystaff/app"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/hteeteepee"
)

type VeterinaryStaffModule struct {
	RestApi http.Handler
}

// NewModule - Initializes the VeterinaryStaff module with its needed dependencies.
func NewModule(
	pgconn *pgxpool.Pool,
	pubSub *eeventdriven.BasePgsqlPubSubProcessor,
) (*VeterinaryStaffModule, error) {
	application, err := app.NewApplication(
		pgconn,
		pubSub,
	)
	if err != nil {
		return nil, err
	}

	apiServer := inbound.NewHttpAdapter(application)

	inbound.NewSqlSubscriberAdapter(application.Events, pubSub).RegisterEventHandlers()

	module := &VeterinaryStaffModule{
		RestApi: inbound.HandlerFromMux(
			apiServer.StrictHandler(),
			hteeteepee.NewRouter(
				"veterinarystaff",
			),
		),
	}

	return module, nil
}
