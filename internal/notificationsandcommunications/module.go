package notificationsandcommunications

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cooperlutz/go-full/internal/notificationsandcommunications/adapters/inbound"
	"github.com/cooperlutz/go-full/internal/notificationsandcommunications/app"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/hteeteepee"
)

type NotificationsAndCommunicationsModule struct {
	RestApi http.Handler
}

// NewModule - Initializes the NotificationsAndCommunications module with its needed dependencies.
func NewModule(
	pgconn *pgxpool.Pool,
	pubSub *eeventdriven.BasePgsqlPubSubProcessor,
) (*NotificationsAndCommunicationsModule, error) {
	application, err := app.NewApplication(
		pgconn,
		pubSub,
	)
	if err != nil {
		return nil, err
	}

	apiServer := inbound.NewHttpAdapter(application)

	inbound.NewSqlSubscriberAdapter(application.Events, pubSub).RegisterEventHandlers()

	module := &NotificationsAndCommunicationsModule{
		RestApi: inbound.HandlerFromMux(
			apiServer.StrictHandler(),
			hteeteepee.NewRouter(
				"notificationsandcommunications",
			),
		),
	}

	return module, nil
}
