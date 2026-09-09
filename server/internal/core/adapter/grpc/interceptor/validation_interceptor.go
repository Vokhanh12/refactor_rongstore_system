package interceptor

import (
	"context"

	"buf.build/go/protovalidate"
	"google.golang.org/grpc"
)

func ValidationUnaryInterceptor(
	validator protovalidate.Validator,
) grpc.UnaryServerInterceptor {

	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {

		violations := validate(req)

		if err := validator.Validate(req); err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}
