package handler

import (
	"context"

	commonv1 "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/core/common/v1/resources"
	authzrs "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/iam/authz/v1/resources"
	corem "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/adapter/mappers"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/adapter/mappers"
	uc "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/usecases"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/logger"
)

type AuthzHandler struct {
	roleMutateUsecase uc.MutateRoleUsecase
	logger            logger.Logger
}

func NewAuthzHandler(roleMutateUc uc.MutateRoleUsecase, logger logger.Logger) *AuthzHandler {
	return &AuthzHandler{
		roleMutateUsecase: roleMutateUc,
		logger:            logger,
	}
}

// RoleMutate implements [grpc.AuthzPort].
func (a *AuthzHandler) RoleMutate(ctx context.Context, req *authzrs.RoleMutateRequest) (*commonv1.MutateResponse, error) {
	batch := mappers.RoleMutateRequestToBatch(req)

	results := a.roleMutateUsecase.Execute(ctx, batch)

	for _, item := range results.Items {

		if item.Error != nil {
			a.logger.Error(ctx, "iam_handler.role_mutate", item.Error.Internal, nil)
		}

	}

	return corem.BuildMutateResponse(ctx, results, mappers.MapRoleActionProto(results.Items)), nil
}
