package inbound

import (
	"github.com/cooperlutz/go-full/internal/examination/app"
	"github.com/cooperlutz/go-full/pkg/workerbee"
)

type WorkerAdapter struct {
	worker *workerbee.Worker
	app    app.Application
}

func NewTaskWorkerAdapter(wrk *workerbee.Worker, app app.Application) *WorkerAdapter {
	examinationWorker := &WorkerAdapter{
		worker: wrk,
		app:    app,
	}

	examinationWorker.worker.AddTask(app.Commands.CompleteExamsPastTimeLimit.Handle)

	return examinationWorker
}
