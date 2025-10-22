package pingpong

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cooperlutz/go-full/internal/pingpong/api/rest"
	"github.com/cooperlutz/go-full/internal/pingpong/app/service"
	"github.com/cooperlutz/go-full/internal/pingpong/infra/persist"
)

// NewModule - Initializes the PingPong module with its needed dependencies.
func NewModule(pgconn *pgxpool.Pool) http.Handler {
	repo := persist.NewPingPongPostgresRepo(pgconn)
	svc := service.NewPingPongService(repo)
	api := rest.NewPingPongAPIRouter(svc)

	return api
}
