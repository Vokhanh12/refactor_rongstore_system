package handler

import (
	"context"

	comv1rs "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/core/common/v1/resources"
	authv1rs "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/iam/auth/v1/resources"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/adapter/mapper"
	uc "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/application/usecase"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/logger"
)

type AuthHandler struct {
	loginUsecase uc.LoginUsecase
	logger       logger.Logger
}

func NewAuthHandler(loginUc uc.LoginUsecase, logger logger.Logger) *AuthHandler {
	return &AuthHandler{
		loginUsecase: loginUc,
		logger:       logger,
	}
}

func (a *AuthHandler) Login(
	ctx context.Context,
	req *authv1rs.LoginRequest,
) (*comv1rs.BaseResponse, error) {

	cmd := mapper.ToLoginCommand(req)

	result, err := a.loginUsecase.Execute(ctx, cmd)
	if err != nil {
		return mapper.FromLoginResult(result), err
	}

	return mapper.FromLoginResult(result), nil
}
