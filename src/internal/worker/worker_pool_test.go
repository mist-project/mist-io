package worker_test

import (
	"bytes"
	"fmt"
	"mist-io/src/internal/logging/logger"
	"mist-io/src/internal/worker"
	"mist-io/src/internal/worker/jobs"
	"mist-io/src/testutil/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestWorkerPool_AddJob(t *testing.T) {

	t.Run("Success:job_added_to_queue", func(t *testing.T) {
		// ARRANGE
		wp := worker.NewWorkerPool(2, 10)
		mockWsc := new(mocks.MockWebSocketConnection)
		notification := jobs.NewNotificationJob(
			[]byte{}, mockWsc,
		)

		// ACT
		wp.AddJob(notification)

		// ASSERT
		assert.Equal(t, wp.GetJobQueueSize(), 1, "Expected job queue size to be 1")
	})

	t.Run("Error:worker_stopped", func(t *testing.T) {
		// ARRANGE
		wp := worker.NewWorkerPool(3, 10)
		wp.StartWorkers()
		wp.Stop() // Cancel the context immediately

		mockWsc := new(mocks.MockWebSocketConnection)
		notification := jobs.NewNotificationJob(
			[]byte{}, mockWsc,
		)

		// ACT
		wp.AddJob(notification)

		// ASSERT
		assert.Equal(t, wp.GetJobQueueSize(), 0, "Expected job queue size to be 0 due to cancelled context")
	})

	t.Run("Error:when_job_execute_fails_it_logs", func(t *testing.T) {
		// ARRANGE
		var buf bytes.Buffer
		logger.SetLogOutput(&buf) // redirect logger output

		wp := worker.NewWorkerPool(2, 10)
		mockWsc := new(mocks.MockWebSocketConnection)

		mockWsc.On("WriteMessage", mock.Anything, mock.Anything).Return(fmt.Errorf("boomed connection"))

		notification := jobs.NewNotificationJob(
			[]byte{}, mockWsc,
		)

		// ACT
		wp.AddJob(notification)
		wp.StartWorkers()

		// Wait for the job to be processed
		wp.Stop()

		// ASSERT
		logOutput := buf.String()
		assert.Contains(t, logOutput, "ERROR", "")
	})
}

func TestWorkerPool_Stop(t *testing.T) {
	t.Run("Success:worker_pool_stops_and_clears_job_queue", func(t *testing.T) {
		// ARRANGE
		wp := worker.NewWorkerPool(2, 10)
		wp.StartWorkers()

		// ACT
		wp.Stop()

		// ASSERT
		assert.Equal(t, wp.GetJobQueueSize(), 0, "Expected job queue size to be 0 after stopping the worker pool")
	})
}
