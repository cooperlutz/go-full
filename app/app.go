package app

import (
	"context"
	"os"

	"github.com/exaring/otelpgx"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cooperlutz/go-full/api/frontend"
	"github.com/cooperlutz/go-full/app/config"
	"github.com/cooperlutz/go-full/internal/pingpong"
	"github.com/cooperlutz/go-full/pkg/hteeteepee"
)

// Application represents the main application structure.
type Application struct {
	conf config.Config
}

// Create a new application instance with the provided configuration.
func NewApplication(conf config.Config) *Application {
	return &Application{
		conf: conf,
	}
}

// this is where all of the wiring happens.
func (a *Application) Run() {
	/* -----------------------------------------------------------------------------------
	System Initializations:
	----------------------------------------------------------------------------------- */
	pgCfg, err := pgxpool.ParseConfig(a.conf.DB.GetDSN())
	if err != nil {
		os.Exit(1)
	}

	pgCfg.ConnConfig.Tracer = otelpgx.NewTracer()

	conn, err := pgxpool.NewWithConfig(context.Background(), pgCfg)
	if err != nil {
		os.Exit(1)
	}

	if err := otelpgx.RecordStats(conn); err != nil {
		os.Exit(1)
	}

	/* -----------------------------------------------------------------------------------
	Modular Service Initializations:

	Create a new instance of the PingPongService, injecting the Postgres repository as a dependency.
	----------------------------------------------------------------------------------- */

	// PingPong
	pingPongModule := pingpong.NewModule(conn)

	/* -----------------------------------------------------------------------------------
	REST API Controller Initialization:

	Create a new Chi router instance to be used by the API controller
	----------------------------------------------------------------------------------- */
	restApiController := hteeteepee.NewRootRouterWithMiddleware()

	/* -----------------------------------------------------------------------------------
	HTTP Server Initialization
	----------------------------------------------------------------------------------- */
	httpServer := hteeteepee.NewHTTPServer(a.conf, restApiController)

	/* -----------------------------------------------------------------------------------
	Setup Web Router
	----------------------------------------------------------------------------------- */
	webRouter := hteeteepee.NewRouter("web")
	webRouter.Handle("/*", frontend.SPAHandler())
	httpServer.RegisterController("/", webRouter)

	/* -----------------------------------------------------------------------------------
	Setup Service Routes

	Each domain's router is created and registered with the main HTTP server handler.
	the resulting mountpoint will be {root}/{service-name}/[routes defined in the service router]
	----------------------------------------------------------------------------------- */
	httpServer.RegisterController("/pingpong", pingPongModule.RestApi) // mounts `/pingpong/api/v1/ping-pong`

	/* -----------------------------------------------------------------------------------
	Run the HTTP server
	----------------------------------------------------------------------------------- */
	httpServer.Run()
}
