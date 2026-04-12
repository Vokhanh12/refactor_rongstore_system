package grpc

import (
	"context"

	commonv1 "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/common/v1"
	iamv1rs "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/iam/v1/resources"
)

type AuthzPort interface {
	RoleMutate(ctx context.Context, req *iamv1rs.StoreOwnerMutateRequest) (*commonv1.BaseResponse, error)
}
