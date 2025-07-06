package faults

import (
	"context"
	"log/slog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	NotFoundMessage               = "Not Found"
	ValidationErrorMessage        = "Validation Error"
	DatabaseErrorMessage          = "Internal Server Error"
	AuthenticationErrorMessage    = "Unauthenticated"
	AuthorizationErrorMessage     = "Unauthorized"
	MessageSubscriberErrorMessage = "Message Subscriber Error"
	MarshallErrorMessage          = "Unprocessable Entity: Marshalling Error"
	UnknownErrorMessage           = "Internal Server Error"
	SocketNotFoundErrorMessage    = "Socket Not Found Error"
	SocketWriteErrorMessage       = "WebSocket Write Error"
	RedisErrorMessage             = "Redis Error"
)

func NotFoundError(root string, debugLevel slog.Level) *CustomError {
	return NewError(NotFoundMessage, root, codes.NotFound, debugLevel)
}

func ValidationError(root string, debugLevel slog.Level) *CustomError {
	return NewError(ValidationErrorMessage, root, codes.InvalidArgument, debugLevel)
}

func DatabaseError(root string, debugLevel slog.Level) *CustomError {
	return NewError(DatabaseErrorMessage, root, codes.Internal, debugLevel)
}

func AuthenticationError(root string, debugLevel slog.Level) *CustomError {
	return NewError(AuthenticationErrorMessage, root, codes.Unauthenticated, debugLevel)
}

func AuthorizationError(root string, debugLevel slog.Level) *CustomError {
	return NewError(AuthorizationErrorMessage, root, codes.PermissionDenied, debugLevel)
}

func UnknownError(root string, debugLevel slog.Level) *CustomError {
	return NewError(UnknownErrorMessage, root, codes.Unknown, debugLevel)
}

func MarshallError(root string, debugLevel slog.Level) *CustomError {
	return NewError(MarshallErrorMessage, root, codes.InvalidArgument, debugLevel)
}

func MessageSubscriberError(root string, debugLevel slog.Level) *CustomError {
	return NewError(MessageSubscriberErrorMessage, root, codes.Unknown, debugLevel)
}

func SocketNotFoundError(root string, debugLevel slog.Level) *CustomError {
	return NewError(SocketNotFoundErrorMessage, root, codes.NotFound, debugLevel)
}

func SocketWriteError(root string, debugLevel slog.Level) *CustomError {
	return NewError(SocketWriteErrorMessage, root, codes.Internal, debugLevel)
}

func RedisError(root string, debugLevel slog.Level) *CustomError {
	return NewError(RedisErrorMessage, root, codes.Internal, debugLevel)
}

func RpcCustomErrorHandler(ctx context.Context, err error) error {
	ce, ok := err.(*CustomError)

	if !ok {
		return status.Errorf(codes.Unknown, "%s", err.Error())
	}

	ce.LogError(ctx)
	return status.Errorf(ce.Code(), "%s", err.Error())
}
