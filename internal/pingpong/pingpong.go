package pingpong

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cooperlutz/go-full/internal/pingpong/api/rest"
	"github.com/cooperlutz/go-full/internal/pingpong/app/service"
	"github.com/cooperlutz/go-full/internal/pingpong/infra/persist"
)

type PingPongModule struct {
	PersistentRepo *persist.PingPongPersistPostgresRepository
	Service        *service.PingPongService
	RestApi        http.Handler
}

// NewModule - Initializes the PingPong module with its needed dependencies.
func NewModule(pgconn *pgxpool.Pool) *PingPongModule {
	repo := persist.NewPingPongPostgresRepo(pgconn)
	svc := service.NewPingPongService(repo)
	api := rest.NewPingPongAPIRouter(svc)

	module := &PingPongModule{
		PersistentRepo: repo,
		Service:        svc,
		RestApi:        api,
	}

	return module
}
