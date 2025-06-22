package jobs

import (
	"context"
)

// ------ STOP WORKER JOB -----
type StopWorkerJob struct {
	ctx context.Context
}

func NewStopWorkerJob(ctx context.Context) *StopWorkerJob {
	return &StopWorkerJob{
		ctx: ctx,
	}
}
func (job *StopWorkerJob) Ctx() context.Context {
	return job.ctx
}

func (job *StopWorkerJob) Execute(worker int) error {
	return nil // No error, just a signal to stop the worker
}
