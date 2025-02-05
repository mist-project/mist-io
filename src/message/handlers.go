package message

import (
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"

	pb "mist-io/src/protos/v1/gen"
)

// ----- auth handlers -----
func (wsc *WsConnection) UpdateJwtToken(message *pb.Input_UpdateJwtToken) {
	wsc.JwtToken = message.UpdateJwtToken.Access
}

// ----- server handlers -----
func (wsc *WsConnection) AppserverListing(message *pb.Input_AppserverListing) ([]byte, error) {
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

func (wsc *WsConnection) AppserverDetails(message *pb.Input_AppserverDetails) ([]byte, error) {
	ctx, cancel := wsc.SetupContext()
	defer cancel()

	response, err := wsc.Client.GetServerClient().GetByIdAppserver(
		ctx, &pb.GetByIdAppserverRequest{Id: message.AppserverDetails.Id},
	)

	if err != nil {
		// TODO: improve this error handling
		return nil, err
	}

	return proto.Marshal(&pb.Output{
		Data: &pb.Output_AppserverDetails{
			AppserverDetails: &pb.GetByIdAppserverResponse{Appserver: response.GetAppserver()},
		},
	})
}

func (wsc *WsConnection) CreateAppserver(message *pb.Input_CreateAppserver) ([]byte, error) {
	ctx, cancel := wsc.SetupContext()
	defer cancel()

	sClient := wsc.Client.GetServerClient()

	_, err := sClient.CreateAppserver(
		ctx, &pb.CreateAppserverRequest{Name: message.CreateAppserver.Name},
	)

	if err != nil {
		// TODO: return notification of failure
		return nil, err
	}

	// get all user listings ( to be used on initial load )
	response, err := sClient.GetUserAppserverSubs(
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

func (wsc *WsConnection) DeleteAppserver(message *pb.Input_DeleteAppserver) ([]byte, error) {
	ctx, cancel := wsc.SetupContext()
	defer cancel()

	sClient := wsc.Client.GetServerClient()

	_, err := sClient.DeleteAppserver(
		ctx, &pb.DeleteAppserverRequest{Id: message.DeleteAppserver.Id},
	)

	if err != nil {
		// TODO: return notification of failure
		return nil, err
	}

	// TODO: replace for individual appserver remove to all serversubs
	response, err := sClient.GetUserAppserverSubs(
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

func (wsc *WsConnection) JoinAppserver(message *pb.Input_JoinAppserver) ([]byte, error) {
	ctx, cancel := wsc.SetupContext()
	defer cancel()

	sClient := wsc.Client.GetServerClient()

	_, err := sClient.CreateAppserverSub(
		ctx, &pb.CreateAppserverSubRequest{AppserverId: message.JoinAppserver.AppserverId},
	)

	if err != nil {
		// TODO: raise error for logging
		return nil, err
	}

	response, err := sClient.GetUserAppserverSubs(
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

// ----- CHANNEL handlers -----
func (wsc *WsConnection) CreateChannel(message *pb.Input_CreateChannel) ([]byte, error) {
	ctx, cancel := wsc.SetupContext()
	defer cancel()

	cClient := wsc.Client.GetChannelClient()
	_, err := cClient.CreateChannel(
		ctx, &pb.CreateChannelRequest{
			Name: message.CreateChannel.Name, AppserverId: message.CreateChannel.AppserverId},
	)

	response, err := wsc.Client.GetChannelClient().ListChannels(
		ctx, &pb.ListChannelsRequest{AppserverId: &wrapperspb.StringValue{Value: message.CreateChannel.AppserverId}},
	)

	if err != nil {
		// TODO: return notification of failure
		fmt.Printf("error: %v\n", err)
		return nil, err
	}

	return proto.Marshal(&pb.Output{
		Data: &pb.Output_ChannelListing{
			ChannelListing: &pb.ListChannelsResponse{Channels: response.GetChannels()},
		},
	})
}

func (wsc *WsConnection) ChanneListing(message *pb.Input_ChannelListing) ([]byte, error) {
	ctx, cancel := wsc.SetupContext()
	defer cancel()

	response, err := wsc.Client.GetChannelClient().ListChannels(
		ctx, &pb.ListChannelsRequest{AppserverId: message.ChannelListing.AppserverId, Name: message.ChannelListing.Name},
	)

	if err != nil {
		// TODO: improve this error handling
		return nil, err
	}

	return proto.Marshal(&pb.Output{
		Data: &pb.Output_ChannelListing{
			ChannelListing: &pb.ListChannelsResponse{Channels: response.GetChannels()},
		},
	})
}
