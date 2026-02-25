package inbound

import (
	"github.com/cooperlutz/go-full/internal/examination/app"
	"github.com/cooperlutz/go-full/pkg/worker"
)

type WorkerAdapter struct {
	worker *worker.Worker
	app    app.Application
}

func NewTaskWorkerAdapter(wrk *worker.Worker, app app.Application) *WorkerAdapter {
	examinationWorker := &WorkerAdapter{
		worker: wrk,
		app:    app,
	}

	examinationWorker.worker.AddTask(app.Commands.CompleteExamsPastTimeLimit.Handle)

	return examinationWorker
}
