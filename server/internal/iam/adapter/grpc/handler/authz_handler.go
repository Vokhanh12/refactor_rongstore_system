package handler

import (
	"context"

	commonv1 "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/common/v1"
	authzrs "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/iam/authz/v1/resources"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/adapter/mappers"
	uc "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/usecases"
	pkgcommonv1 "github.com/vokhanh12/refactor-rongstore-system/server/pkg/common/v1"
)

type AuthzHandler struct {
	roleMutateUsecase uc.MutateRoleUsecase
}

func NewAuthzHandler(roleMutateUc uc.MutateRoleUsecase) *AuthzHandler {
	return &AuthzHandler{
		roleMutateUsecase: roleMutateUc,
	}
}

// RoleMutate implements [grpc.AuthzPort].
func (a *AuthzHandler) RoleMutate(ctx context.Context, req *authzrs.RoleMutateRequest) (*commonv1.BaseResponse, error) {
	batch := mappers.RoleMutateRequestToBatch(req)

	results := a.roleMutateUsecase.Execute(ctx, batch)

	for _, item := range results.Items {
		if item.Error != nil {
			logger.LogBySeverity(ctx, "iam_handler.store_owner_mutate", item.Error)
		}

	}

	mappers.

	return pkgcommonv1.BuildMutateResponse(ctx, results), nil
}
