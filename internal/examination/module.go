package examination

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cooperlutz/go-full/internal/examination/adapters/inbound"
	"github.com/cooperlutz/go-full/internal/examination/app"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/hteeteepee"
)

type ExaminationModule struct {
	RestApi http.Handler
	PubSub  eeventdriven.IPubSubEventProcessor
}

// NewModule - Initializes the Examination module with its needed dependencies.
func NewModule(pgconn *pgxpool.Pool) (*ExaminationModule, error) {
	basePS, err := eeventdriven.NewPubSub(pgconn)
	if err != nil {
		return nil, err
	}

	application := app.NewApplication(
		pgconn,
		basePS,
	)
	router := hteeteepee.NewRouter(
		"examination",
	)
	api := inbound.NewHttpServer(
		application,
	)
	handler := inbound.HandlerFromMux(
		api.StrictHandler(),
		router,
	)
	module := &ExaminationModule{
		RestApi: handler,
		PubSub:  basePS,
	}

	return module, nil
}
