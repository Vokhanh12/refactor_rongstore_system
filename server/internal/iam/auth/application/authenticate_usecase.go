package application

import (
	"context"
)

type AuthenticateCommand struct {
	UserID     string
	TenantID   string
	Roles      []string
	Resource   string
	Action     string
	ResourceID string
}

type AuthenticateCommandResult struct {
	Allowed bool
}

type AuthenticateUsecase struct {
	rolePermissionCache ports.RolePermissionCache
}

func NewAuthenticateUsecase(cache ports.RolePermissionCache) *AuthenticateUsecase {
	return &AuthenticateUsecase{
		rolePermissionCache: cache,
	}
}

func (u *AuthenticateUsecase) Execute(ctx context.Context, cmd AuthenticateCommand) (*AuthenticateResult, *errors.AppError) {
}
