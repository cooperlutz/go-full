package pingpong

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cooperlutz/go-full/internal/pingpong/api/rest"
	"github.com/cooperlutz/go-full/internal/pingpong/app/usecase"
	"github.com/cooperlutz/go-full/internal/pingpong/infra/persist"
)

type PingPongModule struct {
	PersistentRepo *persist.PingPongPersistPostgresRepository
	UseCase        *usecase.PingPongUseCase
	RestApi        http.Handler
}

// NewModule - Initializes the PingPong module with its needed dependencies.
func NewModule(pgconn *pgxpool.Pool) *PingPongModule {
	repo := persist.NewPingPongPostgresRepo(pgconn)
	uc := usecase.NewPingPongUseCase(repo)
	api := rest.NewPingPongAPIRouter(uc)

	module := &PingPongModule{
		PersistentRepo: repo,
		UseCase:        uc,
		RestApi:        api,
	}

	return module
}
