package worker

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cooperlutz/go-full/app/config"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type Task func(ctx context.Context) error

type Worker struct {
	conf config.Telemetry
	cron *time.Ticker
	fns  []Task
}

func NewWorker(config config.Telemetry, interval time.Duration) *Worker {
	return &Worker{
		conf: config,
		cron: time.NewTicker(interval),
	}
}

func (w *Worker) AddTask(fn Task) {
	w.fns = append(w.fns, fn)
}

func (w *Worker) Run() error { //nolint:cyclop,gocyclo,gocognit // worker run function
	ctx := context.Background()

	tp, err := telemetree.InitTracer(ctx, w.conf)
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

	for {
		select {
		case <-w.cron.C:
			if w.fns != nil {
				for _, fn := range w.fns {
					if err := fn(ctx); err != nil {
						return fmt.Errorf("worker error: %w", err)
					}
				}
			}
		case <-ctx.Done():
			return fmt.Errorf("worker stopped: %w", ctx.Err())
		}
	}
}

func (w *Worker) Stop() {
	w.cron.Stop()
}
