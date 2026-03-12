package insuranceandclaims

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cooperlutz/go-full/internal/insuranceandclaims/adapters/inbound"
	"github.com/cooperlutz/go-full/internal/insuranceandclaims/app"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/hteeteepee"
)

type InsuranceAndClaimsModule struct {
	RestApi http.Handler
}

// NewModule - Initializes the InsuranceAndClaims module with its needed dependencies.
func NewModule(
	pgconn *pgxpool.Pool,
	pubSub *eeventdriven.BasePgsqlPubSubProcessor,
) (*InsuranceAndClaimsModule, error) {
	application, err := app.NewApplication(
		pgconn,
		pubSub,
	)
	if err != nil {
		return nil, err
	}

	apiServer := inbound.NewHttpAdapter(application)

	inbound.NewSqlSubscriberAdapter(application.Events, pubSub).RegisterEventHandlers()

	module := &InsuranceAndClaimsModule{
		RestApi: inbound.HandlerFromMux(
			apiServer.StrictHandler(),
			hteeteepee.NewRouter(
				"insuranceandclaims",
			),
		),
	}

	return module, nil
}
