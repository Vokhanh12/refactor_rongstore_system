package handler

import (
	"context"

	commonv1 "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/core/common/v1/resources"
	authzrs "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/iam/authz/v1/resources"
	dp "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/application/dispatcher"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/adapter/mapper"
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
func (a *AuthzHandler) RoleMutate(
	ctx context.Context,
	req *authzrs.RoleMutateRequest,
) (*commonv1.MutateResponse, error) {

	results := make(
		[]*commonv1.MutateResult,
		0,
		len(req.Mutations),
	)
	var multipleErr *dp.MultipleError

	for _, mutation := range req.Mutations {
		op, err := mapper.ToMutateRoleCommand(mutation)
		if err != nil {
			return nil, err
		}

		result, err := a.roleMutateUsecase.Execute(
			ctx,
			op.Action,
			op.Payload,
		)

		if err != nil {
			if multipleErr == nil {
				multipleErr = dp.NewMultipleError(err)
			} else {
				multipleErr.Append(err)
			}
		}

		roleMapResult := mapper.FromMutateRoleResult(op)
		results = append(results, &roleMapResult)

	}

	// for _, r := range results {
	// 	if item != nil {
	// 		a.logger.Error(ctx, "iam_handler.role_mutate", item.Error.Internal, nil)
	// 	}
	// }

	return crm.BuildMutate(ctx, results), nil
}
