package app

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/exaring/otelpgx"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cooperlutz/go-full/api/frontend"
	"github.com/cooperlutz/go-full/app/config"
	"github.com/cooperlutz/go-full/internal/examination"
	"github.com/cooperlutz/go-full/internal/examlibrary"
	"github.com/cooperlutz/go-full/internal/grading"
	"github.com/cooperlutz/go-full/internal/iam"
	"github.com/cooperlutz/go-full/internal/pingpong"
	"github.com/cooperlutz/go-full/internal/reporting"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
	"github.com/cooperlutz/go-full/pkg/hteeteepee"
	"github.com/cooperlutz/go-full/pkg/securitee"
	"github.com/cooperlutz/go-full/pkg/worker"
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
func (a *Application) Run() { //nolint:funlen,cyclop,gocyclo,gocognit // main application run function
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

	pubSub, err := eeventdriven.NewPubSub(conn)
	if err != nil {
		os.Exit(1)
	}

	backgroundWorker := worker.NewWorker(
		a.conf.Telemetry,
		10*time.Second, //nolint: mnd // want the worker to run every 10 seconds
	)

	/* -----------------------------------------------------------------------------------
	Modular Service Initializations:

	Create a new instance of each module, injecting the necessary dependencies.
	----------------------------------------------------------------------------------- */

	// PingPong
	pingPongModule, err := pingpong.NewModule(
		conn,
	)
	if err != nil {
		os.Exit(1)
	}

	// Exam Library
	examLibraryModule, err := examlibrary.NewModule(
		conn,
	)
	if err != nil {
		os.Exit(1)
	}

	// Examination
	examinationModule, err := examination.NewModule(
		conn,
		pubSub,
		examLibraryModule.UseCase,
		backgroundWorker,
	)
	if err != nil {
		os.Exit(1)
	}

	// Identity & Access Management
	iamModule := iam.NewModule(
		conn,
		iam.IamModuleConfig{
			JwtSecret:       a.conf.Security.JWTSecret,
			AccessTokenTTL:  a.conf.Security.AccessTokenTTL,
			RefreshTokenTTL: a.conf.Security.RefreshTokenTTL,
		},
	)

	// Grading
	gradingModule, err := grading.NewModule(
		conn,
		pubSub,
		examLibraryModule.UseCase,
	)
	if err != nil {
		os.Exit(1)
	}

	// Reporting
	reportingModule, err := reporting.NewModule(
		conn,
		pubSub,
	)
	if err != nil {
		os.Exit(1)
	}

	/* -----------------------------------------------------------------------------------
	Protected REST API Controller Initialization:
	----------------------------------------------------------------------------------- */
	protectedRestApiRouter := hteeteepee.NewRootRouterWithMiddleware(
		securitee.AuthMiddleware(iamModule.Service), // Authentication middleware to protect all routes under this router
	)

	/* -----------------------------------------------------------------------------------
	Setup Domain Module Routes

	Each domain module's router is created and registered with the main HTTP server handler.
	the resulting mountpoint will be {root}/api/{service-name}/[routes defined in the service router]
	----------------------------------------------------------------------------------- */
	protectedRestApiRouter.Mount(
		"/pingpong",
		pingPongModule.RestApi,
	)
	protectedRestApiRouter.Mount(
		"/examlibrary",
		examLibraryModule.RestApi,
	)
	protectedRestApiRouter.Mount(
		"/examination",
		examinationModule.RestApi,
	)
	protectedRestApiRouter.Mount(
		"/iam",
		iamModule.UserRestApi,
	)
	protectedRestApiRouter.Mount(
		"/grading",
		gradingModule.RestApi,
	)
	protectedRestApiRouter.Mount(
		"/reporting",
		reportingModule.RestApi,
	)

	/* -----------------------------------------------------------------------------------
	Mount Public Routes
	----------------------------------------------------------------------------------- */
	publicHttpRouter := hteeteepee.NewRootRouterWithMiddleware()
	publicHttpRouter.Mount("/", frontend.SpaRouter())
	publicHttpRouter.Mount("/auth", iamModule.AuthRestApi)

	/* -----------------------------------------------------------------------------------
	HTTP Server Initialization
	----------------------------------------------------------------------------------- */
	httpServer := hteeteepee.NewHTTPServer(
		a.conf,
		publicHttpRouter,
	)
	httpServer.RegisterController(
		"/api",
		protectedRestApiRouter,
	)

	/* -----------------------------------------------------------------------------------
	Run the HTTP server & Pub/Sub processors
	----------------------------------------------------------------------------------- */
	errChannel := make(chan error, 1) // Buffer size of 1 for the error channel

	// run each server in its own goroutine and send any errors to the error channel

	// the HTTP server
	go func() {
		err = httpServer.Run()
		if err != nil {
			errChannel <- err
		}
	}()

	// the Pub/Sub processors
	go func() {
		err = pingPongModule.PubSub.Run()
		if err != nil {
			errChannel <- err
		}
	}()
	go func() {
		err = examLibraryModule.PubSub.Run()
		if err != nil {
			errChannel <- err
		}
	}()
	go func() {
		err = pubSub.Run()
		if err != nil {
			errChannel <- err
		}
	}()

	// the Background Worker
	go func() {
		err = backgroundWorker.Run()
		if err != nil {
			errChannel <- err
		}
	}()

	// Wait for any server to return an error
	if err := <-errChannel; err != nil {
		os.Exit(1)
		log.Fatal(err)
	}
}
