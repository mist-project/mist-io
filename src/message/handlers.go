package message

import (
	pb "mist-io/src/protos/v1/gen"

	"google.golang.org/protobuf/proto"
)

// ----- auth handlers -----
func (wsc *WsConnection) UpdateJwtToken(message *pb.Input_UpdateJwtToken) {
	wsc.JwtToken = message.UpdateJwtToken.Access
}

// ----- server handlers -----

func (wsc *WsConnection) ServerListing(
	message *pb.Input_ServerListing,
) ([]byte, error) {
	ctx, cancel := wsc.SetupContext()
	defer cancel()

	response, err := pb.NewServerServiceClient(wsc.ClientConn).ListAppservers(
		ctx, &pb.ListAppserversRequest{},
	)

	if err != nil {
		// TODO: improve this error handling
		return nil, err
	}

	return proto.Marshal(&pb.Output{
		Data: &pb.Output_ServerListing{
			ServerListing: &pb.ServerListingResponse{Appservers: response.GetAppservers()},
		},
	})
}
