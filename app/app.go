package app

import (
	"context"
	"log/slog"
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
	"github.com/cooperlutz/go-full/pkg/telemetree"
	"github.com/cooperlutz/go-full/pkg/utilitee"
	"github.com/cooperlutz/go-full/pkg/workerbee"
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
	slog.Info("Starting application initializations...")

	ctx := context.Background()

	err := telemetree.InitLogger(ctx, a.conf.Telemetry)
	if err != nil {
		os.Exit(1)
	}

	slog.Info("Logger initialization complete.")

	pgCfg, err := pgxpool.ParseConfig(a.conf.DB.GetDSN())
	if err != nil {
		os.Exit(1)
	}

	pgCfg.ConnConfig.Tracer = otelpgx.NewTracer()

	conn, err := pgxpool.NewWithConfig(ctx, pgCfg)
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

	backgroundWorker := workerbee.NewWorker(
		a.conf.Telemetry,
		10*time.Second, //nolint: mnd // want the worker to run every 10 seconds
	)

	slog.Info("System initializations complete. Starting service initializations...")

	/* -----------------------------------------------------------------------------------
	Modular Service Initializations:

	Create a new instance of each module, injecting the necessary dependencies.
	----------------------------------------------------------------------------------- */
	// PingPong
	pingPongModule, err := pingpong.NewModule(
		conn,
	)
	if err != nil {
		slog.Error("Error initializing ping pong module: " + err.Error())
		os.Exit(1)
	}

	// Exam Library
	examLibraryModule, err := examlibrary.NewModule(
		conn,
	)
	if err != nil {
		slog.Error("Error initializing exam library module: " + err.Error())
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
		slog.Error("Error initializing examination module: " + err.Error())
		os.Exit(1)
	}

	iamModule := iam.NewModule(
		conn,
		iam.IamModuleConfig{
			PrivateKey:      utilitee.MustParseRSAKey(a.conf.Security.Base64EncodedJWTPrivateKeyPEM),
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
		slog.Error("Error initializing grading module: " + err.Error())
		os.Exit(1)
	}

	// Reporting
	reportingModule, err := reporting.NewModule(
		conn,
		pubSub,
	)
	if err != nil {
		slog.Error("Error initializing reporting module: " + err.Error())
		os.Exit(1)
	}

	slog.Info("Service initializations complete. Starting server setup...")

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
		err = httpServer.Run(ctx)
		if err != nil {
			slog.Error("Error running HTTP server: " + err.Error())

			errChannel <- err
		}
	}()

	// the Pub/Sub processors
	go func() {
		err = pingPongModule.PubSub.Run(ctx)
		if err != nil {
			slog.Error("Error running ping pong pub/sub processor: " + err.Error())

			errChannel <- err
		}
	}()
	go func() {
		err = examLibraryModule.PubSub.Run(ctx)
		if err != nil {
			slog.Error("Error running exam library pub/sub processor: " + err.Error())

			errChannel <- err
		}
	}()
	go func() {
		err = pubSub.Run(ctx)
		if err != nil {
			slog.Error("Error running pub/sub processor: " + err.Error())

			errChannel <- err
		}
	}()

	// the Background Worker
	go func() {
		err = backgroundWorker.Run(ctx)
		if err != nil {
			slog.Error("Error running background worker: " + err.Error())

			errChannel <- err
		}
	}()

	// Wait for any server to return an error
	if err := <-errChannel; err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
