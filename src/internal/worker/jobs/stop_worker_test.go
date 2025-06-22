package jobs_test

import (
	"context"
	"mist-io/src/internal/worker/jobs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStopWorkerJob_returns_job_with_context(t *testing.T) {
	// ARRANGE
	ctx := context.WithValue(context.Background(), "key", "value")

	// ACT
	job := jobs.NewStopWorkerJob(ctx)

	// ASSERT
	assert.NotNil(t, job)
	assert.Equal(t, ctx, job.Ctx())
}

func TestStopWorkerJob_Ctx_returns_correct_context(t *testing.T) {
	// ARRANGE
	ctx := context.WithValue(context.Background(), "test", "123")
	job := jobs.NewStopWorkerJob(ctx)

	// ACT
	resultCtx := job.Ctx()

	// ASSERT
	assert.NotNil(t, resultCtx)
	assert.Equal(t, ctx, resultCtx)
}

func TestStopWorkerJob_Execute_returns_nil(t *testing.T) {
	// ARRANGE
	ctx := context.Background()
	job := jobs.NewStopWorkerJob(ctx)

	// ACT
	err := job.Execute(0)

	// ASSERT
	assert.NoError(t, err)
}
