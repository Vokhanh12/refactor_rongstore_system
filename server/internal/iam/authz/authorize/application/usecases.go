package application

import (
	"context"
	"fmt"
	rpdomain "server/internal/iam/authz/role_permission/domain"
	"server/pkg/errors"
)

type AuthorizeCommand struct {
	UserID     string
	TenantID   string
	Roles      []string
	Resource   string
	Action     string
	ResourceID string
}

type AuthorizeResult struct {
	Allowed bool
}

type AuthorizeUsecase struct {
	rolePermissionCache      rpdomain.RolePermissionCache
	rolePermissionRepository rpdomain.RolePermissionRepository
}

func NewAuthorizeUsecase(rpCache rpdomain.RolePermissionCache, rpPepository rpdomain.RolePermissionRepository) *AuthorizeUsecase {
	return &AuthorizeUsecase{
		rolePermissionCache:      rpCache,
		rolePermissionRepository: rpPepository,
	}
}

func (u *AuthorizeUsecase) Execute(ctx context.Context, cmd AuthorizeCommand) (*AuthorizeResult, *errors.AppError) {

	if len(cmd.Roles) == 0 {
		return &AuthorizeResult{Allowed: false},
			errors.New(iamErrors.AUTHORIZATION_ROLE_REQUIRED)
	}

	if cmd.Resource == "" || cmd.Action == "" {
		return &AuthorizeResult{Allowed: false},
			errors.New(iamErrors.AUTHORIZATION_RESOURCE_OR_ACTION_REQUIRED)
	}

	permKey := fmt.Sprintf("%s:%s", cmd.Resource, cmd.Action)

	for _, role := range cmd.Roles {

		perms, found, err := u.rolePermissionCache.GetPermissions(ctx, role)
		if err != nil {
			return nil, err
		}

		if !found {
			continue
		}

		for _, p := range perms {
			if p == permKey {
				return &AuthorizeResult{
					Allowed: true,
				}, nil
			}
		}
	}

	return &AuthorizeResult{
		Allowed: false,
	}, nil
}
