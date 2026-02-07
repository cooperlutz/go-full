package app

import (
	"context"
	"os"
	"sync"

	"github.com/exaring/otelpgx"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cooperlutz/go-full/api/frontend"
	"github.com/cooperlutz/go-full/app/config"
	"github.com/cooperlutz/go-full/internal/examination"
	"github.com/cooperlutz/go-full/internal/examlibrary"
	"github.com/cooperlutz/go-full/internal/iam"
	iam_repo "github.com/cooperlutz/go-full/internal/iam/adapters/outbound"
	iam_handlers "github.com/cooperlutz/go-full/internal/iam/handlers"
	"github.com/cooperlutz/go-full/internal/pingpong"
	"github.com/cooperlutz/go-full/pkg/hteeteepee"
	"github.com/cooperlutz/go-full/pkg/securitee"
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
func (a *Application) Run() { //nolint:funlen // main application run function
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

	Create a new instance of the PingPongService, injecting the Postgres connection as a dependency.
	----------------------------------------------------------------------------------- */

	// PingPong
	pingPongModule, err := pingpong.NewModule(conn)
	if err != nil {
		os.Exit(1)
	}

	// Exam Library
	examLibraryModule, err := examlibrary.NewModule(conn)
	if err != nil {
		os.Exit(1)
	}

	// Examination
	examinationModule, err := examination.NewModule(conn, examLibraryModule.UseCase)
	if err != nil {
		os.Exit(1)
	}

	// User Management
	iamRepo := iam_repo.New(conn)
	iamSvc := iam.NewIamService(
		iamRepo,
		a.conf.Security.JWTSecret,
		a.conf.Security.AccessTokenTTL,
	)

	/* -----------------------------------------------------------------------------------
	REST API Controller Initialization:

	Create a new Chi router instance to be used by the API controller
	----------------------------------------------------------------------------------- */
	// Public Routes
	publicRestApiController := hteeteepee.NewRootRouterWithMiddleware()
	authHandler := iam_handlers.NewIAMApiController(iamSvc)
	publicRestApiController.Mount("/auth", authHandler.IamRouter)

	/* -----------------------------------------------------------------------------------
	Setup Domain Module Routes

	Each domain module's router is created and registered with the main HTTP server handler.
	the resulting mountpoint will be {root}/{service-name}/[routes defined in the service router]
	----------------------------------------------------------------------------------- */
	authMiddleware := securitee.AuthMiddleware(iamSvc)
	protectedRestApiController := hteeteepee.NewRootRouterWithMiddleware(authMiddleware)
	userHandler := iam_handlers.NewUserHandler(iamRepo)
	protectedRestApiController.HandleFunc("/iam/profile", userHandler.Profile)
	protectedRestApiController.Mount("/pingpong", pingPongModule.RestApi)
	protectedRestApiController.Mount("/examlibrary", examLibraryModule.RestApi)
	protectedRestApiController.Mount("/examination", examinationModule.RestApi)

	/* -----------------------------------------------------------------------------------
	Setup Web Router
	----------------------------------------------------------------------------------- */
	webRouter := hteeteepee.NewRouter("web")
	webRouter.Handle("/*", frontend.SPAHandler())

	/* -----------------------------------------------------------------------------------
	HTTP Server Initialization
	----------------------------------------------------------------------------------- */
	httpServer := hteeteepee.NewHTTPServer(a.conf, publicRestApiController)
	httpServer.RegisterController("/api", protectedRestApiController)
	httpServer.RegisterController("/", webRouter)

	/* -----------------------------------------------------------------------------------
	Run the HTTP server & Pub/Sub processors
	----------------------------------------------------------------------------------- */
	var wg sync.WaitGroup
	// We increment the WaitGroup counter by 4 for the four servers we plan to run.
	wg.Add(4) //nolint:mnd // we have four goroutines to wait for

	go func() {
		defer wg.Done()

		httpServer.Run()
	}()

	go func() {
		defer wg.Done()

		pingPongModule.PubSub.Run()
	}()
	go func() {
		defer wg.Done()

		examLibraryModule.PubSub.Run()
	}()
	go func() {
		defer wg.Done()

		examinationModule.SubscriberInterface.Start()
	}()

	wg.Wait() // Wait for both servers to finish (they won't, unless there's an error)
}
