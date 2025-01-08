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
func (wsc *WsConnection) AppserverListing(
	message *pb.Input_AppserverListing,
) ([]byte, error) {
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
		Data: &pb.Output_AppserverListing{
			AppserverListing: &pb.GetUserAppserverSubsResponse{Appservers: response.GetAppservers()},
		},
	})
}

func (wsc *WsConnection) CreateAppserver(
	message *pb.Input_CreateAppserver,
) ([]byte, error) {
	ctx, cancel := wsc.SetupContext()
	defer cancel()

	serverClient := wsc.Client.GetServerClient()

	_, err := serverClient.CreateAppserver(
		ctx, &pb.CreateAppserverRequest{Name: message.CreateAppserver.Name},
	)

	if err != nil {
		// TODO: return notification of failure
		return nil, err
	}

	// update all user listings
	response, err := serverClient.GetUserAppserverSubs(
		ctx, &pb.GetUserAppserverSubsRequest{},
	)

	if err != nil {
		// TODO: raise error for logging
		return nil, err
	}

	return proto.Marshal(&pb.Output{
		Data: &pb.Output_AppserverListing{
			AppserverListing: &pb.GetUserAppserverSubsResponse{Appservers: response.GetAppservers()},
		},
	})
}

func (wsc *WsConnection) DeleteAppserver(
	message *pb.Input_DeleteAppserver,
) ([]byte, error) {
	ctx, cancel := wsc.SetupContext()
	defer cancel()

	serverClient := wsc.Client.GetServerClient()

	_, err := serverClient.DeleteAppserver(
		ctx, &pb.DeleteAppserverRequest{Id: message.DeleteAppserver.Id},
	)

	if err != nil {
		// TODO: return notification of failure
		return nil, err
	}

	// update all user listings
	response, err := serverClient.GetUserAppserverSubs(
		ctx, &pb.GetUserAppserverSubsRequest{},
	)

	if err != nil {
		// TODO: raise error for logging
		return nil, err
	}

	return proto.Marshal(&pb.Output{
		Data: &pb.Output_AppserverListing{
			AppserverListing: &pb.GetUserAppserverSubsResponse{Appservers: response.GetAppservers()},
		},
	})
}
