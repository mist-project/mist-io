package message

import (
	"context"
	"fmt"
	"time"

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
