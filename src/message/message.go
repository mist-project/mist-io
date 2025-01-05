package message

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"

	"mist-io/src/auth"
	"mist-io/src/helpers"
	pb "mist-io/src/protos/v1/gen"
)

type WsConnection struct {
	Conn *websocket.Conn
	// Mutex    *sync.Mutex // TBD if needed
	JwtToken string
	Claims   *auth.CustomJWTClaims
	queue    *helpers.Queue[InternalItem]
	Client   GrpcClient
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

		parsedMessage := &pb.InputMessage{}

		err = proto.Unmarshal(p, parsedMessage)
		if err != nil {
			// TODO: update error handling
			log.Println(err)
			return
		}

		wsc.socketMessageHandler(parsedMessage, messageType)

	}
}

func (wsc *WsConnection) socketMessageHandler(message *pb.InputMessage, messageType int) {
	var response []byte
	var err error

	switch input := message.Input.Data.(type) {
	case *pb.Input_UpdateJwtToken:
		fmt.Printf("<<< JWT token message\n")
		wsc.UpdateJwtToken(input)
		return
	case *pb.Input_AppserverListing:
		fmt.Printf("<<< Server listing message\n")
		response, err = wsc.AppserverListing(input)

	case *pb.Input_CreateAppserver:
		fmt.Printf("<<< Server listing message\n")
		response, err = wsc.CreateAppserver(input)

	default:
		fmt.Println("Unknown type")
	}

	if err != nil {
		fmt.Printf("error processing %v\n", err)
		// TODO: better error handling here
		return
	}

	wsc.queue.Enqueue(&InternalItem{
		internalType: messageType,
		data:         response,
	})

}

func (wsc *WsConnection) messageQueueHandler() {
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
