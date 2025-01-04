package message

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"

	"mist-io/src/auth"
	"mist-io/src/helpers"
	pb_shared "mist-io/src/protos/frontend/v1"
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

	go wsc.messageQueueHandler() // process all messages enqueued to be sent to user

	for {
		messageType, p, err := wsc.Conn.ReadMessage()

		parsedMessage := &pb_shared.InputMessage{}

		err = proto.Unmarshal(p, parsedMessage)
		if err != nil {
			// TODO: update error handling
			log.Println(err)
			return
		}
		fmt.Println("made ithere')")

		// for now echo by adding message to queue
		wsc.queue.Enqueue(&InternalItem{
			internalType: messageType,
			data:         p,
		})
	}

}

func (wsc WsConnection) socketMessageHandler(message *pb_shared.InputMessage) {
	switch input := message.Input.Data.(type) {
	case *pb_shared.Input_UpdateJwtToken:
		fmt.Println("Text:", input.UpdateJwtToken)
	default:
		fmt.Println("Unknown type")
	}
}

func (wsc WsConnection) messageQueueHandler() {
	// This will be the entrypoint used to store all messages to the user
	// One message can be sent at a time only
	for {
		item := wsc.queue.Dequeue()
		wsc.sendMessageToUser(item)
	}
}

func (wsc *WsConnection) sendMessageToUser(item *InternalItem) {
	// Might not need locks since the process messages blocks when sending one message at a time
	// wsc.Mutex.Lock()
	// defer wsc.Mutex.Unlock()
	wsc.Conn.WriteMessage(item.internalType, item.data)
}
