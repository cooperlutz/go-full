package billingandpayments

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cooperlutz/go-full/internal/billingandpayments/adapters/inbound"
	"github.com/cooperlutz/go-full/internal/billingandpayments/app"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/hteeteepee"
)

type BillingAndPaymentsModule struct {
	RestApi http.Handler
}

// NewModule - Initializes the BillingAndPayments module with its needed dependencies.
func NewModule(
	pgconn *pgxpool.Pool,
	pubSub *eeventdriven.BasePgsqlPubSubProcessor,
) (*BillingAndPaymentsModule, error) {
	application, err := app.NewApplication(
		pgconn,
		pubSub,
	)
	if err != nil {
		return nil, err
	}

	apiServer := inbound.NewHttpAdapter(application)

	inbound.NewSqlSubscriberAdapter(application.Events, pubSub).RegisterEventHandlers()

	module := &BillingAndPaymentsModule{
		RestApi: inbound.HandlerFromMux(
			apiServer.StrictHandler(),
			hteeteepee.NewRouter(
				"billingandpayments",
			),
		),
	}

	return module, nil
}
