package handler

import (
	"context"

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

func (a *AuthHandler) Login(ctx context.Context, req *authrs.LoginRequest) (*commonv1.LoginResponse, error) {

}
