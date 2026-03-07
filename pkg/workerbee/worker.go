// Package workerbee provides a simple worker that executes tasks at regular cron intervals
// The worker initializes OpenTelemetry tracing and metrics providers based on the provided configuration.
package workerbee

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

func (w Worker) getTasks() ([]Task, error) {
	if len(w.fns) == 0 {
		return nil, &ErrNoTasks{}
	}

	return w.fns, nil
}

func (w *Worker) AddTask(fn Task) {
	w.fns = append(w.fns, fn)
}

type ErrNoTasks struct{}

func (e *ErrNoTasks) Error() string {
	return "no tasks to execute"
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
			tasks, err := w.getTasks()
			if err != nil {
				return err
			}

			for _, fn := range tasks {
				if err := fn(ctx); err != nil {
					return fmt.Errorf("worker error: %w", err)
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
