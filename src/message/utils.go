package message

import (
	"context"
	"fmt"
	"mist-io/src/protos/v1/appserver"
	"mist-io/src/protos/v1/channel"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (wsc *WsConnection) SetupContext() (context.Context, context.CancelFunc) {
	// TODO: replace timeout in the future
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	grpcMetadata := metadata.Pairs(
		"authorization", fmt.Sprintf("Bearer %s", wsc.JwtToken),
	)

	ctx = metadata.NewOutgoingContext(ctx, grpcMetadata)
	return ctx, cancel
}

type GrpcClient interface {
	GetServerClient() appserver.AppserverServiceClient
	GetChannelClient() channel.ChannelServiceClient
}

type Client struct {
	Conn *grpc.ClientConn
}

func (c Client) GetServerClient() appserver.AppserverServiceClient {
	return appserver.NewAppserverServiceClient(c.Conn)
}

func (c Client) GetChannelClient() channel.ChannelServiceClient {
	return channel.NewChannelServiceClient(c.Conn)
}
