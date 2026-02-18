package hteeteepee

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/cooperlutz/go-full/app/config"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type HTTPServer struct {
	Config config.Config
	Router *chi.Mux
	Server *http.Server
}

func NewHTTPServer(conf config.Config, router *chi.Mux) *HTTPServer {
	return &HTTPServer{
		Config: conf,
		Router: router,
		Server: &http.Server{
			// All Configuration Options Below for reference
			Addr: conf.HTTP.Port,
			// Handler: r,
			// DisableGeneralOptionsHandler: true,
			// TLSConfig: nil,
			// ReadTimeout: 0,
			// WriteTimeout: 0,
			ReadHeaderTimeout: 1 * time.Second,
			// IdleTimeout: 0,
			// MaxHeaderBytes: 0,
			// ErrorLog: nil,
			// BaseContext: nil,
			// ConnContext: nil,
			// ConnState: nil,
			// Protocols: nil,
			// HTTP2: nil,
			// TLSNextProto: nil,
		},
	}
}

func (s *HTTPServer) RegisterController(serviceEndpoint string, handler http.Handler) {
	// Register your routes here
	s.Router.Mount(serviceEndpoint, handler)
}

func (s *HTTPServer) Run() error {
	// Set the server handler
	s.Server.Handler = s.Router

	ctx := context.Background()

	tp, err := telemetree.InitTracer(ctx, s.Config.Telemetry)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := tp.Shutdown(ctx); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	mp, err := telemetree.InitMeter(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := mp.Shutdown(ctx); err != nil {
			log.Printf("Error shutting down meter provider: %v", err)
		}
	}()

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)

	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// Graceful shutdown
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, cancel := context.WithTimeout(serverCtx, 30*time.Second) //nolint:mnd // idk
		defer cancel()

		go func() {
			<-shutdownCtx.Done()

			if errors.Is(shutdownCtx.Err(), context.DeadlineExceeded) {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := s.Server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}

		serverStopCtx()
	}()

	// Run the server
	err = s.Server.ListenAndServe()

	if !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)

		return err
	}

	// Wait for server context to be stopped upon receiving a shutdown signal
	<-serverCtx.Done()

	return nil
}
