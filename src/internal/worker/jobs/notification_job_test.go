package jobs_test

import (
	"context"
	"errors"
	"mist-io/src/internal/faults"
	"mist-io/src/internal/worker/jobs"
	"mist-io/src/testutil/mocks"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestNotificationJob_Execute(t *testing.T) {
	tests := []struct {
		name         string
		mockError    error
		expectError  bool
		expectedType interface{}
	}{
		{
			name:         "success_message_written_successfully",
			mockError:    nil,
			expectError:  false,
			expectedType: nil,
		},
		{
			name:         "error_write_message_fails",
			mockError:    errors.New("write failed"),
			expectError:  true,
			expectedType: &faults.CustomError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// ARRANGE
			mockConn := new(mocks.MockWebSocketConnection)
			content := []byte("hello world")

			mockConn.On("WriteMessage", websocket.BinaryMessage, content).
				Return(tt.mockError)

			job := jobs.NewNotificationJob(content, mockConn)

			// ACT
			err := job.Execute(1)

			// ASSERT
			if tt.expectError {
				assert.Error(t, err)
				assert.IsType(t, tt.expectedType, err)
			} else {
				assert.NoError(t, err)
			}

			mockConn.AssertExpectations(t)
		})
	}
}

func TestNotificationJob_Ctx_returns_default_context(t *testing.T) {
	// ARRANGE
	mockConn := new(mocks.MockWebSocketConnection)
	content := []byte("hello world")
	job := jobs.NewNotificationJob(content, mockConn)

	// ACT
	ctx := job.Ctx()

	// ASSERT
	assert.NotNil(t, ctx)
	assert.Equal(t, context.Background(), ctx)
}
