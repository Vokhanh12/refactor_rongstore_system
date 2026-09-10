package interceptor

import (
	"context"
	"errors"

	"buf.build/go/protovalidate"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/core/adapter/mapper"
	v "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/errors"
	aerr "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
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

		msg, ok := req.(proto.Message)
		if !ok {
			return nil, errors.New(
				"request does not implement proto.Message",
			)
		}

		err := validator.Validate(msg)

		if err == nil {
			return handler(ctx, req)
		}

		var validationErr *protovalidate.ValidationError

		if errors.As(err, &validationErr) {
			violations := mapper.ToValidationError(validationErr)

			return nil, aerr.New(v.VALIDATION_FAILED, aerr.WithAppendViolations(violations))
		}

		return nil, err
	}
}
