package handler

import (
	"context"

	commonv1 "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/core/common/v1/resources"
	authzrs "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/iam/authz/v1/resources"
	cm "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/adapter/mapper"
	dp "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/application/dispatcher"
	im "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/adapter/mapper"
	uc "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/usecase"
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
func (a *AuthzHandler) RoleMutate(
	ctx context.Context,
	req *authzrs.RoleMutateRequest,
) (*commonv1.MutateResponse, error) {

	results := make([]*commonv1.MutateResult, 0, len(req.Mutations))
	var multiErr dp.MultipleError

	for _, mutation := range req.Mutations {
		op, err := im.ToMutateRoleCommand(mutation)
		if err != nil {
			return nil, err
		}

		result, err := a.roleMutateUsecase.Execute(
			ctx,
			op.Action,
			op.Payload,
		)

		if err != nil {
			multiErr.Add(op.OpID, err)
			continue
		}

		mapped, err := im.FromMutateRoleResult(op, result)
		if err != nil {
			multiErr.Add(op.OpID, err)
			continue
		}

		results = append(results, &mapped)
	}

	return cm.BuildMutateResponse(ctx, results, &multiErr)
}
