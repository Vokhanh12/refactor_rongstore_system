package handler

import (
	"context"

	commonv1 "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/common/v1"
	resourcespb "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/iam/v1/resources"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/adapter/grpc"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/adapter/mappers"
)

var _ grpc.AuthzPort = (*AuthzHandler)(nil)

type AuthzHandler struct{}

func NewAuthzHandler() *AuthzHandler {
	return &AuthzHandler{}
}

// RoleMutate implements [grpc.AuthzPort].
func (a *AuthzHandler) RoleMutate(ctx context.Context, req *resourcespb.RoleMutateRequest) (*commonv1.BaseResponse, error) {
	batch := mappers.RoleMutateRequestToBatch(req)

	results := h.RoleMutateUsecase.Execute(ctx, batch)

	for _, item := range results.Items {
		if item.Error != nil {
			logger.LogBySeverity(ctx, "iam_handler.store_owner_mutate", item.Error)
		}

	}

	resDTO := mappers.RoleMutateResultToResponseDTO(results)
	return reshelper.BuildMutateResponse(ctx, resDTO), nil
}
