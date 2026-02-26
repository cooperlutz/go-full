package workerbee

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/app/config"
)

func sampleFunction(ctx context.Context) error {
	fmt.Println("Executing sample function...")
	return nil
}

func sampleFunctionReturningAnError(ctx context.Context) error {
	fmt.Println("Executing sample function that returns an error...")
	return assert.AnError
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

func TestWorker_NoTasks(t *testing.T) {
	t.Parallel()
	worker := NewWorker(config.Telemetry{TraceEndpoint: ""}, 2*time.Second)
	err := worker.Run()
	assert.Error(t, err)
	assert.Len(t, worker.cron.C, 0)
}

func TestWorker_TaskReturnsError(t *testing.T) {
	t.Parallel()
	worker := NewWorker(config.Telemetry{TraceEndpoint: ""}, 2*time.Second)
	worker.AddTask(sampleFunctionReturningAnError)
	go func() {
		if err := worker.Run(); err != nil {
			fmt.Printf("Worker stopped with error: %v\n", err)
			assert.Error(t, err)
		}
	}()
}
