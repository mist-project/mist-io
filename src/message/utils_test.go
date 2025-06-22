package message_test

import (
	"testing"

	"mist-io/src/message"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestGetServerClient(t *testing.T) {
	// ARRANGE
	mockConn := new(grpc.ClientConn)
	client := message.Client{
		Conn: mockConn,
	}

	// ACT
	serverClient := client.GetServerClient()

	// ASSERT
	assert.NotNil(t, serverClient)
}

func TestGetChannelClient(t *testing.T) {
	// ARRANGE
	mockConn := new(grpc.ClientConn)
	client := message.Client{
		Conn: mockConn,
	}

	// ACT
	serverClient := client.GetChannelClient()

	// ASSERT
	assert.NotNil(t, serverClient)
}
