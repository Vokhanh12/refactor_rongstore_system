package application

import (
	"context"
	"fmt"
	"server/internal/iam/authz/domain/ports"
	iamErrors "server/internal/iam/errors"
	errors "server/pkg/errors"
)

type AuthenticateCommand struct {
	UserID     string
	TenantID   string
	Roles      []string
	Resource   string
	Action     string
	ResourceID string
}

type AuthenticateResult struct {
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

	if len(cmd.Roles) == 0 {
		return &AuthenticateResult{Allowed: false}, errors.New(iamErrors.AUTHORIZATION_ROLE_REQUIRED)
	}

	if cmd.Resource == "" || cmd.Action == "" {
		return &AuthenticateResult{Allowed: false}, errors.New(iamErrors.AUTHORIZATION_RESOURCE_OR_ACTION_REQUIRED)
	}

	permKey := fmt.Sprintf("%s:%s", cmd.Resource, cmd.Action)

	for _, role := range cmd.Roles {

		perms, err := u.rolePermissionCache.GetRolePermissions(ctx, role)
		if err != nil {
			return nil, errors.New(iamErrors.REDIS_UNAVAILABLE,
				errors.WithCauseDetail(err),
				errors.WithMessage("failed to load role permissions"))
		}

		for _, p := range perms {
			if p == permKey {
				return &AuthenticateResult{
					Allowed: true,
				}, nil
			}
		}
	}

	return &AuthenticateResult{
		Allowed: false,
	}, nil

}
