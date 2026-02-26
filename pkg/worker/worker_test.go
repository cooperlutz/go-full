package worker

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/cooperlutz/go-full/app/config"
)

func sampleFunction(ctx context.Context) error {
	fmt.Println("Executing sample function...")
	return nil
}

func TestNewWorker(t *testing.T) {
	t.Parallel()
	worker := NewWorker(config.Telemetry{TraceEndpoint: ""}, 2*time.Second)
	worker.AddTask(sampleFunction)
	go func() {
		if err := worker.Run(); err != nil {
			fmt.Printf("Worker stopped with error: %v\n", err)
		}
	}()
}
