package message

import (
	"log"
	"mist-io/src/auth"
	"mist-io/src/helpers"

	"github.com/gorilla/websocket"
)

type WsConnection struct {
	Conn *websocket.Conn
	// Mutex    *sync.Mutex // TBD if needed
	JwtToken string
	Claims   *auth.CustomJWTClaims
	queue    *helpers.Queue[InternalItem]
}

type InternalItem struct {
	internalType int
	data         []byte
}

func (wsc *WsConnection) Manage() {
	// Initialize conditional variable
	wsc.queue = helpers.NewQueue[InternalItem]()

	go wsc.processMessages() // process all messages enqueued to be sent to user

	for {
		messageType, p, err := wsc.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// for now echo by adding message to queue
		wsc.queue.Enqueue(&InternalItem{
			internalType: messageType,
			data:         p,
		})
	}

}

func (wsc WsConnection) processMessages() {
	// This will be the entrypoint used to determine all messages to the user
	// One message can be sent at a time only
	for {
		item := wsc.queue.Dequeue()
		wsc.sendMessage(item)
	}
}

func (wsc *WsConnection) sendMessage(item *InternalItem) {
	// Might not need locks since the process messages blocks when sending one message at a time
	// wsc.Mutex.Lock()
	// defer wsc.Mutex.Unlock()
	wsc.Conn.WriteMessage(item.internalType, item.data)
}
