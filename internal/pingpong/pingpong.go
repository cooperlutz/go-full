package pingpong

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cooperlutz/go-full/internal/pingpong/api/rest"
	"github.com/cooperlutz/go-full/internal/pingpong/app/usecase"
	"github.com/cooperlutz/go-full/internal/pingpong/infra/persist"
	"github.com/cooperlutz/go-full/internal/pingpong/infra/pubsub"
)

// PingPongModule encapsulates the PingPong module's components, dependencies, interfaces
// that is implemented within the core application to provide PingPong functionality.
type PingPongModule struct {
	PersistentRepo *persist.PingPongPersistPostgresRepository
	UseCase        *usecase.PingPongUseCase
	RestApi        http.Handler
	PubSub         *pubsub.PingPongPubSub
}

// NewModule - Initializes the PingPong module with its needed dependencies.
func NewModule(pgconn *pgxpool.Pool) (*PingPongModule, error) {
	repo := persist.NewPingPongPostgresRepo(pgconn)

	ps, err := pubsub.New(pgconn, repo)
	if err != nil {
		return nil, err
	}

	uc := usecase.NewPingPongUseCase(repo, ps)
	api := rest.NewPingPongAPIRouter(uc)

	module := &PingPongModule{
		PersistentRepo: repo,
		UseCase:        uc,
		RestApi:        api,
		PubSub:         ps,
	}

	err = module.PubSub.RegisterSubscriberHandlers()
	if err != nil {
		return nil, err
	}

	return module, nil
}
