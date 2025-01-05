package message

import (
	"google.golang.org/protobuf/proto"

	pb "mist-io/src/protos/v1/gen"
)

// ----- auth handlers -----
func (wsc *WsConnection) UpdateJwtToken(message *pb.Input_UpdateJwtToken) {
	wsc.JwtToken = message.UpdateJwtToken.Access
}

// ----- server handlers -----
func (wsc *WsConnection) ServerListing(
	message *pb.Input_ServerListing,
) ([]byte, error) {
	// ) (int, error) {
	ctx, cancel := wsc.SetupContext()
	defer cancel()

	response, err := wsc.Client.GetServerClient().GetUserAppserverSubs(
		ctx, &pb.GetUserAppserverSubsRequest{},
	)

	if err != nil {
		// TODO: improve this error handling
		return nil, err
	}

	return proto.Marshal(&pb.Output{
		Data: &pb.Output_ServerListing{
			ServerListing: &pb.GetUserAppserverSubsResponse{Appservers: response.GetAppservers()},
		},
	})
}
