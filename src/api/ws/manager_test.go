package ws_test

import (
	"fmt"
	"mist-io/src/api/ws"
	"mist-io/src/testutil"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestWSManager_NewWsManager(t *testing.T) {
	t.Run("returns_ws_manager", func(t *testing.T) {

		// ACT
		wsm := ws.NewWSManager()

		// ASSERT
		assert.NotNil(t, wsm)
	})
}

func TestWSManager_FindSocketConnection(t *testing.T) {
	wsId := "123"

	t.Run("Success:returns_socket_connection_if_it_exists", func(t *testing.T) {

		// ARRANGE
		wsm := ws.NewWSManager()
		conn := &websocket.Conn{}
		wsm.AddSocketConnection(wsId, conn)

		// ACT
		c, err := wsm.FindSocketConnection(wsId)

		// ASSERT
		assert.Equal(t, c, conn)
		assert.NoError(t, err)
	})

	t.Run("Error:returns_error_when_socket_connection_does_not_exist", func(t *testing.T) {
		// ARRANGE
		wsm := ws.NewWSManager()

		// ACT
		c, err := wsm.FindSocketConnection(wsId)

		// ASSERT
		assert.Nil(t, c)
		testutil.AssertCustomErrorContains(t, err, fmt.Sprintf("socket with id (%s) not found", wsId))
	})
}

func TestWSManager_AddSocketConnection(t *testing.T) {
	wsId := "123"
	conn := &websocket.Conn{}

	t.Run("Success:when_socket_connection_does_not_exist_it_adds", func(t *testing.T) {

		// ARRANGE
		wsm := ws.NewWSManager()
		conn := &websocket.Conn{}
		wsm.AddSocketConnection(wsId, conn)

		// ACT
		err := wsm.AddSocketConnection(wsId, conn)
		assert.NoError(t, err)

		// ASSERT
		c, err := wsm.FindSocketConnection(wsId)
		assert.Equal(t, c, conn)
		assert.NoError(t, err)
	})

	t.Run("Success:when_socket_connection_is_the_same_it_does_not_replace", func(t *testing.T) {
		// ARRANGE
		wsm := ws.NewWSManager()

		// ACT
		err := wsm.AddSocketConnection(wsId, conn)
		assert.NoError(t, err)

		// ASSERT
		c, err := wsm.FindSocketConnection(wsId)
		assert.Equal(t, c, conn)

		// ACT
		err = wsm.AddSocketConnection(wsId, conn)
		assert.NoError(t, err)

		// ASSERT
		c, err = wsm.FindSocketConnection(wsId)
	})
}

func TestWSManager_RemoveSocketConnection(t *testing.T) {
	wsId := "123"
	conn := &websocket.Conn{}

	t.Run("Success:when_connection_exists_it_deletes", func(t *testing.T) {

		// ARRANGE
		wsm := ws.NewWSManager()
		wsm.AddSocketConnection(wsId, conn)

		// ACT
		wsm.RemoveSocketConnection(wsId)

		// ASSERT
		c, err := wsm.FindSocketConnection(wsId)
		assert.Nil(t, c)
		testutil.AssertCustomErrorContains(t, err, fmt.Sprintf("socket with id (%s) not found", wsId))

	})

	t.Run("Success:when_connection_exists_it_does_not_panic", func(t *testing.T) {
		// ARRANGE
		wsm := ws.NewWSManager()

		// ACT
		wsm.RemoveSocketConnection(wsId)

		// ASSERT
		c, err := wsm.FindSocketConnection(wsId)
		assert.Nil(t, c)
		testutil.AssertCustomErrorContains(t, err, fmt.Sprintf("socket with id (%s) not found", wsId))
	})

}
