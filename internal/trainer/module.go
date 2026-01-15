package trainer

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cooperlutz/go-full/internal/trainer/app"
	"github.com/cooperlutz/go-full/internal/trainer/ports"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/hteeteepee"
)

// TrainerModule encapsulates the Trainer module's components, dependencies, interfaces
// that is implemented within the core application to provide Trainer functionality.
type TrainerModule struct {
	RestApi http.Handler
	PubSub  eeventdriven.IPubSubEventProcessor
}

// NewModule - Initializes the Trainer module with its needed dependencies.
func NewModule(pgconn *pgxpool.Pool) (*TrainerModule, error) {
	ctx := context.Background()

	basePS, err := eeventdriven.NewPubSub(pgconn)
	if err != nil {
		return nil, err
	}

	application := app.NewApplication(ctx, pgconn, basePS)
	router := hteeteepee.NewRouter("trainer")
	api := ports.NewHttpServer(application)
	handler := ports.HandlerFromMux(
		api,
		router,
	)

	module := &TrainerModule{
		RestApi: handler,
		PubSub:  basePS,
	}

	return module, nil
}
