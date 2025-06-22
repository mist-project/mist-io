package message

import (
	"github.com/gorilla/websocket"

	"mist-io/src/auth"
)

type WsConnection struct {
	Conn *websocket.Conn
	// Mutex    *sync.Mutex // TBD if needed
	JwtToken string
	Claims   *auth.CustomJWTClaims
}

func (wsc *WsConnection) Manage() {
	// TODO: add more features in the future.
	for {
		_, _, err := wsc.Conn.ReadMessage()
		if err != nil {
			return
		}
	}
}
