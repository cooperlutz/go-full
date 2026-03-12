package ownermanagement

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cooperlutz/go-full/internal/ownermanagement/adapters/inbound"
	"github.com/cooperlutz/go-full/internal/ownermanagement/app"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/hteeteepee"
)

type OwnerManagementModule struct {
	RestApi http.Handler
}

// NewModule - Initializes the OwnerManagement module with its needed dependencies.
func NewModule(
	pgconn *pgxpool.Pool,
	pubSub *eeventdriven.BasePgsqlPubSubProcessor,
) (*OwnerManagementModule, error) {
	application, err := app.NewApplication(
		pgconn,
		pubSub,
	)
	if err != nil {
		return nil, err
	}

	apiServer := inbound.NewHttpAdapter(application)

	inbound.NewSqlSubscriberAdapter(application.Events, pubSub).RegisterEventHandlers()

	module := &OwnerManagementModule{
		RestApi: inbound.HandlerFromMux(
			apiServer.StrictHandler(),
			hteeteepee.NewRouter(
				"ownermanagement",
			),
		),
	}

	return module, nil
}
